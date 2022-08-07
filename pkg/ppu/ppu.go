//go:build !nesgo

// Package ppu provides PPU (Picture Processing Unit) functionality.
package ppu

import (
	"github.com/retroenv/nesgo/pkg/bus"
	"github.com/retroenv/nesgo/pkg/ppu/addressing"
	"github.com/retroenv/nesgo/pkg/ppu/control"
	"github.com/retroenv/nesgo/pkg/ppu/mask"
	"github.com/retroenv/nesgo/pkg/ppu/memory"
	"github.com/retroenv/nesgo/pkg/ppu/nmi"
	"github.com/retroenv/nesgo/pkg/ppu/palette"
	"github.com/retroenv/nesgo/pkg/ppu/renderstate"
	"github.com/retroenv/nesgo/pkg/ppu/screen"
	"github.com/retroenv/nesgo/pkg/ppu/sprites"
	"github.com/retroenv/nesgo/pkg/ppu/status"
	"github.com/retroenv/nesgo/pkg/ppu/tiles"
)

const (
	FPS    = 60
	Height = screen.Height
	Width  = screen.Width
)

// PPU represents the Picture Processing Unit.
type PPU struct {
	bus *bus.Bus

	fineX          uint16
	dataReadBuffer byte

	addressing  *addressing.Addressing
	control     *control.Control
	mask        *mask.Mask
	memory      *memory.Memory
	nmi         *nmi.Nmi
	palette     *palette.Palette
	renderState *renderstate.RenderState
	screen      *screen.Screen
	sprites     *sprites.Sprites
	status      *status.Status
	tiles       *tiles.Tiles
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

	p.addressing = addressing.New()
	p.mask = mask.New()
	p.nmi = nmi.New()
	p.palette = palette.New()
	p.renderState = renderstate.New()
	p.screen = screen.New()
	p.status = status.New()

	p.memory = memory.New(p.bus.Mapper, p.bus.NameTable, p.palette)
	p.sprites = sprites.New(p.bus.CPU, p.bus.Mapper, p.renderState, p.status)

	p.tiles = tiles.New(p.addressing, p.memory, p.bus.NameTable)

	p.control = control.New(p.addressing, p.nmi, p.sprites, p.tiles)
}

func (p *PPU) readData() byte {
	address := p.addressing.Address()
	address &= 0x3FFF // valid addresses are $0000-$3FFF; higher addresses will be mirrored down

	// when reading data, the contents of an internal read buffer is returned and the buffer
	// gets updated with the newly read data
	data := p.dataReadBuffer

	p.dataReadBuffer = p.memory.Read(address)

	if address >= 0x3F00 {
		// Palette data reads are unbuffered, $3F00-$3FFF are Palette RAM indexes and mirrors of it
		data = p.dataReadBuffer
	}

	// TODO handle special case of reading during rendering
	p.addressing.Increment(p.control.VRAMIncrement)
	return data
}

func (p *PPU) getStatus() byte {
	p.addressing.ClearLatch()

	p.status.SetVerticalBlank(p.nmi.Occurred())
	p.nmi.SetOccurred(false)

	value := p.status.Value()
	return value
}
