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

var ppu *PPU

// PPU implements the Picture Processing Unit.
type PPU struct {
	ptr uint16

	image *image.RGBA
}

func newPPU() *PPU {
	p := &PPU{}
	p.reset()
	return p
}

func (p *PPU) reset() {
	p.ptr = 0
	p.image = image.NewRGBA(image.Rect(0, 0, width, height))
}

func (p *PPU) readRegister(address uint16) byte {
	switch address {
	case PPU_STATUS:
		b := ram.readMemory(address)
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
		ram.writeMemory(address, value)

	case PPU_ADDR:
		p.ptr = p.ptr<<8 | uint16(value)

	case PPU_DATA:
		if p.ptr > 0x4000 {
			panic(fmt.Sprintf("ppu data address 0x%04X is out of range", p.ptr))
		}

		ram.writeMemory(p.ptr, value)

		// TODO handle special addresses
		// TODO handle vram delta
		p.ptr++

	default:
		panic(fmt.Sprintf("unhandled ppu write at address: 0x%04X", address))
	}
}

// nolint: unused
func (p *PPU) setVBlank() {
	status := ram.readMemory(PPU_STATUS)
	status |= 0x80
	ram.writeMemory(PPU_STATUS, status)
	// TODO handle NMI
}

func (p *PPU) clearVBlank() {
	status := ram.readMemory(PPU_STATUS)
	status &= 0x7f
	ram.writeMemory(PPU_STATUS, status)
}

func (p *PPU) startRender() {
	ppu.setVBlank()
	time.Sleep(time.Second / fps)
}

func (p *PPU) finishRender() {
	status := ram.readMemory(PPU_STATUS)
	status &= 0xbf
	ram.writeMemory(PPU_STATUS, status)
	ppu.clearVBlank()
}

func (p *PPU) renderScreen() {
	mask := ram.readMemory(PPU_MASK)
	if mask&MASK_BG != 0 {
		p.renderBackground()
	}
}

func (p *PPU) renderBackground() {
	idx := int(ram.readMemory(PALETTE_START))
	idx %= len(colors)
	c := colors[idx]

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			p.image.SetRGBA(x, y, c)
		}
	}
}
