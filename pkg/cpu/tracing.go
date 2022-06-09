package cpu

import (
	"fmt"
	"strings"
)

// TracingMode defines a tracing mode.
type TracingMode int

// tracing modes, either disabled, in Go mode or emulator mode.
const (
	NoTracing TracingMode = 1 << iota
	GoTracing
	EmulatorTracing
)

// TraceStep contains all info needed to print a trace step.
type TraceStep struct {
	PC     uint16
	Opcode []byte
}

func (t TraceStep) print(cpu *CPU, instruction *Instruction) {
	var opcodes [3]string
	for i := 0; i < 3; i++ {
		s := "  "
		if i < len(t.Opcode) {
			op := t.Opcode[i]
			s = fmt.Sprintf("%02X", op)
		}

		opcodes[i] = s
	}

	fmt.Printf("%04X  %s %s %s  %s %28s"+
		"A:%02X X:%02X Y:%02X P:%02X SP:%02X CYC:\n",
		t.PC, opcodes[0], opcodes[1], opcodes[2],
		strings.ToUpper(instruction.Name), "",
		cpu.A, cpu.X, cpu.Y, cpu.GetFlags(), cpu.SP)
}

func (c *CPU) trace(instruction *Instruction, params ...interface{}) {
	if c.tracing == GoTracing {
		mode, _ := addressModeFromCall(instruction, params...)
		if !instruction.HasAddressing(mode) {
			panic(fmt.Sprintf("unexpected addressing mode type %T", mode))
		}

		c.TraceStep.Opcode = []byte{instruction.Addressing[mode].Opcode}
		// TODO add parameter opcodes and instruction params
	}

	c.TraceStep.print(c, instruction)
}
