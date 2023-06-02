//go:build !nesgo

package cpu

import (
	"fmt"
	"math"

	. "github.com/retroenv/retrogolib/addressing"
	"github.com/retroenv/retrogolib/arch/cpu/m6502"
)

// Adc - Add with Carry.
func (c *CPU) Adc(params ...any) {
	defer c.instructionHook(m6502.Adc, params...)()

	c.adcInternal(params...)
}

func (c *CPU) adcInternal(params ...any) {
	a := c.A
	value := c.bus.Memory.ReadAddressModes(true, params...)
	sum := int(c.A) + int(c.Flags.C) + int(value)
	c.A = uint8(sum)
	c.setZN(c.A)

	if sum > math.MaxUint8 {
		c.Flags.C = 1
	} else {
		c.Flags.C = 0
	}
	c.setV((a^value)&0x80 == 0 && (a^c.A)&0x80 != 0)
}

// And - AND with accumulator.
func (c *CPU) And(params ...any) {
	defer c.instructionHook(m6502.And, params...)()
	c.andInternal(params...)
}

func (c *CPU) andInternal(params ...any) {
	value := c.bus.Memory.ReadAddressModes(true, params...)
	c.A &= value
	c.setZN(c.A)
}

// Asl - Arithmetic Shift Left.
func (c *CPU) Asl(params ...any) {
	defer c.instructionHook(m6502.Asl, params...)()
	c.aslInternal(params...)
}

func (c *CPU) aslInternal(params ...any) {
	if hasAccumulatorParam(params...) {
		c.Flags.C = (c.A >> 7) & 1
		c.A <<= 1
		c.setZN(c.A)
		return
	}

	val := c.bus.Memory.ReadAddressModes(false, params...)
	c.Flags.C = (val >> 7) & 1
	val <<= 1
	c.setZN(val)
	c.bus.Memory.WriteAddressModes(val, params...)
}

// Bcc - Branch if Carry Clear - returns whether the
// carry flag is clear.
func (c *CPU) Bcc() bool {
	defer c.instructionHook(m6502.Bcc)()
	return c.Flags.C == 0
}

// BccInternal - Branch if Carry Clear.
func (c *CPU) BccInternal(params ...any) {
	defer c.instructionHook(m6502.Bcc, params...)()
	c.branch(c.Flags.C == 0, params[0])
}

// Bcs - Branch if Carry Set - returns whether the carry flag is set.
func (c *CPU) Bcs() bool {
	defer c.instructionHook(m6502.Bcs)()
	return c.Flags.C != 0
}

// BcsInternal - Branch if Carry Set.
func (c *CPU) BcsInternal(params ...any) {
	defer c.instructionHook(m6502.Bcs, params...)()
	c.branch(c.Flags.C != 0, params[0])
}

// Beq - Branch if Equal - returns whether the zero flag is set.
func (c *CPU) Beq() bool {
	defer c.instructionHook(m6502.Beq)()
	return c.Flags.Z != 0
}

// BeqInternal - Branch if Equal.
func (c *CPU) BeqInternal(params ...any) {
	defer c.instructionHook(m6502.Beq, params...)()
	c.branch(c.Flags.Z != 0, params[0])
}

// Bit - Bit Test - set the Z flag by ANDing A with given address content.
func (c *CPU) Bit(params ...any) {
	defer c.instructionHook(m6502.Bit, params...)()

	value := c.bus.Memory.ReadAbsolute(params[0], nil)
	c.setV((value>>6)&1 == 1)
	c.setZ(value & c.A)
	c.setN(value)
}

// Bmi - Branch if Minus - returns whether the negative flag is set.
func (c *CPU) Bmi() bool {
	defer c.instructionHook(m6502.Bmi)()
	return c.Flags.N != 0
}

// BmiInternal - Branch if Minus.
func (c *CPU) BmiInternal(params ...any) {
	defer c.instructionHook(m6502.Bmi, params...)()
	c.branch(c.Flags.N != 0, params[0])
}

