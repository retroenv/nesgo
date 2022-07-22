//go:build !nesgo
// +build !nesgo

package ppu

import (
	"image"
)

// Image returns the rendered image to display.
func (p *PPU) Image() *image.RGBA {
	return p.screen.Image()
}

// Step executes a PPU cycle.
func (p *PPU) Step() {
	p.nmi.Trigger(p.bus.CPU)
	p.renderState.Tick(p.mask)

	if p.mask.RenderBackground() || p.mask.RenderSprites() {
		p.renderBackground()
		// sprite evaluation occurs if either the sprite layer or background layer is enabled
		p.sprites.Render()
	}

	if p.renderState.Cycle() != 1 {
		return
	}

	switch p.renderState.ScanLine() {
	case 241:
		// the vertical blank flag of the PPU is set at tick 1 (the second tick) of scanline 241,
		// where the vertical blank NMI also occurs
		p.screen.FinishRendering()
		p.nmi.SetOccurred(true)

	case 261:
		p.nmi.SetOccurred(false)
		p.status.SetSpriteOverflow(false)
		p.status.SetSpriteZeroHit(false)
	}
}

func (p *PPU) renderBackground() {
	cycle := p.renderState.Cycle()
	scanLine := p.renderState.ScanLine()

	preLine := scanLine == 261
	visibleLine := scanLine < 240
	renderLine := preLine || visibleLine

	// cycle 0 is an idle cycle
	preFetchCycle := cycle >= 321 && cycle <= 336
	visibleCycle := cycle >= 1 && cycle <= 256
	fetchCycle := preFetchCycle || visibleCycle

	if visibleLine && visibleCycle {
		p.renderPixel()
	}

	if renderLine && fetchCycle {
		p.tiles.FetchCycle(cycle)
	}

	if preLine && cycle >= 280 && cycle <= 304 {
		p.addressing.CopyY()
	}

	if renderLine {
		p.renderLine(cycle, fetchCycle)
	}
}

func (p *PPU) renderLine(cycle int, fetchCycle bool) {
	if fetchCycle && cycle%8 == 0 {
		p.addressing.IncrementX()
	}
	if cycle == 256 {
		p.addressing.IncrementY()
	}
	if cycle == 257 {
		p.addressing.CopyX()
	}
}

func (p *PPU) renderPixel() {
	var backgroundColor, spriteColor byte
	if p.mask.RenderBackground() {
		backgroundColor = p.tiles.BackgroundPixel(p.fineX)
	}

	var spritePriority, spriteZeroHit bool
	if p.mask.RenderSprites() {
		spritePriority, spriteZeroHit, spriteColor = p.sprites.Pixel()
	}

	x := p.renderState.Cycle() - 1
	if x < 8 {
		if !p.mask.RenderBackgroundLeft() {
			backgroundColor = 0
		}
		if !p.mask.RenderSpritesLeft() {
			spriteColor = 0
		}
	}

	hasBackground := backgroundColor%4 != 0
	hasSprite := spriteColor%4 != 0
	var paletteIndex byte

	switch {
	case !hasBackground && hasSprite:
		paletteIndex = spriteColor | 0x10

	case hasBackground && !hasSprite:
		paletteIndex = backgroundColor

	case hasBackground && hasSprite:
		if spriteZeroHit && x < 255 {
			p.status.SetSpriteZeroHit(true)
		}

		if spritePriority {
			paletteIndex = spriteColor | 0x10
		} else {
			paletteIndex = backgroundColor
		}
	}

	colorIndex := p.palette.Read(uint16(paletteIndex))
	colorIndex %= 64
	color := colors[colorIndex]
	y := p.renderState.ScanLine()
	p.screen.SetPixel(x, y, color)
}
