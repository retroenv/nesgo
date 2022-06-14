package disasm

import (
	"fmt"

	. "github.com/retroenv/nesgo/pkg/addressing"
	"github.com/retroenv/nesgo/pkg/cpu"
)

// paramConverterCa65 converts the opcode parameters to ca65
// compatible output.
type paramConverterCa65 struct {
}

func (c paramConverterCa65) Immediate(opcode cpu.Opcode, params ...interface{}) string {
	imm := params[0]
	return fmt.Sprintf("#$%02X", imm)
}

func (c paramConverterCa65) Accumulator(opcode cpu.Opcode, params ...interface{}) string {
	return "a"
}

func (c paramConverterCa65) Absolute(opcode cpu.Opcode, params ...interface{}) string {
	address := params[0].(Absolute)

	if _, ok := cpu.BranchingInstructions[opcode.Instruction.Name]; ok {
		return fmt.Sprintf("_label_%04x", address)
	}

	return fmt.Sprintf("$%04X", address)
}

func (c paramConverterCa65) AbsoluteX(opcode cpu.Opcode, params ...interface{}) string {
	address := params[0].(Absolute)
	return fmt.Sprintf("$%04X,X", address)
}

func (c paramConverterCa65) AbsoluteY(opcode cpu.Opcode, params ...interface{}) string {
	address := params[0].(Absolute)
	return fmt.Sprintf("$%04X,Y", address)
}

func (c paramConverterCa65) ZeroPage(opcode cpu.Opcode, params ...interface{}) string {
	address := params[0].(Absolute)
	return fmt.Sprintf("$%02X", address)
}

func (c paramConverterCa65) ZeroPageX(opcode cpu.Opcode, params ...interface{}) string {
	address := params[0].(ZeroPage)
	return fmt.Sprintf("$%02X,X", address)
}

func (c paramConverterCa65) ZeroPageY(opcode cpu.Opcode, params ...interface{}) string {
	address := params[0].(ZeroPage)
	return fmt.Sprintf("$%02X,Y", address)
}

func (c paramConverterCa65) Relative(opcode cpu.Opcode, params ...interface{}) string {
	address := params[0]

	if _, ok := cpu.BranchingInstructions[opcode.Instruction.Name]; ok {
		return fmt.Sprintf("_label_%04x", address)
	}

	return fmt.Sprintf("$%04X", address)
}

func (c paramConverterCa65) Indirect(opcode cpu.Opcode, params ...interface{}) string {
	address := params[0].(Indirect)
	return fmt.Sprintf("($%04X)", address)
}

func (c paramConverterCa65) IndirectX(opcode cpu.Opcode, params ...interface{}) string {
	address := params[0].(Absolute)
	return fmt.Sprintf("($%04X,X)", address)
}

func (c paramConverterCa65) IndirectY(opcode cpu.Opcode, params ...interface{}) string {
	address := params[0].(Absolute)
	return fmt.Sprintf("($%04X),Y", address)
}
