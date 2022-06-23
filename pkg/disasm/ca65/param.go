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
	address := param.(Absolute)
	return fmt.Sprintf("$%04X", address)
}

// AbsoluteX converts the parameters to the assembler implementation compatible string.
func (c ParamConverter) AbsoluteX(param interface{}) string {
	address := param.(Absolute)
	return fmt.Sprintf("$%04X,X", address)
}

// AbsoluteY converts the parameters to the assembler implementation compatible string.
func (c ParamConverter) AbsoluteY(param interface{}) string {
	address := param.(Absolute)
	return fmt.Sprintf("$%04X,Y", address)
}

// ZeroPage converts the parameters to the assembler implementation compatible string.
func (c ParamConverter) ZeroPage(param interface{}) string {
	address := param.(Absolute)
	return fmt.Sprintf("$%02X", address)
}

// ZeroPageX converts the parameters to the assembler implementation compatible string.
func (c ParamConverter) ZeroPageX(param interface{}) string {
	address := param.(ZeroPage)
	return fmt.Sprintf("$%02X,X", address)
}

// ZeroPageY converts the parameters to the assembler implementation compatible string.
func (c ParamConverter) ZeroPageY(param interface{}) string {
	address := param.(ZeroPage)
	return fmt.Sprintf("$%02X,Y", address)
}

// Relative converts the parameters to the assembler implementation compatible string.
func (c ParamConverter) Relative(param interface{}) string {
	return fmt.Sprintf("$%04X", param)
}

// Indirect converts the parameters to the assembler implementation compatible string.
func (c ParamConverter) Indirect(param interface{}) string {
	address := param.(Indirect)
	return fmt.Sprintf("($%04X)", address)
}

// IndirectX converts the parameters to the assembler implementation compatible string.
func (c ParamConverter) IndirectX(param interface{}) string {
	address := param.(Indirect)
	return fmt.Sprintf("($%04X,X)", address)
}

// IndirectY converts the parameters to the assembler implementation compatible string.
func (c ParamConverter) IndirectY(param interface{}) string {
	address := param.(Indirect)
	return fmt.Sprintf("($%04X),Y", address)
}
