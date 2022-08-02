package bus

import "github.com/retroenv/nesgo/pkg/cartridge"

// Mapper represents a mapper memory access interface.
type Mapper interface {
	BasicMemory

	MirrorMode() cartridge.MirrorMode
	Name() string
}
