package disasm

// Options defines options to control the disassembler.
type Options struct {
	Assembler string

	HexComments bool
	ZeroBytes   bool
}
