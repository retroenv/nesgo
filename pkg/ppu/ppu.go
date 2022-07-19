//go:build !nesgo
// +build !nesgo

// Package ppu provides PPU (Picture Processing Unit) functionality.
package ppu

import (
	"fmt"
	"image"

	"github.com/retroenv/nesgo/pkg/bus"
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
	mask    mask
	status  status

	addressLatch bool
	vramAddress  register
	tempAddress  register

	fineX          uint16
	dataReadBuffer byte

	palette     *palette
	renderState *renderState

	back  *image.RGBA // rendering in progress image
	front *image.RGBA // visible image
}

// New returns a new PPU.
func New(bus *bus.Bus) *PPU {
	p := &PPU{
		bus:     bus,
		palette: &palette{},
	}
	p.reset()
	return p
}

func (p *PPU) reset() {
	p.addressLatch = false
	p.vramAddress = register{}
	p.tempAddress = register{}

	p.fineX = 0
	p.dataReadBuffer = 0

	p.back = image.NewRGBA(image.Rect(0, 0, Width, Height))
	p.front = image.NewRGBA(image.Rect(0, 0, Width, Height))

	p.palette.reset()
	p.renderState = newRenderState()

	p.setControl(0x00)
	p.setMask(0x00)
	p.status = status{}
}

func (p *PPU) readData() byte {
	address := p.vramAddress.address()
	address &= 0x3FFF // valid addresses are $0000-$3FFF; higher addresses will be mirrored down

	// when reading data, the contents of an internal read buffer is returned and the buffer
	// gets updated with the newly read data
	data := p.dataReadBuffer

	switch {
	case address >= 0x2000 && address < 0x3F00:
		p.dataReadBuffer = 0 // TODO
	case address >= 0x3F00:
		p.dataReadBuffer = p.palette.read(address)
		// palette data reads are unbuffered, $3F00-$3FFF are Palette RAM indexes and mirrors of it
		data = p.dataReadBuffer
	default:
		panic(fmt.Sprintf("unhandled ppu read at address: 0x%04X", address))
	}

	// TODO handle special case of reading during rendering
	p.vramAddress.increment(p.control.VRAMIncrement)
	return data
}

func (p *PPU) writeData(value byte) {
	address := p.vramAddress.address()
	address &= 0x3FFF // valid addresses are $0000-$3FFF; higher addresses will be mirrored down

	switch {
	case address >= 0x2000 && address < 0x3F00:
		// TODO
	case address >= 0x3F00:
		p.palette.write(address, value)
	default:
		panic(fmt.Sprintf("unhandled ppu write at address: 0x%04X", address))
	}

	p.vramAddress.increment(p.control.VRAMIncrement)
}

func (p *PPU) setScroll(value uint16) {
	if p.addressLatch {
		p.tempAddress.FineY = value & 0x07
		p.tempAddress.CoarseY = value >> 3
	} else {
		p.fineX = value & 0x07
		p.tempAddress.CoarseX = value >> 3
	}

	p.addressLatch = !p.addressLatch
}

func (p *PPU) setAddress(value uint16) {
	if p.addressLatch {
		address := p.tempAddress.address() & 0xFF00
		address |= value
		p.tempAddress.set(address)
		p.vramAddress = p.tempAddress
	} else {
		address := p.tempAddress.address() & 0x00FF
		address |= value << 8
		p.tempAddress.set(address)
	}

	p.addressLatch = !p.addressLatch
}

// mirroredAddressToBase converts the mirrored addresses to the base address.
// PPU registers are mirrored in every 8 bytes from $2008 through $3FFF.
func mirroredAddressToBase(address uint16) uint16 {
	base := 0x2000 + address&0b00000111
	return base
}
