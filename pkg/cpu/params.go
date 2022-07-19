package cpu

import (
	"fmt"

	. "github.com/retroenv/nesgo/pkg/addressing"
)

type paramConverterFunc func(c *CPU, instruction *Instruction, params ...interface{}) string

var paramConverter = map[Mode]paramConverterFunc{
	ImpliedAddressing:     paramConverterImplied,
	ImmediateAddressing:   paramConverterImmediate,
	AccumulatorAddressing: paramConverterAccumulator,
	AbsoluteAddressing:    paramConverterAbsolute,
	AbsoluteXAddressing:   paramConverterAbsoluteX,
	AbsoluteYAddressing:   paramConverterAbsoluteY,
	ZeroPageAddressing:    paramConverterZeroPage,
	ZeroPageXAddressing:   paramConverterZeroPageX,
	ZeroPageYAddressing:   paramConverterZeroPageY,
	RelativeAddressing:    paramConverterRelative,
	IndirectAddressing:    paramConverterIndirect,
	IndirectXAddressing:   paramConverterIndirectX,
	IndirectYAddressing:   paramConverterIndirectY,
}

// ParamString returns the instruction parameters formatted as string.
func (c *CPU) ParamString(instruction *Instruction, params ...interface{}) string {
	fun, ok := paramConverter[c.TraceStep.Addressing]
	if !ok {
		err := fmt.Errorf("unsupported addressing mode %00x", c.TraceStep.Addressing)
		panic(err)
	}

	s := fun(c, instruction, params...)
	return s
}

func paramConverterImplied(_ *CPU, _ *Instruction, _ ...interface{}) string {
	return ""
}

func paramConverterImmediate(_ *CPU, _ *Instruction, params ...interface{}) string {
	imm := params[0]
	return fmt.Sprintf("#$%02X", imm)
}

func paramConverterAccumulator(_ *CPU, _ *Instruction, _ ...interface{}) string {
	return "A"
}

func paramConverterAbsolute(c *CPU, instruction *Instruction, params ...interface{}) string {
	address := params[0].(Absolute)
	if _, ok := BranchingInstructions[instruction.Name]; ok {
		return fmt.Sprintf("$%04X", address)
	}
	if !shouldOutputMemoryContent(uint16(address)) {
		return fmt.Sprintf("$%04X", address)
	}

	b := c.bus.Memory.ReadMemory(uint16(address))
	return fmt.Sprintf("$%04X = %02X", address, b)
}

func paramConverterAbsoluteX(c *CPU, _ *Instruction, params ...interface{}) string {
	address := params[0].(Absolute)
	offset := address + Absolute(c.X)
	b := c.bus.Memory.ReadMemory(uint16(offset))
	return fmt.Sprintf("$%04X,X @ %04X = %02X", address, offset, b)
}

func paramConverterAbsoluteY(c *CPU, _ *Instruction, params ...interface{}) string {
	address := params[0].(Absolute)
	offset := address + Absolute(c.Y)
	b := c.bus.Memory.ReadMemory(uint16(offset))
	return fmt.Sprintf("$%04X,Y @ %04X = %02X", address, offset, b)
}

func paramConverterZeroPage(c *CPU, _ *Instruction, params ...interface{}) string {
	address := params[0].(Absolute)
	b := c.bus.Memory.ReadMemory(uint16(address))
	return fmt.Sprintf("$%02X = %02X", address, b)
}

func paramConverterZeroPageX(c *CPU, _ *Instruction, params ...interface{}) string {
	address := params[0].(ZeroPage)
	offset := uint16(byte(address) + c.X)
	b := c.bus.Memory.ReadMemory(offset)
	return fmt.Sprintf("$%02X,X @ %02X = %02X", address, offset, b)
}

func paramConverterZeroPageY(c *CPU, _ *Instruction, params ...interface{}) string {
	address := params[0].(ZeroPage)
	offset := uint16(byte(address) + c.Y)
	b := c.bus.Memory.ReadMemory(offset)
	return fmt.Sprintf("$%02X,Y @ %02X = %02X", address, offset, b)
}

func paramConverterRelative(_ *CPU, _ *Instruction, params ...interface{}) string {
	address := params[0]
	return fmt.Sprintf("$%04X", address)
}

func paramConverterIndirect(c *CPU, _ *Instruction, params ...interface{}) string {
	address := params[0].(Indirect)
	value := c.bus.Memory.ReadMemory16Bug(uint16(address))
	return fmt.Sprintf("($%02X%02X) = %04X", c.TraceStep.Opcode[2], c.TraceStep.Opcode[1], value)
}

func paramConverterIndirectX(c *CPU, _ *Instruction, params ...interface{}) string {
	var address uint16
	indirectAddress, ok := params[0].(Indirect)
	if ok {
		address = uint16(indirectAddress)
	} else {
		address = uint16(params[0].(IndirectResolved))
	}

	b := c.bus.Memory.ReadMemory(address)
	offset := c.X + c.TraceStep.Opcode[1]
	return fmt.Sprintf("($%02X,X) @ %02X = %04X = %02X", c.TraceStep.Opcode[1], offset, address, b)
}

func paramConverterIndirectY(c *CPU, _ *Instruction, params ...interface{}) string {
	var address uint16
	indirectAddress, ok := params[0].(Indirect)
	if ok {
		address = uint16(indirectAddress)
	} else {
		address = uint16(params[0].(IndirectResolved))
	}

	b := c.bus.Memory.ReadMemory(address)
	offset := address - uint16(c.Y)
	return fmt.Sprintf("($%02X),Y = %04X @ %04X = %02X", c.TraceStep.Opcode[1], offset, address, b)
}
