//go:build !nesgo
// +build !nesgo

package nes

import (
	"fmt"

	"github.com/retroenv/nesgo/pkg/cartridge"
	"github.com/retroenv/nesgo/pkg/cpu"
	"github.com/retroenv/nesgo/pkg/gui"
	"github.com/retroenv/nesgo/pkg/system"
)

// Start is the main entrypoint for a NES program that starts the execution.
// Different options can be passed.
// Following callback function that will be called by NES when different events occur:
// resetHandler: called when the system gets turned on or reset
// nmiHandler:   occurs when the PPU starts preparing the next frame of
//               graphics, 60 times per second
// irqHandler:   can be triggered by the NES sound processor or from
//               certain types of cartridge hardware.
func Start(resetHandlerParam func(), options ...Option) {
	opts := NewOptions(options...)
	sys := InitializeSystem(opts)

	if opts.emulator {
		sys.ResetHandler = func() {
			runEmulatorStep(sys)
		}
	} else {
		sys.ResetHandler = resetHandlerParam
	}

	start(sys)
}

// InitializeSystem initializes the NES system.
// This needs to be called for any unit code that does not use the Start()
// function, for example in unit tests.
func InitializeSystem(opts *Options) *system.System {
	if opts.cartridge == nil {
		opts.cartridge = cartridge.New()
	}

	sys := system.New(opts.cartridge)
	if opts.entrypoint >= 0 {
		sys.PC = uint16(opts.entrypoint)
	}

	setAliases(sys.CPU)
	A = &sys.CPU.A
	X = &sys.CPU.X
	Y = &sys.CPU.Y
	PC = &sys.CPU.PC
	sys.CPU.SetTracing(opts.tracing)
	cpu.LinkInstructionFuncs(sys.CPU)
	sys.Memory.LinkRegisters(&sys.CPU.X, &sys.CPU.Y, X, Y)

	return sys
}

func runEmulatorStep(sys *system.System) {
	for {
		b := sys.ReadMemory(*PC)
		ins, ok := cpu.Opcodes[b]
		if !ok {
			err := fmt.Errorf("unsupported opcode %00x", b)
			panic(err)
		}

		sys.TraceStep = cpu.TraceStep{
			PC:         *PC,
			Addressing: ins.Addressing,
		}
		oldPC := *PC


		if ins.Instruction.NoParamFunc != nil {
			sys.TraceStep.Opcode = []byte{b}
			ins.Instruction.NoParamFunc()
			updatePC(oldPC, 1)
			continue
		}

		params, opcodes := readParams(sys, ins.Addressing)

		sys.TraceStep.Opcode = append([]byte{b}, opcodes...)

		ins.Instruction.ParamFunc(params...)
		updatePC(oldPC, len(sys.TraceStep.Opcode))
	}
}

func updatePC(oldPC uint16, amount int) {
	// update PC only if the instruction execution did not change it
	if oldPC == *PC {
		*PC += uint16(amount)
	}
}

func start(sys *system.System) {
	if err := gui.RunRenderer(sys); err != nil {
		panic(err)
	}
}
