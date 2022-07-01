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
	address, ok := param.(Absolute)
	if ok {
		return fmt.Sprintf("$%04X", address)
	}
	alias := param.(string)
	return alias
}

// AbsoluteX converts the parameters to the assembler implementation compatible string.
func (c ParamConverter) AbsoluteX(param interface{}) string {
	address, ok := param.(Absolute)
	if ok {
		return fmt.Sprintf("$%04X,X", address)
	}
	alias := param.(string)
	return fmt.Sprintf("%s,X", alias)
}

// AbsoluteY converts the parameters to the assembler implementation compatible string.
func (c ParamConverter) AbsoluteY(param interface{}) string {
	address, ok := param.(Absolute)
	if ok {
		return fmt.Sprintf("$%04X,Y", address)
	}
	alias := param.(string)
	return fmt.Sprintf("%s,Y", alias)
}

// ZeroPage converts the parameters to the assembler implementation compatible string.
func (c ParamConverter) ZeroPage(param interface{}) string {
	address, ok := param.(Absolute)
	if ok {
		return fmt.Sprintf("$%02X", address)
	}
	alias := param.(string)
	return alias
}

// ZeroPageX converts the parameters to the assembler implementation compatible string.
func (c ParamConverter) ZeroPageX(param interface{}) string {
	address, ok := param.(ZeroPage)
	if ok {
		return fmt.Sprintf("$%02X,X", address)
	}
	alias := param.(string)
	return fmt.Sprintf("%s,X", alias)
}

// ZeroPageY converts the parameters to the assembler implementation compatible string.
func (c ParamConverter) ZeroPageY(param interface{}) string {
	address, ok := param.(ZeroPage)
	if ok {
		return fmt.Sprintf("$%02X,Y", address)
	}
	alias := param.(string)
	return fmt.Sprintf("%s,Y", alias)
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
