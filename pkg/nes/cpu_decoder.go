//go:build !nesgo
// +build !nesgo

package nes

import "github.com/retroenv/nesgo/internal/ast"

type instruction struct {
	name string // nolint: structcheck

	// instruction has no parameters
	noParamFunc *func()
	// instruction has parameters
	paramFunc *func(params ...interface{})

	addressing ast.AddressingMode
}

var instructions = map[byte]instruction{
	0x00: {
		name:        "brk",
		noParamFunc: &Brk,
	},
	0x01: {
		name:       "ora",
		paramFunc:  &Ora,
		addressing: ast.IndirectXAddressing,
	},
	0x05: {
		name:       "ora",
		paramFunc:  &Ora,
		addressing: ast.ZeroPageAddressing,
	},
	0x06: {
		name:       "asl",
		paramFunc:  &Asl,
		addressing: ast.ZeroPageAddressing,
	},
	0x08: {
		name:        "php",
		noParamFunc: &Php,
	},
	0x09: {
		name:       "ora",
		paramFunc:  &Ora,
		addressing: ast.ImmediateAddressing,
	},
	0x0a: {
		name:       "asl",
		paramFunc:  &Asl,
		addressing: ast.AccumulatorAddressing,
	},
	0x0d: {
		name:       "ora",
		paramFunc:  &Ora,
		addressing: ast.AbsoluteAddressing,
	},
	0x0e: {
		name:       "asl",
		paramFunc:  &Asl,
		addressing: ast.AbsoluteAddressing,
	},
	0x10: {
		name:       "bpl",
		paramFunc:  &bpl,
		addressing: ast.RelativeAddressing,
	},
	0x11: {
		name:       "ora",
		paramFunc:  &Ora,
		addressing: ast.IndirectYAddressing,
	},
	0x15: {
		name:       "ora",
		paramFunc:  &Ora,
		addressing: ast.ZeroPageXAddressing,
	},
	0x16: {
		name:       "asl",
		paramFunc:  &Asl,
		addressing: ast.ZeroPageXAddressing,
	},
	0x18: {
		name:        "clc",
		noParamFunc: &Clc,
	},
	0x19: {
		name:       "ora",
		paramFunc:  &Ora,
		addressing: ast.AbsoluteYAddressing,
	},
	0x1d: {
		name:       "ora",
		paramFunc:  &Ora,
		addressing: ast.AbsoluteXAddressing,
	},
	0x1e: {
		name:       "asl",
		paramFunc:  &Asl,
		addressing: ast.AbsoluteXAddressing,
	},
	0x20: {
		name:       "jsr",
		paramFunc:  &jsr,
		addressing: ast.AbsoluteAddressing,
	},
	0x21: {
		name:       "and",
		paramFunc:  &And,
		addressing: ast.IndirectXAddressing,
	},
	0x24: {
		name:       "bit",
		paramFunc:  &Bit,
		addressing: ast.ZeroPageAddressing,
	},
	0x25: {
		name:       "and",
		paramFunc:  &And,
		addressing: ast.ZeroPageAddressing,
	},
	0x26: {
		name:       "rol",
		paramFunc:  &Rol,
		addressing: ast.ZeroPageAddressing,
	},
	0x28: {
		name:        "plp",
		noParamFunc: &Plp,
	},
	0x29: {
		name:       "and",
		paramFunc:  &And,
		addressing: ast.ImmediateAddressing,
	},
	0x2a: {
		name:       "rol",
		paramFunc:  &Rol,
		addressing: ast.AccumulatorAddressing,
	},
	0x2c: {
		name:       "bit",
		paramFunc:  &Bit,
		addressing: ast.AbsoluteAddressing,
	},
	0x2d: {
		name:       "and",
		paramFunc:  &And,
		addressing: ast.AbsoluteAddressing,
	},
	0x2e: {
		name:       "rol",
		paramFunc:  &Rol,
		addressing: ast.AbsoluteAddressing,
	},
	0x30: {
		name:       "bmi",
		paramFunc:  &bmi,
		addressing: ast.RelativeAddressing,
	},
	0x31: {
		name:       "and",
		paramFunc:  &And,
		addressing: ast.IndirectYAddressing,
	},
	0x35: {
		name:       "and",
		paramFunc:  &And,
		addressing: ast.ZeroPageXAddressing,
	},
	0x36: {
		name:       "rol",
		paramFunc:  &Rol,
		addressing: ast.ZeroPageXAddressing,
	},
	0x38: {
		name:        "sec",
		noParamFunc: &Sec,
	},
	0x39: {
		name:       "and",
		paramFunc:  &And,
		addressing: ast.AbsoluteYAddressing,
	},
	0x3d: {
		name:       "and",
		paramFunc:  &And,
		addressing: ast.AbsoluteXAddressing,
	},
	0x3e: {
		name:       "rol",
		paramFunc:  &Rol,
		addressing: ast.AbsoluteXAddressing,
	},
	0x40: {
		name:        "rti",
		noParamFunc: &rti,
	},
	0x41: {
		name:       "eor",
		paramFunc:  &Eor,
		addressing: ast.IndirectXAddressing,
	},
	0x45: {
		name:       "eor",
		paramFunc:  &Eor,
		addressing: ast.ZeroPageAddressing,
	},
	0x46: {
		name:       "lsr",
		paramFunc:  &Lsr,
		addressing: ast.ZeroPageAddressing,
	},
	0x48: {
		name:        "pha",
		noParamFunc: &Pha,
	},
	0x49: {
		name:       "eor",
		paramFunc:  &Eor,
		addressing: ast.ImmediateAddressing,
	},
	0x4a: {
		name:       "lsr",
		paramFunc:  &Lsr,
		addressing: ast.AccumulatorAddressing,
	},
	0x4c: {
		name:       "jmp",
		paramFunc:  &jmp,
		addressing: ast.AbsoluteAddressing,
	},
	0x4d: {
		name:       "eor",
		paramFunc:  &Eor,
		addressing: ast.AbsoluteAddressing,
	},
	0x4e: {
		name:       "lsr",
		paramFunc:  &Lsr,
		addressing: ast.AbsoluteAddressing,
	},
	0x50: {
		name:       "bvc",
		paramFunc:  &bvc,
		addressing: ast.RelativeAddressing,
	},
	0x51: {
		name:       "eor",
		paramFunc:  &Eor,
		addressing: ast.IndirectYAddressing,
	},
	0x55: {
		name:       "eor",
		paramFunc:  &Eor,
		addressing: ast.ZeroPageXAddressing,
	},
	0x56: {
		name:       "lsr",
		paramFunc:  &Lsr,
		addressing: ast.ZeroPageXAddressing,
	},
	0x58: {
		name:        "cli",
		noParamFunc: &Cli,
	},
	0x59: {
		name:       "eor",
		paramFunc:  &Eor,
		addressing: ast.AbsoluteYAddressing,
	},
	0x5d: {
		name:       "eor",
		paramFunc:  &Eor,
		addressing: ast.AbsoluteXAddressing,
	},
	0x5e: {
		name:       "lsr",
		paramFunc:  &Lsr,
		addressing: ast.AbsoluteXAddressing,
	},
	0x60: {
		name:        "rts",
		noParamFunc: &rts,
	},
	0x61: {
		name:       "adc",
		paramFunc:  &Adc,
		addressing: ast.IndirectXAddressing,
	},
	0x65: {
		name:       "adc",
		paramFunc:  &Adc,
		addressing: ast.ZeroPageAddressing,
	},
	0x66: {
		name:       "ror",
		paramFunc:  &Ror,
		addressing: ast.ZeroPageAddressing,
	},
	0x68: {
		name:        "pla",
		noParamFunc: &Pla,
	},
	0x69: {
		name:       "adc",
		paramFunc:  &Adc,
		addressing: ast.ImmediateAddressing,
	},
	0x6a: {
		name:       "ror",
		paramFunc:  &Ror,
		addressing: ast.AccumulatorAddressing,
	},
	0x6c: {
		name:       "jmp",
		paramFunc:  &jmp,
		addressing: ast.IndirectAddressing,
	},
	0x6d: {
		name:       "adc",
		paramFunc:  &Adc,
		addressing: ast.AbsoluteAddressing,
	},
	0x6e: {
		name:       "ror",
		paramFunc:  &Ror,
		addressing: ast.AbsoluteAddressing,
	},
	0x70: {
		name:       "bvs",
		paramFunc:  &bvs,
		addressing: ast.RelativeAddressing,
	},
	0x71: {
		name:       "adc",
		paramFunc:  &Adc,
		addressing: ast.IndirectYAddressing,
	},
	0x75: {
		name:       "adc",
		paramFunc:  &Adc,
		addressing: ast.ZeroPageXAddressing,
	},
	0x76: {
		name:       "ror",
		paramFunc:  &Ror,
		addressing: ast.ZeroPageXAddressing,
	},
	0x78: {
		name:        "sei",
		noParamFunc: &Sei,
	},
	0x79: {
		name:       "adc",
		paramFunc:  &Adc,
		addressing: ast.AbsoluteYAddressing,
	},
	0x7d: {
		name:       "adc",
		paramFunc:  &Adc,
		addressing: ast.AbsoluteXAddressing,
	},
	0x7e: {
		name:       "ror",
		paramFunc:  &Ror,
		addressing: ast.AbsoluteXAddressing,
	},
	0x81: {
		name:       "sta",
		paramFunc:  &Sta,
		addressing: ast.IndirectXAddressing,
	},
	0x84: {
		name:       "sty",
		paramFunc:  &Sty,
		addressing: ast.ZeroPageAddressing,
	},
	0x85: {
		name:       "sta",
		paramFunc:  &Sta,
		addressing: ast.ZeroPageAddressing,
	},
	0x86: {
		name:       "stx",
		paramFunc:  &Stx,
		addressing: ast.ZeroPageAddressing,
	},
	0x88: {
		name:        "dey",
		noParamFunc: &Dey,
	},
	0x8a: {
		name:        "txa",
		noParamFunc: &Txa,
	},
	0x8c: {
		name:       "sty",
		paramFunc:  &Sty,
		addressing: ast.AbsoluteAddressing,
	},
	0x8d: {
		name:       "sta",
		paramFunc:  &Sta,
		addressing: ast.AbsoluteAddressing,
	},
	0x8e: {
		name:       "stx",
		paramFunc:  &Stx,
		addressing: ast.AbsoluteAddressing,
	},
	0x90: {
		name:       "bcc",
		paramFunc:  &bcc,
		addressing: ast.RelativeAddressing,
	},
	0x91: {
		name:       "sta",
		paramFunc:  &Sta,
		addressing: ast.IndirectYAddressing,
	},
	0x94: {
		name:       "sty",
		paramFunc:  &Sty,
		addressing: ast.ZeroPageXAddressing,
	},
	0x95: {
		name:       "sta",
		paramFunc:  &Sta,
		addressing: ast.ZeroPageXAddressing,
	},
	0x96: {
		name:       "stx",
		paramFunc:  &Stx,
		addressing: ast.ZeroPageYAddressing,
	},
	0x98: {
		name:        "tya",
		noParamFunc: &Tya,
	},
	0x99: {
		name:       "sta",
		paramFunc:  &Sta,
		addressing: ast.AbsoluteYAddressing,
	},
	0x9a: {
		name:        "txs",
		noParamFunc: &Txs,
	},
	0x9d: {
		name:       "sta",
		paramFunc:  &Sta,
		addressing: ast.AbsoluteXAddressing,
	},
	0xa0: {
		name:       "ldy",
		paramFunc:  &Ldy,
		addressing: ast.ImmediateAddressing,
	},
	0xa1: {
		name:       "lda",
		paramFunc:  &Lda,
		addressing: ast.IndirectXAddressing,
	},
	0xa2: {
		name:       "ldx",
		paramFunc:  &Ldx,
		addressing: ast.ImmediateAddressing,
	},
	0xa4: {
		name:       "ldy",
		paramFunc:  &Ldy,
		addressing: ast.ZeroPageAddressing,
	},
	0xa5: {
		name:       "lda",
		paramFunc:  &Lda,
		addressing: ast.ZeroPageAddressing,
	},
	0xa6: {
		name:       "ldx",
		paramFunc:  &Ldx,
		addressing: ast.ZeroPageAddressing,
	},
	0xa8: {
		name:        "tay",
		noParamFunc: &Tay,
	},
	0xa9: {
		name:       "lda",
		paramFunc:  &Lda,
		addressing: ast.ImmediateAddressing,
	},
	0xaa: {
		name:        "tax",
		noParamFunc: &Tax,
	},
	0xac: {
		name:       "ldy",
		paramFunc:  &Ldy,
		addressing: ast.AbsoluteAddressing,
	},
	0xad: {
		name:       "lda",
		paramFunc:  &Lda,
		addressing: ast.AbsoluteAddressing,
	},
	0xae: {
		name:       "ldx",
		paramFunc:  &Ldx,
		addressing: ast.AbsoluteAddressing,
	},
	0xb0: {
		name:       "bcs",
		paramFunc:  &bcs,
		addressing: ast.RelativeAddressing,
	},
	0xb1: {
		name:       "lda",
		paramFunc:  &Lda,
		addressing: ast.IndirectYAddressing,
	},
	0xb4: {
		name:       "ldy",
		paramFunc:  &Ldy,
		addressing: ast.ZeroPageXAddressing,
	},
	0xb5: {
		name:       "lda",
		paramFunc:  &Lda,
		addressing: ast.ZeroPageXAddressing,
	},
	0xb6: {
		name:       "ldx",
		paramFunc:  &Ldx,
		addressing: ast.ZeroPageYAddressing,
	},
	0xb8: {
		name:        "clv",
		noParamFunc: &Clv,
	},
	0xb9: {
		name:       "lda",
		paramFunc:  &Lda,
		addressing: ast.AbsoluteYAddressing,
	},
	0xba: {
		name:        "tsx",
		noParamFunc: &Tsx,
	},
	0xbc: {
		name:       "ldy",
		paramFunc:  &Ldy,
		addressing: ast.AbsoluteXAddressing,
	},
	0xbe: {
		name:       "ldx",
		paramFunc:  &Ldx,
		addressing: ast.AbsoluteYAddressing,
	},
	0xbd: {
		name:       "lda",
		paramFunc:  &Lda,
		addressing: ast.AbsoluteXAddressing,
	},
	0xc0: {
		name:       "cpy",
		paramFunc:  &Cpy,
		addressing: ast.ImmediateAddressing,
	},
	0xc1: {
		name:       "cmp",
		paramFunc:  &Cmp,
		addressing: ast.IndirectXAddressing,
	},
	0xc4: {
		name:       "cpy",
		paramFunc:  &Cpy,
		addressing: ast.ZeroPageAddressing,
	},
	0xc5: {
		name:       "cmp",
		paramFunc:  &Cmp,
		addressing: ast.ZeroPageAddressing,
	},
	0xc6: {
		name:       "dec",
		paramFunc:  &Dec,
		addressing: ast.ZeroPageAddressing,
	},
	0xc8: {
		name:        "iny",
		noParamFunc: &Iny,
	},
	0xc9: {
		name:       "cmp",
		paramFunc:  &Cmp,
		addressing: ast.ImmediateAddressing,
	},
	0xca: {
		name:        "dex",
		noParamFunc: &Dex,
	},
	0xcc: {
		name:       "cpy",
		paramFunc:  &Cpy,
		addressing: ast.AbsoluteAddressing,
	},
	0xcd: {
		name:       "cmp",
		paramFunc:  &Cmp,
		addressing: ast.AbsoluteAddressing,
	},
	0xce: {
		name:       "dec",
		paramFunc:  &Dec,
		addressing: ast.AbsoluteAddressing,
	},
	0xd0: {
		name:       "bne",
		paramFunc:  &bne,
		addressing: ast.RelativeAddressing,
	},
	0xd1: {
		name:       "cmp",
		paramFunc:  &Cmp,
		addressing: ast.IndirectYAddressing,
	},
	0xd5: {
		name:       "cmp",
		paramFunc:  &Cmp,
		addressing: ast.ZeroPageXAddressing,
	},
	0xd6: {
		name:       "dec",
		paramFunc:  &Dec,
		addressing: ast.ZeroPageXAddressing,
	},
	0xd8: {
		name:        "cld",
		noParamFunc: &Cld,
	},
	0xd9: {
		name:       "cmp",
		paramFunc:  &Cmp,
		addressing: ast.AbsoluteYAddressing,
	},
	0xdd: {
		name:       "cmp",
		paramFunc:  &Cmp,
		addressing: ast.AbsoluteXAddressing,
	},
	0xde: {
		name:       "dec",
		paramFunc:  &Dec,
		addressing: ast.AbsoluteXAddressing,
	},
	0xe0: {
		name:       "cpx",
		paramFunc:  &Cpx,
		addressing: ast.ImmediateAddressing,
	},
	0xe1: {
		name:       "sbc",
		paramFunc:  &Sbc,
		addressing: ast.IndirectXAddressing,
	},
	0xe4: {
		name:       "cpx",
		paramFunc:  &Cpx,
		addressing: ast.ZeroPageAddressing,
	},
	0xe5: {
		name:       "sbc",
		paramFunc:  &Sbc,
		addressing: ast.ZeroPageAddressing,
	},
	0xe6: {
		name:       "inc",
		paramFunc:  &Inc,
		addressing: ast.ZeroPageAddressing,
	},
	0xe8: {
		name:        "inx",
		noParamFunc: &Inx,
	},
	0xe9: {
		name:       "sbc",
		paramFunc:  &Sbc,
		addressing: ast.ImmediateAddressing,
	},
	0xea: {
		name:        "nop",
		noParamFunc: &Nop,
	},
	0xec: {
		name:       "cpx",
		paramFunc:  &Cpx,
		addressing: ast.AbsoluteAddressing,
	},
	0xed: {
		name:       "sbc",
		paramFunc:  &Sbc,
		addressing: ast.AbsoluteAddressing,
	},
	0xee: {
		name:       "inc",
		paramFunc:  &Inc,
		addressing: ast.AbsoluteAddressing,
	},
	0xf0: {
		name:       "beq",
		paramFunc:  &beq,
		addressing: ast.RelativeAddressing,
	},
	0xf1: {
		name:       "sbc",
		paramFunc:  &Sbc,
		addressing: ast.IndirectYAddressing,
	},
	0xf5: {
		name:       "sbc",
		paramFunc:  &Sbc,
		addressing: ast.ZeroPageXAddressing,
	},
	0xf6: {
		name:       "inc",
		paramFunc:  &Inc,
		addressing: ast.ZeroPageXAddressing,
	},
	0xf8: {
		name:        "sed",
		noParamFunc: &Sed,
	},
	0xf9: {
		name:       "sbc",
		paramFunc:  &Sbc,
		addressing: ast.AbsoluteYAddressing,
	},
	0xfd: {
		name:       "sbc",
		paramFunc:  &Sbc,
		addressing: ast.AbsoluteXAddressing,
	},
	0xfe: {
		name:       "inc",
		paramFunc:  &Inc,
		addressing: ast.AbsoluteXAddressing,
	},
}
