package cpu

import (
	"fmt"

	"github.com/retroenv/retrogolib/nes/addressing"
)

// ca65ParamConverter converts the opcode parameters to ca65 compatible output.
type ca65ParamConverter struct {
}

// Immediate converts the parameters to the assembler implementation compatible string.
func (c ca65ParamConverter) Immediate(param any) string {
	return fmt.Sprintf("#$%02X", param)
}

// Accumulator converts the parameters to the assembler implementation compatible string.
func (c ca65ParamConverter) Accumulator() string {
	return "a"
}

// Absolute converts the parameters to the assembler implementation compatible string.
func (c ca65ParamConverter) Absolute(param any) string {
	switch val := param.(type) {
	case int, addressing.Absolute:
		return fmt.Sprintf("$%04X", val)
	case string:
		return val
	default:
		panic(fmt.Sprintf("unsupported param type %T", val))
	}
}

// AbsoluteX converts the parameters to the assembler implementation compatible string.
func (c ca65ParamConverter) AbsoluteX(param any) string {
	switch val := param.(type) {
	case int, addressing.Absolute:
		return fmt.Sprintf("$%04X,X", val)
	case string:
		return fmt.Sprintf("%s,X", val)
	default:
		panic(fmt.Sprintf("unsupported param type %T", val))
	}
}

// AbsoluteY converts the parameters to the assembler implementation compatible string.
func (c ca65ParamConverter) AbsoluteY(param any) string {
	switch val := param.(type) {
	case int, addressing.Absolute:
		return fmt.Sprintf("$%04X,Y", val)
	case string:
		return fmt.Sprintf("%s,Y", val)
	default:
		panic(fmt.Sprintf("unsupported param type %T", val))
	}
}

// ZeroPage converts the parameters to the assembler implementation compatible string.
func (c ca65ParamConverter) ZeroPage(param any) string {
	switch val := param.(type) {
	case int, addressing.Absolute, addressing.ZeroPage:
		return fmt.Sprintf("$%02X", val)
	case string:
		return val
	default:
		panic(fmt.Sprintf("unsupported param type %T", val))
	}
}

// ZeroPageX converts the parameters to the assembler implementation compatible string.
func (c ca65ParamConverter) ZeroPageX(param any) string {
	switch val := param.(type) {
	case int, addressing.Absolute, addressing.ZeroPage:
		return fmt.Sprintf("$%02X,X", val)
	case string:
		return fmt.Sprintf("%s,X", val)
	default:
		panic(fmt.Sprintf("unsupported param type %T", val))
	}
}

// ZeroPageY converts the parameters to the assembler implementation compatible string.
func (c ca65ParamConverter) ZeroPageY(param any) string {
	switch val := param.(type) {
	case int, addressing.Absolute, addressing.ZeroPage:
		return fmt.Sprintf("$%02X,Y", val)
	case string:
		return fmt.Sprintf("%s,Y", val)
	default:
		panic(fmt.Sprintf("unsupported param type %T", val))
	}
}

// Relative converts the parameters to the assembler implementation compatible string.
func (c ca65ParamConverter) Relative(param any) string {
	if param == nil {
		return ""
	}
	return fmt.Sprintf("$%04X", param)
}

// Indirect converts the parameters to the assembler implementation compatible string.
func (c ca65ParamConverter) Indirect(param any) string {
	address, ok := param.(addressing.Indirect)
	if ok {
		return fmt.Sprintf("($%04X)", address)
	}
	alias := param.(string)
	return fmt.Sprintf("(%s)", alias)
}

// IndirectX converts the parameters to the assembler implementation compatible string.
func (c ca65ParamConverter) IndirectX(param any) string {
	address, ok := param.(addressing.Indirect)
	if ok {
		return fmt.Sprintf("($%04X,X)", address)
	}
	alias := param.(string)
	return fmt.Sprintf("(%s,X)", alias)
}

// IndirectY converts the parameters to the assembler implementation compatible string.
func (c ca65ParamConverter) IndirectY(param any) string {
	address, ok := param.(addressing.Indirect)
	if ok {
		return fmt.Sprintf("($%04X),Y", address)
	}
	alias := param.(string)
	return fmt.Sprintf("(%s),Y", alias)
}
