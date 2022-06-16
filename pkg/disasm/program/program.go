// Package program represents an NES program.
package program

// Offset defines the content of an offset in a program that
// can represent data or code.
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
	Offsets  []Offset
	Handlers Handlers
}

// New creates a new program initialize with a program code size.
func New(size int) *Program {
	return &Program{
		Offsets: make([]Offset, size),
	}
}
