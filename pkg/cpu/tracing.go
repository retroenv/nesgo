package cpu

import (
	"fmt"
	"strings"

	. "github.com/retroenv/nesgo/pkg/addressing"
	"github.com/retroenv/nesgo/pkg/disasm/param"
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
	PC             uint16
	Opcode         []byte
	Addressing     Mode
	Timing         byte
	PageCrossCycle bool
	PageCrossed    bool
	Unofficial     bool
	Instruction    string
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
	unofficial := " "
	if t.Unofficial {
		unofficial = "*"
	}

	// output the trace in a Nintendulator / nestest.log compatible format
	s := fmt.Sprintf("%04X  %s %s %s %s%-31s A:%02X X:%02X Y:%02X P:%02X SP:%02X CYC:%d\n",
		t.PC, opcodes[0], opcodes[1], opcodes[2], unofficial, t.Instruction,
		cpu.A, cpu.X, cpu.Y, cpu.GetFlags(), cpu.SP, cpu.cycles)
	if cpu.tracingTarget != nil {
		_, _ = fmt.Fprint(cpu.tracingTarget, s)
	} else {
		fmt.Print(s)
	}
}

// Trace logs the trace information of the passed instruction and its parameters.
// Params can be of length 0 to 2.
func (c *CPU) trace(instruction *Instruction, params ...interface{}) {
	var paramsAsString string

	if c.tracing == GoTracing {
		c.TraceStep.Addressing = c.addressModeFromCall(instruction, params...)
		if !instruction.HasAddressing(c.TraceStep.Addressing) {
			panic(fmt.Sprintf("unexpected addressing mode type %T", c.TraceStep.Addressing))
		}

		opcodeByte := instruction.Addressing[c.TraceStep.Addressing].Opcode

		var err error
		var firstParam interface{}
		if len(params) > 0 {
			firstParam = params[0]
		}
		paramsAsString, err = param.String(c.paramConverter, c.TraceStep.Addressing, firstParam)
		if err != nil {
			panic(err)
		}

		c.TraceStep.Opcode = []byte{opcodeByte}
		// TODO add parameter opcodes
	} else {
		paramsAsString = c.ParamString(instruction, params...)
	}

	c.TraceStep.Unofficial = instruction.Unofficial
	c.TraceStep.Instruction = strings.ToUpper(instruction.Name)
	if paramsAsString != "" {
		c.TraceStep.Instruction += " " + paramsAsString
	}
	c.TraceStep.print(c)
}

func shouldOutputMemoryContent(address uint16) bool {
	switch {
	case address < 0x0800:
		return true
	case address >= 0x4000 && address <= 0x4020:
		return true
	case address >= CodeBaseAddress:
		return true
	default:
		return false
	}
}

func addressModeFromCallNoParam(instruction *Instruction) Mode {
	if instruction.HasAddressing(AccumulatorAddressing) {
		return AccumulatorAddressing
	}
	// branches have no target in go mode
	if instruction.HasAddressing(RelativeAddressing) {
		return RelativeAddressing
	}
	return ImpliedAddressing
}
