// Package param contains the parameter to string conversion interface for disassembling.
package param

import (
	"fmt"

	. "github.com/retroenv/nesgo/pkg/addressing"
)

// Converter is an interface for the conversion of the instruction parameters to
// specific assembler implementation outputs.
type Converter interface {
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

// String returns the parameters as a string that is compatible to the
// assembler presented by the converter.
func String(converter Converter, addressing Mode, param interface{}) (string, error) {
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
