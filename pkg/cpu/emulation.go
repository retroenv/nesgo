//go:build !nesgo

package cpu

import (
	"fmt"
	"math"

	. "github.com/retroenv/retrogolib/nes/addressing"
)

// Adc - Add with Carry.
func (c *CPU) Adc(params ...any) {
	defer c.instructionHook(adc, params...)()

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
	defer c.instructionHook(and, params...)()
	c.andInternal(params...)
}

func (c *CPU) andInternal(params ...any) {
	value := c.bus.Memory.ReadAddressModes(true, params...)
	c.A &= value
	c.setZN(c.A)
}

// Asl - Arithmetic Shift Left.
func (c *CPU) Asl(params ...any) {
	defer c.instructionHook(asl, params...)()
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
	defer c.instructionHook(bcc)()
	return c.Flags.C == 0
}

// BccInternal - Branch if Carry Clear.
func (c *CPU) BccInternal(params ...any) {
	defer c.instructionHook(bcc, params...)()
	c.branch(c.Flags.C == 0, params[0])
}

// Bcs - Branch if Carry Set - returns whether the carry flag is set.
func (c *CPU) Bcs() bool {
	defer c.instructionHook(bcs)()
	return c.Flags.C != 0
}

// BcsInternal - Branch if Carry Set.
func (c *CPU) BcsInternal(params ...any) {
	defer c.instructionHook(bcs, params...)()
	c.branch(c.Flags.C != 0, params[0])
}

// Beq - Branch if Equal - returns whether the zero flag is set.
func (c *CPU) Beq() bool {
	defer c.instructionHook(beq)()
	return c.Flags.Z != 0
}

// BeqInternal - Branch if Equal.
func (c *CPU) BeqInternal(params ...any) {
	defer c.instructionHook(beq, params...)()
	c.branch(c.Flags.Z != 0, params[0])
}

// Bit - Bit Test - set the Z flag by ANDing A with given address content.
func (c *CPU) Bit(params ...any) {
	defer c.instructionHook(bit, params...)()

	value := c.bus.Memory.ReadAbsolute(params[0], nil)
	c.setV((value>>6)&1 == 1)
	c.setZ(value & c.A)
	c.setN(value)
}

// Bmi - Branch if Minus - returns whether the negative flag is set.
func (c *CPU) Bmi() bool {
	defer c.instructionHook(bmi)()
	return c.Flags.N != 0
}

// BmiInternal - Branch if Minus.
func (c *CPU) BmiInternal(params ...any) {
	defer c.instructionHook(bmi, params...)()
	c.branch(c.Flags.N != 0, params[0])
}

// Bne - Branch if Not Equal - returns whether the zero flag is clear.
func (c *CPU) Bne() bool {
	defer c.instructionHook(bne)()
	return c.Flags.Z == 0
}

// BneInternal - Branch if Not Equal.
func (c *CPU) BneInternal(params ...any) {
	defer c.instructionHook(bne, params...)()
	c.branch(c.Flags.Z == 0, params[0])
}

// Bpl - Branch if Positive - returns whether the negative flag is clear.
func (c *CPU) Bpl() bool {
	defer c.instructionHook(bpl)()
	return c.Flags.N == 0
}

// BplInternal - Branch if Positive.
func (c *CPU) BplInternal(params ...any) {
	defer c.instructionHook(bpl, params...)()
	c.branch(c.Flags.N == 0, params[0])
}

// Brk - Force Interrupt.
func (c *CPU) Brk() {
	unlock := c.instructionHook(brk)
	unlock()
	c.irq()
}

// Bvc - Branch if Overflow Clear - returns whether the overflow flag is clear.
func (c *CPU) Bvc() bool {
	defer c.instructionHook(bvc)()
	return c.Flags.V == 0
}

// BvcInternal - Branch if Overflow Clear.
func (c *CPU) BvcInternal(params ...any) {
	defer c.instructionHook(bvc, params...)()
	c.branch(c.Flags.V == 0, params[0])
}

// Bvs - Branch if Overflow Set - returns whether the overflow flag is set.
func (c *CPU) Bvs() bool {
	defer c.instructionHook(bvs)()
	return c.Flags.V != 0
}

// BvsInternal - Branch if Overflow Set.
func (c *CPU) BvsInternal(params ...any) {
	defer c.instructionHook(bvs, params...)()
	c.branch(c.Flags.V != 0, params[0])
}

// Clc - Clear Carry Flag.
func (c *CPU) Clc() {
	defer c.instructionHook(clc)()
	c.Flags.C = 0
}

// Cld - Clear Decimal Mode.
func (c *CPU) Cld() {
	defer c.instructionHook(cld)()
	c.Flags.D = 0
}

