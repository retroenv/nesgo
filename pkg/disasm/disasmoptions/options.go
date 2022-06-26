// Package disasmoptions implements disassembler options that are passed down to the actual
// assembly writer.
package disasmoptions

// Options defines options to control the disassembler.
type Options struct {
	Assembler string // what assembler to use

	HexComments    bool
	OffsetComments bool
	ZeroBytes      bool
}

// New returns a new options instance with default options.
func New() Options {
	return Options{
		Assembler: "ca65",

		HexComments:    true,
		OffsetComments: true,
	}
}
