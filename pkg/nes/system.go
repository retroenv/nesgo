//go:build !nesgo
// +build !nesgo

package nes

// System implements a NES system.
type System struct {
	*CPU
	*Memory

	nmiHandler   func()
	irqHandler   func()
	resetHandler func()
}

func newSystem() *System {
	memory := newMemory()
	cpu := newCPU(memory)
	memory.x = &cpu.X
	memory.y = &cpu.Y
	return &System{
		CPU:    cpu,
		Memory: memory,
	}
}

// InitializeSystem initializes the NES system.
// This needs to be called for any unit code that does not use the Start()
// function, for example in unit tests.
func InitializeSystem() *System {
	system := newSystem()
	setAliases(system.CPU)
	A = &system.CPU.A
	X = &system.CPU.X
	Y = &system.CPU.Y
	PC = &system.CPU.PC
	*PC = 0x8000
	return system
}

// nolint: unused
var nmiHandler func()

var irqHandler func()

var resetHandler func()

// Start is the main entrypoint for a NES program that starts the execution.
// It expects 1 to 3 parameters for callback function that will be called
// by NES when different events occur:
// resetHandler: called when the system gets turned on or reset
// nmiHandler:   occurs when the PPU starts preparing the next frame of
//               graphics, 60 times per second
// irqHandler:   can be triggered by the NES sound processor or from
//               certain types of cartridge hardware.
func Start(resetHandlerParam func(), nmiIrqHandlers ...func()) {
	system := InitializeSystem()
	system.resetHandler = resetHandlerParam

	if len(nmiIrqHandlers) > 1 {
		system.irqHandler = nmiIrqHandlers[1]
	}
	if len(nmiIrqHandlers) > 0 {
		system.nmiHandler = nmiIrqHandlers[0]
	}

	start(system)
}

func start(system *System) {
	nmiHandler = system.nmiHandler
	irqHandler = system.irqHandler
	resetHandler = system.resetHandler

	if err := runRenderer(system); err != nil {
		panic(err)
	}
}
