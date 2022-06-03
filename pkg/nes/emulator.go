//go:build !nesgo
// +build !nesgo

package nes

import (
	"fmt"

	"github.com/retroenv/nesgo/internal/ast"
	"github.com/retroenv/nesgo/pkg/ines"
)

// StartEmulator starts emulating the cartridge.
func StartEmulator(cartridge *ines.Cartridge) {
	system := InitializeSystem()
	system.cartridge = cartridge
	system.resetHandler = func() {
		runStep(system)
	}

	start(system)
}

func runStep(system *System) {
	for {
		b := system.readMemory(*PC)
		*PC++

		// TODO add debug tracing

		ins, ok := instructions[b]
		if !ok {
			err := fmt.Errorf("unsupported opcode %00x", b)
			panic(err)
		}

		if ins.noParamFunc != nil {
			f := *ins.noParamFunc
			f()
			continue
		}

		params := readParams(system, ins)
		f := *ins.paramFunc
		f(params...)
	}
}

func readParams(system *System, ins instruction) []interface{} {
	var params []interface{}

	switch ins.addressing {
	case ast.ImmediateAddressing:
		b := system.readMemory(*PC)
		*PC++
		params = append(params, int(b))

	case ast.AbsoluteAddressing, ast.AbsoluteXAddressing, ast.AbsoluteYAddressing:
		b1 := uint16(system.readMemory(*PC))
		*PC++
		b2 := uint16(system.readMemory(*PC))
		*PC++

		params = append(params, Absolute(b2<<8|b1))

	case ast.ZeroPageAddressing, ast.ZeroPageXAddressing:
		b := system.readMemory(*PC)
		*PC++
		params = append(params, Absolute(b))

	case ast.RelativeAddressing:
		offset := uint16(system.readMemory(*PC))
		*PC++

		var address uint16
		if offset < 0x80 {
			address = *PC + offset
		} else {
			address = *PC + offset - 0x100
		}

		params = append(params, Absolute(address))

	default:
		err := fmt.Errorf("unsupported addressing %00x", ins.addressing)
		panic(err)
	}

	switch ins.addressing {
	case ast.AbsoluteXAddressing, ast.ZeroPageXAddressing:
		params = append(params, *X)
	case ast.AbsoluteYAddressing:
		params = append(params, *Y)
	}

	return params
}
