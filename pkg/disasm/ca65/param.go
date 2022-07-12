// Package ca65 provides helpers to create ca65 assembler compatible asm output.
package ca65

import (
	"fmt"

	. "github.com/retroenv/nesgo/pkg/addressing"
)

const Name = "ca65"

// ParamConverter converts the opcode parameters to ca65
// compatible output.
type ParamConverter struct {
}

// Immediate converts the parameters to the assembler implementation compatible string.
func (c ParamConverter) Immediate(param interface{}) string {
	return fmt.Sprintf("#$%02X", param)
}

// Accumulator converts the parameters to the assembler implementation compatible string.
func (c ParamConverter) Accumulator() string {
	return "a"
}

// Absolute converts the parameters to the assembler implementation compatible string.
func (c ParamConverter) Absolute(param interface{}) string {
	switch val := param.(type) {
	case int, Absolute:
		return fmt.Sprintf("$%04X", val)
	case string:
		return val
	default:
		panic(fmt.Sprintf("unsupported param type %T", val))
	}
}

// AbsoluteX converts the parameters to the assembler implementation compatible string.
func (c ParamConverter) AbsoluteX(param interface{}) string {
	switch val := param.(type) {
	case int, Absolute:
		return fmt.Sprintf("$%04X,X", val)
	case string:
		return fmt.Sprintf("%s,X", val)
	default:
		panic(fmt.Sprintf("unsupported param type %T", val))
	}
}

// AbsoluteY converts the parameters to the assembler implementation compatible string.
func (c ParamConverter) AbsoluteY(param interface{}) string {
	switch val := param.(type) {
	case int, Absolute:
		return fmt.Sprintf("$%04X,Y", val)
	case string:
		return fmt.Sprintf("%s,Y", val)
	default:
		panic(fmt.Sprintf("unsupported param type %T", val))
	}
}

// ZeroPage converts the parameters to the assembler implementation compatible string.
func (c ParamConverter) ZeroPage(param interface{}) string {
	switch val := param.(type) {
	case int, Absolute, ZeroPage:
		return fmt.Sprintf("$%02X", val)
	case string:
		return val
	default:
		panic(fmt.Sprintf("unsupported param type %T", val))
	}
}

// ZeroPageX converts the parameters to the assembler implementation compatible string.
func (c ParamConverter) ZeroPageX(param interface{}) string {
	switch val := param.(type) {
	case int, Absolute, ZeroPage:
		return fmt.Sprintf("$%02X,X", val)
	case string:
		return fmt.Sprintf("%s,X", val)
	default:
		panic(fmt.Sprintf("unsupported param type %T", val))
	}
}

// ZeroPageY converts the parameters to the assembler implementation compatible string.
func (c ParamConverter) ZeroPageY(param interface{}) string {
	switch val := param.(type) {
	case int, Absolute, ZeroPage:
		return fmt.Sprintf("$%02X,Y", val)
	case string:
		return fmt.Sprintf("%s,Y", val)
	default:
		panic(fmt.Sprintf("unsupported param type %T", val))
	}
}

// Relative converts the parameters to the assembler implementation compatible string.
func (c ParamConverter) Relative(param interface{}) string {
	return fmt.Sprintf("$%04X", param)
}

// Indirect converts the parameters to the assembler implementation compatible string.
func (c ParamConverter) Indirect(param interface{}) string {
	address, ok := param.(Indirect)
	if ok {
		return fmt.Sprintf("($%04X)", address)
	}
	alias := param.(string)
	return fmt.Sprintf("(%s)", alias)
}

// IndirectX converts the parameters to the assembler implementation compatible string.
func (c ParamConverter) IndirectX(param interface{}) string {
	address, ok := param.(Indirect)
	if ok {
		return fmt.Sprintf("($%04X,X)", address)
	}
	alias := param.(string)
	return fmt.Sprintf("(%s,X)", alias)
}

// IndirectY converts the parameters to the assembler implementation compatible string.
func (c ParamConverter) IndirectY(param interface{}) string {
	address, ok := param.(Indirect)
	if ok {
		return fmt.Sprintf("($%04X),Y", address)
	}
	alias := param.(string)
	return fmt.Sprintf("(%s),Y", alias)
}
