package bus

import "github.com/retroenv/retrogolib/arch/nes/cartridge"

// MapperState contains the current state of the mapper.
type MapperState struct {
	ID   byte   `json:"id"`
	Name string `json:"name"`

	ChrWindows []int `json:"chrWindows"`
	PrgWindows []int `json:"prgWindows"`
}

// Mapper represents a mapper memory access interface.
type Mapper interface {
	BasicMemory

	MirrorMode() cartridge.MirrorMode
	State() MapperState
}