// Cli - Clear Interrupt Disable.
func (c *CPU) Cli() {
	defer c.instructionHook(cli)()
	c.Flags.I = 0
}

// Clv - Clear Overflow Flag.
func (c *CPU) Clv() {
	defer c.instructionHook(clv)()
	c.Flags.V = 0
}

// Cmp - Compare - compares the contents of A.
func (c *CPU) Cmp(params ...any) {
	defer c.instructionHook(cmp, params...)()

	val := c.bus.Memory.ReadAddressModes(true, params...)
	c.compare(c.A, val)
}

// Cpx - Compare X Register - compares the contents of X.
func (c *CPU) Cpx(params ...any) {
	defer c.instructionHook(cpx, params...)()

	val := c.bus.Memory.ReadAddressModes(true, params[0])
	c.compare(c.X, val)
}

// Cpy - Compare Y Register - compares the contents of Y.
func (c *CPU) Cpy(params ...any) {
	defer c.instructionHook(cpy, params...)()

	val := c.bus.Memory.ReadAddressModes(true, params[0])
	c.compare(c.Y, val)
}

// Dec - Decrement memory.
func (c *CPU) Dec(params ...any) {
	defer c.instructionHook(dec, params...)()
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
	defer c.instructionHook(dex)()

	c.X--
	c.setZN(c.X)
}

// Dey - Decrement Y Register.
func (c *CPU) Dey() {
	defer c.instructionHook(dey)()

	c.Y--
	c.setZN(c.Y)
}

// Eor - Exclusive OR - XOR.
func (c *CPU) Eor(params ...any) {
	defer c.instructionHook(eor, params...)()
	c.eorInternal(params...)
}

func (c *CPU) eorInternal(params ...any) {
	value := c.bus.Memory.ReadAddressModes(true, params...)
	c.A ^= value
	c.setZN(c.A)
}

// Inc - Increments memory.
func (c *CPU) Inc(params ...any) {
	defer c.instructionHook(inc, params...)()
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
	defer c.instructionHook(inx)()

	c.X++
	c.setZN(c.X)
}

// Iny - Increment Y Register.
func (c *CPU) Iny() {
	defer c.instructionHook(iny)()

	c.Y++
	c.setZN(c.Y)
}

// Jmp - jump to address.
func (c *CPU) Jmp(params ...any) {
	defer c.instructionHook(jmp, params...)()

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
	defer c.instructionHook(jsr, params...)()

	c.Push16(c.PC + 2)

	addr := params[0].(Absolute)
	c.PC = uint16(addr)
}

// Lda - Load Accumulator - load a byte into A.
func (c *CPU) Lda(params ...any) {
	defer c.instructionHook(lda, params...)()

	c.A = c.bus.Memory.ReadAddressModes(true, params...)
	c.setZN(c.A)
}

// Ldx - Load X Register - load a byte into X.
func (c *CPU) Ldx(params ...any) {
	defer c.instructionHook(ldx, params...)()

	c.X = c.bus.Memory.ReadAddressModes(true, params...)
	c.setZN(c.X)
}

// Ldy - Load Y Register - load a byte into Y.
func (c *CPU) Ldy(params ...any) {
	defer c.instructionHook(ldy, params...)()

	c.Y = c.bus.Memory.ReadAddressModes(true, params...)
	c.setZN(c.Y)
}

// Lsr - Logical Shift Right.
func (c *CPU) Lsr(params ...any) {
	defer c.instructionHook(lsr, params...)()
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
	defer c.instructionHook(nop)()
}

// Ora - OR with Accumulator.
func (c *CPU) Ora(params ...any) {
	defer c.instructionHook(ora, params...)()
	c.oraInternal(params...)
}

func (c *CPU) oraInternal(params ...any) {
	value := c.bus.Memory.ReadAddressModes(true, params...)
	c.A |= value
	c.setZN(c.A)
}

// Pha - Push Accumulator - push A content to stack.
func (c *CPU) Pha() {
	defer c.instructionHook(pha)()
	c.push(c.A)
}

// Php - Push Processor Status - push status flags to stack.
func (c *CPU) Php() {
	defer c.instructionHook(php)()
	c.phpInternal()
}

func (c *CPU) phpInternal() {
	f := c.GetFlags()
	f |= 0b0001_0000 // break is set to 1
	c.push(f)
}

// Pla - Pull Accumulator - pull A content from stack.
func (c *CPU) Pla() {
	defer c.instructionHook(pla)()

	c.A = c.Pop()
	c.setZN(c.A)
}

// Plp - Pull Processor Status - pull status flags from stack.
func (c *CPU) Plp() {
	defer c.instructionHook(plp)()

	f := c.Pop()
	f &= 0b1110_1111 // break flag is ignored
	f |= 0b0010_0000 // unused flag is set
	c.setFlags(f)
}

