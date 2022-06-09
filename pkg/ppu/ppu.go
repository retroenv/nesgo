//go:build !nesgo
// +build !nesgo

// Package ppu provides PPU (Picture Processing Unit) functionality.
package ppu

import (
	"fmt"
	"image"
	"time"

	"github.com/retroenv/nesgo/pkg/mapper"
)

const (
	Width  = 256
	Height = 240
	fps    = 60
)

// PPU implements the Picture Processing Unit.
type PPU struct {
	ptr    uint16
	ram    ram
	mapper mapper.Mapper

	image *image.RGBA
}

// New returns a new PPU.
func New(ram ram, mapper mapper.Mapper) *PPU {
	p := &PPU{
		ram:    ram,
		mapper: mapper,
	}
	p.reset()
	return p
}

func (p *PPU) reset() {
	p.ptr = 0
	p.ram.Reset()
	p.image = image.NewRGBA(image.Rect(0, 0, Width, Height))
}

// Image returns the rendered image to display.
func (p *PPU) Image() *image.RGBA {
	return p.image
}

// ReadRegister reads from a PPU memory address.
func (p *PPU) ReadRegister(address uint16) byte {
	switch {
	case address < 0x2000:
		return p.mapper.ReadMemory(address)

	case address == PPU_STATUS:
		b := p.ram.ReadMemory(address)
		p.ptr = 0
		p.clearVBlank()
		return b

	default:
		panic(fmt.Sprintf("unhandled ppu read at address: 0x%04X", address))
	}
}

// WriteRegister writes to a PPU memory address.
func (p *PPU) WriteRegister(address uint16, value byte) {
	switch {
	case address < 0x2000:
		p.mapper.WriteMemory(address, value)

	case address == PPU_CTRL, address == PPU_MASK:
		p.ram.WriteMemory(address, value)

	case address == PPU_SCROLL:
		// TODO support

	case address == PPU_ADDR:
		p.ptr = p.ptr<<8 | uint16(value)

	case address == PPU_DATA:
		if p.ptr > 0x4000 {
			panic(fmt.Sprintf("ppu data address 0x%04X is out of range", p.ptr))
		}

		p.ram.WriteMemory(p.ptr, value)

		// TODO handle special addresses
		// TODO handle vram delta
		p.ptr++

	default:
		panic(fmt.Sprintf("unhandled ppu write at address: 0x%04X", address))
	}
}

func (p *PPU) setVBlank() {
	status := p.ram.ReadMemory(PPU_STATUS)
	status |= 0x80
	p.ram.WriteMemory(PPU_STATUS, status)
	// TODO handle NMI
}

func (p *PPU) clearVBlank() {
	status := p.ram.ReadMemory(PPU_STATUS)
	status &= 0x7f
	p.ram.WriteMemory(PPU_STATUS, status)
}

// StartRender starts the rendering process.
func (p *PPU) StartRender() {
	p.setVBlank()
	time.Sleep(time.Second / fps)
}

// FinishRender finishes the rendering process.
func (p *PPU) FinishRender() {
	status := p.ram.ReadMemory(PPU_STATUS)
	status &= 0xbf
	p.ram.WriteMemory(PPU_STATUS, status)
	p.clearVBlank()
}

// RenderScreen renders the screen into the internal image.
func (p *PPU) RenderScreen() {
	mask := p.ram.ReadMemory(PPU_MASK)
	if mask&MASK_BG != 0 {
		p.renderBackground()
	}
}

func (p *PPU) renderBackground() {
	idx := int(p.ram.ReadMemory(PALETTE_START))
	idx %= len(colors)
	c := colors[idx]

	for y := 0; y < Height; y++ {
		for x := 0; x < Width; x++ {
			p.image.SetRGBA(x, y, c)
		}
	}
}
