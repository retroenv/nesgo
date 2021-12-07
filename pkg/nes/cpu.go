//go:build !nesgo
// +build !nesgo

package nes

import "math"

// CPU registers that can be used as parameter for instructions that support
// absolute or indirect indexing using X or Y register.
var (
	X *uint8 // x register
	Y *uint8 // y register
)

var notImplemented = "instruction is not implemented yet"

// Adc - Add with Carry.
func Adc(param interface{}, reg ...interface{}) {
	timeInstructionExecution()
	value := readMemoryAddressModes(param, reg...)
	sum := int(cpu.A) + int(cpu.Flags.C) + int(value)
	cpu.A = uint8(sum)
	cpu.setZN(cpu.A)

	if sum > math.MaxUint8 {
		cpu.Flags.C = 1
	} else {
		cpu.Flags.C = 0
	}

	// TODO support decimal mode
}

// And - AND with accumulator.
func And(param interface{}, reg ...interface{}) {
	timeInstructionExecution()
	value := readMemoryAddressModes(param, reg...)
	cpu.A &= value
	cpu.setZN(cpu.A)
}

// Asl - Arithmetic Shift Left.
func Asl(param ...interface{}) {
	timeInstructionExecution()

	if param == nil { // A implied
		cpu.Flags.C = (cpu.A >> 7) & 1
		cpu.A <<= 1
		cpu.setZN(cpu.A)
		return
	}

	val := readMemoryAddressModes(param)
	cpu.Flags.C = (val >> 7) & 1
	val <<= 1
	cpu.setZN(val)
	writeMemoryAddressModes(param, val)
}

// Bcc - Branch if Carry Clear - returns whether the
// carry flag is clear.
func Bcc() bool {
	timeInstructionExecution()
	return cpu.Flags.C == 0
}

// Bcs - Branch if Carry Set - returns whether the carry flag is set.
func Bcs() bool {
	timeInstructionExecution()
	return cpu.Flags.C != 0
}

// Beq - Branch if Equal - returns whether the zero flag is set.
func Beq() bool {
	timeInstructionExecution()
	return cpu.Flags.Z != 0
}

// Bit - Bit Test - set the Z flag by ANDing A with given address content.
func Bit(address uint16) {
	timeInstructionExecution()
	value := readMemoryAbsolute(address)
	cpu.Flags.V = (value >> 6) & 1
	cpu.setZ(value & cpu.A)
	cpu.setN(value)
}

// Bmi - Branch if Minus - returns whether the negative flag is set.
func Bmi() bool {
	timeInstructionExecution()
	return cpu.Flags.N != 0
}

// Bne - Branch if Not Equal - returns whether the zero flag is clear.
func Bne() bool {
	timeInstructionExecution()
	return cpu.Flags.Z == 0
}

// Bpl - Branch if Positive - returns whether the negative flag is clear.
func Bpl() bool {
	timeInstructionExecution()
	return cpu.Flags.N == 0
}

// Brk - Force Interrupt.
func Brk() {
	timeInstructionExecution()
	panic(notImplemented) // TODO: implement
}

// Bvc - Branch if Overflow Clear - returns whether the overflow flag is clear.
func Bvc() bool {
	timeInstructionExecution()
	return cpu.Flags.V == 0
}

// Bvs - Branch if Overflow Set - returns whether the overflow flag is set.
func Bvs() bool {
	timeInstructionExecution()
	return cpu.Flags.V != 0
}

// Clc - Clear Carry Flag.
func Clc() {
	timeInstructionExecution()
	cpu.Flags.C = 0
}

// Cld - Clear Decimal Mode.
func Cld() {
	timeInstructionExecution()
	cpu.Flags.D = 0
}

// Cli - Clear Interrupt Disable.
func Cli() {
	timeInstructionExecution()
	cpu.Flags.I = 0
}

// Clv - Clear Overflow Flag.
func Clv() {
	timeInstructionExecution()
	cpu.Flags.V = 0
}

// Cmp - Compare - compares the contents of A.
func Cmp(param interface{}) {
	timeInstructionExecution()
	val := readMemoryAddressModes(param)
	cpu.compare(cpu.A, val)
}

// Cpx - Compare X Register - compares the contents of X.
func Cpx(param interface{}) {
	timeInstructionExecution()
	val := readMemoryAddressModes(param)
	cpu.compare(cpu.X, val)
}

// Cpy - Compare Y Register - compares the contents of Y.
func Cpy(param interface{}) {
	timeInstructionExecution()
	val := readMemoryAddressModes(param)
	cpu.compare(cpu.Y, val)
}

// Dex - Decrement X Register.
func Dex() {
	timeInstructionExecution()
	cpu.X--
	cpu.setZN(cpu.X)
}

// Dey - Decrement Y Register.
func Dey() {
	timeInstructionExecution()
	cpu.Y--
	cpu.setZN(cpu.Y)
}

// Eor - Exclusive OR - XOR.
func Eor(param interface{}, reg ...interface{}) {
	timeInstructionExecution()
	value := readMemoryAddressModes(param, reg...)
	cpu.A ^= value
	cpu.setZN(cpu.A)
}

// Inx - Increment X Register.
func Inx() {
	timeInstructionExecution()
	cpu.X++
	cpu.setZN(cpu.X)
}

// Iny - Increment Y Register.
func Iny() {
	timeInstructionExecution()
	cpu.Y++
	cpu.setZN(cpu.Y)
}

// Lda - Load Accumulator - load a byte into A.
func Lda(param interface{}, reg ...interface{}) {
	timeInstructionExecution()
	cpu.A = readMemoryAddressModes(param, reg...)
	cpu.setZN(cpu.A)
}

