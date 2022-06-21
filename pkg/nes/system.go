//go:build !nesgo
// +build !nesgo

package nes

import (
	"fmt"

	"github.com/retroenv/nesgo/pkg/cartridge"
	"github.com/retroenv/nesgo/pkg/cpu"
	"github.com/retroenv/nesgo/pkg/system"
)

// GuiStarter will be set by the chosen and imported GUI renderer.
var GuiStarter func(sys *system.System) (guiRender func() (bool, error), guiCleanup func(), err error)

// Start is the main entrypoint for a NES program that starts the execution.
// Different options can be passed.
// Following callback function that will be called by NES when different events occur:
// resetHandler: called when the system gets turned on or reset
// nmiHandler:   occurs when the PPU starts preparing the next frame of
//               graphics, 60 times per second
// irqHandler:   can be triggered by the NES sound processor or from
//               certain types of cartridge hardware.
func Start(resetHandlerParam func(), options ...Option) {
	sys := InitializeSystem(options...)

	opts := NewOptions(options...)
	if opts.emulator {
		sys.ResetHandler = func() {
			RunEmulatorSteps(sys)
		}
	} else {
		sys.ResetHandler = resetHandlerParam
	}

	if GuiStarter == nil {
		GuiStarter = setupNoGui
	}
	if err := runRenderer(sys); err != nil {
		panic(err)
	}
}

// InitializeSystem initializes the NES system.
// This needs to be called for any unit code that does not use the Start()
// function, for example in unit tests.
func InitializeSystem(options ...Option) *system.System {
	opts := NewOptions(options...)
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
	oldPC := *PC
	opcode := DecodePCInstruction(sys)

	ins := opcode.Instruction
	if ins.NoParamFunc != nil {
		ins.NoParamFunc()
		updatePC(sys, ins, oldPC, 1)
		return
	}

	params, opcodes, pageCrossed := ReadOpParams(sys.Memory, opcode.Addressing, true)
	sys.TraceStep.Opcode = append(sys.TraceStep.Opcode, opcodes...)
	sys.TraceStep.PageCrossed = pageCrossed

	ins.ParamFunc(params...)
	updatePC(sys, ins, oldPC, len(sys.TraceStep.Opcode))
}

// DecodePCInstruction decodes the current instruction that
// the program counter points to.
func DecodePCInstruction(sys *system.System) cpu.Opcode {
	b := sys.ReadMemory(*PC)
	opcode, ok := cpu.Opcodes[b]
	if !ok {
		err := fmt.Errorf("unsupported opcode %00x", b)
		panic(err)
	}

	sys.TraceStep = cpu.TraceStep{
		PC:             *PC,
		Opcode:         []byte{b},
		Addressing:     opcode.Addressing,
		Timing:         opcode.Timing,
		PageCrossCycle: opcode.PageCrossCycle,
		PageCrossed:    false,
	}
	return opcode
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

// runRenderer starts the chosen GUI renderer.
func runRenderer(sys *system.System) error {
	render, cleanup, err := GuiStarter(sys)
	if err != nil {
		return err
	}
	defer cleanup()

	go func() {
		sys.ResetHandler()
		for { // forever loop in case reset handler returns
		}
	}()

	running := true
	for running {
		sys.PPU.StartRender()

		sys.PPU.RenderScreen()

		running, err = render()
		if err != nil {
			return err
		}

		sys.PPU.FinishRender()
	}
	return nil
}
