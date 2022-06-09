//go:build !nesgo
// +build !nesgo

package nes

import (
	"github.com/retroenv/nesgo/pkg/cartridge"
	"github.com/retroenv/nesgo/pkg/cpu"
)

// Options contains options for the nesgo system.
type Options struct {
	emulator  bool
	cartridge *cartridge.Cartridge
	tracing   cpu.TracingMode

	nmiHandler func()
	irqHandler func()
}

// Option defines a Start parameter.
type Option func(*Options)

// NewOptions creates a new options instance from the passed options.
func NewOptions(options ...Option) *Options {
	opts := &Options{}
	for _, option := range options {
		option(opts)
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

// WithTracing enables tracing for the program.
func WithTracing() func(*Options) {
	return func(options *Options) {
		options.tracing = cpu.GoTracing
	}
}
