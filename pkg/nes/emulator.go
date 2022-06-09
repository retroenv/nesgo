//go:build !nesgo
// +build !nesgo

package nes

import (
	"fmt"

	"github.com/retroenv/nesgo/pkg/cartridge"
	"github.com/retroenv/nesgo/pkg/cpu"
	"github.com/retroenv/nesgo/pkg/system"
)

// StartEmulator starts emulating the cartridge.
func StartEmulator(cartridge *cartridge.Cartridge, tracing bool) {
	sys := InitializeSystem(cartridge)
	sys.CPU.SetTracing(cpu.EmulatorTracing)
	sys.ResetHandler = func() {
		runStep(sys)
	}

	start(sys)
}

func runStep(sys *system.System) {
	for {
		sys.TraceStep = cpu.TraceStep{
			PC: *PC,
		}

		b := sys.ReadMemory(*PC)
		*PC++

		ins, ok := cpu.Opcodes[b]
		if !ok {
			err := fmt.Errorf("unsupported opcode %00x", b)
			panic(err)
		}

		if ins.Instruction.NoParamFunc != nil {
			sys.TraceStep.Opcode = []byte{b}
			ins.Instruction.NoParamFunc()
			continue
		}

		params, opcodes := readParams(sys, ins.Addressing)

		sys.TraceStep.Opcode = append([]byte{b}, opcodes...)

		ins.Instruction.ParamFunc(params...)
	}
}
