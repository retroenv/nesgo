//go:build !nesgo
// +build !nesgo

package ppu

import (
	"image"
	"time"
)

// StartRender starts the rendering process.
func (p *PPU) StartRender() {
	time.Sleep(time.Second / FPS)
}

// FinishRender finishes the rendering process.
func (p *PPU) FinishRender() {
}

// RenderScreen renders the screen into the internal image.
func (p *PPU) RenderScreen() {
	// TODO sync with cpu
}

// Image returns the rendered image to display.
func (p *PPU) Image() *image.RGBA {
	return p.front
}
