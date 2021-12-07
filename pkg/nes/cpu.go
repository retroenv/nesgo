//go:build !nesgo
// +build !nesgo

package nes

import "math"

var notImplemented = "instruction is not implemented yet"

// Adc - Add with Carry.
func (c *CPU) Adc(param interface{}, reg ...interface{}) {
	timeInstructionExecution()
	value := c.memory.readMemoryAddressModes(param, reg...)
	sum := int(c.A) + int(c.Flags.C) + int(value)
	c.A = uint8(sum)
	c.setZN(c.A)

	if sum > math.MaxUint8 {
		c.Flags.C = 1
	} else {
		c.Flags.C = 0
	}

	if c.Flags.D == 1 {
		panic(notImplemented) // TODO: support decimal mode
	}
}

// And - AND with accumulator.
func (c *CPU) And(param interface{}, reg ...interface{}) {
	timeInstructionExecution()
	value := c.memory.readMemoryAddressModes(param, reg...)
	c.A &= value
	c.setZN(c.A)
}

// Asl - Arithmetic Shift Left.
func (c *CPU) Asl(param ...interface{}) {
	timeInstructionExecution()

	if param == nil { // A implied
		c.Flags.C = (c.A >> 7) & 1
		c.A <<= 1
		c.setZN(c.A)
		return
	}

	val := c.memory.readMemoryAddressModes(param)
	c.Flags.C = (val >> 7) & 1
	val <<= 1
	c.setZN(val)
	c.memory.writeMemoryAddressModes(param, val)
}

// Bcc - Branch if Carry Clear - returns whether the
// carry flag is clear.
func (c *CPU) Bcc() bool {
	timeInstructionExecution()
	return c.Flags.C == 0
}

// Bcs - Branch if Carry Set - returns whether the carry flag is set.
func (c *CPU) Bcs() bool {
	timeInstructionExecution()
	return c.Flags.C != 0
}

// Beq - Branch if Equal - returns whether the zero flag is set.
func (c *CPU) Beq() bool {
	timeInstructionExecution()
	return c.Flags.Z != 0
}

// Bit - Bit Test - set the Z flag by ANDing A with given address content.
func (c *CPU) Bit(address uint16) {
	timeInstructionExecution()
	value := c.memory.readMemoryAbsolute(address)
	c.Flags.V = (value >> 6) & 1
	c.setZ(value & c.A)
	c.setN(value)
}

// Bmi - Branch if Minus - returns whether the negative flag is set.
func (c *CPU) Bmi() bool {
	timeInstructionExecution()
	return c.Flags.N != 0
}

// Bne - Branch if Not Equal - returns whether the zero flag is clear.
func (c *CPU) Bne() bool {
	timeInstructionExecution()
	return c.Flags.Z == 0
}

// Bpl - Branch if Positive - returns whether the negative flag is clear.
func (c *CPU) Bpl() bool {
	timeInstructionExecution()
	return c.Flags.N == 0
}

// Brk - Force Interrupt.
func (c *CPU) Brk() {
	timeInstructionExecution()
	panic(notImplemented) // TODO: implement
}

// Bvc - Branch if Overflow Clear - returns whether the overflow flag is clear.
func (c *CPU) Bvc() bool {
	timeInstructionExecution()
	return c.Flags.V == 0
}

// Bvs - Branch if Overflow Set - returns whether the overflow flag is set.
func (c *CPU) Bvs() bool {
	timeInstructionExecution()
	return c.Flags.V != 0
}

// Clc - Clear Carry Flag.
func (c *CPU) Clc() {
	timeInstructionExecution()
	c.Flags.C = 0
}

// Cld - Clear Decimal Mode.
func (c *CPU) Cld() {
	timeInstructionExecution()
	c.Flags.D = 0
}

// Cli - Clear Interrupt Disable.
func (c *CPU) Cli() {
	timeInstructionExecution()
	c.Flags.I = 0
}

// Clv - Clear Overflow Flag.
func (c *CPU) Clv() {
	timeInstructionExecution()
	c.Flags.V = 0
}

