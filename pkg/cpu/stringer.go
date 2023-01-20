package cpu

import (
	"fmt"

	. "github.com/retroenv/retrogolib/nes/addressing"
)

// ParamConverter is an interface for the conversion of the instruction parameters to
// specific assembler implementation outputs.
type ParamConverter interface {
	Absolute(param any) string
	AbsoluteX(param any) string
	AbsoluteY(param any) string
	Accumulator() string
	Immediate(param any) string
	Indirect(param any) string
	IndirectX(param any) string
	IndirectY(param any) string
	Relative(param any) string
	ZeroPage(param any) string
	ZeroPageX(param any) string
	ZeroPageY(param any) string
}

// ParamString returns the parameters as a string that is compatible to the
// assembler presented by the converter.
// nolint:cyclop
func ParamString(converter ParamConverter, addressing Mode, param any) (string, error) {
	switch addressing {
	case ImpliedAddressing:
		return "", nil
	case ImmediateAddressing:
		return converter.Immediate(param), nil
	case AccumulatorAddressing:
		return converter.Accumulator(), nil
	case AbsoluteAddressing:
		return converter.Absolute(param), nil
	case AbsoluteXAddressing:
		return converter.AbsoluteX(param), nil
	case AbsoluteYAddressing:
		return converter.AbsoluteY(param), nil
	case ZeroPageAddressing:
		return converter.ZeroPage(param), nil
	case ZeroPageXAddressing:
		return converter.ZeroPageX(param), nil
	case ZeroPageYAddressing:
		return converter.ZeroPageY(param), nil
	case RelativeAddressing:
		return converter.Relative(param), nil
	case IndirectAddressing:
		return converter.Indirect(param), nil
	case IndirectXAddressing:
		return converter.IndirectX(param), nil
	case IndirectYAddressing:
		return converter.IndirectY(param), nil
	default:
		return "", fmt.Errorf("unsupported addressing mode %d", addressing)
	}
}
