//go:build !nesgo
// +build !nesgo

package cpu

import (
	"math"

	"github.com/retroenv/nesgo/pkg/addressing"
)

// Adc - Add with Carry.
func (c *CPU) Adc(params ...interface{}) {
	timeInstructionExecution()
	value := c.memory.ReadMemoryAddressModes(true, params...)
	sum := int(c.A) + int(c.Flags.C) + int(value)
	c.A = uint8(sum)
	c.setZN(c.A)

	if sum > math.MaxUint8 {
		c.Flags.C = 1
	} else {
		c.Flags.C = 0
	}
}

// And - AND with accumulator.
func (c *CPU) And(params ...interface{}) {
	timeInstructionExecution()
	value := c.memory.ReadMemoryAddressModes(true, params...)
	c.A &= value
	c.setZN(c.A)
}

// Asl - Arithmetic Shift Left.
func (c *CPU) Asl(params ...interface{}) {
	timeInstructionExecution()

	if params == nil { // A implied
		c.Flags.C = (c.A >> 7) & 1
		c.A <<= 1
		c.setZN(c.A)
		return
	}

	val := c.memory.ReadMemoryAddressModes(false, params...)
	c.Flags.C = (val >> 7) & 1
	val <<= 1
	c.setZN(val)
	c.memory.WriteMemoryAddressModes(val, params...)
}

// Bcc - Branch if Carry Clear - returns whether the
// carry flag is clear.
func (c *CPU) Bcc() bool {
	timeInstructionExecution()
	return c.Flags.C == 0
}

// BccInternal - Branch if Carry Clear.
func (c *CPU) BccInternal(params ...interface{}) {
	timeInstructionExecution()
	c.branch(c.Bcc, params[0])
}

// Bcs - Branch if Carry Set - returns whether the carry flag is set.
func (c *CPU) Bcs() bool {
	timeInstructionExecution()
	return c.Flags.C != 0
}

// BcsInternal - Branch if Carry Set.
func (c *CPU) BcsInternal(params ...interface{}) {
	timeInstructionExecution()
	c.branch(c.Bcs, params[0])
}

// Beq - Branch if Equal - returns whether the zero flag is set.
func (c *CPU) Beq() bool {
	timeInstructionExecution()
	return c.Flags.Z != 0
}

// BeqInternal - Branch if Equal.
func (c *CPU) BeqInternal(params ...interface{}) {
	timeInstructionExecution()
	c.branch(c.Beq, params[0])
}

// Bit - Bit Test - set the Z flag by ANDing A with given address content.
func (c *CPU) Bit(params ...interface{}) {
	timeInstructionExecution()
	value := c.memory.ReadMemoryAbsolute(params[0], nil)
	c.Flags.V = (value >> 6) & 1
	c.setZ(value & c.A)
	c.setN(value)
}

// Bmi - Branch if Minus - returns whether the negative flag is set.
func (c *CPU) Bmi() bool {
	timeInstructionExecution()
	return c.Flags.N != 0
}

// BmiInternal - Branch if Minus.
func (c *CPU) BmiInternal(params ...interface{}) {
	timeInstructionExecution()
	c.branch(c.Bmi, params[0])
}

// Bne - Branch if Not Equal - returns whether the zero flag is clear.
func (c *CPU) Bne() bool {
	timeInstructionExecution()
	return c.Flags.Z == 0
}

// BneInternal - Branch if Not Equal.
func (c *CPU) BneInternal(params ...interface{}) {
	timeInstructionExecution()
	c.branch(c.Bne, params[0])
}

// Bpl - Branch if Positive - returns whether the negative flag is clear.
func (c *CPU) Bpl() bool {
	timeInstructionExecution()
	return c.Flags.N == 0
}

// BplInternal - Branch if Positive.
func (c *CPU) BplInternal(params ...interface{}) {
	timeInstructionExecution()
	c.branch(c.Bpl, params[0])
}

// Brk - Force Interrupt.
func (c *CPU) Brk() {
	timeInstructionExecution()
	if *c.irqHandler != nil {
		f := *c.irqHandler
		f()
	}
}

// Bvc - Branch if Overflow Clear - returns whether the overflow flag is clear.
func (c *CPU) Bvc() bool {
	timeInstructionExecution()
	return c.Flags.V == 0
}

// BvcInternal - Branch if Overflow Clear.
func (c *CPU) BvcInternal(params ...interface{}) {
	timeInstructionExecution()
	c.branch(c.Bvc, params[0])
}

// Bvs - Branch if Overflow Set - returns whether the overflow flag is set.
func (c *CPU) Bvs() bool {
	timeInstructionExecution()
	return c.Flags.V != 0
}