// Bne - Branch if Not Equal - returns whether the zero flag is clear.
func (c *CPU) Bne() bool {
	defer c.instructionHook(m6502.Bne)()
	return c.Flags.Z == 0
}

// BneInternal - Branch if Not Equal.
func (c *CPU) BneInternal(params ...any) {
	defer c.instructionHook(m6502.Bne, params...)()
	c.branch(c.Flags.Z == 0, params[0])
}

// Bpl - Branch if Positive - returns whether the negative flag is clear.
func (c *CPU) Bpl() bool {
	defer c.instructionHook(m6502.Bpl)()
	return c.Flags.N == 0
}

// BplInternal - Branch if Positive.
func (c *CPU) BplInternal(params ...any) {
	defer c.instructionHook(m6502.Bpl, params...)()
	c.branch(c.Flags.N == 0, params[0])
}

// Brk - Force Interrupt.
func (c *CPU) Brk() {
	unlock := c.instructionHook(m6502.Brk)
	unlock()
	c.irq()
}

// Bvc - Branch if Overflow Clear - returns whether the overflow flag is clear.
func (c *CPU) Bvc() bool {
	defer c.instructionHook(m6502.Bvc)()
	return c.Flags.V == 0
}

// BvcInternal - Branch if Overflow Clear.
func (c *CPU) BvcInternal(params ...any) {
	defer c.instructionHook(m6502.Bvc, params...)()
	c.branch(c.Flags.V == 0, params[0])
}

// Bvs - Branch if Overflow Set - returns whether the overflow flag is set.
func (c *CPU) Bvs() bool {
	defer c.instructionHook(m6502.Bvs)()
	return c.Flags.V != 0
}

// BvsInternal - Branch if Overflow Set.
func (c *CPU) BvsInternal(params ...any) {
	defer c.instructionHook(m6502.Bvs, params...)()
	c.branch(c.Flags.V != 0, params[0])
}

// Clc - Clear Carry Flag.
func (c *CPU) Clc() {
	defer c.instructionHook(m6502.Clc)()
	c.Flags.C = 0
}

// Cld - Clear Decimal Mode.
func (c *CPU) Cld() {
	defer c.instructionHook(m6502.Cld)()
	c.Flags.D = 0
}

// Cli - Clear Interrupt Disable.
func (c *CPU) Cli() {
	defer c.instructionHook(m6502.Cli)()
	c.Flags.I = 0
}

// Clv - Clear Overflow Flag.
func (c *CPU) Clv() {
	defer c.instructionHook(m6502.Clv)()
	c.Flags.V = 0
}

// Cmp - Compare - compares the contents of A.
func (c *CPU) Cmp(params ...any) {
	defer c.instructionHook(m6502.Cmp, params...)()

	val := c.bus.Memory.ReadAddressModes(true, params...)
	c.compare(c.A, val)
}

// Cpx - Compare X Register - compares the contents of X.
func (c *CPU) Cpx(params ...any) {
	defer c.instructionHook(m6502.Cpx, params...)()

	val := c.bus.Memory.ReadAddressModes(true, params[0])
	c.compare(c.X, val)
}

// Cpy - Compare Y Register - compares the contents of Y.
func (c *CPU) Cpy(params ...any) {
	defer c.instructionHook(m6502.Cpy, params...)()

	val := c.bus.Memory.ReadAddressModes(true, params[0])
	c.compare(c.Y, val)
}

// Dec - Decrement memory.
func (c *CPU) Dec(params ...any) {
	defer c.instructionHook(m6502.Dec, params...)()
	c.decInternal(params...)
}

func (c *CPU) decInternal(params ...any) {
	val := c.bus.Memory.ReadAddressModes(false, params...)
	val--
	c.bus.Memory.WriteAddressModes(val, params...)
	c.setZN(val)
}

// Dex - Decrement X Register.
func (c *CPU) Dex() {
	defer c.instructionHook(m6502.Dex)()

	c.X--
	c.setZN(c.X)
}

// Dey - Decrement Y Register.
func (c *CPU) Dey() {
	defer c.instructionHook(m6502.Dey)()

	c.Y--
	c.setZN(c.Y)
}

