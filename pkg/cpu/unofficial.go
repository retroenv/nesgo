//go:build !nesgo
// +build !nesgo

// This file contains support for unofficial CPU instructions.
// https://www.nesdev.org/wiki/Programming_with_unofficial_opcodes

package cpu

import (
	. "github.com/retroenv/nesgo/pkg/addressing"
)

func linkUnofficialInstructionFuncs(cpu *CPU) {
	unofficialDcp.ParamFunc = cpu.unofficialDcp
	unofficialIsb.ParamFunc = cpu.unofficialIsb
	unofficialLax.ParamFunc = cpu.unofficialLax
	unofficialNop.ParamFunc = cpu.unofficialNop
	unofficialRla.ParamFunc = cpu.unofficialRla
	unofficialRra.ParamFunc = cpu.unofficialRra
	unofficialSax.ParamFunc = cpu.unofficialSax
	unofficialSbc.ParamFunc = cpu.unofficialSbc
	unofficialSlo.ParamFunc = cpu.unofficialSlo
	unofficialSre.ParamFunc = cpu.unofficialSre
}

var unofficialDcp = &Instruction{
	Name:       "dcp",
	Unofficial: true,
	Addressing: map[Mode]AddressingInfo{
		ZeroPageAddressing:  {Opcode: 0xc7},
		ZeroPageXAddressing: {Opcode: 0xd7},
		AbsoluteAddressing:  {Opcode: 0xcf},
		AbsoluteXAddressing: {Opcode: 0xdf},
		AbsoluteYAddressing: {Opcode: 0xdb},
		IndirectXAddressing: {Opcode: 0xc3},
		IndirectYAddressing: {Opcode: 0xd3},
	},
}

var unofficialIsb = &Instruction{
	Name:       "isb",
	Unofficial: true,
	Addressing: map[Mode]AddressingInfo{
		ZeroPageAddressing:  {Opcode: 0xe7},
		ZeroPageXAddressing: {Opcode: 0xf7},
		AbsoluteAddressing:  {Opcode: 0xef},
		AbsoluteXAddressing: {Opcode: 0xff},
		AbsoluteYAddressing: {Opcode: 0xfb},
		IndirectXAddressing: {Opcode: 0xe3},
		IndirectYAddressing: {Opcode: 0xf3},
	},
}

var unofficialLax = &Instruction{
	Name:       "lax",
	Unofficial: true,
	Addressing: map[Mode]AddressingInfo{
		ZeroPageAddressing:  {Opcode: 0xa7},
		ZeroPageYAddressing: {Opcode: 0xb7},
		AbsoluteAddressing:  {Opcode: 0xaf},
		AbsoluteYAddressing: {Opcode: 0xbf},
		IndirectXAddressing: {Opcode: 0xa3},
		IndirectYAddressing: {Opcode: 0xb3},
	},
}

var unofficialNop = &Instruction{
	Name:       "nop",
	Unofficial: true,
}

var unofficialRla = &Instruction{
	Name:       "rla",
	Unofficial: true,
	Addressing: map[Mode]AddressingInfo{
		ZeroPageAddressing:  {Opcode: 0x27},
		ZeroPageXAddressing: {Opcode: 0x37},
		AbsoluteAddressing:  {Opcode: 0x2f},
		AbsoluteXAddressing: {Opcode: 0x3f},
		AbsoluteYAddressing: {Opcode: 0x3b},
		IndirectXAddressing: {Opcode: 0x23},
		IndirectYAddressing: {Opcode: 0x33},
	},
}

var unofficialRra = &Instruction{
	Name:       "rra",
	Unofficial: true,
	Addressing: map[Mode]AddressingInfo{
		ZeroPageAddressing:  {Opcode: 0x67},
		ZeroPageXAddressing: {Opcode: 0x77},
		AbsoluteAddressing:  {Opcode: 0x6f},
		AbsoluteXAddressing: {Opcode: 0x7f},
		AbsoluteYAddressing: {Opcode: 0x7b},
		IndirectXAddressing: {Opcode: 0x63},
		IndirectYAddressing: {Opcode: 0x73},
	},
}

var unofficialSax = &Instruction{
	Name:       "sax",
	Unofficial: true,
	Addressing: map[Mode]AddressingInfo{
		ZeroPageAddressing:  {Opcode: 0x87},
		ZeroPageYAddressing: {Opcode: 0x97},
		AbsoluteAddressing:  {Opcode: 0x8f},
		IndirectXAddressing: {Opcode: 0x83},
	},
}

var unofficialSbc = &Instruction{
	Name:       "sbc",
	Unofficial: true,
	Addressing: map[Mode]AddressingInfo{
		ImmediateAddressing: {Opcode: 0xeb},
	},
}

var unofficialSlo = &Instruction{
	Name:       "slo",
	Unofficial: true,
	Addressing: map[Mode]AddressingInfo{
		ZeroPageAddressing:  {Opcode: 0x07},
		ZeroPageXAddressing: {Opcode: 0x17},
		AbsoluteAddressing:  {Opcode: 0x0f},
		AbsoluteXAddressing: {Opcode: 0x1f},
		AbsoluteYAddressing: {Opcode: 0x1b},
		IndirectXAddressing: {Opcode: 0x03},
		IndirectYAddressing: {Opcode: 0x13},
	},
}

var unofficialSre = &Instruction{
	Name:       "sre",
	Unofficial: true,
	Addressing: map[Mode]AddressingInfo{
		ZeroPageAddressing:  {Opcode: 0x47},
		ZeroPageXAddressing: {Opcode: 0x57},
		AbsoluteAddressing:  {Opcode: 0x4f},
		AbsoluteXAddressing: {Opcode: 0x5f},
		AbsoluteYAddressing: {Opcode: 0x5b},
		IndirectXAddressing: {Opcode: 0x43},
		IndirectYAddressing: {Opcode: 0x53},
	},
}
