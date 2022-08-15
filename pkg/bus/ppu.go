package bus

import "image"

// PPU represents the Picture Processing Unit.
type PPU interface {
	BasicMemory

	Image() *image.RGBA
	Palette() Palette
	Step(cycles int)
}

// Palette returns the PPU palette.
type Palette interface {
	BasicMemory

	Data() [32]byte
}
