//go:build !nesgo
// +build !nesgo

package nes

import (
	"fmt"
	"image"
	"time"
)

const (
	width  = 256
	height = 240
	fps    = 60
)

// PPU implements the Picture Processing Unit.
type PPU struct {
	ptr uint16
	ram *RAM

	image *image.RGBA
}

func newPPU() *PPU {
	p := &PPU{
		ram: newRAM(0x2000),
	}
	p.reset()
	return p
}

func (p *PPU) reset() {
	p.ptr = 0
	p.ram.reset()
	p.image = image.NewRGBA(image.Rect(0, 0, width, height))
}

func (p *PPU) readRegister(address uint16) byte {
	switch address {
	case PPU_STATUS:
		b := p.ram.readMemory(address)
		p.ptr = 0
		p.clearVBlank()
		return b

	default:
		panic(fmt.Sprintf("unhandled ppu read at address: 0x%04X", address))
	}
}

func (p *PPU) writeRegister(address uint16, value byte) {
	switch address {
	case PPU_CTRL, PPU_MASK:
		p.ram.writeMemory(address, value)

	case PPU_ADDR:
		p.ptr = p.ptr<<8 | uint16(value)

	case PPU_DATA:
		if p.ptr > 0x4000 {
			panic(fmt.Sprintf("ppu data address 0x%04X is out of range", p.ptr))
		}

		p.ram.writeMemory(p.ptr, value)

		// TODO handle special addresses
		// TODO handle vram delta
		p.ptr++

	default:
		panic(fmt.Sprintf("unhandled ppu write at address: 0x%04X", address))
	}
}

// nolint: unused
func (p *PPU) setVBlank() {
	status := p.ram.readMemory(PPU_STATUS)
	status |= 0x80
	p.ram.writeMemory(PPU_STATUS, status)
	// TODO handle NMI
}

func (p *PPU) clearVBlank() {
	status := p.ram.readMemory(PPU_STATUS)
	status &= 0x7f
	p.ram.writeMemory(PPU_STATUS, status)
}

func (p *PPU) startRender() {
	p.setVBlank()
	time.Sleep(time.Second / fps)
}

func (p *PPU) finishRender() {
	status := p.ram.readMemory(PPU_STATUS)
	status &= 0xbf
	p.ram.writeMemory(PPU_STATUS, status)
	p.clearVBlank()
}

func (p *PPU) renderScreen() {
	mask := p.ram.readMemory(PPU_MASK)
	if mask&MASK_BG != 0 {
		p.renderBackground()
	}
}

func (p *PPU) renderBackground() {
	idx := int(p.ram.readMemory(PALETTE_START))
	idx %= len(colors)
	c := colors[idx]

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			p.image.SetRGBA(x, y, c)
		}
	}
}
