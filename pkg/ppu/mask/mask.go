//go:build !nesgo

// Package mask contains the PPU mask.
package mask

// Mask implements a PPU mask fields manager.
type Mask struct {
	value byte // cached value since the fields are never modified directly

	Grayscale            bool
	renderBackgroundLeft bool
	renderSpritesLeft    bool
	renderBackground     bool
	renderSprites        bool
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
	m.renderBackgroundLeft = value&MASK_BG_CLIP != 0
	m.renderSpritesLeft = value&MASK_SPR_CLIP != 0
	m.renderBackground = value&MASK_BG != 0
	m.renderSprites = value&MASK_SPR != 0
	m.EnhanceRed = value&MASK_TINT_RED != 0
	m.EnhanceGreen = value&MASK_TINT_GREEN != 0
	m.EnhanceBlue = value&MASK_TINT_BLUE != 0
}

// Value returns the mask fields encoded as byte.
func (m Mask) Value() byte {
	return m.value
}

// RenderBackground returns a flag whether the background should be rendered.
func (m Mask) RenderBackground() bool {
	return m.renderBackground
}

// RenderSprites returns a flag whether the sprites should be rendered.
func (m Mask) RenderSprites() bool {
	return m.renderSprites
}

// RenderBackgroundLeft returns a flag whether the background should be rendered in leftmost 8 pixels of screen.
func (m Mask) RenderBackgroundLeft() bool {
	return m.renderBackgroundLeft
}

// RenderSpritesLeft returns a flag whether the sprites should be rendered in leftmost 8 pixels of screen.
func (m Mask) RenderSpritesLeft() bool {
	return m.renderSpritesLeft
}
