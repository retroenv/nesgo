package disasm

import (
	"fmt"

	. "github.com/retroenv/nesgo/pkg/addressing"
	"github.com/retroenv/nesgo/pkg/cpu"
)

// paramConverter is an interface for the conversion of opcode parameters to
// specific assembler implementations.
type paramConverter interface {
	Absolute(params ...interface{}) string
	AbsoluteX(params ...interface{}) string
	AbsoluteY(params ...interface{}) string
	Accumulator(params ...interface{}) string
	Immediate(params ...interface{}) string
	Indirect(params ...interface{}) string
	IndirectX(params ...interface{}) string
	IndirectY(params ...interface{}) string
	Relative(params ...interface{}) string
	ZeroPage(params ...interface{}) string
	ZeroPageX(params ...interface{}) string
	ZeroPageY(params ...interface{}) string
}

// paramStrings returns the parameters as a string that is compatible to the
// assembler presented by the converter.
func paramStrings(converter paramConverter, opcode cpu.Opcode, params ...interface{}) (string, error) {
	switch opcode.Addressing {
	case ImpliedAddressing:
		return "", nil
	case ImmediateAddressing:
		return converter.Immediate(params...), nil
	case AccumulatorAddressing:
		return converter.Accumulator(params...), nil
	case AbsoluteAddressing:
		return converter.Absolute(params...), nil
	case AbsoluteXAddressing:
		return converter.AbsoluteX(params...), nil
	case AbsoluteYAddressing:
		return converter.AbsoluteY(params...), nil
	case ZeroPageAddressing:
		return converter.ZeroPage(params...), nil
	case ZeroPageXAddressing:
		return converter.ZeroPageX(params...), nil
	case ZeroPageYAddressing:
		return converter.ZeroPageY(params...), nil
	case RelativeAddressing:
		return converter.Relative(params...), nil
	case IndirectAddressing:
		return converter.Indirect(params...), nil
	case IndirectXAddressing:
		return converter.IndirectX(params...), nil
	case IndirectYAddressing:
		return converter.IndirectY(params...), nil
	default:
		return "", fmt.Errorf("unsupported addressing mode %d", opcode.Addressing)
	}
}