// Cmp - Compare - compares the contents of A.
func (c *CPU) Cmp(param interface{}) {
	timeInstructionExecution()
	val := c.memory.readMemoryAddressModes(param)
	c.compare(c.A, val)
}

// Cpx - Compare X Register - compares the contents of X.
func (c *CPU) Cpx(param interface{}) {
	timeInstructionExecution()
	val := c.memory.readMemoryAddressModes(param)
	c.compare(c.X, val)
}

// Cpy - Compare Y Register - compares the contents of Y.
func (c *CPU) Cpy(param interface{}) {
	timeInstructionExecution()
	val := c.memory.readMemoryAddressModes(param)
	c.compare(c.Y, val)
}

// Dex - Decrement X Register.
func (c *CPU) Dex() {
	timeInstructionExecution()
	c.X--
	c.setZN(c.X)
}

// Dey - Decrement Y Register.
func (c *CPU) Dey() {
	timeInstructionExecution()
	c.Y--
	c.setZN(c.Y)
}

// Eor - Exclusive OR - XOR.
func (c *CPU) Eor(param interface{}, reg ...interface{}) {
	timeInstructionExecution()
	value := c.memory.readMemoryAddressModes(param, reg...)
	c.A ^= value
	c.setZN(c.A)
}

// Inx - Increment X Register.
func (c *CPU) Inx() {
	timeInstructionExecution()
	c.X++
	c.setZN(c.X)
}

// Iny - Increment Y Register.
func (c *CPU) Iny() {
	timeInstructionExecution()
	c.Y++
	c.setZN(c.Y)
}

// Lda - Load Accumulator - load a byte into A.
func (c *CPU) Lda(param interface{}, reg ...interface{}) {
	timeInstructionExecution()
	c.A = c.memory.readMemoryAddressModes(param, reg...)
	c.setZN(c.A)
}

// Ldx - Load X Register - load a byte into X.
func (c *CPU) Ldx(param interface{}, reg ...interface{}) {
	timeInstructionExecution()
	c.X = c.memory.readMemoryAddressModes(param, reg...)
	c.setZN(c.X)
}

// Ldy - Load Y Register - load a byte into Y.
func (c *CPU) Ldy(param interface{}, reg ...interface{}) {
	timeInstructionExecution()
	c.Y = c.memory.readMemoryAddressModes(param, reg...)
	c.setZN(c.Y)
}

// Lsr - Logical Shift Right.
func (c *CPU) Lsr(param ...interface{}) {
	timeInstructionExecution()

	if param == nil { // A implied
		c.Flags.C = c.A & 1
		c.A >>= 1
		c.setZN(c.A)
		return
	}

	val := c.memory.readMemoryAddressModes(param)
	c.Flags.C = val & 1
	val >>= 1
	c.setZN(val)
	c.memory.writeMemoryAddressModes(param, val)
}

// Nop - No Operation.
func (c *CPU) Nop() {
	timeInstructionExecution()
}

// Ora - OR with Accumulator.
func (c *CPU) Ora(param interface{}, reg ...interface{}) {
	timeInstructionExecution()
	value := c.memory.readMemoryAddressModes(param, reg...)
	c.A |= value
	c.setZN(c.A)
}

// Pha - Push Accumulator - push A content to stack.
func (c *CPU) Pha() {
	timeInstructionExecution()
	c.push(c.A)
}

// Php - Push Processor Status - push status flags to stack.
func (c *CPU) Php() {
	timeInstructionExecution()
	f := c.flags()
	f |= 0b11000 // bit 4 and 5 are set to 1
	c.push(f)
}

// Pla - Pull Accumulator - pull A content from stack.
func (c *CPU) Pla() {
	timeInstructionExecution()
	c.A = c.pop()
	c.setZN(c.A)
}

// Plp - Pull Processor Status - pull status flags from stack.
func (c *CPU) Plp() {
	timeInstructionExecution()
	f := c.pop()
	f &^= 0b11000 // bit 4 and 5 are cleared
	c.setFlags(f)
}

