//go:build !nesgo

package cpu

import (
	. "github.com/retroenv/nesgo/pkg/addressing"
)

// Opcode is a NES CPU opcode that contains the instruction info and used addressing mode.
type Opcode struct {
	Instruction    *Instruction
	Addressing     Mode
	Timing         byte
	PageCrossCycle bool
}

// Opcodes maps first opcode bytes to NES CPU instruction information.
// https://www.masswerk.at/6502/6502_instruction_set.html
var Opcodes = [256]Opcode{
	{Instruction: brk, Addressing: ImpliedAddressing, Timing: 7},   // 0x00
	{Instruction: ora, Addressing: IndirectXAddressing, Timing: 6}, // 0x01
	{}, // 0x02
	{Instruction: unofficialSlo, Addressing: IndirectXAddressing, Timing: 8}, // 0x03
	{Instruction: unofficialNop, Addressing: ZeroPageAddressing, Timing: 3},  // 0x04
	{Instruction: ora, Addressing: ZeroPageAddressing, Timing: 3},            // 0x05
	{Instruction: asl, Addressing: ZeroPageAddressing, Timing: 5},            // 0x06
	{Instruction: unofficialSlo, Addressing: ZeroPageAddressing, Timing: 5},  // 0x07
	{Instruction: php, Addressing: ImpliedAddressing, Timing: 3},             // 0x08
	{Instruction: ora, Addressing: ImmediateAddressing, Timing: 2},           // 0x09
	{Instruction: asl, Addressing: AccumulatorAddressing, Timing: 2},         // 0x0a
	{}, // 0x0b
	{Instruction: unofficialNop, Addressing: AbsoluteAddressing, Timing: 4},                        // 0x0c
	{Instruction: ora, Addressing: AbsoluteAddressing, Timing: 4},                                  // 0x0d
	{Instruction: asl, Addressing: AbsoluteAddressing, Timing: 6},                                  // 0x0e
	{Instruction: unofficialSlo, Addressing: AbsoluteAddressing, Timing: 6},                        // 0x0f
	{Instruction: bpl, Addressing: RelativeAddressing, Timing: 2},                                  // 0x10
	{Instruction: ora, Addressing: IndirectYAddressing, Timing: 5, PageCrossCycle: true},           // 0x11
	{Instruction: unofficialSlo, Addressing: IndirectYAddressing, Timing: 8},                       // 0x13
	{Instruction: unofficialNop, Addressing: ZeroPageXAddressing, Timing: 4},                       // 0x14
	{Instruction: ora, Addressing: ZeroPageXAddressing, Timing: 4},                                 // 0x15
	{Instruction: asl, Addressing: ZeroPageXAddressing, Timing: 6},                                 // 0x16
	{Instruction: unofficialSlo, Addressing: ZeroPageXAddressing, Timing: 6},                       // 0x17
	{Instruction: clc, Addressing: ImpliedAddressing, Timing: 2},                                   // 0x18
	{Instruction: ora, Addressing: AbsoluteYAddressing, Timing: 4, PageCrossCycle: true},           // 0x19
	{Instruction: unofficialNop, Addressing: ImpliedAddressing, Timing: 2},                         // 0x1a
	{Instruction: unofficialSlo, Addressing: AbsoluteYAddressing, Timing: 7},                       // 0x1b
	{Instruction: unofficialNop, Addressing: AbsoluteXAddressing, Timing: 4, PageCrossCycle: true}, // 0x1c
	{}, // 0x1c
	{Instruction: ora, Addressing: AbsoluteXAddressing, Timing: 4, PageCrossCycle: true}, // 0x1d
	{Instruction: asl, Addressing: AbsoluteXAddressing, Timing: 7},                       // 0x1e
	{Instruction: unofficialSlo, Addressing: AbsoluteXAddressing, Timing: 7},             // 0x1f
	{Instruction: jsr, Addressing: AbsoluteAddressing, Timing: 6},                        // 0x20
	{Instruction: and, Addressing: IndirectXAddressing, Timing: 6},                       // 0x21
	{}, // 0x22
	{Instruction: unofficialRla, Addressing: IndirectXAddressing, Timing: 8}, // 0x23
	{Instruction: bit, Addressing: ZeroPageAddressing, Timing: 3},            // 0x24
	{Instruction: and, Addressing: ZeroPageAddressing, Timing: 3},            // 0x25
	{Instruction: rol, Addressing: ZeroPageAddressing, Timing: 5},            // 0x26
	{Instruction: unofficialRla, Addressing: ZeroPageAddressing, Timing: 5},  // 0x27
	{Instruction: plp, Addressing: ImpliedAddressing, Timing: 4},             // 0x28
	{Instruction: and, Addressing: ImmediateAddressing, Timing: 2},           // 0x29
	{Instruction: rol, Addressing: AccumulatorAddressing, Timing: 2},         // 0x2a
	{}, // 0x2b
	{Instruction: bit, Addressing: AbsoluteAddressing, Timing: 4},                        // 0x2c
	{Instruction: and, Addressing: AbsoluteAddressing, Timing: 4},                        // 0x2d
	{Instruction: rol, Addressing: AbsoluteAddressing, Timing: 6},                        // 0x2e
	{Instruction: unofficialRla, Addressing: AbsoluteAddressing, Timing: 6},              // 0x2f
	{Instruction: bmi, Addressing: RelativeAddressing, Timing: 2},                        // 0x30
	{Instruction: and, Addressing: IndirectYAddressing, Timing: 5, PageCrossCycle: true}, // 0x31
	{}, // 0x32
	{Instruction: unofficialRla, Addressing: IndirectYAddressing, Timing: 8},                       // 0x33
	{Instruction: unofficialNop, Addressing: ZeroPageXAddressing, Timing: 4},                       // 0x34
	{Instruction: and, Addressing: ZeroPageXAddressing, Timing: 4},                                 // 0x35
	{Instruction: rol, Addressing: ZeroPageXAddressing, Timing: 6},                                 // 0x36
	{Instruction: unofficialRla, Addressing: ZeroPageXAddressing, Timing: 6},                       // 0x37
	{Instruction: sec, Addressing: ImpliedAddressing, Timing: 2},                                   // 0x38
	{Instruction: and, Addressing: AbsoluteYAddressing, Timing: 4, PageCrossCycle: true},           // 0x39
	{Instruction: unofficialNop, Addressing: ImpliedAddressing, Timing: 2},                         // 0x3a
	{Instruction: unofficialRla, Addressing: AbsoluteYAddressing, Timing: 7},                       // 0x3b
	{Instruction: unofficialNop, Addressing: AbsoluteXAddressing, Timing: 4, PageCrossCycle: true}, // 0x3c
	{Instruction: and, Addressing: AbsoluteXAddressing, Timing: 4, PageCrossCycle: true},           // 0x3d
	{Instruction: rol, Addressing: AbsoluteXAddressing, Timing: 7},                                 // 0x3e
	{Instruction: unofficialRla, Addressing: AbsoluteXAddressing, Timing: 7},                       // 0x3f
	{Instruction: rti, Addressing: ImpliedAddressing, Timing: 6},                                   // 0x40
	{Instruction: eor, Addressing: IndirectXAddressing, Timing: 6},                                 // 0x41
	{}, // 0x42
	{Instruction: unofficialSre, Addressing: IndirectXAddressing, Timing: 8}, // 0x43
	{Instruction: unofficialNop, Addressing: ZeroPageAddressing, Timing: 3},  // 0x44
	{Instruction: eor, Addressing: ZeroPageAddressing, Timing: 3},            // 0x45
	{Instruction: lsr, Addressing: ZeroPageAddressing, Timing: 5},            // 0x46
	{Instruction: unofficialSre, Addressing: ZeroPageAddressing, Timing: 5},  // 0x47
	{Instruction: pha, Addressing: ImpliedAddressing, Timing: 3},             // 0x48
	{Instruction: eor, Addressing: ImmediateAddressing, Timing: 2},           // 0x49
	{Instruction: lsr, Addressing: AccumulatorAddressing, Timing: 2},         // 0x4a
	{}, // 0x4b
	{Instruction: jmp, Addressing: AbsoluteAddressing, Timing: 3},                        // 0x4c
	{Instruction: eor, Addressing: AbsoluteAddressing, Timing: 4},                        // 0x4d
	{Instruction: lsr, Addressing: AbsoluteAddressing, Timing: 6},                        // 0x4e
	{Instruction: unofficialSre, Addressing: AbsoluteAddressing, Timing: 6},              // 0x4f
	{Instruction: bvc, Addressing: RelativeAddressing, Timing: 2},                        // 0x50
	{Instruction: eor, Addressing: IndirectYAddressing, Timing: 5, PageCrossCycle: true}, // 0x51
	{}, // 0x52
	{Instruction: unofficialSre, Addressing: IndirectYAddressing, Timing: 8},                       // 0x53
	{Instruction: unofficialNop, Addressing: ZeroPageXAddressing, Timing: 4},                       // 0x54
	{Instruction: eor, Addressing: ZeroPageXAddressing, Timing: 4},                                 // 0x55
	{Instruction: lsr, Addressing: ZeroPageXAddressing, Timing: 6},                                 // 0x56
	{Instruction: unofficialSre, Addressing: ZeroPageXAddressing, Timing: 6},                       // 0x57
	{Instruction: cli, Addressing: ImpliedAddressing, Timing: 2},                                   // 0x58
	{Instruction: eor, Addressing: AbsoluteYAddressing, Timing: 4, PageCrossCycle: true},           // 0x59
	{Instruction: unofficialNop, Addressing: ImpliedAddressing, Timing: 2},                         // 0x5a
	{Instruction: unofficialSre, Addressing: AbsoluteYAddressing, Timing: 7},                       // 0x5b
	{Instruction: unofficialNop, Addressing: AbsoluteXAddressing, Timing: 4, PageCrossCycle: true}, // 0x5c
	{Instruction: eor, Addressing: AbsoluteXAddressing, Timing: 4, PageCrossCycle: true},           // 0x5d
	{Instruction: lsr, Addressing: AbsoluteXAddressing, Timing: 7, PageCrossCycle: true},           // 0x5e
	{Instruction: unofficialSre, Addressing: AbsoluteXAddressing, Timing: 7},                       // 0x5f
	{Instruction: rts, Addressing: ImpliedAddressing, Timing: 6},                                   // 0x60
	{Instruction: adc, Addressing: IndirectXAddressing, Timing: 6},                                 // 0x61
	{}, // 0x62
	{Instruction: unofficialRra, Addressing: IndirectXAddressing, Timing: 8}, // 0x63
	{Instruction: unofficialNop, Addressing: ZeroPageAddressing, Timing: 3},  // 0x64
	{Instruction: adc, Addressing: ZeroPageAddressing, Timing: 3},            // 0x65
	{Instruction: ror, Addressing: ZeroPageAddressing, Timing: 5},            // 0x66
	{Instruction: unofficialRra, Addressing: ZeroPageAddressing, Timing: 5},  // 0x67
	{Instruction: pla, Addressing: ImpliedAddressing, Timing: 4},             // 0x68
	{Instruction: adc, Addressing: ImmediateAddressing, Timing: 2},           // 0x69
	{Instruction: ror, Addressing: AccumulatorAddressing, Timing: 2},         // 0x6a
	{}, // 0x6b
	{Instruction: jmp, Addressing: IndirectAddressing, Timing: 5},                        // 0x6c
	{Instruction: adc, Addressing: AbsoluteAddressing, Timing: 4},                        // 0x6d
	{Instruction: ror, Addressing: AbsoluteAddressing, Timing: 6},                        // 0x6e
	{Instruction: unofficialRra, Addressing: AbsoluteAddressing, Timing: 6},              // 0x6f
	{Instruction: bvs, Addressing: RelativeAddressing, Timing: 2},                        // 0x70
	{Instruction: adc, Addressing: IndirectYAddressing, Timing: 5, PageCrossCycle: true}, // 0x71
	{}, // 0x72
	{Instruction: unofficialRra, Addressing: IndirectYAddressing, Timing: 8},                       // 0x73
	{Instruction: unofficialNop, Addressing: ZeroPageXAddressing, Timing: 4},                       // 0x74
	{Instruction: adc, Addressing: ZeroPageXAddressing, Timing: 4},                                 // 0x75
	{Instruction: ror, Addressing: ZeroPageXAddressing, Timing: 6},                                 // 0x76
	{Instruction: unofficialRra, Addressing: ZeroPageXAddressing, Timing: 6},                       // 0x77
	{Instruction: sei, Addressing: ImpliedAddressing, Timing: 2},                                   // 0x78
	{Instruction: adc, Addressing: AbsoluteYAddressing, Timing: 4, PageCrossCycle: true},           // 0x79
	{Instruction: unofficialNop, Addressing: ImpliedAddressing, Timing: 2},                         // 0x7a
	{Instruction: unofficialRra, Addressing: AbsoluteYAddressing, Timing: 7},                       // 0x7b
	{Instruction: unofficialNop, Addressing: AbsoluteXAddressing, Timing: 4, PageCrossCycle: true}, // 0x7c
	{Instruction: adc, Addressing: AbsoluteXAddressing, Timing: 4, PageCrossCycle: true},           // 0x7d
	{Instruction: ror, Addressing: AbsoluteXAddressing, Timing: 7},                                 // 0x7e
	{Instruction: unofficialRra, Addressing: AbsoluteXAddressing, Timing: 7},                       // 0x7f
	{Instruction: unofficialNop, Addressing: ImmediateAddressing, Timing: 2},                       // 0x80
	{Instruction: sta, Addressing: IndirectXAddressing, Timing: 6},                                 // 0x81
	{Instruction: unofficialNop, Addressing: ImmediateAddressing, Timing: 2},                       // 0x82
	{Instruction: unofficialSax, Addressing: IndirectXAddressing, Timing: 6},                       // 0x83
	{Instruction: sty, Addressing: ZeroPageAddressing, Timing: 3},                                  // 0x84
	{Instruction: sta, Addressing: ZeroPageAddressing, Timing: 3},                                  // 0x85
	{Instruction: stx, Addressing: ZeroPageAddressing, Timing: 3},                                  // 0x86
	{Instruction: unofficialSax, Addressing: ZeroPageAddressing, Timing: 3},                        // 0x87
	{Instruction: dey, Addressing: ImpliedAddressing, Timing: 2},                                   // 0x88
	{Instruction: unofficialNop, Addressing: ImmediateAddressing, Timing: 2},                       // 0x89
	{Instruction: txa, Addressing: ImpliedAddressing, Timing: 2},                                   // 0x8a
	{}, // 0x8b
	{Instruction: sty, Addressing: AbsoluteAddressing, Timing: 4},           // 0x8c
	{Instruction: sta, Addressing: AbsoluteAddressing, Timing: 4},           // 0x8d
	{Instruction: stx, Addressing: AbsoluteAddressing, Timing: 4},           // 0x8e
	{Instruction: unofficialSax, Addressing: AbsoluteAddressing, Timing: 4}, // 0x8f
	{Instruction: bcc, Addressing: RelativeAddressing, Timing: 2},           // 0x90
	{Instruction: sta, Addressing: IndirectYAddressing, Timing: 6},          // 0x91
	{}, // 0x92
	{}, // 0x93
	{Instruction: sty, Addressing: ZeroPageXAddressing, Timing: 4},           // 0x94
	{Instruction: sta, Addressing: ZeroPageXAddressing, Timing: 4},           // 0x95
	{Instruction: stx, Addressing: ZeroPageYAddressing, Timing: 4},           // 0x96
	{Instruction: unofficialSax, Addressing: ZeroPageYAddressing, Timing: 4}, // 0x97
	{Instruction: tya, Addressing: ImpliedAddressing, Timing: 2},             // 0x98
	{Instruction: sta, Addressing: AbsoluteYAddressing, Timing: 5},           // 0x99
	{Instruction: txs, Addressing: ImpliedAddressing, Timing: 2},             // 0x9a
	{}, // 0x9b
	{}, // 0x9c
	{Instruction: sta, Addressing: AbsoluteXAddressing, Timing: 5}, // 0x9d
	{}, // 0x9e
	{}, // 0x9f
	{Instruction: ldy, Addressing: ImmediateAddressing, Timing: 2},           // 0xa0
	{Instruction: lda, Addressing: IndirectXAddressing, Timing: 6},           // 0xa1
	{Instruction: ldx, Addressing: ImmediateAddressing, Timing: 2},           // 0xa2
	{Instruction: unofficialLax, Addressing: IndirectXAddressing, Timing: 6}, // 0xa3
	{Instruction: ldy, Addressing: ZeroPageAddressing, Timing: 3},            // 0xa4
	{Instruction: lda, Addressing: ZeroPageAddressing, Timing: 3},            // 0xa5
	{Instruction: ldx, Addressing: ZeroPageAddressing, Timing: 3},            // 0xa6
	{Instruction: unofficialLax, Addressing: ZeroPageAddressing, Timing: 3},  // 0xa7
	{Instruction: tay, Addressing: ImpliedAddressing, Timing: 2},             // 0xa8
	{Instruction: lda, Addressing: ImmediateAddressing, Timing: 2},           // 0xa9
	{Instruction: tax, Addressing: ImpliedAddressing, Timing: 2},             // 0xaa
	{}, // 0xab
	{Instruction: ldy, Addressing: AbsoluteAddressing, Timing: 4},                        // 0xac
	{Instruction: lda, Addressing: AbsoluteAddressing, Timing: 4},                        // 0xad
	{Instruction: ldx, Addressing: AbsoluteAddressing, Timing: 4},                        // 0xae
	{Instruction: unofficialLax, Addressing: AbsoluteAddressing, Timing: 4},              // 0xaf
	{Instruction: bcs, Addressing: RelativeAddressing, Timing: 2},                        // 0xb0
	{Instruction: lda, Addressing: IndirectYAddressing, Timing: 5, PageCrossCycle: true}, // 0xb1
	{}, // 0xb2
	{Instruction: unofficialLax, Addressing: IndirectYAddressing, Timing: 5, PageCrossCycle: true}, // 0xb3
	{Instruction: ldy, Addressing: ZeroPageXAddressing, Timing: 4},                                 // 0xb4
	{Instruction: lda, Addressing: ZeroPageXAddressing, Timing: 4},                                 // 0xb5
	{Instruction: ldx, Addressing: ZeroPageYAddressing, Timing: 4},                                 // 0xb6
	{Instruction: unofficialLax, Addressing: ZeroPageYAddressing, Timing: 4},                       // 0xb7
	{Instruction: clv, Addressing: ImpliedAddressing, Timing: 2},                                   // 0xb8
	{Instruction: lda, Addressing: AbsoluteYAddressing, Timing: 4, PageCrossCycle: true},           // 0xb9
	{Instruction: tsx, Addressing: ImpliedAddressing, Timing: 2},                                   // 0xba
	{}, // 0xbb
	{Instruction: ldy, Addressing: AbsoluteXAddressing, Timing: 4, PageCrossCycle: true}, // 0xbc
	{Instruction: lda, Addressing: AbsoluteXAddressing, Timing: 4, PageCrossCycle: true}, // 0xbd
	{Instruction: ldx, Addressing: AbsoluteYAddressing, Timing: 4, PageCrossCycle: true}, // 0xbe
	{Instruction: unofficialLax, Addressing: AbsoluteYAddressing, Timing: 4},             // 0xbf
	{Instruction: cpy, Addressing: ImmediateAddressing, Timing: 2},                       // 0xc0
	{Instruction: cmp, Addressing: IndirectXAddressing, Timing: 6},                       // 0xc1
	{Instruction: unofficialNop, Addressing: ImmediateAddressing, Timing: 2},             // 0xc2
	{Instruction: unofficialDcp, Addressing: IndirectXAddressing, Timing: 8},             // 0xc3
	{Instruction: cpy, Addressing: ZeroPageAddressing, Timing: 3},                        // 0xc4
	{Instruction: cmp, Addressing: ZeroPageAddressing, Timing: 3},                        // 0xc5
	{Instruction: dec, Addressing: ZeroPageAddressing, Timing: 5},                        // 0xc6
	{Instruction: unofficialDcp, Addressing: ZeroPageAddressing, Timing: 5},              // 0xc7
	{Instruction: iny, Addressing: ImpliedAddressing, Timing: 2},                         // 0xc8
	{Instruction: cmp, Addressing: ImmediateAddressing, Timing: 2},                       // 0xc9
	{Instruction: dex, Addressing: ImpliedAddressing, Timing: 2},                         // 0xca
	{}, // 0xcb
	{Instruction: cpy, Addressing: AbsoluteAddressing, Timing: 4},                        // 0xcc
	{Instruction: cmp, Addressing: AbsoluteAddressing, Timing: 4},                        // 0xcd
	{Instruction: dec, Addressing: AbsoluteAddressing, Timing: 6},                        // 0xce
	{Instruction: unofficialDcp, Addressing: AbsoluteAddressing, Timing: 6},              // 0xcf
	{Instruction: bne, Addressing: RelativeAddressing, Timing: 2},                        // 0xd0
	{Instruction: cmp, Addressing: IndirectYAddressing, Timing: 5, PageCrossCycle: true}, // 0xd1
	{}, // 0xd2
	{Instruction: unofficialDcp, Addressing: IndirectYAddressing, Timing: 8},                       // 0xd3
	{Instruction: unofficialNop, Addressing: ZeroPageXAddressing, Timing: 4},                       // 0xd4
	{Instruction: cmp, Addressing: ZeroPageXAddressing, Timing: 4},                                 // 0xd5
	{Instruction: dec, Addressing: ZeroPageXAddressing, Timing: 6},                                 // 0xd6
	{Instruction: unofficialDcp, Addressing: ZeroPageXAddressing, Timing: 6},                       // 0xd7
	{Instruction: cld, Addressing: ImpliedAddressing, Timing: 2},                                   // 0xd8
	{Instruction: cmp, Addressing: AbsoluteYAddressing, Timing: 4, PageCrossCycle: true},           // 0xd9
	{Instruction: unofficialNop, Addressing: ImpliedAddressing, Timing: 2},                         // 0xda
	{Instruction: unofficialDcp, Addressing: AbsoluteYAddressing, Timing: 7},                       // 0xdb
	{Instruction: unofficialNop, Addressing: AbsoluteXAddressing, Timing: 4, PageCrossCycle: true}, // 0xdc
	{Instruction: cmp, Addressing: AbsoluteXAddressing, Timing: 4, PageCrossCycle: true},           // 0xdd
	{Instruction: dec, Addressing: AbsoluteXAddressing, Timing: 7},                                 // 0xde
	{Instruction: unofficialDcp, Addressing: AbsoluteXAddressing, Timing: 7},                       // 0xdf
	{Instruction: cpx, Addressing: ImmediateAddressing, Timing: 2},                                 // 0xe0
	{Instruction: sbc, Addressing: IndirectXAddressing, Timing: 6},                                 // 0xe1
	{Instruction: unofficialNop, Addressing: ImmediateAddressing, Timing: 2},                       // 0xe2
	{Instruction: unofficialIsb, Addressing: IndirectXAddressing, Timing: 8},                       // 0xe3
	{Instruction: cpx, Addressing: ZeroPageAddressing, Timing: 3},                                  // 0xe4
	{Instruction: sbc, Addressing: ZeroPageAddressing, Timing: 3},                                  // 0xe5
	{Instruction: inc, Addressing: ZeroPageAddressing, Timing: 5},                                  // 0xe6
	{Instruction: unofficialIsb, Addressing: ZeroPageAddressing, Timing: 5},                        // 0xe7
	{Instruction: inx, Addressing: ImpliedAddressing, Timing: 2},                                   // 0xe8
	{Instruction: sbc, Addressing: ImmediateAddressing, Timing: 2},                                 // 0xe9
	{Instruction: nop, Addressing: ImpliedAddressing, Timing: 2},                                   // 0xea
	{Instruction: unofficialSbc, Addressing: ImmediateAddressing, Timing: 2},                       // 0xeb
	{Instruction: cpx, Addressing: AbsoluteAddressing, Timing: 4},                                  // 0xec
	{Instruction: sbc, Addressing: AbsoluteAddressing, Timing: 4},                                  // 0xed
	{Instruction: inc, Addressing: AbsoluteAddressing, Timing: 6},                                  // 0xee
	{Instruction: unofficialIsb, Addressing: AbsoluteAddressing, Timing: 6},                        // 0xef
	{Instruction: beq, Addressing: RelativeAddressing, Timing: 2},                                  // 0xf0
	{Instruction: sbc, Addressing: IndirectYAddressing, Timing: 5, PageCrossCycle: true},           // 0xf1
	{}, // 0xf2
	{Instruction: unofficialIsb, Addressing: IndirectYAddressing, Timing: 8},                       // 0xf3
	{Instruction: unofficialNop, Addressing: ZeroPageXAddressing, Timing: 4},                       // 0xf4
	{Instruction: sbc, Addressing: ZeroPageXAddressing, Timing: 4},                                 // 0xf5
	{Instruction: inc, Addressing: ZeroPageXAddressing, Timing: 6},                                 // 0xf6
	{Instruction: unofficialIsb, Addressing: ZeroPageXAddressing, Timing: 6},                       // 0xf7
	{Instruction: sed, Addressing: ImpliedAddressing, Timing: 2},                                   // 0xf8
	{Instruction: sbc, Addressing: AbsoluteYAddressing, Timing: 4, PageCrossCycle: true},           // 0xf9
	{Instruction: unofficialNop, Addressing: ImpliedAddressing, Timing: 2},                         // 0xfa
	{Instruction: unofficialIsb, Addressing: AbsoluteYAddressing, Timing: 7},                       // 0xfb
	{Instruction: unofficialNop, Addressing: AbsoluteXAddressing, Timing: 4, PageCrossCycle: true}, // 0xfc
	{Instruction: sbc, Addressing: AbsoluteXAddressing, Timing: 4, PageCrossCycle: true},           // 0xfd
	{Instruction: inc, Addressing: AbsoluteXAddressing, Timing: 7, PageCrossCycle: true},           // 0xfe
	{Instruction: unofficialIsb, Addressing: AbsoluteXAddressing, Timing: 7},                       // 0xff
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