// BvsInternal - Branch if Overflow Set.
func (c *CPU) BvsInternal(params ...interface{}) {
	timeInstructionExecution()
	c.branch(c.Bvs, params[0])
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
func (c *CPU) Cmp(params ...interface{}) {
	timeInstructionExecution()
	val := c.memory.ReadMemoryAddressModes(true, params[0])
	c.compare(c.A, val)
}

// Cpx - Compare X Register - compares the contents of X.
func (c *CPU) Cpx(params ...interface{}) {
	timeInstructionExecution()
	val := c.memory.ReadMemoryAddressModes(true, params[0])
	c.compare(c.X, val)
}

// Cpy - Compare Y Register - compares the contents of Y.
func (c *CPU) Cpy(params ...interface{}) {
	timeInstructionExecution()
	val := c.memory.ReadMemoryAddressModes(true, params[0])
	c.compare(c.Y, val)
}

// Dec - Decrement memory.
func (c *CPU) Dec(params ...interface{}) {
	timeInstructionExecution()
	val := c.memory.ReadMemoryAddressModes(false, params...)
	val--
	c.memory.WriteMemoryAddressModes(val, params...)
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
func (c *CPU) Eor(params ...interface{}) {
	timeInstructionExecution()
	value := c.memory.ReadMemoryAddressModes(true, params...)
	c.A ^= value
	c.setZN(c.A)
}

// Inc - Increments memory.
func (c *CPU) Inc(params ...interface{}) {
	timeInstructionExecution()
	val := c.memory.ReadMemoryAddressModes(false, params...)
	val++
	c.memory.WriteMemoryAddressModes(val, params...)
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

// Jmp - jump to address.
func (c *CPU) Jmp(params ...interface{}) {
	timeInstructionExecution()
	// TODO implement
}

// Jsr - jump to subroutine.
func (c *CPU) Jsr(params ...interface{}) {
	timeInstructionExecution()

	c.push16(c.PC - 1)

	addr := params[0].(addressing.Absolute)
	c.PC = uint16(addr)
}

// Lda - Load Accumulator - load a byte into A.
func (c *CPU) Lda(params ...interface{}) {
	timeInstructionExecution()
	c.A = c.memory.ReadMemoryAddressModes(true, params...)
	c.setZN(c.A)
}

// Ldx - Load X Register - load a byte into X.
func (c *CPU) Ldx(params ...interface{}) {
	timeInstructionExecution()
	c.X = c.memory.ReadMemoryAddressModes(true, params...)
	c.setZN(c.X)
}

// Ldy - Load Y Register - load a byte into Y.
func (c *CPU) Ldy(params ...interface{}) {
	timeInstructionExecution()
	c.Y = c.memory.ReadMemoryAddressModes(true, params...)
	c.setZN(c.Y)
}

// Lsr - Logical Shift Right.
func (c *CPU) Lsr(params ...interface{}) {
	timeInstructionExecution()

	if params == nil { // A implied
		c.Flags.C = c.A & 1
		c.A >>= 1
		c.setZN(c.A)
		return
	}

	val := c.memory.ReadMemoryAddressModes(false, params...)
	c.Flags.C = val & 1
	val >>= 1
	c.setZN(val)
	c.memory.WriteMemoryAddressModes(val, params...)
}

// Nop - No Operation.
func (c *CPU) Nop() {
	timeInstructionExecution()
}

// Ora - OR with Accumulator.
func (c *CPU) Ora(params ...interface{}) {
	timeInstructionExecution()
	value := c.memory.ReadMemoryAddressModes(true, params...)
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
	f := c.GetFlags()
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
func (c *CPU) Rol(params ...interface{}) {
	timeInstructionExecution()

	cFlag := c.Flags.C
	if params == nil { // A implied
		c.Flags.C = (c.A >> 7) & 1
		c.A = (c.A << 1) | cFlag
		c.setZN(c.A)
		return
	}

	val := c.memory.ReadMemoryAddressModes(false, params...)
	c.Flags.C = (val >> 7) & 1
	val = (val << 1) | cFlag
	c.setZN(val)
	c.memory.WriteMemoryAddressModes(val, params...)
}

// Ror - Rotate Right.
func (c *CPU) Ror(params ...interface{}) {
	timeInstructionExecution()

	cFlag := c.Flags.C
	if params == nil { // A implied
		c.Flags.C = c.A & 1
		c.A = (c.A >> 1) | (cFlag << 7)
		c.setZN(c.A)
		return
	}

	val := c.memory.ReadMemoryAddressModes(false, params...)
	c.Flags.C = val & 1
	val = (val >> 1) | (cFlag << 7)
	c.setZN(val)
	c.memory.WriteMemoryAddressModes(val, params...)
}

// Rti - Return from Interrupt.
func (c *CPU) Rti() {
	timeInstructionExecution()
}

// RtiInternal - Return from Interrupt.
func (c *CPU) RtiInternal() {
	timeInstructionExecution()

	c.PC = c.pop16()
}

// Rts - return from subroutine.
func (c *CPU) Rts() {
	timeInstructionExecution()

	c.PC = c.pop16() + 1
}

// Sbc - subtract with Carry.
func (c *CPU) Sbc(params ...interface{}) {
	timeInstructionExecution()

	value := c.memory.ReadMemoryAddressModes(true, params...)
	sub := int(c.A) - int(value) - (1 - int(c.Flags.C))
	c.A = uint8(sub)
	c.setZN(c.A)

	if sub >= 0 {
		c.Flags.C = 1
	} else {
		c.Flags.C = 0
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
func (c *CPU) Sta(params ...interface{}) {
	timeInstructionExecution()
	c.memory.WriteMemoryAddressModes(c.A, params...)
}

// Stx - Store X Register - store content of X at address Addr and
// add an optional register to the address.
func (c *CPU) Stx(params ...interface{}) {
	timeInstructionExecution()
	c.memory.WriteMemoryAddressModes(c.X, params...)
}

// Sty - Store Y Register - store content of Y at address Addr and
// add an optional register to the address.
func (c *CPU) Sty(params ...interface{}) {
	timeInstructionExecution()
	c.memory.WriteMemoryAddressModes(c.Y, params...)
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
