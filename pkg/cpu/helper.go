//go:build !nesgo

package cpu

import (
	. "github.com/retroenv/nesgo/pkg/addressing"
)

// execute branch jump if the branching op result is true.
func (c *CPU) branch(branchTo bool, param any) {
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
func hasAccumulatorParam(params ...any) bool {
	if params == nil {
		return true
	}
	param := params[0]
	_, ok := param.(Accumulator)
	return ok
}

// push a value to the stack and update the stack pointer.
func (c *CPU) push(value byte) {
	c.bus.Memory.Write(uint16(StackBase+int(c.SP)), value)
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
	return c.bus.Memory.Read(uint16(StackBase + int(c.SP)))
}

// Pop16 pops a word from the stack and updates the stack pointer.
func (c *CPU) Pop16() uint16 {
	low := uint16(c.Pop())
	high := uint16(c.Pop())
	return high<<8 | low
}