// Eor - Exclusive OR - XOR.
func (c *CPU) Eor(params ...any) {
	defer c.instructionHook(m6502.Eor, params...)()
	c.eorInternal(params...)
}

func (c *CPU) eorInternal(params ...any) {
	value := c.bus.Memory.ReadAddressModes(true, params...)
	c.A ^= value
	c.setZN(c.A)
}

// Inc - Increments memory.
func (c *CPU) Inc(params ...any) {
	defer c.instructionHook(m6502.Inc, params...)()
	c.incInternal(params...)
}

func (c *CPU) incInternal(params ...any) {
	val := c.bus.Memory.ReadAddressModes(false, params...)
	val++
	c.bus.Memory.WriteAddressModes(val, params...)
	c.setZN(val)
}

// Inx - Increment X Register.
func (c *CPU) Inx() {
	defer c.instructionHook(m6502.Inx)()

	c.X++
	c.setZN(c.X)
}

// Iny - Increment Y Register.
func (c *CPU) Iny() {
	defer c.instructionHook(m6502.Iny)()

	c.Y++
	c.setZN(c.Y)
}

// Jmp - jump to address.
func (c *CPU) Jmp(params ...any) {
	defer c.instructionHook(m6502.Jmp, params...)()

	param := params[0]
	switch address := param.(type) {
	case Absolute:
		c.PC = uint16(address)
	case Indirect:
		c.PC = c.bus.Memory.ReadWordBug(uint16(address))

	default:
		panic(fmt.Sprintf("unsupported jmp mode type %T", param))
	}
}

// Jsr - jump to subroutine.
func (c *CPU) Jsr(params ...any) {
	defer c.instructionHook(m6502.Jsr, params...)()

	c.Push16(c.PC + 2)

	addr := params[0].(Absolute)
	c.PC = uint16(addr)
}

// Lda - Load Accumulator - load a byte into A.
func (c *CPU) Lda(params ...any) {
	defer c.instructionHook(m6502.Lda, params...)()

	c.A = c.bus.Memory.ReadAddressModes(true, params...)
	c.setZN(c.A)
}

// Ldx - Load X Register - load a byte into X.
func (c *CPU) Ldx(params ...any) {
	defer c.instructionHook(m6502.Ldx, params...)()

	c.X = c.bus.Memory.ReadAddressModes(true, params...)
	c.setZN(c.X)
}

// Ldy - Load Y Register - load a byte into Y.
func (c *CPU) Ldy(params ...any) {
	defer c.instructionHook(m6502.Ldy, params...)()

	c.Y = c.bus.Memory.ReadAddressModes(true, params...)
	c.setZN(c.Y)
}

// Lsr - Logical Shift Right.
func (c *CPU) Lsr(params ...any) {
	defer c.instructionHook(m6502.Lsr, params...)()
	c.lsrInternal(params...)
}

func (c *CPU) lsrInternal(params ...any) {
	if hasAccumulatorParam(params...) {
		c.Flags.C = c.A & 1
		c.A >>= 1
		c.setZN(c.A)
		return
	}

	val := c.bus.Memory.ReadAddressModes(false, params...)
	c.Flags.C = val & 1
	val >>= 1
	c.setZN(val)
	c.bus.Memory.WriteAddressModes(val, params...)
}

// Nop - No Operation.
func (c *CPU) Nop() {
	defer c.instructionHook(m6502.Nop)()
}

// Ora - OR with Accumulator.
func (c *CPU) Ora(params ...any) {
	defer c.instructionHook(m6502.Ora, params...)()
	c.oraInternal(params...)
}

func (c *CPU) oraInternal(params ...any) {
	value := c.bus.Memory.ReadAddressModes(true, params...)
	c.A |= value
	c.setZN(c.A)
}

// Pha - Push Accumulator - push A content to stack.
func (c *CPU) Pha() {
	defer c.instructionHook(m6502.Pha)()
	c.push(c.A)
}

// Php - Push Processor Status - push status flags to stack.
func (c *CPU) Php() {
	defer c.instructionHook(m6502.Php)()
	c.phpInternal()
}

