//go:build !nesgo
// +build !nesgo

// This file contains support for unofficial CPU instructions.
// https://www.nesdev.com/undocumented_opcodes.txt
// https://www.nesdev.org/wiki/Programming_with_unofficial_opcodes

package cpu

import (
	. "github.com/retroenv/nesgo/pkg/addressing"
)

func linkUnofficialInstructionFuncs(cpu *CPU) {
	unofficialAso.ParamFunc = cpu.unofficialAso
	unofficialDcp.ParamFunc = cpu.unofficialDcp
	unofficialIns.ParamFunc = cpu.unofficialIns
	unofficialLax.ParamFunc = cpu.unofficialLax
	unofficialLse.ParamFunc = cpu.unofficialLse
	unofficialNop.ParamFunc = cpu.unofficialNop
	unofficialRla.ParamFunc = cpu.unofficialRla
	unofficialRra.ParamFunc = cpu.unofficialRra
	unofficialSax.ParamFunc = cpu.unofficialSax
	unofficialSbc.ParamFunc = cpu.unofficialSbc
}

var unofficialAso = &Instruction{
	Name:       "aso",
	unofficial: true,
	Addressing: map[Mode]AddressingInfo{
		ZeroPageAddressing:  {Opcode: 0x07, Timing: 5},
		ZeroPageXAddressing: {Opcode: 0x17, Timing: 6},
		AbsoluteAddressing:  {Opcode: 0x0f, Timing: 6},
		AbsoluteXAddressing: {Opcode: 0x1f, Timing: 7},
		AbsoluteYAddressing: {Opcode: 0x1b, Timing: 7},
		IndirectXAddressing: {Opcode: 0x03, Timing: 8},
		IndirectYAddressing: {Opcode: 0x13, Timing: 8},
	},
}

var unofficialDcp = &Instruction{
	Name:       "dcp",
	unofficial: true,
	Addressing: map[Mode]AddressingInfo{
		ZeroPageAddressing:  {Opcode: 0xc7, Timing: 5},
		ZeroPageXAddressing: {Opcode: 0xd7, Timing: 6},
		AbsoluteAddressing:  {Opcode: 0xcf, Timing: 6},
		AbsoluteXAddressing: {Opcode: 0xdf, Timing: 7},
		AbsoluteYAddressing: {Opcode: 0xdb, Timing: 7},
		IndirectXAddressing: {Opcode: 0xc3, Timing: 8},
		IndirectYAddressing: {Opcode: 0xd3, Timing: 8},
	},
}

var unofficialIns = &Instruction{
	Name:       "ins",
	unofficial: true,
	Addressing: map[Mode]AddressingInfo{
		ZeroPageAddressing:  {Opcode: 0xe7, Timing: 5},
		ZeroPageXAddressing: {Opcode: 0xf7, Timing: 6},
		AbsoluteAddressing:  {Opcode: 0xef, Timing: 6},
		AbsoluteXAddressing: {Opcode: 0xff, Timing: 7},
		AbsoluteYAddressing: {Opcode: 0xfb, Timing: 7},
		IndirectXAddressing: {Opcode: 0xe3, Timing: 8},
		IndirectYAddressing: {Opcode: 0xf3, Timing: 8},
	},
}

var unofficialLax = &Instruction{
	Name:       "lax",
	unofficial: true,
	Addressing: map[Mode]AddressingInfo{
		ZeroPageAddressing:  {Opcode: 0xa7, Timing: 3},
		ZeroPageYAddressing: {Opcode: 0xb7, Timing: 4},
		AbsoluteAddressing:  {Opcode: 0xaf, Timing: 4},
		AbsoluteYAddressing: {Opcode: 0xbf, Timing: 4},
		IndirectXAddressing: {Opcode: 0xa3, Timing: 6},
		IndirectYAddressing: {Opcode: 0xb3, Timing: 5},
	},
}

var unofficialLse = &Instruction{
	Name:       "lse",
	unofficial: true,
	Addressing: map[Mode]AddressingInfo{
		ZeroPageAddressing:  {Opcode: 0x47, Timing: 5},
		ZeroPageXAddressing: {Opcode: 0x57, Timing: 6},
		AbsoluteAddressing:  {Opcode: 0x4f, Timing: 6},
		AbsoluteXAddressing: {Opcode: 0x5f, Timing: 7},
		AbsoluteYAddressing: {Opcode: 0x5b, Timing: 7},
		IndirectXAddressing: {Opcode: 0x43, Timing: 8},
		IndirectYAddressing: {Opcode: 0x53, Timing: 8},
	},
}

var unofficialNop = &Instruction{
	Name:       "nop",
	unofficial: true,
}

var unofficialRla = &Instruction{
	Name:       "rla",
	unofficial: true,
	Addressing: map[Mode]AddressingInfo{
		ZeroPageAddressing:  {Opcode: 0x27, Timing: 5},
		ZeroPageXAddressing: {Opcode: 0x37, Timing: 6},
		AbsoluteAddressing:  {Opcode: 0x2f, Timing: 6},
		AbsoluteXAddressing: {Opcode: 0x3f, Timing: 7},
		AbsoluteYAddressing: {Opcode: 0x3b, Timing: 7},
		IndirectXAddressing: {Opcode: 0x23, Timing: 8},
		IndirectYAddressing: {Opcode: 0x33, Timing: 8},
	},
}

var unofficialRra = &Instruction{
	Name: "rra",
	Addressing: map[Mode]AddressingInfo{
		ZeroPageAddressing:  {Opcode: 0x67, Timing: 5},
		ZeroPageXAddressing: {Opcode: 0x77, Timing: 6},
		AbsoluteAddressing:  {Opcode: 0x6f, Timing: 6},
		AbsoluteXAddressing: {Opcode: 0x7f, Timing: 7},
		AbsoluteYAddressing: {Opcode: 0x7b, Timing: 7},
		IndirectXAddressing: {Opcode: 0x63, Timing: 8},
		IndirectYAddressing: {Opcode: 0x73, Timing: 8},
	},
}

var unofficialSax = &Instruction{
	Name:       "sax",
	unofficial: true,
	Addressing: map[Mode]AddressingInfo{
		ZeroPageAddressing:  {Opcode: 0x87, Timing: 3},
		ZeroPageYAddressing: {Opcode: 0x97, Timing: 4},
		AbsoluteAddressing:  {Opcode: 0x8f, Timing: 4},
		IndirectXAddressing: {Opcode: 0x83, Timing: 6},
	},
}

var unofficialSbc = &Instruction{
	Name: "sbc",
	Addressing: map[Mode]AddressingInfo{
		ImmediateAddressing: {Opcode: 0xeb, Timing: 2},
	},
}
