//go:build !nesgo

package nes

import (
	"io"

	"github.com/retroenv/nesgo/pkg/cpu"
	"github.com/retroenv/retrogolib/arch/nes/cartridge"
)

// Options contains options for the nesgo system.
type Options struct {
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
type Option func(*Options)

// NewOptions creates a new options instance from the passed options.
func NewOptions(optionList ...Option) *Options {
	opts := &Options{
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
func WithCartridge(cart *cartridge.Cartridge) func(*Options) {
	return func(options *Options) {
		options.cartridge = cart
	}
}

// WithEmulator sets the emulator mode.
func WithEmulator() func(*Options) {
	return func(options *Options) {
		options.emulator = true
	}
}

// WithIrqHandler sets an Irq Handler for the program.
func WithIrqHandler(f func()) func(*Options) {
	return func(options *Options) {
		options.irqHandler = f
	}
}

// WithNmiHandler sets a Nmi Handler for the program.
func WithNmiHandler(f func()) func(*Options) {
	return func(options *Options) {
		options.nmiHandler = f
	}
}

// WithDebug enables the debugging mode and webserver.
func WithDebug(debugAddress string) func(*Options) {
	return func(options *Options) {
		options.debug = true
		options.debugAddress = debugAddress
	}
}

// WithTracing enables tracing for the program.
func WithTracing() func(*Options) {
	return func(options *Options) {
		options.tracing = cpu.GoTracing
	}
}

// WithTracingTarget set the tracing target io writer.
func WithTracingTarget(target io.Writer) func(*Options) {
	return func(options *Options) {
		options.tracing = cpu.GoTracing
		options.tracingTarget = target
	}
}

// WithEntrypoint enables tracing for the program.
func WithEntrypoint(address int) func(*Options) {
	return func(options *Options) {
		options.entrypoint = address
	}
}

// WithStopAt stops execution of the program at a specific address.
func WithStopAt(address int) func(*Options) {
	return func(options *Options) {
		options.stopAt = address
	}
}

// WithDisabledGUI disabled the GUI.
func WithDisabledGUI() func(*Options) {
	return func(options *Options) {
		options.noGui = true
	}
}
