//go:build !nesgo
// +build !nesgo

// Package ppu provides PPU (Picture Processing Unit) functionality.
package ppu

import (
	"fmt"
	"image"

	"github.com/retroenv/nesgo/pkg/bus"
	"github.com/retroenv/nesgo/pkg/ppu/addressing"
	"github.com/retroenv/nesgo/pkg/ppu/mask"
	"github.com/retroenv/nesgo/pkg/ppu/nametable"
	"github.com/retroenv/nesgo/pkg/ppu/palette"
	"github.com/retroenv/nesgo/pkg/ppu/renderstate"
	"github.com/retroenv/nesgo/pkg/ppu/sprites"
	"github.com/retroenv/nesgo/pkg/ppu/status"
)

const (
	Width  = 256
	Height = 240
	FPS    = 60
)

// PPU represents the Picture Processing Unit.
type PPU struct {
	bus *bus.Bus

	control control

	fineX          uint16
	dataReadBuffer byte

	addressing  *addressing.Addressing
	mask        *mask.Mask
	nameTable   *nametable.NameTable
	nmi         *nmi
	palette     *palette.Palette
	renderState *renderstate.RenderState
	sprites     *sprites.Sprites
	status      *status.Status

	attributeTableByte byte
	lowTileByte        byte
	highTileByte       byte
	tileData           uint64

	back  *image.RGBA // rendering in progress image
	front *image.RGBA // visible image
}

// New returns a new PPU.
func New(bus *bus.Bus) *PPU {
	p := &PPU{
		bus: bus,
	}
	p.reset()
	return p
}

func (p *PPU) reset() {
	p.fineX = 0
	p.dataReadBuffer = 0

	p.back = image.NewRGBA(image.Rect(0, 0, Width, Height))
	p.front = image.NewRGBA(image.Rect(0, 0, Width, Height))

	p.addressing = addressing.New()
	p.mask = mask.New()
	p.nameTable = nametable.New(p.bus.Mapper)
	p.nmi = &nmi{}
	p.palette = palette.New()
	p.renderState = renderstate.New()
	p.sprites = sprites.New(p.bus.CPU, p.bus.Mapper, p.renderState, p.status)
	p.status = status.New()

	p.setControl(0x00)
	p.mask.Set(0x00)
}

func (p *PPU) readData() byte {
	address := p.addressing.Address()
	address &= 0x3FFF // valid addresses are $0000-$3FFF; higher addresses will be mirrored down

	// when reading data, the contents of an internal read buffer is returned and the buffer
	// gets updated with the newly read data
	data := p.dataReadBuffer

	switch {
	case address >= 0x2000 && address < 0x3F00:
		p.dataReadBuffer = p.nameTable.Read(address, p.bus.Cartridge.Mirror)

	case address >= 0x3F00:
		p.dataReadBuffer = p.palette.Read(address)
		// Palette data reads are unbuffered, $3F00-$3FFF are Palette RAM indexes and mirrors of it
		data = p.dataReadBuffer

	default:
		panic(fmt.Sprintf("unhandled ppu read at address: 0x%04X", address))
	}

	// TODO handle special case of reading during rendering
	p.addressing.Increment(p.control.VRAMIncrement)
	return data
}

func (p *PPU) writeData(value byte) {
	address := p.addressing.Address()
	address &= 0x3FFF // valid addresses are $0000-$3FFF; higher addresses will be mirrored down

	switch {
	case address >= 0x2000 && address < 0x3F00:
		p.nameTable.Write(address, value, p.bus.Cartridge.Mirror)

	case address >= 0x3F00:
		p.palette.Write(address, value)

	default:
		panic(fmt.Sprintf("unhandled ppu write at address: 0x%04X", address))
	}

	p.addressing.Increment(p.control.VRAMIncrement)
}

func (p *PPU) setScroll(value byte) {
	if !p.addressing.Latch() {
		p.fineX = uint16(value) & 0x07
	}
	p.addressing.SetScroll(value)
}

func (p *PPU) getStatus() byte {
	p.addressing.ClearLatch()

	p.status.SetVerticalBlank(p.nmi.occurred)
	p.nmi.occurred = false
	p.nmi.change()

	value := p.status.Value()
	return value
}

// mirroredAddressToBase converts the mirrored addresses to the base address.
// PPU registers are mirrored in every 8 bytes from $2008 through $3FFF.
func mirroredAddressToBase(address uint16) uint16 {
	base := 0x2000 + address&0b00000111
	return base
}
