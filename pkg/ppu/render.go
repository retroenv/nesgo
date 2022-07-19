//go:build !nesgo
// +build !nesgo

package ppu

import (
	"image"
	"time"
)

// StartRender starts the rendering process.
func (p *PPU) StartRender() {
	p.setVBlank()
	time.Sleep(time.Second / FPS)
}

// FinishRender finishes the rendering process.
func (p *PPU) FinishRender() {
	status := p.ram.Read(PPU_STATUS)
	status &= 0xbf
	p.ram.Write(PPU_STATUS, status)
	p.clearVBlank()
}

// RenderScreen renders the screen into the internal image.
func (p *PPU) RenderScreen() {
	if p.mask.RenderBackground {
		p.renderBackground()
	}
}

// Image returns the rendered image to display.
func (p *PPU) Image() *image.RGBA {
	return p.image
}

func (p *PPU) renderBackground() {
	idx := int(p.ram.Read(PALETTE_START))
	idx %= len(colors)
	c := colors[idx]

	for y := 0; y < Height; y++ {
		for x := 0; x < Width; x++ {
			p.image.SetRGBA(x, y, c)
		}
	}
}
