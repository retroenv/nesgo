//go:build !nesgo
// +build !nesgo

package cpu

import (
	. "github.com/retroenv/nesgo/pkg/addressing"
)

// execute branch jump if the branching op result is true.
func (c *CPU) branch(branchFunc func() bool, param interface{}) {
	if branchFunc() {
		addr := param.(Absolute)
		c.PC = uint16(addr)
	}
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

// push a value to the stack and update the stack pointer.
func (c *CPU) push16(value uint16) {
	high := byte(value >> 8)
	low := byte(value & 0xFF)
	c.push(high)
	c.push(low)
}

// pop a value from the stack and update the stack pointer.
func (c *CPU) pop() byte {
	c.SP++
	return c.memory.ReadMemory(uint16(StackBase + int(c.SP)))
}

// pop a value from the stack and update the stack pointer.
func (c *CPU) pop16() uint16 {
	low := uint16(c.pop())
	high := uint16(c.pop())
	return high<<8 | low
}
