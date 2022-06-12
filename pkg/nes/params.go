//go:build !nesgo
// +build !nesgo

package nes

import (
	"fmt"

	. "github.com/retroenv/nesgo/pkg/addressing"
	"github.com/retroenv/nesgo/pkg/system"
)

type paramReaderFunc func(sys *system.System) ([]interface{}, []byte)

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

// readParams reads the opcode parameters after the first opcode byte
// and translates it into emulator specific types.
// It returns the total size of the opcode in bytes.
func readParams(sys *system.System, addressing Mode) ([]interface{}, []byte) {
	fun, ok := paramReader[addressing]
	if !ok {
		err := fmt.Errorf("unsupported addressing mode %00x", addressing)
		panic(err)
	}

	params, opcodes := fun(sys)
	return params, opcodes
}

func paramReaderImplied(sys *system.System) ([]interface{}, []byte) {
	return nil, nil
}

func paramReaderImmediate(sys *system.System) ([]interface{}, []byte) {
	b := sys.ReadMemory(*PC + 1)
	params := []interface{}{int(b)}
	opcodes := []byte{b}
	return params, opcodes
}

func paramReaderAccumulator(sys *system.System) ([]interface{}, []byte) {
	params := []interface{}{Accumulator(0)}
	return params, nil
}

func paramReaderAbsolute(sys *system.System) ([]interface{}, []byte) {
	b1 := uint16(sys.ReadMemory(*PC + 1))
	b2 := uint16(sys.ReadMemory(*PC + 2))

	params := []interface{}{Absolute(b2<<8 | b1)}
	opcodes := []byte{byte(b1), byte(b2)}
	return params, opcodes
}

func paramReaderAbsoluteX(sys *system.System) ([]interface{}, []byte) {
	b1 := uint16(sys.ReadMemory(*PC + 1))
	b2 := uint16(sys.ReadMemory(*PC + 2))
	w := b2<<8 | b1
	offsetAddress(sys, w, *X)

	params := []interface{}{Absolute(w), *X}
	opcodes := []byte{byte(b1), byte(b2)}
	return params, opcodes
}

func paramReaderAbsoluteY(sys *system.System) ([]interface{}, []byte) {
	b1 := uint16(sys.ReadMemory(*PC + 1))
	b2 := uint16(sys.ReadMemory(*PC + 2))
	w := b2<<8 | b1
	offsetAddress(sys, w, *Y)

	params := []interface{}{Absolute(w), *Y}
	opcodes := []byte{byte(b1), byte(b2)}
	return params, opcodes
}

func paramReaderZeroPage(sys *system.System) ([]interface{}, []byte) {
	b := sys.ReadMemory(*PC + 1)

	params := []interface{}{Absolute(b)}
	opcodes := []byte{b}
	return params, opcodes
}

func paramReaderZeroPageX(sys *system.System) ([]interface{}, []byte) {
	b := sys.ReadMemory(*PC + 1)

	params := []interface{}{ZeroPage(b), *X}
	opcodes := []byte{b}
	return params, opcodes
}

func paramReaderZeroPageY(sys *system.System) ([]interface{}, []byte) {
	b := sys.ReadMemory(*PC + 1)

	params := []interface{}{ZeroPage(b), *Y}
	opcodes := []byte{b}
	return params, opcodes
}

func paramReaderRelative(sys *system.System) ([]interface{}, []byte) {
	offset := uint16(sys.ReadMemory(*PC + 1))

	var address uint16
	if offset < 0x80 {
		address = *PC + 2 + offset
	} else {
		address = *PC + 2 + offset - 0x100
	}

	params := []interface{}{Absolute(address)}
	opcodes := []byte{byte(offset)}
	return params, opcodes
}

func paramReaderIndirect(sys *system.System) ([]interface{}, []byte) {
	address := sys.ReadMemory16Bug(*PC + 1)
	b1 := uint16(sys.ReadMemory(*PC + 1))
	b2 := uint16(sys.ReadMemory(*PC + 2))

	params := []interface{}{Indirect(address)}
	opcodes := []byte{byte(b1), byte(b2)}
	return params, opcodes
}

func paramReaderIndirectX(sys *system.System) ([]interface{}, []byte) {
	b := sys.ReadMemory(*PC + 1)
	offset := uint16(b + *X)
	address := sys.ReadMemory16Bug(offset)

	params := []interface{}{Absolute(address)}
	opcodes := []byte{b}
	return params, opcodes
}

func paramReaderIndirectY(sys *system.System) ([]interface{}, []byte) {
	b := sys.ReadMemory(*PC + 1)
	address := sys.ReadMemory16Bug(uint16(b))
	address = offsetAddress(sys, address, *Y)

	params := []interface{}{Absolute(address)}
	opcodes := []byte{b}
	return params, opcodes
}

// offsetAddress returns the offset address and accounts for the resulting
// address crossing a page boundary.
func offsetAddress(sys *system.System, address uint16, offset byte) uint16 {
	newAddress := address + uint16(offset)
	sys.TraceStep.PageCrossed = newAddress&0xff00 != address&0xff00
	return newAddress
}
