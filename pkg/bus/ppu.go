package bus

import "image"

// PPU represents the Picture Processing Unit.
type PPU interface {
	BasicMemory

	Image() *image.RGBA
	Step()
}
