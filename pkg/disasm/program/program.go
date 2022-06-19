// Package program represents an NES program.
package program

import "github.com/retroenv/nesgo/pkg/cartridge"

// Offset defines the content of an offset in a program that can represent data or code.
type Offset struct {
	IsCallTarget bool // opcode is target of a jsr call, indicating a subroutine

	Label  string // name of label or subroutine if identified as a jump target
	Output string
}

// Handlers defines the handlers that the NES can jump to.
type Handlers struct {
	NMI   string
	Reset string
	IRQ   string
}

// Program defines an NES program that contains code or data.
type Program struct {
	PRG       []Offset // PRG-ROM banks
	CHR       []byte   // CHR-ROM banks
	Handlers  Handlers
	Mapper    byte
	Constants map[string]uint16
}

// New creates a new program initialize with a program code size.
func New(cart *cartridge.Cartridge) *Program {
	return &Program{
		PRG:       make([]Offset, len(cart.PRG)),
		CHR:       cart.CHR,
		Mapper:    cart.Mapper,
		Constants: map[string]uint16{},
	}
}