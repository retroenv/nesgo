//go:build !nesgo
// +build !nesgo

package cpu

import (
	. "github.com/retroenv/nesgo/pkg/addressing"
)

type Opcode struct {
	Instruction *Instruction
	Addressing  Mode
}

var Opcodes = map[byte]Opcode{
	0x00: {Instruction: brk, Addressing: ImpliedAddressing},
	0x01: {Instruction: ora, Addressing: IndirectXAddressing},
	0x05: {Instruction: ora, Addressing: ZeroPageAddressing},
	0x06: {Instruction: asl, Addressing: ZeroPageAddressing},
	0x08: {Instruction: php, Addressing: ImpliedAddressing},
	0x09: {Instruction: ora, Addressing: ImmediateAddressing},
	0x0a: {Instruction: asl, Addressing: AccumulatorAddressing},
	0x0d: {Instruction: ora, Addressing: AbsoluteAddressing},
	0x0e: {Instruction: asl, Addressing: AbsoluteAddressing},
	0x10: {Instruction: bpl, Addressing: RelativeAddressing},
	0x11: {Instruction: ora, Addressing: IndirectYAddressing},
	0x15: {Instruction: ora, Addressing: ZeroPageXAddressing},
	0x16: {Instruction: asl, Addressing: ZeroPageXAddressing},
	0x18: {Instruction: clc, Addressing: ImpliedAddressing},
	0x19: {Instruction: ora, Addressing: AbsoluteYAddressing},
	0x1d: {Instruction: ora, Addressing: AbsoluteXAddressing},
	0x1e: {Instruction: asl, Addressing: AbsoluteXAddressing},
	0x20: {Instruction: jsr, Addressing: AbsoluteAddressing},
	0x21: {Instruction: and, Addressing: IndirectXAddressing},
	0x24: {Instruction: bit, Addressing: ZeroPageAddressing},
	0x25: {Instruction: and, Addressing: ZeroPageAddressing},
	0x26: {Instruction: rol, Addressing: ZeroPageAddressing},
	0x28: {Instruction: plp, Addressing: ImpliedAddressing},
	0x29: {Instruction: and, Addressing: ImmediateAddressing},
	0x2a: {Instruction: rol, Addressing: AccumulatorAddressing},
	0x2c: {Instruction: bit, Addressing: AbsoluteAddressing},
	0x2d: {Instruction: and, Addressing: AbsoluteAddressing},
	0x2e: {Instruction: rol, Addressing: AbsoluteAddressing},
	0x30: {Instruction: bmi, Addressing: RelativeAddressing},
	0x31: {Instruction: and, Addressing: IndirectYAddressing},
	0x35: {Instruction: and, Addressing: ZeroPageXAddressing},
	0x36: {Instruction: rol, Addressing: ZeroPageXAddressing},
	0x38: {Instruction: sec, Addressing: ImpliedAddressing},
	0x39: {Instruction: and, Addressing: AbsoluteYAddressing},
	0x3d: {Instruction: and, Addressing: AbsoluteXAddressing},
	0x3e: {Instruction: rol, Addressing: AbsoluteXAddressing},
	0x40: {Instruction: rti, Addressing: ImpliedAddressing},
	0x41: {Instruction: eor, Addressing: IndirectXAddressing},
	0x45: {Instruction: eor, Addressing: ZeroPageAddressing},
	0x46: {Instruction: lsr, Addressing: ZeroPageAddressing},
	0x48: {Instruction: pha, Addressing: ImpliedAddressing},
	0x49: {Instruction: eor, Addressing: ImmediateAddressing},
	0x4a: {Instruction: lsr, Addressing: AccumulatorAddressing},
	0x4c: {Instruction: jmp, Addressing: AbsoluteAddressing},
	0x4d: {Instruction: eor, Addressing: AbsoluteAddressing},
	0x4e: {Instruction: lsr, Addressing: AbsoluteAddressing},
	0x50: {Instruction: bvc, Addressing: RelativeAddressing},
	0x51: {Instruction: eor, Addressing: IndirectYAddressing},
	0x55: {Instruction: eor, Addressing: ZeroPageXAddressing},
	0x56: {Instruction: lsr, Addressing: ZeroPageXAddressing},
	0x58: {Instruction: cli, Addressing: ImpliedAddressing},
	0x59: {Instruction: eor, Addressing: AbsoluteYAddressing},
	0x5d: {Instruction: eor, Addressing: AbsoluteXAddressing},
	0x5e: {Instruction: lsr, Addressing: AbsoluteXAddressing},
	0x60: {Instruction: rts, Addressing: ImpliedAddressing},
	0x61: {Instruction: adc, Addressing: IndirectXAddressing},
	0x65: {Instruction: adc, Addressing: ZeroPageAddressing},
	0x66: {Instruction: ror, Addressing: ZeroPageAddressing},
	0x68: {Instruction: pla, Addressing: ImpliedAddressing},
	0x69: {Instruction: adc, Addressing: ImmediateAddressing},
	0x6a: {Instruction: ror, Addressing: AccumulatorAddressing},
	0x6c: {Instruction: jmp, Addressing: IndirectAddressing},
	0x6d: {Instruction: adc, Addressing: AbsoluteAddressing},
	0x6e: {Instruction: ror, Addressing: AbsoluteAddressing},
	0x70: {Instruction: bvs, Addressing: RelativeAddressing},
	0x71: {Instruction: adc, Addressing: IndirectYAddressing},
	0x75: {Instruction: adc, Addressing: ZeroPageXAddressing},
	0x76: {Instruction: ror, Addressing: ZeroPageXAddressing},
	0x78: {Instruction: sei, Addressing: ImpliedAddressing},
	0x79: {Instruction: adc, Addressing: AbsoluteYAddressing},
	0x7d: {Instruction: adc, Addressing: AbsoluteXAddressing},
	0x7e: {Instruction: ror, Addressing: AbsoluteXAddressing},
	0x81: {Instruction: sta, Addressing: IndirectXAddressing},
	0x84: {Instruction: sty, Addressing: ZeroPageAddressing},
	0x85: {Instruction: sta, Addressing: ZeroPageAddressing},
	0x86: {Instruction: stx, Addressing: ZeroPageAddressing},
	0x88: {Instruction: dey, Addressing: ImpliedAddressing},
	0x8a: {Instruction: txa, Addressing: ImpliedAddressing},
	0x8c: {Instruction: sty, Addressing: AbsoluteAddressing},
	0x8d: {Instruction: sta, Addressing: AbsoluteAddressing},
	0x8e: {Instruction: stx, Addressing: AbsoluteAddressing},
	0x90: {Instruction: bcc, Addressing: RelativeAddressing},
	0x91: {Instruction: sta, Addressing: IndirectYAddressing},
	0x00: {Instruction: brk, Addressing: ImpliedAddressing},
	0x00: {Instruction: brk, Addressing: ImpliedAddressing},
	0x00: {Instruction: brk, Addressing: ImpliedAddressing},
	0x00: {Instruction: brk, Addressing: ImpliedAddressing},
	0x00: {Instruction: brk, Addressing: ImpliedAddressing},
	0x00: {Instruction: brk, Addressing: ImpliedAddressing},
	0x00: {Instruction: brk, Addressing: ImpliedAddressing},
	0x00: {Instruction: brk, Addressing: ImpliedAddressing},
	0x00: {Instruction: brk, Addressing: ImpliedAddressing},

	/*
		0x94: {
			name:       "sty",
			paramFunc:  &Sty,
			addressing: ZeroPageXAddressing,
		},
		0x95: {
			name:       "sta",
			paramFunc:  &Sta,
			addressing: ZeroPageXAddressing,
		},
		0x96: {
			name:       "stx",
			paramFunc:  &Stx,
			addressing: ZeroPageYAddressing,
		},
		0x98: {
			name:        "tya",
			noParamFunc: &Tya,
		},
		0x99: {
			name:       "sta",
			paramFunc:  &Sta,
			addressing: AbsoluteYAddressing,
		},
		0x9a: {
			name:        "txs",
			noParamFunc: &Txs,
		},
		0x9d: {
			name:       "sta",
			paramFunc:  &Sta,
			addressing: AbsoluteXAddressing,
		},
		0xa0: {
			name:       "ldy",
			paramFunc:  &Ldy,
			addressing: ImmediateAddressing,
		},
		0xa1: {
			name:       "lda",
			paramFunc:  &Lda,
			addressing: IndirectXAddressing,
		},
		0xa2: {
			name:       "ldx",
			paramFunc:  &Ldx,
			addressing: ImmediateAddressing,
		},
		0xa4: {
			name:       "ldy",
			paramFunc:  &Ldy,
			addressing: ZeroPageAddressing,
		},
		0xa5: {
			name:       "lda",
			paramFunc:  &Lda,
			addressing: ZeroPageAddressing,
		},
		0xa6: {
			name:       "ldx",
			paramFunc:  &Ldx,
			addressing: ZeroPageAddressing,
		},
		0xa8: {
			name:        "tay",
			noParamFunc: &Tay,
		},
		0xa9: {
			name:       "lda",
			paramFunc:  &Lda,
			addressing: ImmediateAddressing,
		},
		0xaa: {
			name:        "tax",
			noParamFunc: &Tax,
		},
		0xac: {
			name:       "ldy",
			paramFunc:  &Ldy,
			addressing: AbsoluteAddressing,
		},
		0xad: {
			name:       "lda",
			paramFunc:  &Lda,
			addressing: AbsoluteAddressing,
		},
		0xae: {
			name:       "ldx",
			paramFunc:  &Ldx,
			addressing: AbsoluteAddressing,
		},
		0xb0: {
			name:       "bcs",
			paramFunc:  &bcs,
			addressing: RelativeAddressing,
		},
		0xb1: {
			name:       "lda",
			paramFunc:  &Lda,
			addressing: IndirectYAddressing,
		},
		0xb4: {
			name:       "ldy",
			paramFunc:  &Ldy,
			addressing: ZeroPageXAddressing,
		},
		0xb5: {
			name:       "lda",
			paramFunc:  &Lda,
			addressing: ZeroPageXAddressing,
		},
		0xb6: {
			name:       "ldx",
			paramFunc:  &Ldx,
			addressing: ZeroPageYAddressing,
		},
		0xb8: {
			name:        "clv",
			noParamFunc: &Clv,
		},
		0xb9: {
			name:       "lda",
			paramFunc:  &Lda,
			addressing: AbsoluteYAddressing,
		},
		0xba: {
			name:        "tsx",
			noParamFunc: &Tsx,
		},
		0xbc: {
			name:       "ldy",
			paramFunc:  &Ldy,
			addressing: AbsoluteXAddressing,
		},
		0xbe: {
			name:       "ldx",
			paramFunc:  &Ldx,
			addressing: AbsoluteYAddressing,
		},
		0xbd: {
			name:       "lda",
			paramFunc:  &Lda,
			addressing: AbsoluteXAddressing,
		},
		0xc0: {
			name:       "cpy",
			paramFunc:  &Cpy,
			addressing: ImmediateAddressing,
		},
		0xc1: {
			name:       "cmp",
			paramFunc:  &Cmp,
			addressing: IndirectXAddressing,
		},
		0xc4: {
			name:       "cpy",
			paramFunc:  &Cpy,
			addressing: ZeroPageAddressing,
		},
		0xc5: {
			name:       "cmp",
			paramFunc:  &Cmp,
			addressing: ZeroPageAddressing,
		},
		0xc6: {
			name:       "dec",
			paramFunc:  &Dec,
			addressing: ZeroPageAddressing,
		},
		0xc8: {
			name:        "iny",
			noParamFunc: &Iny,
		},
		0xc9: {
			name:       "cmp",
			paramFunc:  &Cmp,
			addressing: ImmediateAddressing,
		},
		0xca: {
			name:        "dex",
			noParamFunc: &Dex,
		},
		0xcc: {
			name:       "cpy",
			paramFunc:  &Cpy,
			addressing: AbsoluteAddressing,
		},
		0xcd: {
			name:       "cmp",
			paramFunc:  &Cmp,
			addressing: AbsoluteAddressing,
		},
		0xce: {
			name:       "dec",
			paramFunc:  &Dec,
			addressing: AbsoluteAddressing,
		},
		0xd0: {
			name:       "bne",
			paramFunc:  &bne,
			addressing: RelativeAddressing,
		},
		0xd1: {
			name:       "cmp",
			paramFunc:  &Cmp,
			addressing: IndirectYAddressing,
		},
		0xd5: {
			name:       "cmp",
			paramFunc:  &Cmp,
			addressing: ZeroPageXAddressing,
		},
		0xd6: {
			name:       "dec",
			paramFunc:  &Dec,
			addressing: ZeroPageXAddressing,
		},
		0xd8: {
			name:        "cld",
			noParamFunc: &Cld,
		},
		0xd9: {
			name:       "cmp",
			paramFunc:  &Cmp,
			addressing: AbsoluteYAddressing,
		},
		0xdd: {
			name:       "cmp",
			paramFunc:  &Cmp,
			addressing: AbsoluteXAddressing,
		},
		0xde: {
			name:       "dec",
			paramFunc:  &Dec,
			addressing: AbsoluteXAddressing,
		},
		0xe0: {
			name:       "cpx",
			paramFunc:  &Cpx,
			addressing: ImmediateAddressing,
		},
		0xe1: {
			name:       "sbc",
			paramFunc:  &Sbc,
			addressing: IndirectXAddressing,
		},
		0xe4: {
			name:       "cpx",
			paramFunc:  &Cpx,
			addressing: ZeroPageAddressing,
		},
		0xe5: {
			name:       "sbc",
			paramFunc:  &Sbc,
			addressing: ZeroPageAddressing,
		},
		0xe6: {
			name:       "inc",
			paramFunc:  &Inc,
			addressing: ZeroPageAddressing,
		},
		0xe8: {
			name:        "inx",
			noParamFunc: &Inx,
		},
		0xe9: {
			name:       "sbc",
			paramFunc:  &Sbc,
			addressing: ImmediateAddressing,
		},
		0xea: {
			name:        "nop",
			noParamFunc: &Nop,
		},
		0xec: {
			name:       "cpx",
			paramFunc:  &Cpx,
			addressing: AbsoluteAddressing,
		},
		0xed: {
			name:       "sbc",
			paramFunc:  &Sbc,
			addressing: AbsoluteAddressing,
		},
		0xee: {
			name:       "inc",
			paramFunc:  &Inc,
			addressing: AbsoluteAddressing,
		},
		0xf0: {
			name:       "beq",
			paramFunc:  &beq,
			addressing: RelativeAddressing,
		},
		0xf1: {
			name:       "sbc",
			paramFunc:  &Sbc,
			addressing: IndirectYAddressing,
		},
		0xf5: {
			name:       "sbc",
			paramFunc:  &Sbc,
			addressing: ZeroPageXAddressing,
		},
		0xf6: {
			name:       "inc",
			paramFunc:  &Inc,
			addressing: ZeroPageXAddressing,
		},
		0xf8: {
			name:        "sed",
			noParamFunc: &Sed,
		},
		0xf9: {
			name:       "sbc",
			paramFunc:  &Sbc,
			addressing: AbsoluteYAddressing,
		},
		0xfd: {
			name:       "sbc",
			paramFunc:  &Sbc,
			addressing: AbsoluteXAddressing,
		},
		0xfe: {
			name:       "inc",
			paramFunc:  &Inc,
			addressing: AbsoluteXAddressing,
		},
	*/
}
