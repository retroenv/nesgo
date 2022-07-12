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
// nolint: cyclop
func String(converter Converter, addressing Mode, params ...interface{}) (string, error) {
	switch addressing {
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
		if len(params) == 0 {
			return "", nil
		}
		return converter.Relative(params[0]), nil
	case IndirectAddressing:
		return converter.Indirect(params[0]), nil
	case IndirectXAddressing:
		return converter.IndirectX(params[0]), nil
	case IndirectYAddressing:
		return converter.IndirectY(params[0]), nil
	default:
		return "", fmt.Errorf("unsupported addressing mode %d", addressing)
	}
}
