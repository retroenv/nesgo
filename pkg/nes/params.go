//go:build !nesgo

package nes

import (
	"fmt"

	. "github.com/retroenv/nesgo/pkg/addressing"
)

type paramMemReader interface {
	Read(address uint16) byte
	ReadWordBug(address uint16) uint16
}

type paramReaderFunc func(mem paramMemReader, resolveIndirect bool) ([]any, []byte, bool)

var paramReader = map[Mode]paramReaderFunc{
	ImpliedAddressing:     paramReaderImplied,
	ImmediateAddressing:   paramReaderImmediate,
	AccumulatorAddressing: paramReaderAccumulator,
	AbsoluteAddressing:    paramReaderAbsolute,
	AbsoluteXAddressing:   paramReaderAbsoluteX,
	AbsoluteYAddressing:   paramReaderAbsoluteY,
	ZeroPageAddressing:    paramReaderZeroPage,
	ZeroPageXAddressing:   paramReaderZeroPageX,
	ZeroPageYAddressing:   paramReaderZeroPageY,
	RelativeAddressing:    paramReaderRelative,
	IndirectAddressing:    paramReaderIndirect,
	IndirectXAddressing:   paramReaderIndirectX,
	IndirectYAddressing:   paramReaderIndirectY,
}

// ReadOpParams reads the opcode parameters after the first opcode byte
// and translates it into emulator specific types.
// resolveIndirect specifies if indirect addresses should be resolved,
// for Disassembler usage this is not wanted but for Emulator usage.
func ReadOpParams(mem paramMemReader, addressing Mode, resolveIndirect bool) ([]any, []byte, bool) {
	fun, ok := paramReader[addressing]
	if !ok {
		err := fmt.Errorf("unsupported addressing mode %00x", addressing)
		panic(err)
	}

	params, opcodes, pageCrossed := fun(mem, resolveIndirect)
	return params, opcodes, pageCrossed
}

func paramReaderImplied(mem paramMemReader, resolveIndirect bool) ([]any, []byte, bool) {
	return nil, nil, false
}

func paramReaderImmediate(mem paramMemReader, resolveIndirect bool) ([]any, []byte, bool) {
	b := mem.Read(*PC + 1)
	params := []any{int(b)}
	opcodes := []byte{b}
	return params, opcodes, false
}

func paramReaderAccumulator(mem paramMemReader, resolveIndirect bool) ([]any, []byte, bool) {
	params := []any{Accumulator(0)}
	return params, nil, false
}

func paramReaderAbsolute(mem paramMemReader, resolveIndirect bool) ([]any, []byte, bool) {
	b1 := uint16(mem.Read(*PC + 1))
	b2 := uint16(mem.Read(*PC + 2))

	params := []any{Absolute(b2<<8 | b1)}
	opcodes := []byte{byte(b1), byte(b2)}
	return params, opcodes, false
}

func paramReaderAbsoluteX(mem paramMemReader, resolveIndirect bool) ([]any, []byte, bool) {
	b1 := uint16(mem.Read(*PC + 1))
	b2 := uint16(mem.Read(*PC + 2))
	w := b2<<8 | b1
	_, pageCrossed := offsetAddress(w, *X)

	params := []any{Absolute(w), *X}
	opcodes := []byte{byte(b1), byte(b2)}
	return params, opcodes, pageCrossed
}

func paramReaderAbsoluteY(mem paramMemReader, resolveIndirect bool) ([]any, []byte, bool) {
	b1 := uint16(mem.Read(*PC + 1))
	b2 := uint16(mem.Read(*PC + 2))
	w := b2<<8 | b1
	_, pageCrossed := offsetAddress(w, *Y)

	params := []any{Absolute(w), *Y}
	opcodes := []byte{byte(b1), byte(b2)}
	return params, opcodes, pageCrossed
}

func paramReaderZeroPage(mem paramMemReader, resolveIndirect bool) ([]any, []byte, bool) {
	b := mem.Read(*PC + 1)

	params := []any{Absolute(b)}
	opcodes := []byte{b}
	return params, opcodes, false
}

func paramReaderZeroPageX(mem paramMemReader, resolveIndirect bool) ([]any, []byte, bool) {
	b := mem.Read(*PC + 1)

	params := []any{ZeroPage(b), X}
	opcodes := []byte{b}
	return params, opcodes, false
}

func paramReaderZeroPageY(mem paramMemReader, resolveIndirect bool) ([]any, []byte, bool) {
	b := mem.Read(*PC + 1)

	params := []any{ZeroPage(b), Y}
	opcodes := []byte{b}
	return params, opcodes, false
}

func paramReaderRelative(mem paramMemReader, resolveIndirect bool) ([]any, []byte, bool) {
	offset := uint16(mem.Read(*PC + 1))

	var address uint16
	if offset < 0x80 {
		address = *PC + 2 + offset
	} else {
		address = *PC + 2 + offset - 0x100
	}

	params := []any{Absolute(address)}
	opcodes := []byte{byte(offset)}
	return params, opcodes, false
}

func paramReaderIndirect(mem paramMemReader, resolveIndirect bool) ([]any, []byte, bool) {
	address := mem.ReadWordBug(*PC + 1)
	b1 := uint16(mem.Read(*PC + 1))
	b2 := uint16(mem.Read(*PC + 2))

	params := []any{Indirect(address)}
	opcodes := []byte{byte(b1), byte(b2)}
	return params, opcodes, false
}

func paramReaderIndirectX(mem paramMemReader, resolveIndirect bool) ([]any, []byte, bool) {
	b := mem.Read(*PC + 1)
	offset := uint16(b + *X)

	address := uint16(b)
	var params []any
	if resolveIndirect {
		address = mem.ReadWordBug(offset)
		params = []any{IndirectResolved(address), X}
	} else {
		params = []any{Indirect(address), X}
	}

	opcodes := []byte{b}
	return params, opcodes, false
}

func paramReaderIndirectY(mem paramMemReader, resolveIndirect bool) ([]any, []byte, bool) {
	b := mem.Read(*PC + 1)

	var pageCrossed bool
	address := uint16(b)
	var params []any
	if resolveIndirect {
		address = mem.ReadWordBug(uint16(b))
		address, pageCrossed = offsetAddress(address, *Y)
		params = []any{IndirectResolved(address), Y}
	} else {
		params = []any{Indirect(address), Y}
	}

	opcodes := []byte{b}
	return params, opcodes, pageCrossed
}

// offsetAddress returns the offset address and whether it crosses a page boundary.
func offsetAddress(address uint16, offset byte) (uint16, bool) {
	newAddress := address + uint16(offset)
	pageCrossed := newAddress&0xff00 != address&0xff00
	return newAddress, pageCrossed
}
