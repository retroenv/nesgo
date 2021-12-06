//go:build !nesgo
// +build !nesgo

package nes

import "math"

// CPU register and flags
var (
	SP uint8 // stack pointer

	A *uint8 // accumulator
	X *uint8 // x register
	Y *uint8 // y register

	C uint8 // carry flag
	Z uint8 // zero flag
	I uint8 // interrupt disable flag
	D uint8 // decimal mode flag
	B uint8 // break command flag
	U uint8 // unused flag
	V uint8 // overflow flag
	N uint8 // negative flag
)

var notImplemented = "instruction is not implemented yet"

// Adc - Add with Carry.
func Adc(param interface{}, reg ...interface{}) {
	timeInstructionExecution()
	value := readMemoryAddressModes(param, reg...)
	sum := int(cpu.A) + int(C) + int(value)
	cpu.A = uint8(sum)
	setZN(cpu.A)

	if sum > math.MaxUint8 {
		C = 1
	} else {
		C = 0
	}

	// TODO support decimal mode
}

// And - AND with accumulator.
func And(param interface{}, reg ...interface{}) {
	timeInstructionExecution()
	value := readMemoryAddressModes(param, reg...)
	cpu.A &= value
	setZN(cpu.A)
}

// Asl - Arithmetic Shift Left - shift left Accumulator.
func Asl() {
	timeInstructionExecution()
	panic(notImplemented) // TODO: implement
}

// Bcc - Branch if Carry Clear - returns whether the
// carry flag is clear.
func Bcc() bool {
	timeInstructionExecution()
	return C == 0
}

// Bcs - Branch if Carry Set - returns whether the carry flag is set.
func Bcs() bool {
	timeInstructionExecution()
	return C != 0
}

// Beq - Branch if Equal - returns whether the zero flag is set.
func Beq() bool {
	timeInstructionExecution()
	return Z != 0
}

// Bit - Bit Test - set the Z flag by ANDing A with given address content.
func Bit(address uint16) {
	timeInstructionExecution()
	value := readMemoryAbsolute(address)
	V = (value >> 6) & 1
	setZ(value & cpu.A)
	setN(value)
}

// Bmi - Branch if Minus - returns whether the negative flag is set.
func Bmi() bool {
	timeInstructionExecution()
	return N != 0
}

// Bne - Branch if Not Equal - returns whether the zero flag is clear.
func Bne() bool {
	timeInstructionExecution()
	return Z == 0
}

// Bpl - Branch if Positive - returns whether the negative flag is clear.
func Bpl() bool {
	timeInstructionExecution()
	return N == 0
}

// Brk - Force Interrupt.
func Brk() {
	timeInstructionExecution()
	panic(notImplemented) // TODO: implement
}

// Bvc - Branch if Overflow Clear - returns whether the overflow flag is clear.
func Bvc() bool {
	timeInstructionExecution()
	return V == 0
}

// Bvs - Branch if Overflow Set - returns whether the overflow flag is set.
func Bvs() bool {
	timeInstructionExecution()
	return V != 0
}

// Clc - Clear Carry Flag.
func Clc() {
	timeInstructionExecution()
	C = 0
}

// Cld - Clear Decimal Mode.
func Cld() {
	timeInstructionExecution()
	D = 0
}

// Cli - Clear Interrupt Disable.
func Cli() {
	timeInstructionExecution()
	I = 0
}

// Clv - Clear Overflow Flag.
func Clv() {
	timeInstructionExecution()
	V = 0
}

// Cmp - Compare - compares the contents of A.
func Cmp(param interface{}) {
	timeInstructionExecution()
	val := readMemoryAddressModes(param)
	compare(cpu.A, val)
}

// Cpx - Compare X Register - compares the contents of X.
func Cpx(param interface{}) {
	timeInstructionExecution()
	val := readMemoryAddressModes(param)
	compare(cpu.X, val)
}

// Cpy - Compare Y Register - compares the contents of Y.
func Cpy(param interface{}) {
	timeInstructionExecution()
	val := readMemoryAddressModes(param)
	compare(cpu.Y, val)
}

// Dex - Decrement X Register.
func Dex() {
	timeInstructionExecution()
	cpu.X--
	setZN(cpu.X)
}

// Dey - Decrement Y Register.
func Dey() {
	timeInstructionExecution()
	cpu.Y--
	setZN(cpu.Y)
}

