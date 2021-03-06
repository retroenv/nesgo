//go:build !nesgo
// +build !nesgo

package cpu

import (
	. "github.com/retroenv/nesgo/pkg/addressing"
)

// Opcode is a NES CPU opcode that contains the instruction info and used
// addressing mode.
type Opcode struct {
	Instruction    *Instruction
	Addressing     Mode
	Timing         byte
	PageCrossCycle bool
}

// Opcodes maps first opcode bytes to NES CPU instruction information.
// https://www.masswerk.at/6502/6502_instruction_set.html
var Opcodes = map[byte]Opcode{
	// Official instructions
	0x00: {Instruction: brk, Addressing: ImpliedAddressing, Timing: 7},
	0x01: {Instruction: ora, Addressing: IndirectXAddressing, Timing: 6},
	0x05: {Instruction: ora, Addressing: ZeroPageAddressing, Timing: 3},
	0x06: {Instruction: asl, Addressing: ZeroPageAddressing, Timing: 5},
	0x08: {Instruction: php, Addressing: ImpliedAddressing, Timing: 3},
	0x09: {Instruction: ora, Addressing: ImmediateAddressing, Timing: 2},
	0x0a: {Instruction: asl, Addressing: AccumulatorAddressing, Timing: 2},
	0x0d: {Instruction: ora, Addressing: AbsoluteAddressing, Timing: 4},
	0x0e: {Instruction: asl, Addressing: AbsoluteAddressing, Timing: 6},
	0x10: {Instruction: bpl, Addressing: RelativeAddressing, Timing: 2},
	0x11: {Instruction: ora, Addressing: IndirectYAddressing, Timing: 5, PageCrossCycle: true},
	0x15: {Instruction: ora, Addressing: ZeroPageXAddressing, Timing: 4},
	0x16: {Instruction: asl, Addressing: ZeroPageXAddressing, Timing: 6},
	0x18: {Instruction: clc, Addressing: ImpliedAddressing, Timing: 2},
	0x19: {Instruction: ora, Addressing: AbsoluteYAddressing, Timing: 4, PageCrossCycle: true},
	0x1d: {Instruction: ora, Addressing: AbsoluteXAddressing, Timing: 4, PageCrossCycle: true},
	0x1e: {Instruction: asl, Addressing: AbsoluteXAddressing, Timing: 7},
	0x20: {Instruction: jsr, Addressing: AbsoluteAddressing, Timing: 6},
	0x21: {Instruction: and, Addressing: IndirectXAddressing, Timing: 6},
	0x24: {Instruction: bit, Addressing: ZeroPageAddressing, Timing: 3},
	0x25: {Instruction: and, Addressing: ZeroPageAddressing, Timing: 3},
	0x26: {Instruction: rol, Addressing: ZeroPageAddressing, Timing: 5},
	0x28: {Instruction: plp, Addressing: ImpliedAddressing, Timing: 4},
	0x29: {Instruction: and, Addressing: ImmediateAddressing, Timing: 2},
	0x2a: {Instruction: rol, Addressing: AccumulatorAddressing, Timing: 2},
	0x2c: {Instruction: bit, Addressing: AbsoluteAddressing, Timing: 4},
	0x2d: {Instruction: and, Addressing: AbsoluteAddressing, Timing: 4},
	0x2e: {Instruction: rol, Addressing: AbsoluteAddressing, Timing: 6},
	0x30: {Instruction: bmi, Addressing: RelativeAddressing, Timing: 2},
	0x31: {Instruction: and, Addressing: IndirectYAddressing, Timing: 5, PageCrossCycle: true},
	0x35: {Instruction: and, Addressing: ZeroPageXAddressing, Timing: 4},
	0x36: {Instruction: rol, Addressing: ZeroPageXAddressing, Timing: 6},
	0x38: {Instruction: sec, Addressing: ImpliedAddressing, Timing: 2},
	0x39: {Instruction: and, Addressing: AbsoluteYAddressing, Timing: 4, PageCrossCycle: true},
	0x3d: {Instruction: and, Addressing: AbsoluteXAddressing, Timing: 4, PageCrossCycle: true},
	0x3e: {Instruction: rol, Addressing: AbsoluteXAddressing, Timing: 7},
	0x40: {Instruction: rti, Addressing: ImpliedAddressing, Timing: 6},
	0x41: {Instruction: eor, Addressing: IndirectXAddressing, Timing: 6},
	0x45: {Instruction: eor, Addressing: ZeroPageAddressing, Timing: 3},
	0x46: {Instruction: lsr, Addressing: ZeroPageAddressing, Timing: 5},
	0x48: {Instruction: pha, Addressing: ImpliedAddressing, Timing: 3},
	0x49: {Instruction: eor, Addressing: ImmediateAddressing, Timing: 2},
	0x4a: {Instruction: lsr, Addressing: AccumulatorAddressing, Timing: 2},
	0x4c: {Instruction: jmp, Addressing: AbsoluteAddressing, Timing: 3},
	0x4d: {Instruction: eor, Addressing: AbsoluteAddressing, Timing: 4},
	0x4e: {Instruction: lsr, Addressing: AbsoluteAddressing, Timing: 6},
	0x50: {Instruction: bvc, Addressing: RelativeAddressing, Timing: 2},
	0x51: {Instruction: eor, Addressing: IndirectYAddressing, Timing: 5, PageCrossCycle: true},
	0x55: {Instruction: eor, Addressing: ZeroPageXAddressing, Timing: 4},
	0x56: {Instruction: lsr, Addressing: ZeroPageXAddressing, Timing: 6},
	0x58: {Instruction: cli, Addressing: ImpliedAddressing, Timing: 2},
	0x59: {Instruction: eor, Addressing: AbsoluteYAddressing, Timing: 4, PageCrossCycle: true},
	0x5d: {Instruction: eor, Addressing: AbsoluteXAddressing, Timing: 4, PageCrossCycle: true},
	0x5e: {Instruction: lsr, Addressing: AbsoluteXAddressing, Timing: 7, PageCrossCycle: true},
	0x60: {Instruction: rts, Addressing: ImpliedAddressing, Timing: 6},
	0x61: {Instruction: adc, Addressing: IndirectXAddressing, Timing: 6},
	0x65: {Instruction: adc, Addressing: ZeroPageAddressing, Timing: 3},
	0x66: {Instruction: ror, Addressing: ZeroPageAddressing, Timing: 5},
	0x68: {Instruction: pla, Addressing: ImpliedAddressing, Timing: 4},
	0x69: {Instruction: adc, Addressing: ImmediateAddressing, Timing: 2},
	0x6a: {Instruction: ror, Addressing: AccumulatorAddressing, Timing: 2},
	0x6c: {Instruction: jmp, Addressing: IndirectAddressing, Timing: 5},
	0x6d: {Instruction: adc, Addressing: AbsoluteAddressing, Timing: 4},
	0x6e: {Instruction: ror, Addressing: AbsoluteAddressing, Timing: 6},
	0x70: {Instruction: bvs, Addressing: RelativeAddressing, Timing: 2},
	0x71: {Instruction: adc, Addressing: IndirectYAddressing, Timing: 5, PageCrossCycle: true},
	0x75: {Instruction: adc, Addressing: ZeroPageXAddressing, Timing: 4},
	0x76: {Instruction: ror, Addressing: ZeroPageXAddressing, Timing: 6},
	0x78: {Instruction: sei, Addressing: ImpliedAddressing, Timing: 2},
	0x79: {Instruction: adc, Addressing: AbsoluteYAddressing, Timing: 4, PageCrossCycle: true},
	0x7d: {Instruction: adc, Addressing: AbsoluteXAddressing, Timing: 4, PageCrossCycle: true},
	0x7e: {Instruction: ror, Addressing: AbsoluteXAddressing, Timing: 7},
	0x81: {Instruction: sta, Addressing: IndirectXAddressing, Timing: 6},
	0x84: {Instruction: sty, Addressing: ZeroPageAddressing, Timing: 3},
	0x85: {Instruction: sta, Addressing: ZeroPageAddressing, Timing: 3},
	0x86: {Instruction: stx, Addressing: ZeroPageAddressing, Timing: 3},
	0x88: {Instruction: dey, Addressing: ImpliedAddressing, Timing: 2},
	0x8a: {Instruction: txa, Addressing: ImpliedAddressing, Timing: 2},
	0x8c: {Instruction: sty, Addressing: AbsoluteAddressing, Timing: 4},
	0x8d: {Instruction: sta, Addressing: AbsoluteAddressing, Timing: 4},
	0x8e: {Instruction: stx, Addressing: AbsoluteAddressing, Timing: 4},
	0x90: {Instruction: bcc, Addressing: RelativeAddressing, Timing: 2},
	0x91: {Instruction: sta, Addressing: IndirectYAddressing, Timing: 6},
	0x94: {Instruction: sty, Addressing: ZeroPageXAddressing, Timing: 4},
	0x95: {Instruction: sta, Addressing: ZeroPageXAddressing, Timing: 4},
	0x96: {Instruction: stx, Addressing: ZeroPageYAddressing, Timing: 4},
	0x98: {Instruction: tya, Addressing: ImpliedAddressing, Timing: 2},
	0x99: {Instruction: sta, Addressing: AbsoluteYAddressing, Timing: 5},
	0x9a: {Instruction: txs, Addressing: ImpliedAddressing, Timing: 2},
	0x9d: {Instruction: sta, Addressing: AbsoluteXAddressing, Timing: 5},
	0xa0: {Instruction: ldy, Addressing: ImmediateAddressing, Timing: 2},
	0xa1: {Instruction: lda, Addressing: IndirectXAddressing, Timing: 6},
	0xa2: {Instruction: ldx, Addressing: ImmediateAddressing, Timing: 2},
	0xa4: {Instruction: ldy, Addressing: ZeroPageAddressing, Timing: 3},
	0xa5: {Instruction: lda, Addressing: ZeroPageAddressing, Timing: 3},
	0xa6: {Instruction: ldx, Addressing: ZeroPageAddressing, Timing: 3},
	0xa8: {Instruction: tay, Addressing: ImpliedAddressing, Timing: 2},
	0xa9: {Instruction: lda, Addressing: ImmediateAddressing, Timing: 2},
	0xaa: {Instruction: tax, Addressing: ImpliedAddressing, Timing: 2},
	0xac: {Instruction: ldy, Addressing: AbsoluteAddressing, Timing: 4},
	0xad: {Instruction: lda, Addressing: AbsoluteAddressing, Timing: 4},
	0xae: {Instruction: ldx, Addressing: AbsoluteAddressing, Timing: 4},
	0xb0: {Instruction: bcs, Addressing: RelativeAddressing, Timing: 2},
	0xb1: {Instruction: lda, Addressing: IndirectYAddressing, Timing: 5, PageCrossCycle: true},
	0xb4: {Instruction: ldy, Addressing: ZeroPageXAddressing, Timing: 4},
	0xb5: {Instruction: lda, Addressing: ZeroPageXAddressing, Timing: 4},
	0xb6: {Instruction: ldx, Addressing: ZeroPageYAddressing, Timing: 4},
	0xb8: {Instruction: clv, Addressing: ImpliedAddressing, Timing: 2},
	0xb9: {Instruction: lda, Addressing: AbsoluteYAddressing, Timing: 4, PageCrossCycle: true},
	0xba: {Instruction: tsx, Addressing: ImpliedAddressing, Timing: 2},
	0xbc: {Instruction: ldy, Addressing: AbsoluteXAddressing, Timing: 4, PageCrossCycle: true},
	0xbd: {Instruction: lda, Addressing: AbsoluteXAddressing, Timing: 4, PageCrossCycle: true},
	0xbe: {Instruction: ldx, Addressing: AbsoluteYAddressing, Timing: 4, PageCrossCycle: true},
	0xc0: {Instruction: cpy, Addressing: ImmediateAddressing, Timing: 2},
	0xc1: {Instruction: cmp, Addressing: IndirectXAddressing, Timing: 6},
	0xc4: {Instruction: cpy, Addressing: ZeroPageAddressing, Timing: 3},
	0xc5: {Instruction: cmp, Addressing: ZeroPageAddressing, Timing: 3},
	0xc6: {Instruction: dec, Addressing: ZeroPageAddressing, Timing: 5},
	0xc8: {Instruction: iny, Addressing: ImpliedAddressing, Timing: 2},
	0xc9: {Instruction: cmp, Addressing: ImmediateAddressing, Timing: 2},
	0xca: {Instruction: dex, Addressing: ImpliedAddressing, Timing: 2},
	0xcc: {Instruction: cpy, Addressing: AbsoluteAddressing, Timing: 4},
	0xcd: {Instruction: cmp, Addressing: AbsoluteAddressing, Timing: 4},
	0xce: {Instruction: dec, Addressing: AbsoluteAddressing, Timing: 6},
	0xd0: {Instruction: bne, Addressing: RelativeAddressing, Timing: 2},
	0xd1: {Instruction: cmp, Addressing: IndirectYAddressing, Timing: 5, PageCrossCycle: true},
	0xd5: {Instruction: cmp, Addressing: ZeroPageXAddressing, Timing: 4},
	0xd6: {Instruction: dec, Addressing: ZeroPageXAddressing, Timing: 6},
	0xd8: {Instruction: cld, Addressing: ImpliedAddressing, Timing: 2},
	0xd9: {Instruction: cmp, Addressing: AbsoluteYAddressing, Timing: 4, PageCrossCycle: true},
	0xdd: {Instruction: cmp, Addressing: AbsoluteXAddressing, Timing: 4, PageCrossCycle: true},
	0xde: {Instruction: dec, Addressing: AbsoluteXAddressing, Timing: 7},
	0xe0: {Instruction: cpx, Addressing: ImmediateAddressing, Timing: 2},
	0xe1: {Instruction: sbc, Addressing: IndirectXAddressing, Timing: 6},
	0xe4: {Instruction: cpx, Addressing: ZeroPageAddressing, Timing: 3},
	0xe5: {Instruction: sbc, Addressing: ZeroPageAddressing, Timing: 3},
	0xe6: {Instruction: inc, Addressing: ZeroPageAddressing, Timing: 5},
	0xe8: {Instruction: inx, Addressing: ImpliedAddressing, Timing: 2},
	0xe9: {Instruction: sbc, Addressing: ImmediateAddressing, Timing: 2},
	0xea: {Instruction: nop, Addressing: ImpliedAddressing, Timing: 2},
	0xec: {Instruction: cpx, Addressing: AbsoluteAddressing, Timing: 4},
	0xed: {Instruction: sbc, Addressing: AbsoluteAddressing, Timing: 4},
	0xee: {Instruction: inc, Addressing: AbsoluteAddressing, Timing: 6},
	0xf0: {Instruction: beq, Addressing: RelativeAddressing, Timing: 2},
	0xf1: {Instruction: sbc, Addressing: IndirectYAddressing, Timing: 5, PageCrossCycle: true},
	0xf5: {Instruction: sbc, Addressing: ZeroPageXAddressing, Timing: 4},
	0xf6: {Instruction: inc, Addressing: ZeroPageXAddressing, Timing: 6},
	0xf8: {Instruction: sed, Addressing: ImpliedAddressing, Timing: 2},
	0xf9: {Instruction: sbc, Addressing: AbsoluteYAddressing, Timing: 4, PageCrossCycle: true},
	0xfd: {Instruction: sbc, Addressing: AbsoluteXAddressing, Timing: 4, PageCrossCycle: true},
	0xfe: {Instruction: inc, Addressing: AbsoluteXAddressing, Timing: 7, PageCrossCycle: true},

	// Unofficial instructions
	0x03: {Instruction: unofficialSlo, Addressing: IndirectXAddressing, Timing: 8},
	0x04: {Instruction: unofficialNop, Addressing: ZeroPageAddressing, Timing: 3},
	0x07: {Instruction: unofficialSlo, Addressing: ZeroPageAddressing, Timing: 5},
	0x0c: {Instruction: unofficialNop, Addressing: AbsoluteAddressing, Timing: 4},
	0x0f: {Instruction: unofficialSlo, Addressing: AbsoluteAddressing, Timing: 6},
	0x13: {Instruction: unofficialSlo, Addressing: IndirectYAddressing, Timing: 8},
	0x14: {Instruction: unofficialNop, Addressing: ZeroPageXAddressing, Timing: 4},
	0x17: {Instruction: unofficialSlo, Addressing: ZeroPageXAddressing, Timing: 6},
	0x1a: {Instruction: unofficialNop, Addressing: ImpliedAddressing, Timing: 2},
	0x1b: {Instruction: unofficialSlo, Addressing: AbsoluteYAddressing, Timing: 7},
	0x1c: {Instruction: unofficialNop, Addressing: AbsoluteXAddressing, Timing: 4, PageCrossCycle: true},
	0x1f: {Instruction: unofficialSlo, Addressing: AbsoluteXAddressing, Timing: 7},
	0x23: {Instruction: unofficialRla, Addressing: IndirectXAddressing, Timing: 8},
	0x27: {Instruction: unofficialRla, Addressing: ZeroPageAddressing, Timing: 5},
	0x2f: {Instruction: unofficialRla, Addressing: AbsoluteAddressing, Timing: 6},
	0x33: {Instruction: unofficialRla, Addressing: IndirectYAddressing, Timing: 8},
	0x34: {Instruction: unofficialNop, Addressing: ZeroPageXAddressing, Timing: 4},
	0x37: {Instruction: unofficialRla, Addressing: ZeroPageXAddressing, Timing: 6},
	0x3a: {Instruction: unofficialNop, Addressing: ImpliedAddressing, Timing: 2},
	0x3b: {Instruction: unofficialRla, Addressing: AbsoluteYAddressing, Timing: 7},
	0x3c: {Instruction: unofficialNop, Addressing: AbsoluteXAddressing, Timing: 4, PageCrossCycle: true},
	0x3f: {Instruction: unofficialRla, Addressing: AbsoluteXAddressing, Timing: 7},
	0x43: {Instruction: unofficialSre, Addressing: IndirectXAddressing, Timing: 8},
	0x44: {Instruction: unofficialNop, Addressing: ZeroPageAddressing, Timing: 3},
	0x47: {Instruction: unofficialSre, Addressing: ZeroPageAddressing, Timing: 5},
	0x4f: {Instruction: unofficialSre, Addressing: AbsoluteAddressing, Timing: 6},
	0x53: {Instruction: unofficialSre, Addressing: IndirectYAddressing, Timing: 8},
	0x54: {Instruction: unofficialNop, Addressing: ZeroPageXAddressing, Timing: 4},
	0x57: {Instruction: unofficialSre, Addressing: ZeroPageXAddressing, Timing: 6},
	0x5a: {Instruction: unofficialNop, Addressing: ImpliedAddressing, Timing: 2},
	0x5b: {Instruction: unofficialSre, Addressing: AbsoluteYAddressing, Timing: 7},
	0x5c: {Instruction: unofficialNop, Addressing: AbsoluteXAddressing, Timing: 4, PageCrossCycle: true},
	0x5f: {Instruction: unofficialSre, Addressing: AbsoluteXAddressing, Timing: 7},
	0x63: {Instruction: unofficialRra, Addressing: IndirectXAddressing, Timing: 8},
	0x64: {Instruction: unofficialNop, Addressing: ZeroPageAddressing, Timing: 3},
	0x67: {Instruction: unofficialRra, Addressing: ZeroPageAddressing, Timing: 5},
	0x6f: {Instruction: unofficialRra, Addressing: AbsoluteAddressing, Timing: 6},
	0x73: {Instruction: unofficialRra, Addressing: IndirectYAddressing, Timing: 8},
	0x74: {Instruction: unofficialNop, Addressing: ZeroPageXAddressing, Timing: 4},
	0x77: {Instruction: unofficialRra, Addressing: ZeroPageXAddressing, Timing: 6},
	0x7a: {Instruction: unofficialNop, Addressing: ImpliedAddressing, Timing: 2},
	0x7b: {Instruction: unofficialRra, Addressing: AbsoluteYAddressing, Timing: 7},
	0x7c: {Instruction: unofficialNop, Addressing: AbsoluteXAddressing, Timing: 4, PageCrossCycle: true},
	0x7f: {Instruction: unofficialRra, Addressing: AbsoluteXAddressing, Timing: 7},
	0x80: {Instruction: unofficialNop, Addressing: ImmediateAddressing, Timing: 2},
	0x82: {Instruction: unofficialNop, Addressing: ImmediateAddressing, Timing: 2},
	0x83: {Instruction: unofficialSax, Addressing: IndirectXAddressing, Timing: 6},
	0x87: {Instruction: unofficialSax, Addressing: ZeroPageAddressing, Timing: 3},
	0x89: {Instruction: unofficialNop, Addressing: ImmediateAddressing, Timing: 2},
	0x8f: {Instruction: unofficialSax, Addressing: AbsoluteAddressing, Timing: 4},
	0x97: {Instruction: unofficialSax, Addressing: ZeroPageYAddressing, Timing: 4},
	0xa3: {Instruction: unofficialLax, Addressing: IndirectXAddressing, Timing: 6},
	0xa7: {Instruction: unofficialLax, Addressing: ZeroPageAddressing, Timing: 3},
	0xaf: {Instruction: unofficialLax, Addressing: AbsoluteAddressing, Timing: 4},
	0xb3: {Instruction: unofficialLax, Addressing: IndirectYAddressing, Timing: 5, PageCrossCycle: true},
	0xb7: {Instruction: unofficialLax, Addressing: ZeroPageYAddressing, Timing: 4},
	0xbf: {Instruction: unofficialLax, Addressing: AbsoluteYAddressing, Timing: 4},
	0xc2: {Instruction: unofficialNop, Addressing: ImmediateAddressing, Timing: 2},
	0xc3: {Instruction: unofficialDcp, Addressing: IndirectXAddressing, Timing: 8},
	0xc7: {Instruction: unofficialDcp, Addressing: ZeroPageAddressing, Timing: 5},
	0xcf: {Instruction: unofficialDcp, Addressing: AbsoluteAddressing, Timing: 6},
	0xd3: {Instruction: unofficialDcp, Addressing: IndirectYAddressing, Timing: 8},
	0xd4: {Instruction: unofficialNop, Addressing: ZeroPageXAddressing, Timing: 4},
	0xd7: {Instruction: unofficialDcp, Addressing: ZeroPageXAddressing, Timing: 6},
	0xda: {Instruction: unofficialNop, Addressing: ImpliedAddressing, Timing: 2},
	0xdb: {Instruction: unofficialDcp, Addressing: AbsoluteYAddressing, Timing: 7},
	0xdc: {Instruction: unofficialNop, Addressing: AbsoluteXAddressing, Timing: 4, PageCrossCycle: true},
	0xdf: {Instruction: unofficialDcp, Addressing: AbsoluteXAddressing, Timing: 7},
	0xe2: {Instruction: unofficialNop, Addressing: ImmediateAddressing, Timing: 2},
	0xe3: {Instruction: unofficialIsb, Addressing: IndirectXAddressing, Timing: 8},
	0xe7: {Instruction: unofficialIsb, Addressing: ZeroPageAddressing, Timing: 5},
	0xeb: {Instruction: unofficialSbc, Addressing: ImmediateAddressing, Timing: 2},
	0xef: {Instruction: unofficialIsb, Addressing: AbsoluteAddressing, Timing: 6},
	0xf3: {Instruction: unofficialIsb, Addressing: IndirectYAddressing, Timing: 8},
	0xf4: {Instruction: unofficialNop, Addressing: ZeroPageXAddressing, Timing: 4},
	0xf7: {Instruction: unofficialIsb, Addressing: ZeroPageXAddressing, Timing: 6},
	0xfa: {Instruction: unofficialNop, Addressing: ImpliedAddressing, Timing: 2},
	0xfb: {Instruction: unofficialIsb, Addressing: AbsoluteYAddressing, Timing: 7},
	0xfc: {Instruction: unofficialNop, Addressing: AbsoluteXAddressing, Timing: 4, PageCrossCycle: true},
	0xff: {Instruction: unofficialIsb, Addressing: AbsoluteXAddressing, Timing: 7},
}

// ReadsMemory returns whether the instruction accesses memory reading.
func (opcode Opcode) ReadsMemory() bool {
	switch opcode.Addressing {
	case ImmediateAddressing, ImpliedAddressing, RelativeAddressing:
		return false
	}

	_, ok := MemoryReadInstructions[opcode.Instruction.Name]
	return ok
}

// WritesMemory returns whether the instruction accesses memory writing.
func (opcode Opcode) WritesMemory() bool {
	switch opcode.Addressing {
	case ImmediateAddressing, ImpliedAddressing, RelativeAddressing:
		return false
	}

	_, ok := MemoryWriteInstructions[opcode.Instruction.Name]
	return ok
}

// ReadWritesMemory returns whether the instruction accesses memory reading and writing.
func (opcode Opcode) ReadWritesMemory() bool {
	switch opcode.Addressing {
	case ImmediateAddressing, ImpliedAddressing, RelativeAddressing:
		return false
	}

	_, ok := MemoryReadWriteInstructions[opcode.Instruction.Name]
	return ok
}
