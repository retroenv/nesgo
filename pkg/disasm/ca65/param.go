// Package ca65 provides helpers to create ca65 assembler compatible asm output.
package ca65

import (
	"fmt"

	. "github.com/retroenv/nesgo/pkg/addressing"
	"github.com/retroenv/nesgo/pkg/cpu"
)

const Name = "ca65"

// ParamConverter converts the opcode parameters to ca65
// compatible output.
type ParamConverter struct {
}

// Immediate converts the parameters to the assembler implementation compatible string.
func (c ParamConverter) Immediate(opcode cpu.Opcode, params ...interface{}) string {
	imm := params[0]
	return fmt.Sprintf("#$%02X", imm)
}

// Accumulator converts the parameters to the assembler implementation compatible string.
func (c ParamConverter) Accumulator(opcode cpu.Opcode, params ...interface{}) string {
	return "a"
}

// Absolute converts the parameters to the assembler implementation compatible string.
func (c ParamConverter) Absolute(opcode cpu.Opcode, params ...interface{}) string {
	address := params[0].(Absolute)
	return fmt.Sprintf("$%04X", address)
}

// AbsoluteX converts the parameters to the assembler implementation compatible string.
func (c ParamConverter) AbsoluteX(opcode cpu.Opcode, params ...interface{}) string {
	address := params[0].(Absolute)
	return fmt.Sprintf("$%04X,X", address)
}

// AbsoluteY converts the parameters to the assembler implementation compatible string.
func (c ParamConverter) AbsoluteY(opcode cpu.Opcode, params ...interface{}) string {
	address := params[0].(Absolute)
	return fmt.Sprintf("$%04X,Y", address)
}

// ZeroPage converts the parameters to the assembler implementation compatible string.
func (c ParamConverter) ZeroPage(opcode cpu.Opcode, params ...interface{}) string {
	address := params[0].(Absolute)
	return fmt.Sprintf("$%02X", address)
}

// ZeroPageX converts the parameters to the assembler implementation compatible string.
func (c ParamConverter) ZeroPageX(opcode cpu.Opcode, params ...interface{}) string {
	address := params[0].(ZeroPage)
	return fmt.Sprintf("$%02X,X", address)
}

// ZeroPageY converts the parameters to the assembler implementation compatible string.
func (c ParamConverter) ZeroPageY(opcode cpu.Opcode, params ...interface{}) string {
	address := params[0].(ZeroPage)
	return fmt.Sprintf("$%02X,Y", address)
}

// Relative converts the parameters to the assembler implementation compatible string.
func (c ParamConverter) Relative(opcode cpu.Opcode, params ...interface{}) string {
	address := params[0]
	return fmt.Sprintf("$%04X", address)
}

// Indirect converts the parameters to the assembler implementation compatible string.
func (c ParamConverter) Indirect(opcode cpu.Opcode, params ...interface{}) string {
	address := params[0].(Indirect)
	return fmt.Sprintf("($%04X)", address)
}

// IndirectX converts the parameters to the assembler implementation compatible string.
func (c ParamConverter) IndirectX(opcode cpu.Opcode, params ...interface{}) string {
	address := params[0].(Absolute)
	return fmt.Sprintf("($%04X,X)", address)
}

// IndirectY converts the parameters to the assembler implementation compatible string.
func (c ParamConverter) IndirectY(opcode cpu.Opcode, params ...interface{}) string {
	address := params[0].(Absolute)
	return fmt.Sprintf("($%04X),Y", address)
}
