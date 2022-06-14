package disasm

import (
	"fmt"

	. "github.com/retroenv/nesgo/pkg/addressing"
	"github.com/retroenv/nesgo/pkg/cpu"
)

// ParamConverter is an interface for the conversion of opcode parameters to
// specific assembler implementations.
type ParamConverter interface {
	Absolute(opcode cpu.Opcode, params ...interface{}) string
	AbsoluteX(opcode cpu.Opcode, params ...interface{}) string
	AbsoluteY(opcode cpu.Opcode, params ...interface{}) string
	Accumulator(opcode cpu.Opcode, params ...interface{}) string
	Immediate(opcode cpu.Opcode, params ...interface{}) string
	Indirect(opcode cpu.Opcode, params ...interface{}) string
	IndirectX(opcode cpu.Opcode, params ...interface{}) string
	IndirectY(opcode cpu.Opcode, params ...interface{}) string
	Relative(opcode cpu.Opcode, params ...interface{}) string
	ZeroPage(opcode cpu.Opcode, params ...interface{}) string
	ZeroPageX(opcode cpu.Opcode, params ...interface{}) string
	ZeroPageY(opcode cpu.Opcode, params ...interface{}) string
}

// ParamStrings returns the parameters as a string that is compatible to the
// assembler presented by the converter.
func ParamStrings(converter ParamConverter, opcode cpu.Opcode, params ...interface{}) (string, error) {
	switch opcode.Addressing {
	case ImpliedAddressing:
		return "", nil
	case ImmediateAddressing:
		return converter.Immediate(opcode, params...), nil
	case AccumulatorAddressing:
		return converter.Accumulator(opcode, params...), nil
	case AbsoluteAddressing:
		return converter.Absolute(opcode, params...), nil
	case AbsoluteXAddressing:
		return converter.AbsoluteX(opcode, params...), nil
	case AbsoluteYAddressing:
		return converter.AbsoluteY(opcode, params...), nil
	case ZeroPageAddressing:
		return converter.ZeroPage(opcode, params...), nil
	case ZeroPageXAddressing:
		return converter.ZeroPageX(opcode, params...), nil
	case ZeroPageYAddressing:
		return converter.ZeroPageY(opcode, params...), nil
	case RelativeAddressing:
		return converter.Relative(opcode, params...), nil
	case IndirectAddressing:
		return converter.Indirect(opcode, params...), nil
	case IndirectXAddressing:
		return converter.IndirectX(opcode, params...), nil
	case IndirectYAddressing:
		return converter.IndirectY(opcode, params...), nil
	default:
		return "", fmt.Errorf("unsupported addressing mode %d", opcode.Addressing)
	}
}
