package cpu

import (
	"fmt"
	"math"
	"strings"

	. "github.com/retroenv/nesgo/pkg/addressing"
)

// TracingMode defines a tracing mode.
type TracingMode int

// tracing modes, either disabled, in Go mode or emulator mode.
const (
	NoTracing TracingMode = iota
	GoTracing
	EmulatorTracing
)

// TraceStep contains all info needed to print a trace step.
type TraceStep struct {
	PC          uint16
	Opcode      []byte
	Addressing  Mode
	Instruction string
}

func (t TraceStep) print(cpu *CPU) {
	var opcodes [3]string
	for i := 0; i < 3; i++ {
		s := "  "
		if i < len(t.Opcode) {
			op := t.Opcode[i]
			s = fmt.Sprintf("%02X", op)
		}

		opcodes[i] = s
	}

	fmt.Printf("%04X  %s %s %s  %-31s A:%02X X:%02X Y:%02X P:%02X SP:%02X\n",
		t.PC, opcodes[0], opcodes[1], opcodes[2], t.Instruction,
		cpu.A, cpu.X, cpu.Y, cpu.GetFlags(), cpu.SP)
}

func (c *CPU) trace(instruction *Instruction, params ...interface{}) {
	var paramsAsString string

	if c.tracing == GoTracing {
		c.TraceStep.Addressing, paramsAsString = addressModeFromCall(instruction, params...)
		if !instruction.HasAddressing(c.TraceStep.Addressing) {
			panic(fmt.Sprintf("unexpected addressing mode type %T", c.TraceStep.Addressing))
		}

		c.TraceStep.Opcode = []byte{instruction.Addressing[c.TraceStep.Addressing].Opcode}
		// TODO add parameter opcodes
	} else {
		paramsAsString = c.paramString(instruction, params...)
	}

	c.TraceStep.Instruction = strings.ToUpper(instruction.Name)
	if paramsAsString != "" {
		c.TraceStep.Instruction += " " + paramsAsString
	}
	c.TraceStep.print(c)
}

// addressModeFromCall gets the addressing mode from the passed params
// TODO format in other func
func addressModeFromCall(instruction *Instruction, params ...interface{}) (Mode, string) {
	if len(params) == 0 {
		return addressModeFromCallNoParam(instruction)
	}

	param := params[0]
	var register interface{}
	if len(params) > 1 {
		register = params[1]
	}

	switch address := param.(type) {
	case int:
		if instruction.HasAddressing(ImmediateAddressing) && register == nil && address <= math.MaxUint8 {
			return ImmediateAddressing, fmt.Sprintf("#$%02X", address)
		}
		if register == nil {
			return AbsoluteAddressing, fmt.Sprintf("$%04X", address)
		}
		panic("X/Y support not implemented") // TODO

	case uint8:
		return ImmediateAddressing, fmt.Sprintf("#$%02X", address)

	case *uint8: // variable
		return ImmediateAddressing, fmt.Sprintf("#$%02X", *address)

	case Absolute:
		// branches in emulation mode
		if instruction.HasAddressing(RelativeAddressing) {
			return RelativeAddressing, fmt.Sprintf("$%04X", address)
		}

		return AbsoluteAddressing, fmt.Sprintf("$%04X", address)

	case Indirect:
		if register == nil {
			return IndirectAddressing, fmt.Sprintf("$%04X", address)
		}
		panic("X/Y support not implemented") // TODO

	default:
		panic(fmt.Sprintf("unsupported addressing mode type %T", param))
	}
}

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

// paramString returns the params formatted as string.
func (c *CPU) paramString(instruction *Instruction, params ...interface{}) string {
	fun, ok := paramConverter[c.TraceStep.Addressing]
	if !ok {
		err := fmt.Errorf("unsupported addressing mode %00x", c.TraceStep.Addressing)
		panic(err)
	}

	s := fun(c, instruction, params...)
	return s
}

func paramConverterImplied(c *CPU, instruction *Instruction, params ...interface{}) string {
	return ""
}

func paramConverterImmediate(c *CPU, instruction *Instruction, params ...interface{}) string {
	imm := params[0]
	return fmt.Sprintf("#$%02X", imm)
}

