//go:build !nesgo

package nes

import (
	"io"

	"github.com/retroenv/nesgo/pkg/cartridge"
	"github.com/retroenv/nesgo/pkg/cpu"
)

// options contains options for the nesgo system.
type options struct {
	entrypoint int
	stopAt     int

	debug        bool
	debugAddress string

	emulator  bool
	noGui     bool
	cartridge *cartridge.Cartridge

	tracing       cpu.TracingMode
	tracingTarget io.Writer

	nmiHandler func()
	irqHandler func()
}

// Option defines a Start parameter.
type Option func(*options)

// newOptions creates a new options instance from the passed options.
func newOptions(optionList ...Option) *options {
	opts := &options{
		entrypoint: -1,
		stopAt:     -1,
	}
	for _, option := range optionList {
		option(opts)
	}

	if opts.emulator && opts.tracing != cpu.NoTracing {
		opts.tracing = cpu.EmulatorTracing
	}

	return opts
}

// WithCartridge sets a cartridge to load.
func WithCartridge(cart *cartridge.Cartridge) func(*options) {
	return func(options *options) {
		options.cartridge = cart
	}
}

// WithEmulator sets the emulator mode.
func WithEmulator() func(*options) {
	return func(options *options) {
		options.emulator = true
	}
}

// WithIrqHandler sets an Irq Handler for the program.
func WithIrqHandler(f func()) func(*options) {
	return func(options *options) {
		options.irqHandler = f
	}
}

// WithNmiHandler sets a Nmi Handler for the program.
func WithNmiHandler(f func()) func(*options) {
	return func(options *options) {
		options.nmiHandler = f
	}
}

// WithDebug enables the debugging mode and webserver.
func WithDebug(debugAddress string) func(*options) {
	return func(options *options) {
		options.debug = true
		options.debugAddress = debugAddress
	}
}

// WithTracing enables tracing for the program.
func WithTracing() func(*options) {
	return func(options *options) {
		options.tracing = cpu.GoTracing
	}
}

// WithTracingTarget set the tracing target io writer.
func WithTracingTarget(target io.Writer) func(*options) {
	return func(options *options) {
		options.tracing = cpu.GoTracing
		options.tracingTarget = target
	}
}

// WithEntrypoint enables tracing for the program.
func WithEntrypoint(address int) func(*options) {
	return func(options *options) {
		options.entrypoint = address
	}
}

// WithStopAt stops execution of the program at a specific address.
func WithStopAt(address int) func(*options) {
	return func(options *options) {
		options.stopAt = address
	}
}

// WithDisabledGUI disabled the GUI.
func WithDisabledGUI() func(*options) {
	return func(options *options) {
		options.noGui = true
	}
}
