//go:build !nesgo
// +build !nesgo

package ppu

type mask struct {
	value byte

	Grayscale            bool
	RenderBackgroundLeft bool
	RenderSpritesLeft    bool
	RenderBackground     bool
	RenderSprites        bool
	EnhanceRed           bool
	EnhanceGreen         bool
	EnhanceBlue          bool
}

func (p *PPU) setMask(value byte) {
	p.mask.value = value

	p.mask.Grayscale = value&MASK_MONO != 0
	p.mask.RenderBackgroundLeft = value&MASK_BG_CLIP != 0
	p.mask.RenderSpritesLeft = value&MASK_SPR_CLIP != 0
	p.mask.RenderBackground = value&MASK_BG != 0
	p.mask.RenderSprites = value&MASK_SPR != 0
	p.mask.EnhanceRed = value&MASK_TINT_RED != 0
	p.mask.EnhanceGreen = value&MASK_TINT_GREEN != 0
	p.mask.EnhanceBlue = value&MASK_TINT_BLUE != 0
}
