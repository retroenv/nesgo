package disasm

import (
	"fmt"

	. "github.com/retroenv/nesgo/pkg/addressing"
	"github.com/retroenv/nesgo/pkg/cpu"
)

// paramConverter is an interface for the conversion of opcode parameters to
// specific assembler implementations.
type paramConverter interface {
	Absolute(param interface{}) string
	AbsoluteX(param interface{}) string
	AbsoluteY(param interface{}) string
	Accumulator() string
	Immediate(param interface{}) string
	Indirect(param interface{}) string
	IndirectX(param interface{}) string
	IndirectY(param interface{}) string
	Relative(param interface{}) string
	ZeroPage(param interface{}) string
	ZeroPageX(param interface{}) string
	ZeroPageY(param interface{}) string
}

// paramString returns the parameters as a string that is compatible to the
// assembler presented by the converter.
func paramString(converter paramConverter, opcode cpu.Opcode, params ...interface{}) (string, error) {
	switch opcode.Addressing {
	case ImpliedAddressing:
		return "", nil
	case ImmediateAddressing:
		return converter.Immediate(params[0]), nil
	case AccumulatorAddressing:
		return converter.Accumulator(), nil
	case AbsoluteAddressing:
		return converter.Absolute(params[0]), nil
	case AbsoluteXAddressing:
		return converter.AbsoluteX(params[0]), nil
	case AbsoluteYAddressing:
		return converter.AbsoluteY(params[0]), nil
	case ZeroPageAddressing:
		return converter.ZeroPage(params[0]), nil
	case ZeroPageXAddressing:
		return converter.ZeroPageX(params[0]), nil
	case ZeroPageYAddressing:
		return converter.ZeroPageY(params[0]), nil
	case RelativeAddressing:
		return converter.Relative(params[0]), nil
	case IndirectAddressing:
		return converter.Indirect(params[0]), nil
	case IndirectXAddressing:
		return converter.IndirectX(params[0]), nil
	case IndirectYAddressing:
		return converter.IndirectY(params[0]), nil
	default:
		return "", fmt.Errorf("unsupported addressing mode %d", opcode.Addressing)
	}
}
