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
	ImmediateAddressing: paramReaderImmediate,
	AbsoluteAddressing:  paramReaderAbsolute,
	AbsoluteXAddressing: paramReaderAbsolute,
	AbsoluteYAddressing: paramReaderAbsolute,
	ZeroPageAddressing:  paramReaderZeroPage,
	ZeroPageXAddressing: paramReaderZeroPage,
	RelativeAddressing:  paramReaderRelative,
	IndirectXAddressing: paramReaderIndirectX,
	IndirectYAddressing: paramReaderIndirectY,
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

	switch addressing {
	case AbsoluteXAddressing, ZeroPageXAddressing:
		params = append(params, *X)
	case AbsoluteYAddressing:
		params = append(params, *Y)
	}

	return params, opcodes
}

func paramReaderImmediate(sys *system.System) ([]interface{}, []byte) {
	b := sys.ReadMemory(*PC)
	*PC++

	params := []interface{}{int(b)}
	opcodes := []byte{b}
	return params, opcodes
}

func paramReaderAbsolute(sys *system.System) ([]interface{}, []byte) {
	b1 := uint16(sys.ReadMemory(*PC))
	*PC++
	b2 := uint16(sys.ReadMemory(*PC))
	*PC++

	params := []interface{}{Absolute(b2<<8 | b1)}
	opcodes := []byte{byte(b1), byte(b2)}
	return params, opcodes
}

func paramReaderZeroPage(sys *system.System) ([]interface{}, []byte) {
	b := sys.ReadMemory(*PC)
	*PC++

	params := []interface{}{Absolute(b)}
	opcodes := []byte{b}
	return params, opcodes
}

func paramReaderRelative(sys *system.System) ([]interface{}, []byte) {
	offset := uint16(sys.ReadMemory(*PC))
	*PC++

	var address uint16
	if offset < 0x80 {
		address = *PC + offset
	} else {
		address = *PC + offset - 0x100
	}

	params := []interface{}{Absolute(address)}
	opcodes := []byte{byte(offset)}
	return params, opcodes
}

func paramReaderIndirectX(sys *system.System) ([]interface{}, []byte) {
	b := sys.ReadMemory(*PC)
	*PC++
	offset := uint16(b + *X)
	addr := sys.ReadMemory16(offset)

	params := []interface{}{Absolute(addr)}
	opcodes := []byte{b}
	return params, opcodes
}

func paramReaderIndirectY(sys *system.System) ([]interface{}, []byte) {
	b := sys.ReadMemory(*PC)
	*PC++
	addr := sys.ReadMemory16(uint16(b))
	addr += uint16(*Y)

	params := []interface{}{Absolute(addr)}
	opcodes := []byte{b}
	return params, opcodes
}
