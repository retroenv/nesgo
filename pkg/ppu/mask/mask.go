//go:build !nesgo
// +build !nesgo

// Package mask contains the PPU mask.
package mask

// Mask implements a PPU mask fields manager.
type Mask struct {
	value byte // cached value since the fields are never modified directly

	Grayscale            bool
	RenderBackgroundLeft bool
	RenderSpritesLeft    bool
	RenderBackground     bool
	RenderSprites        bool
	EnhanceRed           bool
	EnhanceGreen         bool
	EnhanceBlue          bool
}

// New returns a new mask manager.
func New() *Mask {
	return &Mask{}
}

// Set and extract the mask fields from given byte value.
func (m *Mask) Set(value byte) {
	m.value = value

	m.Grayscale = value&MASK_MONO != 0
	m.RenderBackgroundLeft = value&MASK_BG_CLIP != 0
	m.RenderSpritesLeft = value&MASK_SPR_CLIP != 0
	m.RenderBackground = value&MASK_BG != 0
	m.RenderSprites = value&MASK_SPR != 0
	m.EnhanceRed = value&MASK_TINT_RED != 0
	m.EnhanceGreen = value&MASK_TINT_GREEN != 0
	m.EnhanceBlue = value&MASK_TINT_BLUE != 0
}

// Value returns the mask fields encoded as byte.
func (m Mask) Value() byte {
	return m.value
}