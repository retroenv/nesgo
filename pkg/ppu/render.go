//go:build !nesgo
// +build !nesgo

package ppu

import (
	"image"
)

// Image returns the rendered image to display.
func (p *PPU) Image() *image.RGBA {
	return p.front
}

// Step executes a PPU cycle.
func (p *PPU) Step() {
	p.nmi.checkTrigger(p.bus.CPU)
	p.renderState.Tick(p.mask)

	if p.mask.RenderBackground || p.mask.RenderSprites {
		p.renderBackground()
		p.sprites.Render()
	}

	if p.renderState.Cycle() != 1 {
		return
	}

	switch p.renderState.ScanLine() {
	case 241:
		// the vertical blank flag of the PPU is set at tick 1 (the second tick) of scanline 241,
		// where the vertical blank NMI also occurs
		p.setVerticalBlank()

	case 261:
		p.clearVerticalBlank()
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
		p.fetchCycle(cycle)
	}

	if preLine && cycle >= 280 && cycle <= 304 {
		p.addressing.CopyY()
	}

	if renderLine {
		p.renderLine(cycle, fetchCycle)
	}
}

func (p *PPU) fetchCycle(cycle int) {
	p.tileData <<= 4
	switch cycle % 8 {
	case 0:
		p.storeTileData()
	case 1:
		p.nameTable.Fetch(p.addressing.Address())
	case 3:
		p.fetchAttributeTableByte()
	case 5:
		p.fetchLowTileByte()
	case 7:
		p.fetchHighTileByte()
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
	x := p.renderState.Cycle() - 1
	y := p.renderState.ScanLine()

	background := p.backgroundPixel()
	sprite, spriteZeroHit, spriteColor := p.sprites.Pixel(p.mask)

	if x < 8 {
		if !p.mask.RenderBackgroundLeft {
			background = 0
		}
		if !p.mask.RenderSpritesLeft {
			spriteColor = 0
		}
	}

	b := background%4 != 0
	s := spriteColor%4 != 0
	var color byte

	switch {
	case !b && !s:
		color = 0
	case !b && s:
		color = spriteColor | 0x10
	case b && !s:
		color = background
	default:
		if spriteZeroHit && x < 255 {
			p.status.SetSpriteZeroHit(true)
		}
		priority := sprite.Priority()
		if priority == 0 {
			color = spriteColor | 0x10
		} else {
			color = background
		}
	}

	colorIndex := p.palette.Read(uint16(color))
	colorIndex %= 64
	c := colors[colorIndex]
	p.back.SetRGBA(x, y, c)
}