// Ldx - Load X Register - load a byte into X.
func Ldx(param interface{}, reg ...interface{}) {
	timeInstructionExecution()
	cpu.X = readMemoryAddressModes(param, reg...)
	cpu.setZN(cpu.X)
}

// Ldy - Load Y Register - load a byte into Y.
func Ldy(param interface{}, reg ...interface{}) {
	timeInstructionExecution()
	cpu.Y = readMemoryAddressModes(param, reg...)
	cpu.setZN(cpu.Y)
}

// Lsr - Logical Shift Right.
func Lsr(param ...interface{}) {
	timeInstructionExecution()

	if param == nil { // A implied
		cpu.Flags.C = cpu.A & 1
		cpu.A >>= 1
		cpu.setZN(cpu.A)
		return
	}

	val := readMemoryAddressModes(param)
	cpu.Flags.C = val & 1
	val >>= 1
	cpu.setZN(val)
	writeMemoryAddressModes(param, val)
}

// Nop - No Operation.
func Nop() {
	timeInstructionExecution()
}

// Ora - OR with Accumulator.
func Ora(param interface{}, reg ...interface{}) {
	timeInstructionExecution()
	value := readMemoryAddressModes(param, reg...)
	cpu.A |= value
	cpu.setZN(cpu.A)
}

// Pha - Push Accumulator - push A content to stack.
func Pha() {
	timeInstructionExecution()
	push(cpu.A)
}

// Php - Push Processor Status - push status flags to stack.
func Php() {
	timeInstructionExecution()
	f := cpu.flags()
	f |= 0b11000 // bit 4 and 5 are set to 1
	push(f)
}

// Pla - Pull Accumulator - pull A content from stack.
func Pla() {
	timeInstructionExecution()
	cpu.A = pop()
	cpu.setZN(cpu.A)
}

// Plp - Pull Processor Status - pull status flags from stack.
func Plp() {
	timeInstructionExecution()
	f := pop()
	f &^= 0b11000 // bit 4 and 5 are cleared
	cpu.setFlags(f)
}

// Rol - Rotate Left.
func Rol(param ...interface{}) {
	timeInstructionExecution()

	c := cpu.Flags.C
	if param == nil { // A implied
		cpu.Flags.C = (cpu.A >> 7) & 1
		cpu.A = (cpu.A << 1) | c
		cpu.setZN(cpu.A)
		return
	}

	val := readMemoryAddressModes(param)
	cpu.Flags.C = (val >> 7) & 1
	val = (val << 1) | c
	cpu.setZN(val)
	writeMemoryAddressModes(param, val)
}

// Ror - Rotate Right.
func Ror(param ...interface{}) {
	timeInstructionExecution()

	c := cpu.Flags.C
	if param == nil { // A implied
		cpu.Flags.C = cpu.A & 1
		cpu.A = (cpu.A >> 1) | (c << 7)
		cpu.setZN(cpu.A)
		return
	}

	val := readMemoryAddressModes(param)
	cpu.Flags.C = val & 1
	val = (val >> 1) | (c << 7)
	cpu.setZN(val)
	writeMemoryAddressModes(param, val)
}

// Rti - Return from Interrupt.
func Rti() {
	timeInstructionExecution()
}

// Sbc - subtract with Carry.
func Sbc(param interface{}, reg ...interface{}) {
	timeInstructionExecution()

	value := readMemoryAddressModes(param, reg...)
	sub := int(cpu.A) - int(value) - (1 - int(cpu.Flags.C))
	cpu.A = uint8(sub)
	cpu.setZN(cpu.A)

	if sub >= 0 {
		cpu.Flags.C = 1
	} else {
		cpu.Flags.C = 0
	}

	// TODO support decimal mode
}

// Sec - Set Carry Flag.
func Sec() {
	timeInstructionExecution()
	cpu.Flags.C = 1
}

// Sed - Set Decimal Flag.
func Sed() {
	timeInstructionExecution()
	cpu.Flags.D = 1
}

// Sei - Set Interrupt Disable.
func Sei() {
	timeInstructionExecution()
	cpu.Flags.I = 1
}

// Sta - Store Accumulator - store content of A at address Addr and
// add an optional register to the address.
func Sta(param interface{}, reg ...interface{}) {
	timeInstructionExecution()
	writeMemoryAddressModes(param, cpu.A, reg...)
}

// Stx - Store X Register - store content of X at address Addr and
// add an optional register to the address.
func Stx(param interface{}, reg ...interface{}) {
	timeInstructionExecution()
	writeMemoryAddressModes(param, cpu.X, reg...)
}

// Sty - Store Y Register - store content of Y at address Addr and
// add an optional register to the address.
func Sty(param interface{}, reg ...interface{}) {
	timeInstructionExecution()
	writeMemoryAddressModes(param, cpu.Y, reg...)
}

// Tax - Transfer Accumulator to X.
func Tax() {
	timeInstructionExecution()
	cpu.X = cpu.A
	cpu.setZN(cpu.X)
}

// Tay - Transfer Accumulator to Y.
func Tay() {
	timeInstructionExecution()
	cpu.Y = cpu.A
	cpu.setZN(cpu.Y)
}

// Tsx - Transfer Stack Pointer to X.
func Tsx() {
	timeInstructionExecution()
	cpu.X = cpu.SP
	cpu.setZN(cpu.X)
}

// Txa - Transfer X to Accumulator.
func Txa() {
	timeInstructionExecution()
	cpu.A = cpu.X
	cpu.setZN(cpu.A)
}

// Txs - Transfer X to Stack Pointer.
func Txs() {
	timeInstructionExecution()
	cpu.SP = cpu.X
}

// Tya - Transfer Y to Accumulator.
func Tya() {
	timeInstructionExecution()
	cpu.A = cpu.Y
	cpu.setZN(cpu.A)
}