func paramConverterAccumulator(c *CPU, instruction *Instruction, params ...interface{}) string {
	return "A"
}

func paramConverterAbsolute(c *CPU, instruction *Instruction, params ...interface{}) string {
	address := params[0].(Absolute)
	if _, ok := BranchingInstructions[instruction.Name]; ok {
		return fmt.Sprintf("$%04X", address)
	}
	if !outputMemoryContent(uint16(address)) {
		return fmt.Sprintf("$%04X", address)
	}

	b := c.memory.ReadMemory(uint16(address))
	return fmt.Sprintf("$%04X = %02X", address, b)
}

func paramConverterAbsoluteX(c *CPU, instruction *Instruction, params ...interface{}) string {
	address := params[0].(Absolute)
	offset := address + Absolute(c.X)
	b := c.memory.ReadMemory(uint16(offset))
	return fmt.Sprintf("$%04X,X @ %04X = %02X", address, offset, b)
}

func paramConverterAbsoluteY(c *CPU, instruction *Instruction, params ...interface{}) string {
	address := params[0].(Absolute)
	offset := address + Absolute(c.Y)
	b := c.memory.ReadMemory(uint16(offset))
	return fmt.Sprintf("$%04X,Y @ %04X = %02X", address, offset, b)
}

func paramConverterZeroPage(c *CPU, instruction *Instruction, params ...interface{}) string {
	address := params[0].(Absolute)
	b := c.memory.ReadMemory(uint16(address))
	return fmt.Sprintf("$%02X = %02X", address, b)
}

func paramConverterZeroPageX(c *CPU, instruction *Instruction, params ...interface{}) string {
	address := params[0].(ZeroPage)
	offset := uint16(byte(address) + c.X)
	b := c.memory.ReadMemory(offset)
	return fmt.Sprintf("$%02X,X @ %02X = %02X", address, offset, b)
}

func paramConverterZeroPageY(c *CPU, instruction *Instruction, params ...interface{}) string {
	address := params[0].(ZeroPage)
	offset := uint16(byte(address) + c.Y)
	b := c.memory.ReadMemory(offset)
	return fmt.Sprintf("$%02X,Y @ %02X = %02X", address, offset, b)
}

func paramConverterRelative(c *CPU, instruction *Instruction, params ...interface{}) string {
	address := params[0]
	return fmt.Sprintf("$%04X", address)
}

func paramConverterIndirect(c *CPU, instruction *Instruction, params ...interface{}) string {
	address := params[0].(Indirect)
	value := c.memory.ReadMemory16Bug(uint16(address))
	return fmt.Sprintf("($%02X%02X) = %04X", c.TraceStep.Opcode[2], c.TraceStep.Opcode[1], value)
}

func paramConverterIndirectX(c *CPU, instruction *Instruction, params ...interface{}) string {
	offset := c.X + c.TraceStep.Opcode[1]
	address := params[0].(Absolute)
	b := c.memory.ReadMemory(uint16(address))
	return fmt.Sprintf("($%02X,X) @ %02X = %04X = %02X", c.TraceStep.Opcode[1], offset, address, b)
}

func paramConverterIndirectY(c *CPU, instruction *Instruction, params ...interface{}) string {
	address := params[0].(Absolute)
	offset := address - Absolute(c.Y)
	b := c.memory.ReadMemory(uint16(address))
	return fmt.Sprintf("($%02X),Y = %04X @ %04X = %02X", c.TraceStep.Opcode[1], offset, address, b)
}

func outputMemoryContent(address uint16) bool {
	switch {
	case address < 0x2000:
		return true
	case address >= 0x8000:
		return true
	default:
		return false
	}
}

func addressModeFromCallNoParam(instruction *Instruction) (Mode, string) {
	if instruction.HasAddressing(AccumulatorAddressing) {
		return AccumulatorAddressing, ""
	}
	// branches have no target in go mode
	if instruction.HasAddressing(RelativeAddressing) {
		return RelativeAddressing, ""
	}
	return ImpliedAddressing, ""
}
