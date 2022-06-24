//go:build !nesgo
// +build !nesgo

package nes

import (
	"fmt"
	"sync/atomic"

	"github.com/retroenv/nesgo/pkg/cartridge"
	"github.com/retroenv/nesgo/pkg/cpu"
	"github.com/retroenv/nesgo/pkg/system"
)

type guiInitializer func(sys *system.System) (guiRender func() (bool, error), guiCleanup func(), err error)

// GuiStarter will be set by the chosen and imported GUI renderer.
var GuiStarter guiInitializer

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
			runEmulatorSteps(sys, opts.stopAt)
		}
	} else {
		sys.ResetHandler = resetHandlerParam
	}

	guiStarter := setupNoGui
	if GuiStarter != nil && !opts.noGui {
		guiStarter = GuiStarter
	}
	if err := runRenderer(sys, opts, guiStarter); err != nil {
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

// runEmulatorSteps runs the emulator until it is quit or reaches the given stop address.
func runEmulatorSteps(sys *system.System, stopAt int) {
	for {
		if stopAt >= 0 && sys.PC == uint16(stopAt) {
			return
		}

		oldPC := *PC
		opcode := DecodePCInstruction(sys)

		ins := opcode.Instruction
		if ins.NoParamFunc != nil {
			ins.NoParamFunc()
			updatePC(sys, ins, oldPC, 1)
			continue
		}

		params, opcodes, pageCrossed := ReadOpParams(sys.Memory, opcode.Addressing, true)
		sys.TraceStep.Opcode = append(sys.TraceStep.Opcode, opcodes...)
		sys.TraceStep.PageCrossed = pageCrossed

		ins.ParamFunc(params...)
		updatePC(sys, ins, oldPC, len(sys.TraceStep.Opcode))
	}
}

// DecodePCInstruction decodes the current instruction that the program counter points to.
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
func runRenderer(sys *system.System, opts *Options, guiStarter guiInitializer) error {
	render, cleanup, err := guiStarter(sys)
	if err != nil {
		return err
	}
	defer cleanup()

	running := uint64(1)
	go func() {
		sys.ResetHandler()
		if opts.stopAt >= 0 {
			atomic.StoreUint64(&running, 0)
			return
		}
		for { // forever loop in case reset handler returns
		}
	}()

	for atomic.LoadUint64(&running) == 1 {
		sys.PPU.StartRender()

		sys.PPU.RenderScreen()

		continueRunning, err := render()
		if err != nil {
			return err
		}
		if !continueRunning {
			atomic.StoreUint64(&running, 0)
		}

		sys.PPU.FinishRender()
	}
	return nil
}
