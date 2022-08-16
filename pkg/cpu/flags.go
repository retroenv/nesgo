package cpu

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
