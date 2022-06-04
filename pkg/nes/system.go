//go:build !nesgo
// +build !nesgo

package nes

import (
	"github.com/retroenv/nesgo/pkg/cartridge"
	"github.com/retroenv/nesgo/pkg/cpu"
	"github.com/retroenv/nesgo/pkg/gui"
	"github.com/retroenv/nesgo/pkg/system"
)

// InitializeSystem initializes the NES system.
// This needs to be called for any unit code that does not use the Start()
// function, for example in unit tests.
func InitializeSystem(cart *cartridge.Cartridge) *system.System {
	if cart == nil {
		cart = cartridge.New()
	}

	sys := system.New(cart)

	setAliases(sys.CPU)
	A = &sys.CPU.A
	X = &sys.CPU.X
	Y = &sys.CPU.Y
	PC = &sys.CPU.PC

	cpu.LinkInstructionFuncs(sys.CPU)
	sys.Memory.LinkRegisters(&sys.CPU.X, &sys.CPU.Y, X, Y)

	return sys
}

// Start is the main entrypoint for a NES program that starts the execution.
// It expects 1 to 3 parameters for callback function that will be called
// by NES when different events occur:
// resetHandler: called when the system gets turned on or reset
// nmiHandler:   occurs when the PPU starts preparing the next frame of
//               graphics, 60 times per second
// irqHandler:   can be triggered by the NES sound processor or from
//               certain types of cartridge hardware.
func Start(resetHandlerParam func(), nmiIrqHandlers ...func()) {
	sys := InitializeSystem(nil)
	sys.ResetHandler = resetHandlerParam

	if len(nmiIrqHandlers) > 1 {
		sys.IrqHandler = nmiIrqHandlers[1]
	}
	if len(nmiIrqHandlers) > 0 {
		sys.NmiHandler = nmiIrqHandlers[0]
	}

	start(sys)
}

func start(sys *system.System) {
	if err := gui.RunRenderer(sys); err != nil {
		panic(err)
	}
}
