// Package cpu provides CPU (Central  Processing Unit) functionality.
package cpu

import "io"

const (
	StackBase = 0x100

	initialFlags = 0b00100100 // I and U flags are 1, the rest 0
	InitialStack = 0xFD
)

// CPU implements a MOS Technology 650 CPU.
type CPU struct {
	A     uint8  // accumulator
	X     uint8  // x register
	Y     uint8  // y register
	PC    uint16 // program counter
	SP    uint8  // stack pointer
	Flags flags

	irqHandler    *func()
	memory        memory
	tracing       TracingMode
	tracingTarget io.Writer
	TraceStep     TraceStep
}

// Bit No.   7   6   5   4   3   2   1   0
// Flag      S   V       B   D   I   Z   C
type flags struct {
	C uint8 // carry flag
	Z uint8 // zero flag
	I uint8 // interrupt disable flag
	D uint8 // decimal mode flag
	B uint8 // break command flag
	U uint8 // unused flag
	V uint8 // overflow flag
	N uint8 // negative flag
}

// New creates a new CPU.
func New(memory memory, irqHandler *func()) *CPU {
	c := &CPU{
		SP:         InitialStack,
		memory:     memory,
		irqHandler: irqHandler,
	}

	// read reset interrupt handler address
	c.PC = memory.ReadMemory16(0xFFFC)

	c.setFlags(initialFlags)
	return c
}

// SetTracing switches CPU tracing.
func (c *CPU) SetTracing(mode TracingMode, target io.Writer) {
	c.tracing = mode
	c.tracingTarget = target
}
