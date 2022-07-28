//go:build !nesgo
// +build !nesgo

package nes

import (
	"fmt"
	"sync/atomic"
	"time"

	"github.com/retroenv/nesgo/pkg/bus"
	"github.com/retroenv/nesgo/pkg/cartridge"
	"github.com/retroenv/nesgo/pkg/controller"
	"github.com/retroenv/nesgo/pkg/cpu"
	"github.com/retroenv/nesgo/pkg/mapper"
	"github.com/retroenv/nesgo/pkg/memory"
	"github.com/retroenv/nesgo/pkg/ppu"
	"github.com/retroenv/nesgo/pkg/ppu/nametable"
)

// System implements a NES system.
type System struct {
	*cpu.CPU

	Bus *bus.Bus

	NmiHandler   func()
	IrqHandler   func()
	ResetHandler func()
}

// NewSystem creates a new NES system.
func NewSystem(cart *cartridge.Cartridge) *System {
	if cart == nil {
		cart = cartridge.New()
	}

	systemBus := &bus.Bus{
		Cartridge:   cart,
		Controller1: controller.New(),
		Controller2: controller.New(),
		NameTable:   nametable.New(cart.Mirror),
	}
	systemBus.Memory = memory.New(systemBus)

	var err error
	systemBus.Mapper, err = mapper.New(systemBus)
	if err != nil {
		panic(err)
	}

	systemBus.PPU = ppu.New(systemBus)

	sys := &System{
		Bus: systemBus,
	}

	sys.CPU = cpu.New(systemBus, &sys.IrqHandler)
	systemBus.CPU = sys.CPU
	return sys
}

// LinkAliases links the register and CPU instruction globals to the actual instance.
// Can not be used in tests in combination with t.Parallel().
func (sys *System) LinkAliases() {
	setAliases(sys.CPU)
	A = &sys.CPU.A
	X = &sys.CPU.X
	Y = &sys.CPU.Y
	PC = &sys.CPU.PC
	cpu.LinkInstructionFuncs(sys.CPU)
	sys.Bus.Memory.LinkRegisters(&sys.CPU.X, &sys.CPU.Y, X, Y)
}

// DecodeInstructionAtPC decodes the current instruction at the program counter.
func (sys *System) DecodeInstructionAtPC() (cpu.Opcode, error) {
	b := sys.Bus.Memory.Read(*PC)
	opcode, ok := cpu.Opcodes[b]
	if !ok {
		return cpu.Opcode{}, fmt.Errorf("unsupported opcode %00x", b)
	}

	sys.TraceStep = cpu.TraceStep{
		PC:             *PC,
		Opcode:         []byte{b},
		Addressing:     opcode.Addressing,
		Timing:         opcode.Timing,
		PageCrossCycle: opcode.PageCrossCycle,
		PageCrossed:    false,
	}
	return opcode, nil
}

// runEmulatorSteps runs the emulator until it is quit or reaches the given stop address.
func (sys *System) runEmulatorSteps(stopAt int) {
	for {
		if stopAt >= 0 && sys.PC == uint16(stopAt) {
			return
		}

		oldPC := *PC
		opcode, err := sys.DecodeInstructionAtPC()
		if err != nil {
			panic(err)
		}

		ins := opcode.Instruction
		startCycles := sys.CPU.Cycles()
		if ins.NoParamFunc != nil {
			ins.NoParamFunc()
			sys.updatePC(ins, oldPC, 1)
			sys.runPPUSteps(startCycles)
			continue
		}

		params, opcodes, pageCrossed := ReadOpParams(sys.Bus.Memory, opcode.Addressing, true)
		sys.TraceStep.Opcode = append(sys.TraceStep.Opcode, opcodes...)
		sys.TraceStep.PageCrossed = pageCrossed

		ins.ParamFunc(params...)
		sys.updatePC(ins, oldPC, len(sys.TraceStep.Opcode))
		sys.runPPUSteps(startCycles)
	}
}

func (sys *System) runPPUSteps(startCycles uint64) {
	cpuCycles := sys.CPU.Cycles() - startCycles

	ppuCycles := int(cpuCycles) * 3
	for i := 0; i < ppuCycles; i++ {
		sys.Bus.PPU.Step()
	}
}

func (sys *System) updatePC(ins *cpu.Instruction, oldPC uint16, amount int) {
	// update PC only if the instruction execution did not change it
	if oldPC == *PC {
		if ins.Name == cpu.JmpInstruction {
			return // endless loop detected
		}

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
func (sys *System) runRenderer(opts *options, guiStarter guiInitializer) error {
	render, cleanup, err := guiStarter(sys.Bus)
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
		continueRunning, err := render()
		if err != nil {
			return err
		}
		if !continueRunning {
			atomic.StoreUint64(&running, 0)
		}

		// TODO replace with better solution
		time.Sleep(time.Second / ppu.FPS)
	}
	return nil
}
