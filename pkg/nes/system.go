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
			RunEmulatorSteps(sys)
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
	sys.CPU.SetTracing(opts.tracing, opts.tracingTarget)
	cpu.LinkInstructionFuncs(sys.CPU)
	sys.Memory.LinkRegisters(&sys.CPU.X, &sys.CPU.Y, X, Y)

	return sys
}

// RunEmulatorSteps runs the emulator until it is quit.
func RunEmulatorSteps(sys *system.System) {
	for {
		runEmulatorStep(sys)
	}
}

// RunEmulatorUntil runs the emulator until the given address.
func RunEmulatorUntil(sys *system.System, address uint16) {
	for {
		if sys.PC == address {
			return
		}
		runEmulatorStep(sys)
	}
}

func runEmulatorStep(sys *system.System) {
	b := sys.ReadMemory(*PC)
	opcode, ok := cpu.Opcodes[b]
	if !ok {
		err := fmt.Errorf("unsupported opcode %00x", b)
		panic(err)
	}

	oldPC := *PC
	sys.TraceStep = cpu.TraceStep{
		PC:             *PC,
		Opcode:         []byte{b},
		Addressing:     opcode.Addressing,
		Timing:         opcode.Timing,
		PageCrossCycle: opcode.PageCrossCycle,
		PageCrossed:    false,
	}

	ins := opcode.Instruction
	if ins.NoParamFunc != nil {
		ins.NoParamFunc()
		updatePC(sys, ins, oldPC, 1)
		return
	}

	params, opcodes, pageCrossed := ReadOpParams(sys.Memory, opcode.Addressing)
	sys.TraceStep.Opcode = append(sys.TraceStep.Opcode, opcodes...)
	sys.TraceStep.PageCrossed = pageCrossed

	ins.ParamFunc(params...)
	updatePC(sys, ins, oldPC, len(sys.TraceStep.Opcode))
}

func updatePC(sys *system.System, ins *cpu.Instruction, oldPC uint16, amount int) {
	// update PC only if the instruction execution did not change it
	if oldPC == *PC {
		*PC += uint16(amount)
	} else {
		// page crossing is measured based on the start of the instruction that follows the
		// current instruction
		nextAddress := oldPC + uint16(len(sys.TraceStep.Opcode))
		pageCrossed := *PC&0xff00 != nextAddress&0xff00
		if pageCrossed {
			sys.CPU.AccountBranchingPageCrossCycle(ins)
		}
	}
}

func start(sys *system.System) {
	if err := gui.RunRenderer(sys); err != nil {
		panic(err)
	}
}