func (c *CPU) phpInternal() {
	f := c.GetFlags()
	f |= 0b0001_0000 // break is set to 1
	c.push(f)
}

// Pla - Pull Accumulator - pull A content from stack.
func (c *CPU) Pla() {
	defer c.instructionHook(m6502.Pla)()

	c.A = c.Pop()
	c.setZN(c.A)
}

// Plp - Pull Processor Status - pull status flags from stack.
func (c *CPU) Plp() {
	defer c.instructionHook(m6502.Plp)()

	f := c.Pop()
	f &= 0b1110_1111 // break flag is ignored
	f |= 0b0010_0000 // unused flag is set
	c.setFlags(f)
}

// Rol - Rotate Left.
func (c *CPU) Rol(params ...any) {
	defer c.instructionHook(m6502.Rol, params...)()
	c.rolInternal(params...)
}

func (c *CPU) rolInternal(params ...any) {
	cFlag := c.Flags.C
	if hasAccumulatorParam(params...) {
		c.Flags.C = (c.A >> 7) & 1
		c.A = (c.A << 1) | cFlag
		c.setZN(c.A)
		return
	}

	val := c.bus.Memory.ReadAddressModes(false, params...)
	c.Flags.C = (val >> 7) & 1
	val = (val << 1) | cFlag
	c.setZN(val)
	c.bus.Memory.WriteAddressModes(val, params...)
}

// Ror - Rotate Right.
func (c *CPU) Ror(params ...any) {
	defer c.instructionHook(m6502.Ror, params...)()
	c.rorInternal(params...)
}

func (c *CPU) rorInternal(params ...any) {
	cFlag := c.Flags.C
	if hasAccumulatorParam(params...) {
		c.Flags.C = c.A & 1
		c.A = (c.A >> 1) | (cFlag << 7)
		c.setZN(c.A)
		return
	}

	val := c.bus.Memory.ReadAddressModes(false, params...)
	c.Flags.C = val & 1
	val = (val >> 1) | (cFlag << 7)
	c.setZN(val)
	c.bus.Memory.WriteAddressModes(val, params...)
}

// Rti - Return from Interrupt.
func (c *CPU) Rti() {
	defer c.instructionHook(m6502.Rti)()

	b := c.Pop()
	b &= 0b1110_1111 // break flag is ignored
	b |= 0b0010_0000 // unused flag is set
	c.setFlags(b)
	c.PC = c.Pop16()

	// lock is already taken
	c.irqRunning = false
	c.nmiRunning = false
}

// Rts - return from subroutine.
func (c *CPU) Rts() {
	defer c.instructionHook(m6502.Rts)()

	if c.emulator {
		c.PC = c.Pop16() + 1
	}
}

// Sbc - subtract with Carry.
func (c *CPU) Sbc(params ...any) {
	defer c.instructionHook(m6502.Sbc, params...)()
	c.sbcInternal(params...)
}

func (c *CPU) sbcInternal(params ...any) {
	a := c.A
	value := c.bus.Memory.ReadAddressModes(true, params...)
	sub := int(c.A) - int(value) - (1 - int(c.Flags.C))
	c.A = uint8(sub)
	c.setZN(c.A)

	if sub >= 0 {
		c.Flags.C = 1
	} else {
		c.Flags.C = 0
	}
	c.setV((a^value)&0x80 != 0 && (a^c.A)&0x80 != 0)
}

// Sec - Set Carry Flag.
func (c *CPU) Sec() {
	defer c.instructionHook(m6502.Sec)()
	c.Flags.C = 1
}

// Sed - Set Decimal Flag.
func (c *CPU) Sed() {
	defer c.instructionHook(m6502.Sed)()
	c.Flags.D = 1
}

// Sei - Set Interrupt Disable.
func (c *CPU) Sei() {
	defer c.instructionHook(m6502.Sei)()
	c.Flags.I = 1
}

