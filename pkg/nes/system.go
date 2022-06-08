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

// Option defines a Start parameter.
type Option func(*system.System)

// WithIrqHandler sets an Irq Handler for the program.
func WithIrqHandler(f func()) func(*system.System) {
	return func(sys *system.System) {
		sys.IrqHandler = f
	}
}

// WithNmiHandler sets a Nmi Handler for the program.
func WithNmiHandler(f func()) func(*system.System) {
	return func(sys *system.System) {
		sys.NmiHandler = f
	}
}

// WithTracing enables tracing for the program.
func WithTracing() func(*system.System) {
	return func(sys *system.System) {
		sys.SetTracing(true)
	}
}

// Start is the main entrypoint for a NES program that starts the execution.
// Different options can be passed.
// Following callback function that will be called by NES when different events occur:
// resetHandler: called when the system gets turned on or reset
// nmiHandler:   occurs when the PPU starts preparing the next frame of
//               graphics, 60 times per second
// irqHandler:   can be triggered by the NES sound processor or from
//               certain types of cartridge hardware.
func Start(resetHandlerParam func(), options ...Option) {
	sys := InitializeSystem(nil)
	sys.ResetHandler = resetHandlerParam

	for _, option := range options {
		option(sys)
	}

	start(sys)
}

func start(sys *system.System) {
	if err := gui.RunRenderer(sys); err != nil {
		panic(err)
	}
}
