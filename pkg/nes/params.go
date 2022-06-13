//go:build !nesgo
// +build !nesgo

package nes

import (
	"fmt"

	. "github.com/retroenv/nesgo/pkg/addressing"
)

type paramMemReader interface {
	ReadMemory(address uint16) byte
	ReadMemory16Bug(address uint16) uint16
}

type paramReaderFunc func(mem paramMemReader) ([]interface{}, []byte, bool)

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
func ReadOpParams(mem paramMemReader, addressing Mode) ([]interface{}, []byte, bool) {
	fun, ok := paramReader[addressing]
	if !ok {
		err := fmt.Errorf("unsupported addressing mode %00x", addressing)
		panic(err)
	}

	params, opcodes, pageCrossed := fun(mem)
	return params, opcodes, pageCrossed
}

func paramReaderImplied(mem paramMemReader) ([]interface{}, []byte, bool) {
	return nil, nil, false
}

func paramReaderImmediate(mem paramMemReader) ([]interface{}, []byte, bool) {
	b := mem.ReadMemory(*PC + 1)
	params := []interface{}{int(b)}
	opcodes := []byte{b}
	return params, opcodes, false
}

func paramReaderAccumulator(mem paramMemReader) ([]interface{}, []byte, bool) {
	params := []interface{}{Accumulator(0)}
	return params, nil, false
}

func paramReaderAbsolute(mem paramMemReader) ([]interface{}, []byte, bool) {
	b1 := uint16(mem.ReadMemory(*PC + 1))
	b2 := uint16(mem.ReadMemory(*PC + 2))

	params := []interface{}{Absolute(b2<<8 | b1)}
	opcodes := []byte{byte(b1), byte(b2)}
	return params, opcodes, false
}

func paramReaderAbsoluteX(mem paramMemReader) ([]interface{}, []byte, bool) {
	b1 := uint16(mem.ReadMemory(*PC + 1))
	b2 := uint16(mem.ReadMemory(*PC + 2))
	w := b2<<8 | b1
	_, pageCrossed := offsetAddress(w, *X)

	params := []interface{}{Absolute(w), *X}
	opcodes := []byte{byte(b1), byte(b2)}
	return params, opcodes, pageCrossed
}

func paramReaderAbsoluteY(mem paramMemReader) ([]interface{}, []byte, bool) {
	b1 := uint16(mem.ReadMemory(*PC + 1))
	b2 := uint16(mem.ReadMemory(*PC + 2))
	w := b2<<8 | b1
	_, pageCrossed := offsetAddress(w, *Y)

	params := []interface{}{Absolute(w), *Y}
	opcodes := []byte{byte(b1), byte(b2)}
	return params, opcodes, pageCrossed
}

func paramReaderZeroPage(mem paramMemReader) ([]interface{}, []byte, bool) {
	b := mem.ReadMemory(*PC + 1)

	params := []interface{}{Absolute(b)}
	opcodes := []byte{b}
	return params, opcodes, false
}

func paramReaderZeroPageX(mem paramMemReader) ([]interface{}, []byte, bool) {
	b := mem.ReadMemory(*PC + 1)

	params := []interface{}{ZeroPage(b), *X}
	opcodes := []byte{b}
	return params, opcodes, false
}

func paramReaderZeroPageY(mem paramMemReader) ([]interface{}, []byte, bool) {
	b := mem.ReadMemory(*PC + 1)

	params := []interface{}{ZeroPage(b), *Y}
	opcodes := []byte{b}
	return params, opcodes, false
}

func paramReaderRelative(mem paramMemReader) ([]interface{}, []byte, bool) {
	offset := uint16(mem.ReadMemory(*PC + 1))

	var address uint16
	if offset < 0x80 {
		address = *PC + 2 + offset
	} else {
		address = *PC + 2 + offset - 0x100
	}

	params := []interface{}{Absolute(address)}
	opcodes := []byte{byte(offset)}
	return params, opcodes, false
}

func paramReaderIndirect(mem paramMemReader) ([]interface{}, []byte, bool) {
	address := mem.ReadMemory16Bug(*PC + 1)
	b1 := uint16(mem.ReadMemory(*PC + 1))
	b2 := uint16(mem.ReadMemory(*PC + 2))

	params := []interface{}{Indirect(address)}
	opcodes := []byte{byte(b1), byte(b2)}
	return params, opcodes, false
}

func paramReaderIndirectX(mem paramMemReader) ([]interface{}, []byte, bool) {
	b := mem.ReadMemory(*PC + 1)
	offset := uint16(b + *X)
	address := mem.ReadMemory16Bug(offset)

	params := []interface{}{Absolute(address)}
	opcodes := []byte{b}
	return params, opcodes, false
}

func paramReaderIndirectY(mem paramMemReader) ([]interface{}, []byte, bool) {
	b := mem.ReadMemory(*PC + 1)
	address := mem.ReadMemory16Bug(uint16(b))
	address, pageCrossed := offsetAddress(address, *Y)

	params := []interface{}{Absolute(address)}
	opcodes := []byte{b}
	return params, opcodes, pageCrossed
}

// offsetAddress returns the offset address and whether it crosses a page boundary.
func offsetAddress(address uint16, offset byte) (uint16, bool) {
	newAddress := address + uint16(offset)
	pageCrossed := newAddress&0xff00 != address&0xff00
	return newAddress, pageCrossed
}
