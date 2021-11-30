//go:build !nesgo
// +build !nesgo

package nes

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
func Adc() {
	timeInstructionExecution()
	panic(notImplemented)
}

// And - AND with accumulator.
func And() {
	timeInstructionExecution()
	panic(notImplemented)
}

// Asl - Arithmetic Shift Left - shift left Accumulator.
func Asl() {
	timeInstructionExecution()
	panic(notImplemented)
}

// AslAddr - Arithmetic Shift Left - of address Addr and add
// an optional register to the address.
func AslAddr(address uint16, reg ...interface{}) {
	timeInstructionExecution()
	panic(notImplemented)
}

// Bcc - Branch if Carry Clear - returns whether the
// carry flag is clear.
func Bcc() bool {
	timeInstructionExecution()
	b := C == 0
	return b
}

// Bcs - Branch if Carry Set - returns whether the carry flag is set.
func Bcs() bool {
	timeInstructionExecution()
	b := C != 0
	return b
}

// Beq - Branch if Equal - returns whether the zero flag is set.
func Beq() bool {
	timeInstructionExecution()
	b := Z != 0
	return b
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
	b := N != 0
	return b
}

// Bne - Branch if Not Equal - returns whether the zero flag is clear.
func Bne() bool {
	timeInstructionExecution()
	b := Z == 0
	return b
}

// Bpl - Branch if Positive - returns whether the negative flag is clear.
func Bpl() bool {
	timeInstructionExecution()
	b := N == 0
	return b
}

// Brk - Force Interrupt.
func Brk() {
	timeInstructionExecution()
	panic(notImplemented)
}

// Bvc - Branch if Overflow Clear - returns whether the overflow flag is clear.
func Bvc() bool {
	timeInstructionExecution()
	b := V == 0
	return b
}

// Bvs - Branch if Overflow Set - returns whether the overflow flag is set.
func Bvs() bool {
	timeInstructionExecution()
	b := V != 0
	return b
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

// Cmp - Compare - compares the contents of A with given address content.
func Cmp(address uint16) {
	timeInstructionExecution()
	panic(notImplemented)
}

// Cpx - Compare X Register - compares the contents of X with given immediate.
func Cpx(immediate byte) {
	timeInstructionExecution()
	compare(cpu.X, immediate)
}

// CpxAddr - Compare X Register - compares the contents of X with given
// address content.
func CpxAddr(address interface{}) {
	timeInstructionExecution()
	value := readMemoryAbsolute(address)
	compare(cpu.X, value)
}

// Cpy - Compare Y Register - compares the contents of Y with given immediate.
func Cpy(immediate byte) {
	timeInstructionExecution()
	compare(cpu.Y, immediate)
}

// CpyAddr - Compare Y Register - compares the contents of Y with given
// address content.
func CpyAddr(address interface{}) {
	timeInstructionExecution()
	value := readMemoryAbsolute(address)
	compare(cpu.Y, value)
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

// Eor - Exclusive OR - XOR the Accumulator.
func Eor(immediate byte) {
	timeInstructionExecution()
	panic(notImplemented)
}

// EorAddr - Exclusive OR - of address Addr and
// add an optional register to the address.
func EorAddr(address uint16, reg ...interface{}) {
	timeInstructionExecution()
	panic(notImplemented)
}

// EorInd - Exclusive OR - of address located in Addr and
// add an optional register to the address.
func EorInd(address uint8, reg ...interface{}) {
	timeInstructionExecution()
	panic(notImplemented)
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
func Lda(i uint8) {
	timeInstructionExecution()
	cpu.A = i
	setZN(cpu.A)
}

// LdaAddr - loading Accumulator from address Addr and
// add an optional register to the address.
func LdaAddr(address interface{}, reg ...interface{}) {
	timeInstructionExecution()
	cpu.A = readMemoryAbsolute(address, reg...)
	setZN(cpu.A)
}

// LdaInd - loading Accumulator of address located in Addr and
// add an optional register to the address.
// Indirect addressing.
func LdaInd(address uint16, reg ...interface{}) {
	timeInstructionExecution()
	panic(notImplemented)
}

// Ldx - Load X Register - load a byte into X.
func Ldx(i uint8) {
	timeInstructionExecution()
	cpu.X = i
	setZN(cpu.X)
}

// LdxAddr - loading X from address Addr and
// add an optional register to the address.
func LdxAddr(address interface{}, reg ...interface{}) {
	timeInstructionExecution()
	cpu.X = readMemoryAbsolute(address, reg...)
	setZN(cpu.X)
}

// Ldy - Load Y Register - load a byte into Y.
func Ldy(i uint8) {
	timeInstructionExecution()
	cpu.Y = i
	setZN(cpu.Y)
}

// LdyAddr - loading Y from address Addr and
// add an optional register to the address.
func LdyAddr(address interface{}, reg ...interface{}) {
	timeInstructionExecution()
	cpu.Y = readMemoryAbsolute(address, reg...)
	setZN(cpu.Y)
}

// Lsr - Logical Shift Right - shift right.
func Lsr(reg ...interface{}) {
	timeInstructionExecution()
	panic(notImplemented)
}

// Nop - No Operation.
func Nop() {
	timeInstructionExecution()
}

// Ora - OR with Accumulator.
func Ora() {
	timeInstructionExecution()
	panic(notImplemented)
}

// Pha - Push Accumulator - push A content to stack.
func Pha() {
	timeInstructionExecution()
	panic(notImplemented)
}

// Php - Push Processor Status - push status flags to stack.
func Php() {
	timeInstructionExecution()
	panic(notImplemented)
}

// Pla - Pull Accumulator - pull A content from stack.
func Pla() {
	timeInstructionExecution()
	panic(notImplemented)
}

// Plp - Pull Processor Status - pull status flags from stack.
func Plp() {
	timeInstructionExecution()
	panic(notImplemented)
}

// Rol - Rotate Left - rotate Accumulator left.
func Rol() {
	timeInstructionExecution()
	panic(notImplemented)
}

// Ror - Rotate Right - rotate Accumulator right.
func Ror() {
	timeInstructionExecution()
	panic(notImplemented)
}

// Rti - Return from Interrupt.
func Rti() {
	timeInstructionExecution()
	panic(notImplemented)
}

// Sbc - subtract with Carry
func Sbc() {
	timeInstructionExecution()
	panic(notImplemented)
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
	timeInstructionExecution()
	I = 1
}

// Sta - Store Accumulator - store content of A at address Addr and
// add an optional register to the address.
// Zero Page/Absolute addressing.
func Sta(address interface{}, reg ...interface{}) {
	timeInstructionExecution()
	writeMemoryAbsolute(address, cpu.A, reg...)
}

// StaInd - Store Accumulator - store content of A at address Addr and
// add an optional register to the address.
// Indirect addressing.
func StaInd(address uint16, reg ...interface{}) {
	timeInstructionExecution()
	writeMemoryIndirect(address, cpu.A, reg...)
}

// Stx - Store X Register - store content of X at address Addr and
// add an optional register to the address.
func Stx(address interface{}, reg ...interface{}) {
	timeInstructionExecution()
	writeMemoryAbsolute(address, cpu.X, reg...)
}

// Sty - Store Y Register - store content of Y at address Addr and
// add an optional register to the address.
func Sty(address interface{}, reg ...interface{}) {
	timeInstructionExecution()
	writeMemoryAbsolute(address, cpu.Y, reg...)
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
