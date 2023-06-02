package bus

import (
	"image"

	"github.com/retroenv/retrogolib/arch/nes/cartridge"
)

// PPU represents the Picture Processing Unit.
type PPU interface {
	BasicMemory

	Image() *image.RGBA
	Palette() Palette
	Step(cycles int)
}

// Palette represents the PPU palette.
type Palette interface {
	BasicMemory

	Data() [32]byte
}

// NameTable represents a name table interface.
type NameTable interface {
	BasicMemory

	Data() [4][]byte
	MirrorMode() cartridge.MirrorMode
	SetMirrorMode(mirrorMode cartridge.MirrorMode)
	SetVRAM(vram []byte)

	Fetch(address uint16)
	Value() byte
}