// Sta - Store Accumulator - store content of A at address Addr and
// add an optional register to the address.
func (c *CPU) Sta(params ...any) {
	defer c.instructionHook(m6502.Sta, params...)()
	c.bus.Memory.WriteAddressModes(c.A, params...)
}

// Stx - Store X Register - store content of X at address Addr and
// add an optional register to the address.
func (c *CPU) Stx(params ...any) {
	defer c.instructionHook(m6502.Stx, params...)()
	c.bus.Memory.WriteAddressModes(c.X, params...)
}

// Sty - Store Y Register - store content of Y at address Addr and
// add an optional register to the address.
func (c *CPU) Sty(params ...any) {
	defer c.instructionHook(m6502.Sty, params...)()
	c.bus.Memory.WriteAddressModes(c.Y, params...)
}

// Tax - Transfer Accumulator to X.
func (c *CPU) Tax() {
	defer c.instructionHook(m6502.Tax)()

	c.X = c.A
	c.setZN(c.X)
}

// Tay - Transfer Accumulator to Y.
func (c *CPU) Tay() {
	defer c.instructionHook(m6502.Tay)()

	c.Y = c.A
	c.setZN(c.Y)
}

// Tsx - Transfer Stack Pointer to X.
func (c *CPU) Tsx() {
	defer c.instructionHook(m6502.Tsx)()

	c.X = c.SP
	c.setZN(c.X)
}

// Txa - Transfer X to Accumulator.
func (c *CPU) Txa() {
	defer c.instructionHook(m6502.Txa)()

	c.A = c.X
	c.setZN(c.A)
}

// Txs - Transfer X to Stack Pointer.
func (c *CPU) Txs() {
	defer c.instructionHook(m6502.Txs)()
	c.SP = c.X
}

// Tya - Transfer Y to Accumulator.
func (c *CPU) Tya() {
	defer c.instructionHook(m6502.Tya)()

	c.A = c.Y
	c.setZN(c.A)
}

// Dcp ...
func (c *CPU) Dcp(params ...any) {
	defer c.instructionHook(m6502.Dcp, params...)()

	c.decInternal(params...)
	val := c.bus.Memory.ReadAddressModes(false, params...)
	c.compare(c.A, val)
}

// Isc ...
func (c *CPU) Isc(params ...any) {
	defer c.instructionHook(m6502.Isc, params...)()

	c.incInternal(params...)
	c.sbcInternal(params...)
}

// Lax ...
func (c *CPU) Lax(params ...any) {
	defer c.instructionHook(m6502.Lax, params...)()

	c.A = c.bus.Memory.ReadAddressModes(false, params...)
	c.X = c.A
	c.setZN(c.A)
}

// NopUnofficial ...
func (c *CPU) NopUnofficial(params ...any) {
	defer c.instructionHook(m6502.NopUnofficial, params...)()

	if len(params) > 0 {
		c.bus.Memory.ReadAddressModes(false, params...)
	}
}

// Rla ...
func (c *CPU) Rla(params ...any) {
	defer c.instructionHook(m6502.Rla, params...)()

	c.rolInternal(params...)
	c.andInternal(params...)
}

// Rra ...
func (c *CPU) Rra(params ...any) {
	defer c.instructionHook(m6502.Rra, params...)()

	c.rorInternal(params...)
	c.adcInternal(params...)
}

// Sax ...
func (c *CPU) Sax(params ...any) {
	defer c.instructionHook(m6502.Sax, params...)()

	val := c.A & c.X
	c.bus.Memory.WriteAddressModes(val, params...)
}

// SbcUnofficial ...
func (c *CPU) SbcUnofficial(params ...any) {
	defer c.instructionHook(m6502.SbcUnofficial, params...)()
	c.sbcInternal(params...)
}

// Slo ...
func (c *CPU) Slo(params ...any) {
	defer c.instructionHook(m6502.Slo, params...)()

	c.aslInternal(params...)
	c.oraInternal(params...)
}

// Sre ...
func (c *CPU) Sre(params ...any) {
	defer c.instructionHook(m6502.Sre, params...)()

	c.lsrInternal(params...)
	c.eorInternal(params...)
}
