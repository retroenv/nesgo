// Package cpu provides CPU (Central Processing Unit) functionality.
package cpu

import (
	"io"
	"sync"

	"github.com/retroenv/nesgo/pkg/bus"
	"github.com/retroenv/nesgo/pkg/disasm/ca65"
	"github.com/retroenv/nesgo/pkg/disasm/param"
)

const (
	StackBase = 0x100

	initialCycles = 7
	initialFlags  = 0b0010_0100 // I and U flags are 1, the rest 0
	InitialStack  = 0xFD
)

// CPU implements a MOS Technology 6502 CPU.
type CPU struct {
	mu sync.RWMutex

	A     uint8  // accumulator
	X     uint8  // x register
	Y     uint8  // y register
	PC    uint16 // program counter
	SP    uint8  // stack pointer
	Flags flags

	bus        *bus.Bus
	irqHandler *func()

	cycles      uint64
	stallCycles uint16 // TODO stall cycles, use a Step() function

	tracing        TracingMode
	tracingTarget  io.Writer
	TraceStep      TraceStep
	paramConverter param.Converter
	lastFunction   string

	triggerNmi bool // TODO use
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
func New(bus *bus.Bus, irqHandler *func()) *CPU {
	c := &CPU{
		SP:             InitialStack,
		bus:            bus,
		irqHandler:     irqHandler,
		cycles:         initialCycles,
		paramConverter: ca65.ParamConverter{},
	}

	// read reset interrupt handler address
	c.PC = bus.Memory.ReadWord(0xFFFC)

	c.setFlags(initialFlags)
	return c
}

// SetTracing sets the CPU tracing options.
func (c *CPU) SetTracing(mode TracingMode, target io.Writer) {
	c.tracing = mode
	c.tracingTarget = target
}

// ResetCycles sets the cycle counter to 0.
// This is useful for counting used CPU cycles for a function.
func (c *CPU) ResetCycles() {
	c.cycles = 0
}

// Cycles returns the amount of CPU cycles executed since system start.
func (c *CPU) Cycles() uint64 {
	return c.cycles
}

// StallCycles stalls the CPU for the given amount of cycles. This is used for DMA transfer in the PPU.
func (c *CPU) StallCycles(cycles uint16) {
	c.stallCycles = cycles
}

// TriggerNMI causes a non-maskable interrupt to occur on the next cycle.
func (c *CPU) TriggerNMI() {
	c.triggerNmi = true
}

// writeLock takes the mutex write lock and returns a function to write unlock to allow an easy use for
// the lock/unlock mechanism in the form of defer c.writeLock()()
func (c *CPU) writeLock() func() {
	c.mu.Lock()
	return c.writeUnlock
}

func (c *CPU) writeUnlock() {
	c.mu.Unlock()
}
