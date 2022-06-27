// Package program represents an NES program.
package program

import (
	"github.com/retroenv/nesgo/pkg/cartridge"
)

// OffsetType defines the type of a program offset.
type OffsetType uint8

// addressing modes.
const (
	UnknownOffset OffsetType = 0
	CodeOffset    OffsetType = 1 << iota
	DataOffset
	CallTarget // opcode is target of a jsr call, indicating a subroutine
)

// Offset defines the content of an offset in a program that can represent data or code.
type Offset struct {
	OpcodeBytes []byte // data byte or all opcode bytes that are part of the instruction

	Type OffsetType

	Label   string // name of label or subroutine if identified as a jump target
	Code    string // asm output of this instruction
	Comment string
}

// Handlers defines the handlers that the NES can jump to.
type Handlers struct {
	NMI   string
	Reset string
	IRQ   string
}

// Program defines an NES program that contains code or data.
type Program struct {
	PRG     []Offset // PRG-ROM banks
	CHR     []byte   // CHR-ROM banks
	RAM     byte     // PRG-RAM banks
	Trainer []byte

	Handlers    Handlers
	Battery     byte
	Mirror      byte
	Mapper      byte
	VideoFormat byte

	Constants map[string]uint16
}

// New creates a new program initialize with a program code size.
func New(cart *cartridge.Cartridge) *Program {
	return &Program{
		PRG:       make([]Offset, len(cart.PRG)),
		CHR:       cart.CHR,
		RAM:       cart.RAM,
		Battery:   cart.Battery,
		Mapper:    cart.Mapper,
		Mirror:    cart.Mirror,
		Trainer:   cart.Trainer,
		Constants: map[string]uint16{},
	}
}