// Rol - Rotate Left.
func (c *CPU) Rol(params ...any) {
	defer c.instructionHook(rol, params...)()
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
	defer c.instructionHook(ror, params...)()
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
	defer c.instructionHook(rti)()

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
	defer c.instructionHook(rts)()

	if c.emulator {
		c.PC = c.Pop16() + 1
	}
}

// Sbc - subtract with Carry.
func (c *CPU) Sbc(params ...any) {
	defer c.instructionHook(sbc, params...)()
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
	defer c.instructionHook(sec)()
	c.Flags.C = 1
}

// Sed - Set Decimal Flag.
func (c *CPU) Sed() {
	defer c.instructionHook(sed)()
	c.Flags.D = 1
}

// Sei - Set Interrupt Disable.
func (c *CPU) Sei() {
	defer c.instructionHook(sei)()
	c.Flags.I = 1
}

// Sta - Store Accumulator - store content of A at address Addr and
// add an optional register to the address.
func (c *CPU) Sta(params ...any) {
	defer c.instructionHook(sta, params...)()
	c.bus.Memory.WriteAddressModes(c.A, params...)
}

// Stx - Store X Register - store content of X at address Addr and
// add an optional register to the address.
func (c *CPU) Stx(params ...any) {
	defer c.instructionHook(stx, params...)()
	c.bus.Memory.WriteAddressModes(c.X, params...)
}

// Sty - Store Y Register - store content of Y at address Addr and
// add an optional register to the address.
func (c *CPU) Sty(params ...any) {
	defer c.instructionHook(sty, params...)()
	c.bus.Memory.WriteAddressModes(c.Y, params...)
}

// Tax - Transfer Accumulator to X.
func (c *CPU) Tax() {
	defer c.instructionHook(tax)()

	c.X = c.A
	c.setZN(c.X)
}

// Tay - Transfer Accumulator to Y.
func (c *CPU) Tay() {
	defer c.instructionHook(tay)()

	c.Y = c.A
	c.setZN(c.Y)
}

// Tsx - Transfer Stack Pointer to X.
func (c *CPU) Tsx() {
	defer c.instructionHook(tsx)()

	c.X = c.SP
	c.setZN(c.X)
}

// Txa - Transfer X to Accumulator.
func (c *CPU) Txa() {
	defer c.instructionHook(txa)()

	c.A = c.X
	c.setZN(c.A)
}

// Txs - Transfer X to Stack Pointer.
func (c *CPU) Txs() {
	defer c.instructionHook(txs)()
	c.SP = c.X
}

// Tya - Transfer Y to Accumulator.
func (c *CPU) Tya() {
	defer c.instructionHook(tya)()

	c.A = c.Y
	c.setZN(c.A)
}

func (c *CPU) unofficialDcp(params ...any) {
	defer c.instructionHook(unofficialDcp, params...)()

	c.decInternal(params...)
	val := c.bus.Memory.ReadAddressModes(false, params...)
	c.compare(c.A, val)
}

func (c *CPU) unofficialIsb(params ...any) {
	defer c.instructionHook(unofficialIsb, params...)()

	c.incInternal(params...)
	c.sbcInternal(params...)
}

func (c *CPU) unofficialLax(params ...any) {
	defer c.instructionHook(unofficialLax, params...)()

	c.A = c.bus.Memory.ReadAddressModes(false, params...)
	c.X = c.A
	c.setZN(c.A)
}

func (c *CPU) unofficialNop(params ...any) {
	defer c.instructionHook(unofficialNop, params...)()

	if len(params) > 0 {
		c.bus.Memory.ReadAddressModes(false, params...)
	}
}

func (c *CPU) unofficialRla(params ...any) {
	defer c.instructionHook(unofficialRla, params...)()

	c.rolInternal(params...)
	c.andInternal(params...)
}

func (c *CPU) unofficialRra(params ...any) {
	defer c.instructionHook(unofficialRra, params...)()

	c.rorInternal(params...)
	c.adcInternal(params...)
}

func (c *CPU) unofficialSax(params ...any) {
	defer c.instructionHook(unofficialSax, params...)()

	val := c.A & c.X
	c.bus.Memory.WriteAddressModes(val, params...)
}

func (c *CPU) unofficialSbc(params ...any) {
	defer c.instructionHook(unofficialSbc, params...)()
	c.sbcInternal(params...)
}

func (c *CPU) unofficialSlo(params ...any) {
	defer c.instructionHook(unofficialSlo, params...)()

	c.aslInternal(params...)
	c.oraInternal(params...)
}

func (c *CPU) unofficialSre(params ...any) {
	defer c.instructionHook(unofficialSre, params...)()

	c.lsrInternal(params...)
	c.eorInternal(params...)
}
