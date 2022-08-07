//go:build !nesgo

package nes

import "github.com/retroenv/nesgo/pkg/cpu"

// CPU registers that can be used as parameter for instructions that support
// absolute or indirect indexing using X or Y register.
var (
	A  *uint8  // accumulator
	X  *uint8  // x register
	Y  *uint8  // y register
	PC *uint16 // program counter
)

var docCPU *cpu.CPU

// All CPU instructions that can be used when writing NES programs in Golang.
// They are aliased to the emulator implementation to allow an easy code
// browsing.
var (
	// Adc - Add with Carry.
	Adc = docCPU.Adc
	// And - AND with accumulator.
	And = docCPU.And
	// Asl - Arithmetic Shift Left.
	Asl = docCPU.Asl
	// Bcc - Branch if Carry Clear - returns whether the
	// carry flag is clear.
	Bcc = docCPU.Bcc
	// Bcs - Branch if Carry Set - returns whether the carry flag is set.
	Bcs = docCPU.Bcs
	// Beq - Branch if Equal - returns whether the zero flag is set.
	Beq = docCPU.Beq
	// Bit - Bit Test - set the Z flag by ANDing A with given address content.
	Bit = docCPU.Bit
	// Bmi - Branch if Minus - returns whether the negative flag is set.
	Bmi = docCPU.Bmi
	// Bne - Branch if Not Equal - returns whether the zero flag is clear.
	Bne = docCPU.Bne
	// Bpl - Branch if Positive - returns whether the negative flag is clear.
	Bpl = docCPU.Bpl
	// Brk - Force Interrupt.
	Brk = docCPU.Brk
	// Bvc - Branch if Overflow Clear - returns whether the overflow flag is clear.
	Bvc = docCPU.Bvc
	// Bvs - Branch if Overflow Set - returns whether the overflow flag is set.
	Bvs = docCPU.Bvs
	// Clc - Clear Carry Flag.
	Clc = docCPU.Clc
	// Cld - Clear Decimal Mode.
	Cld = docCPU.Cld
	// Cli - Clear Interrupt Disable.
	Cli = docCPU.Cli
	// Clv - Clear Overflow Flag.
	Clv = docCPU.Clv
	// Cmp - Compare - compares the contents of A.
	Cmp = docCPU.Cmp
	// Cpx - Compare X Register - compares the contents of X.
	Cpx = docCPU.Cpx
	// Cpy - Compare Y Register - compares the contents of Y.
	Cpy = docCPU.Cpy
	// Dec - Decrement memory.
	Dec = docCPU.Dec
	// Dex - Decrement X Register.
	Dex = docCPU.Dex
	// Dey - Decrement Y Register.
	Dey = docCPU.Dey
	// Eor - Exclusive OR - XOR.
	Eor = docCPU.Eor
	// Inc - Increments memory.
	Inc = docCPU.Inc
	// Inx - Increment X Register.
	Inx = docCPU.Inx
	// Iny - Increment Y Register.
	Iny = docCPU.Iny
	// Lda - Load Accumulator - load a byte into A.
	Lda = docCPU.Lda
	// Ldx - Load X Register - load a byte into X.
	Ldx = docCPU.Ldx
	// Ldy - Load Y Register - load a byte into Y.
	Ldy = docCPU.Ldy
	// Lsr - Logical Shift Right.
	Lsr = docCPU.Lsr
	// Nop - No Operation.
	Nop = docCPU.Nop
	// Ora - OR with Accumulator.
	Ora = docCPU.Ora
	// Pha - Push Accumulator - push A content to stack.
	Pha = docCPU.Pha
	// Php - Push Processor Status - push status flags to stack.
	Php = docCPU.Php
	// Pla - Pull Accumulator - pull A content from stack.
	Pla = docCPU.Pla
	// Plp - Pull Processor Status - pull status flags from stack.
	Plp = docCPU.Plp
	// Rol - Rotate Left.
	Rol = docCPU.Rol
	// Ror - Rotate Right.
	Ror = docCPU.Ror
	// Rti - Return from Interrupt.
	Rti = docCPU.Rti
	// Sbc - subtract with Carry.
	Sbc = docCPU.Sbc
	// Sec - Set Carry Flag.
	Sec = docCPU.Sec
	// Sed - Set Decimal Flag.
	Sed = docCPU.Sed
	// Sei - Set Interrupt Disable.
	Sei = docCPU.Sei
	// Sta - Store Accumulator - store content of A at address Addr and
	// add an optional register to the address.
	Sta = docCPU.Sta
	// Stx - Store X Register - store content of X at address Addr and
	// add an optional register to the address.
	Stx = docCPU.Stx
	// Sty - Store Y Register - store content of Y at address Addr and
	// add an optional register to the address.
	Sty = docCPU.Sty
	// Tax - Transfer Accumulator to X.
	Tax = docCPU.Tax
	// Tay - Transfer Accumulator to Y.
	Tay = docCPU.Tay
	// Tsx - Transfer Stack Pointer to X.
	Tsx = docCPU.Tsx
	// Txa - Transfer X to Accumulator.
	Txa = docCPU.Txa
	// Txs - Transfer X to Stack Pointer.
	Txs = docCPU.Txs
	// Tya - Transfer Y to Accumulator.
	Tya = docCPU.Tya
)

// setAliases links the CPU instructions to the actual CPU instance.
// nolint: funlen
func setAliases(cpu *cpu.CPU) {
	Adc = cpu.Adc
	And = cpu.And
	Asl = cpu.Asl
	Bcc = cpu.Bcc
	Bcs = cpu.Bcs
	Beq = cpu.Beq
	Bit = cpu.Bit
	Bmi = cpu.Bmi
	Bne = cpu.Bne
	Bpl = cpu.Bpl
	Brk = cpu.Brk
	Bvc = cpu.Bvc
	Bvs = cpu.Bvs
	Clc = cpu.Clc
	Cld = cpu.Cld
	Cli = cpu.Cli
	Clv = cpu.Clv
	Cmp = cpu.Cmp
	Cpx = cpu.Cpx
	Cpy = cpu.Cpy
	Dec = cpu.Dec
	Dex = cpu.Dex
	Dey = cpu.Dey
	Eor = cpu.Eor
	Inc = cpu.Inc
	Inx = cpu.Inx
	Iny = cpu.Iny
	Lda = cpu.Lda
	Ldx = cpu.Ldx
	Ldy = cpu.Ldy
	Lsr = cpu.Lsr
	Nop = cpu.Nop
	Ora = cpu.Ora
	Pha = cpu.Pha
	Php = cpu.Php
	Pla = cpu.Pla
	Plp = cpu.Plp
	Rol = cpu.Rol
	Ror = cpu.Ror
	Rti = cpu.Rti
	Sbc = cpu.Sbc
	Sec = cpu.Sec
	Sed = cpu.Sed
	Sei = cpu.Sei
	Sta = cpu.Sta
	Stx = cpu.Stx
	Sty = cpu.Sty
	Tax = cpu.Tax
	Tay = cpu.Tay
	Tsx = cpu.Tsx
	Txa = cpu.Txa
	Txs = cpu.Txs
	Tya = cpu.Tya
}
