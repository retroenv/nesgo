//go:build !nesgo
// +build !nesgo

package cpu

import (
	. "github.com/retroenv/nesgo/pkg/addressing"
)

// execute branch jump if the branching op result is true.
func (c *CPU) branch(branchTo bool, param interface{}) {
	// disable trace while calling the go mode branch code
	// TODO refactor to avoid this
	trace := c.tracing
	c.tracing = NoTracing


	if branchTo {
		addr := param.(Absolute)

		c.PC = uint16(addr)
		c.cycles++
	}

	c.tracing = trace
}

// hasAccumulatorParam returns whether the passed or missing parameter
// indicates usage of the accumulator register.
func hasAccumulatorParam(params ...interface{}) bool {
	if params == nil {
		return true
	}
	param := params[0]
	_, ok := param.(Accumulator)
	return ok
}

func (c *CPU) setFlags(flags uint8) {
	c.Flags.C = (flags >> 0) & 1
	c.Flags.Z = (flags >> 1) & 1
	c.Flags.I = (flags >> 2) & 1
	c.Flags.D = (flags >> 3) & 1
	c.Flags.B = (flags >> 4) & 1
	c.Flags.U = (flags >> 5) & 1
	c.Flags.V = (flags >> 6) & 1
	c.Flags.N = (flags >> 7) & 1
}

// GetFlags returns the current state of flags as byte.
func (c *CPU) GetFlags() uint8 {
	var f byte
	f |= c.Flags.C << 0
	f |= c.Flags.Z << 1
	f |= c.Flags.I << 2
	f |= c.Flags.D << 3
	f |= c.Flags.B << 4
	f |= c.Flags.U << 5
	f |= c.Flags.V << 6
	f |= c.Flags.N << 7
	return f
}

// setZ - set the zero flag if the argument is zero.
func (c *CPU) setZ(value uint8) {
	if value == 0 {
		c.Flags.Z = 1
	} else {
		c.Flags.Z = 0
	}
}

// setN - set the negative flag if the argument is negative (high bit is set).
func (c *CPU) setN(value uint8) {
	if value&0x80 != 0 {
		c.Flags.N = 1
	} else {
		c.Flags.N = 0
	}
}

// setV - set the overflow flag.
func (c *CPU) setV(set bool) {
	if set {
		c.Flags.V = 1
	} else {
		c.Flags.V = 0
	}
}

func (c *CPU) setZN(value uint8) {
	c.setZ(value)
	c.setN(value)
}

func (c *CPU) compare(a, b byte) {
	c.setZN(a - b)
	if a >= b {
		c.Flags.C = 1
	} else {
		c.Flags.C = 0
	}
}

// push a value to the stack and update the stack pointer.
func (c *CPU) push(value byte) {
	c.memory.WriteMemory(uint16(StackBase+int(c.SP)), value)
	c.SP--
}

// Push16 a word to the stack and update the stack pointer.
func (c *CPU) Push16(value uint16) {
	high := byte(value >> 8)
	low := byte(value)
	c.push(high)
	c.push(low)
}

// Pop pops a byte from the stack and update the stack pointer.
func (c *CPU) Pop() byte {
	c.SP++
	return c.memory.ReadMemory(uint16(StackBase + int(c.SP)))
}

// Pop16 pops a word from the stack and updates the stack pointer.
func (c *CPU) Pop16() uint16 {
	low := uint16(c.Pop())
	high := uint16(c.Pop())
	return high<<8 | low
}
