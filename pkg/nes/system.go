//go:build !nesgo
// +build !nesgo

package nes

// System implements a NES system.
type System struct {
	*CPU
}

// nolint: unused
var nmiHandler func()

// nolint: unused
var irqHandler func()

// nolint: unused
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
	nmiHandler = nil
	irqHandler = nil

	if len(nmiIrqHandlers) > 1 {
		irqHandler = nmiIrqHandlers[1]
	}
	if len(nmiIrqHandlers) > 0 {
		nmiHandler = nmiIrqHandlers[0]
	}

	resetHandler = resetHandlerParam
	if err := runRenderer(); err != nil {
		panic(err)
	}
}
