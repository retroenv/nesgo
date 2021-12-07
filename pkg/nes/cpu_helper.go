//go:build !nesgo
// +build !nesgo

package nes

const stackBase = 0x100

// CPU implements a MOS Technology 650 CPU.
type CPU struct {
	A     uint8 // accumulator
	X     uint8 // x register
	Y     uint8 // y register
	SP    uint8 // stack pointer
	Flags flags
}

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

const (
	initialFlags = 0x24 // I and U flags are 1, the rest 0
	initialStack = 0xFD
)

func newCPU() *CPU {
	c := &CPU{
		SP: initialStack,
	}
	c.setFlags(initialFlags)
	return c
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

func (c *CPU) flags() uint8 {
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
	writeMemory(uint16(stackBase+int(c.SP)), value)
	c.SP--
}

// pop a value from the stack and update the stack pointer.
func (c *CPU) pop() byte {
	c.SP++
	return readMemory(uint16(stackBase + int(c.SP)))
}