// Rol - Rotate Left.
func (c *CPU) Rol(param ...interface{}) {
	timeInstructionExecution()

	cFlag := c.Flags.C
	if param == nil { // A implied
		c.Flags.C = (c.A >> 7) & 1
		c.A = (c.A << 1) | cFlag
		c.setZN(c.A)
		return
	}

	val := c.memory.readMemoryAddressModes(param)
	c.Flags.C = (val >> 7) & 1
	val = (val << 1) | cFlag
	c.setZN(val)
	c.memory.writeMemoryAddressModes(param, val)
}

// Ror - Rotate Right.
func (c *CPU) Ror(param ...interface{}) {
	timeInstructionExecution()

	cFlag := c.Flags.C
	if param == nil { // A implied
		c.Flags.C = c.A & 1
		c.A = (c.A >> 1) | (cFlag << 7)
		c.setZN(c.A)
		return
	}

	val := c.memory.readMemoryAddressModes(param)
	c.Flags.C = val & 1
	val = (val >> 1) | (cFlag << 7)
	c.setZN(val)
	c.memory.writeMemoryAddressModes(param, val)
}

// Rti - Return from Interrupt.
func (c *CPU) Rti() {
	timeInstructionExecution()
}

// Sbc - subtract with Carry.
func (c *CPU) Sbc(param interface{}, reg ...interface{}) {
	timeInstructionExecution()

	value := c.memory.readMemoryAddressModes(param, reg...)
	sub := int(c.A) - int(value) - (1 - int(c.Flags.C))
	c.A = uint8(sub)
	c.setZN(c.A)

	if sub >= 0 {
		c.Flags.C = 1
	} else {
		c.Flags.C = 0
	}

	if c.Flags.D == 1 {
		panic(notImplemented) // TODO: support decimal mode
	}
}

// Sec - Set Carry Flag.
func (c *CPU) Sec() {
	timeInstructionExecution()
	c.Flags.C = 1
}

// Sed - Set Decimal Flag.
func (c *CPU) Sed() {
	timeInstructionExecution()
	c.Flags.D = 1
}

// Sei - Set Interrupt Disable.
func (c *CPU) Sei() {
	timeInstructionExecution()
	c.Flags.I = 1
}

// Sta - Store Accumulator - store content of A at address Addr and
// add an optional register to the address.
func (c *CPU) Sta(param interface{}, reg ...interface{}) {
	timeInstructionExecution()
	c.memory.writeMemoryAddressModes(param, c.A, reg...)
}

// Stx - Store X Register - store content of X at address Addr and
// add an optional register to the address.
func (c *CPU) Stx(param interface{}, reg ...interface{}) {
	timeInstructionExecution()
	c.memory.writeMemoryAddressModes(param, c.X, reg...)
}

// Sty - Store Y Register - store content of Y at address Addr and
// add an optional register to the address.
func (c *CPU) Sty(param interface{}, reg ...interface{}) {
	timeInstructionExecution()
	c.memory.writeMemoryAddressModes(param, c.Y, reg...)
}

// Tax - Transfer Accumulator to X.
func (c *CPU) Tax() {
	timeInstructionExecution()
	c.X = c.A
	c.setZN(c.X)
}

// Tay - Transfer Accumulator to Y.
func (c *CPU) Tay() {
	timeInstructionExecution()
	c.Y = c.A
	c.setZN(c.Y)
}

// Tsx - Transfer Stack Pointer to X.
func (c *CPU) Tsx() {
	timeInstructionExecution()
	c.X = c.SP
	c.setZN(c.X)
}

// Txa - Transfer X to Accumulator.
func (c *CPU) Txa() {
	timeInstructionExecution()
	c.A = c.X
	c.setZN(c.A)
}

// Txs - Transfer X to Stack Pointer.
func (c *CPU) Txs() {
	timeInstructionExecution()
	c.SP = c.X
}

// Tya - Transfer Y to Accumulator.
func (c *CPU) Tya() {
	timeInstructionExecution()
	c.A = c.Y
	c.setZN(c.A)
}
