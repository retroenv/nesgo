//go:build !nesgo
// +build !nesgo

package nes

// Start is the main entrypoint for a NES program that starts the execution.
// Different options can be passed.
// Following callback function that will be called by NES when different events occur:
// resetHandler: called when the system gets turned on or reset
// nmiHandler:   occurs when the PPU starts preparing the next frame of
//               graphics, 60 times per second
// irqHandler:   can be triggered by the NES sound processor or from
//               certain types of cartridge hardware.
func Start(resetHandlerParam func(), options ...Option) {
	opts := newOptions(options...)
	sys := NewSystem(opts.cartridge)
	if opts.entrypoint >= 0 {
		sys.PC = uint16(opts.entrypoint)
	}

	sys.LinkAliases()

	sys.CPU.SetTracing(opts.tracing, opts.tracingTarget)

	if opts.emulator {
		sys.ResetHandler = func() {
			sys.runEmulatorSteps(opts.stopAt)
		}
	} else {
		sys.ResetHandler = resetHandlerParam
		sys.CPU.SetResetHandlerTraceInfo(resetHandlerParam)
	}

	guiStarter := setupNoGui
	if GuiStarter != nil && !opts.noGui {
		guiStarter = GuiStarter
	}
	if err := sys.runRenderer(opts, guiStarter); err != nil {
		panic(err)
	}
}
