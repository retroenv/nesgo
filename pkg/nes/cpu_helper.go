//go:build !nesgo
// +build !nesgo

package nes

import (
	"fmt"
	"math"
)

type cPU struct {
	A uint8 // accumulator
	X uint8 // x register
	Y uint8 // y register
}

func setFlags(flags uint8) {
	C = (flags >> 0) & 1
	Z = (flags >> 1) & 1
	I = (flags >> 2) & 1
	D = (flags >> 3) & 1
	B = (flags >> 4) & 1
	U = (flags >> 5) & 1
	V = (flags >> 6) & 1
	N = (flags >> 7) & 1
}

// setZ - set the zero flag if the argument is zero.
func setZ(value uint8) {
	if value == 0 {
		Z = 1
	} else {
		Z = 0
	}
}

// setN - set the negative flag if the argument is negative (high bit is set).
func setN(value uint8) {
	if value&0x80 != 0 {
		N = 1
	} else {
		N = 0
	}
}

func setZN(value uint8) {
	setZ(value)
	setN(value)
}

func compare(a, b byte) {
	setZN(a - b)
	if a >= b {
		C = 1
	} else {
		C = 0
	}
}

func readMemoryAddressModes(param interface{}, reg ...interface{}) byte {
	switch val := param.(type) {
	case int:
		if val <= math.MaxUint8 {
			return uint8(val) // immediate, not an address
		}
		return readMemoryAbsolute(val, reg...)
	case uint8:
		return val // immediate, not an address
	case *uint8: // variable
		return *val
	case Absolute:
		return readMemoryAbsolute(val, reg...)
	}
	panic(fmt.Sprintf("unsupported memory read addressing mode type %T", param))
}

func writeMemoryAddressModes(param interface{}, value byte, reg ...interface{}) {
	switch val := param.(type) {
	case int:
		writeMemoryAbsolute(val, value, reg...)
	case Absolute:
		writeMemoryAbsolute(val, value, reg...)
	case Indirect:
		writeMemoryIndirect(val, value, reg...)
	default:
		panic(fmt.Sprintf("unsupported memory write addressing mode type %T", param))
	}
}
