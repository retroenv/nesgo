// Package disasmoptions implements disassembler options that are passed down to the actual
// assembly writer.
package disasmoptions

// Options defines options to control the disassembler.
type Options struct {
	Assembler string

	HexComments bool
	ZeroBytes   bool
}
