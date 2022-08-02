package bus

import "github.com/retroenv/nesgo/pkg/cartridge"

// NameTable represents a name table interface.
type NameTable interface {
	BasicMemory

	MirrorMode() cartridge.MirrorMode
	SetMirrorMode(mirrorMode cartridge.MirrorMode)
	SetVRAM(vram []byte)

	Fetch(address uint16)
	Value() byte
}