// Eor - Exclusive OR - XOR.
func Eor(param interface{}, reg ...interface{}) {
	timeInstructionExecution()
	value := readMemoryAddressModes(param, reg...)
	cpu.A ^= value
	setZN(cpu.A)
}

// Inx - Increment X Register.
func Inx() {
	timeInstructionExecution()
	cpu.X++
	setZN(cpu.X)
}

// Iny - Increment Y Register.
func Iny() {
	timeInstructionExecution()
	cpu.Y++
	setZN(cpu.Y)
}

// Lda - Load Accumulator - load a byte into A.
func Lda(param interface{}, reg ...interface{}) {
	timeInstructionExecution()
	cpu.A = readMemoryAddressModes(param, reg...)
	setZN(cpu.A)
}

// Ldx - Load X Register - load a byte into X.
func Ldx(param interface{}, reg ...interface{}) {
	timeInstructionExecution()
	cpu.X = readMemoryAddressModes(param, reg...)
	setZN(cpu.X)
}

// Ldy - Load Y Register - load a byte into Y.
func Ldy(param interface{}, reg ...interface{}) {
	timeInstructionExecution()
	cpu.Y = readMemoryAddressModes(param, reg...)
	setZN(cpu.Y)
}

// Lsr - Logical Shift Right - shift right.
func Lsr(param ...interface{}) {
	timeInstructionExecution()

	if param == nil { // A implied
		C = cpu.A & 1
		cpu.A >>= 1
		setZN(cpu.A)
		return
	}

	val := readMemoryAddressModes(param)
	C = val & 1
	val >>= 1
	setZN(val)
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
	setZN(cpu.A)
}

// Pha - Push Accumulator - push A content to stack.
func Pha() {
	timeInstructionExecution()
	push(cpu.A)
}

// Php - Push Processor Status - push status flags to stack.
func Php() {
	timeInstructionExecution()
	panic(notImplemented) // TODO: implement
}

// Pla - Pull Accumulator - pull A content from stack.
func Pla() {
	timeInstructionExecution()
	cpu.A = pop()
	setZN(cpu.A)
}

// Plp - Pull Processor Status - pull status flags from stack.
func Plp() {
	timeInstructionExecution()
	panic(notImplemented) // TODO: implement
}

// Rol - Rotate Left.
func Rol(param ...interface{}) {
	timeInstructionExecution()

	c := C
	if param == nil { // A implied
		C = (cpu.A >> 7) & 1
		cpu.A = (cpu.A << 1) | c
		setZN(cpu.A)
		return
	}

	val := readMemoryAddressModes(param)
	C = (val >> 7) & 1
	val = (val << 1) | c
	setZN(val)
	writeMemoryAddressModes(param, val)
}

// Ror - Rotate Right.
func Ror(param ...interface{}) {
	timeInstructionExecution()

	c := C
	if param == nil { // A implied
		C = cpu.A & 1
		cpu.A = (cpu.A >> 1) | (c << 7)
		setZN(cpu.A)
		return
	}

	val := readMemoryAddressModes(param)
	C = val & 1
	val = (val >> 1) | (c << 7)
	setZN(val)
	writeMemoryAddressModes(param, val)
}

// Rti - Return from Interrupt.
func Rti() {
	timeInstructionExecution()
}

// Sbc - subtract with Carry.
func Sbc() {
	timeInstructionExecution()
	panic(notImplemented) // TODO: implement
}

// Sec - Set Carry Flag.
func Sec() {
	timeInstructionExecution()
	C = 1
}

// Sed - Set Decimal Flag.
func Sed() {
	timeInstructionExecution()
	D = 1
}

// Sei - Set Interrupt Disable.
func Sei() {
	timeInstructionExecution()
	I = 1
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
	setZN(cpu.X)
}

// Tay - Transfer Accumulator to Y.
func Tay() {
	timeInstructionExecution()
	cpu.Y = cpu.A
	setZN(cpu.Y)
}

// Tsx - Transfer Stack Pointer to X.
func Tsx() {
	timeInstructionExecution()
	cpu.X = SP
	setZN(cpu.X)
}

// Txa - Transfer X to Accumulator.
func Txa() {
	timeInstructionExecution()
	cpu.A = cpu.X
	setZN(cpu.A)
}

// Txs - Transfer X to Stack Pointer.
func Txs() {
	timeInstructionExecution()
	SP = cpu.X
}

// Tya - Transfer Y to Accumulator.
func Tya() {
	timeInstructionExecution()
	cpu.A = cpu.Y
	setZN(cpu.A)
}
