//go:build !nesgo
// +build !nesgo

package nes

import (
	"fmt"

	. "github.com/retroenv/nesgo/pkg/addressing"
	"github.com/retroenv/nesgo/pkg/cartridge"
	"github.com/retroenv/nesgo/pkg/system"
)

// StartEmulator starts emulating the cartridge.
func StartEmulator(cartridge *cartridge.Cartridge) {
	sys := InitializeSystem(cartridge)
	sys.ResetHandler = func() {
		runStep(sys)
	}

	start(sys)
}

func runStep(sys *system.System) {
	for {
		b := sys.ReadMemory(*PC)
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

		params := readParams(sys, ins)
		f := *ins.paramFunc
		f(params...)
	}
}

func readParams(sys *system.System, ins instruction) []interface{} {
	var params []interface{}

	switch ins.addressing {
	case ImmediateAddressing:
		b := sys.ReadMemory(*PC)
		*PC++
		params = append(params, int(b))

	case AbsoluteAddressing, AbsoluteXAddressing, AbsoluteYAddressing:
		b1 := uint16(sys.ReadMemory(*PC))
		*PC++
		b2 := uint16(sys.ReadMemory(*PC))
		*PC++

		params = append(params, Absolute(b2<<8|b1))

	case ZeroPageAddressing, ZeroPageXAddressing:
		b := sys.ReadMemory(*PC)
		*PC++
		params = append(params, Absolute(b))

	case RelativeAddressing:
		offset := uint16(sys.ReadMemory(*PC))
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
	case AbsoluteXAddressing, ZeroPageXAddressing:
		params = append(params, *X)
	case AbsoluteYAddressing:
		params = append(params, *Y)
	}

	return params
}
