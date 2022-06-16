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
func (c ParamConverter) Immediate(params ...interface{}) string {
	imm := params[0]
	return fmt.Sprintf("#$%02X", imm)
}

// Accumulator converts the parameters to the assembler implementation compatible string.
func (c ParamConverter) Accumulator(params ...interface{}) string {
	return "a"
}

// Absolute converts the parameters to the assembler implementation compatible string.
func (c ParamConverter) Absolute(params ...interface{}) string {
	address := params[0].(Absolute)
	return fmt.Sprintf("$%04X", address)
}

// AbsoluteX converts the parameters to the assembler implementation compatible string.
func (c ParamConverter) AbsoluteX(params ...interface{}) string {
	address := params[0].(Absolute)
	return fmt.Sprintf("$%04X,X", address)
}

// AbsoluteY converts the parameters to the assembler implementation compatible string.
func (c ParamConverter) AbsoluteY(params ...interface{}) string {
	address := params[0].(Absolute)
	return fmt.Sprintf("$%04X,Y", address)
}

// ZeroPage converts the parameters to the assembler implementation compatible string.
func (c ParamConverter) ZeroPage(params ...interface{}) string {
	address := params[0].(Absolute)
	return fmt.Sprintf("$%02X", address)
}

// ZeroPageX converts the parameters to the assembler implementation compatible string.
func (c ParamConverter) ZeroPageX(params ...interface{}) string {
	address := params[0].(ZeroPage)
	return fmt.Sprintf("$%02X,X", address)
}

// ZeroPageY converts the parameters to the assembler implementation compatible string.
func (c ParamConverter) ZeroPageY(params ...interface{}) string {
	address := params[0].(ZeroPage)
	return fmt.Sprintf("$%02X,Y", address)
}

// Relative converts the parameters to the assembler implementation compatible string.
func (c ParamConverter) Relative(params ...interface{}) string {
	address := params[0]
	return fmt.Sprintf("$%04X", address)
}

// Indirect converts the parameters to the assembler implementation compatible string.
func (c ParamConverter) Indirect(params ...interface{}) string {
	address := params[0].(Indirect)
	return fmt.Sprintf("($%04X)", address)
}

// IndirectX converts the parameters to the assembler implementation compatible string.
func (c ParamConverter) IndirectX(params ...interface{}) string {
	address := params[0].(Absolute)
	return fmt.Sprintf("($%04X,X)", address)
}

// IndirectY converts the parameters to the assembler implementation compatible string.
func (c ParamConverter) IndirectY(params ...interface{}) string {
	address := params[0].(Absolute)
	return fmt.Sprintf("($%04X),Y", address)
}
