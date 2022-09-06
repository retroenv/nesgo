// Package main implements a NES ROM emulator
package main

import (
	"flag"
	"fmt"
	"os"

	_ "github.com/retroenv/nesgo/pkg/gui"
	"github.com/retroenv/nesgo/pkg/nes"
	"github.com/retroenv/retrogolib/buildinfo"
	"github.com/retroenv/retrogolib/nes/cartridge"
)

type optionFlags struct {
	input string

	debug        bool
	debugAddress string

	entrypoint int
	noGui      bool
	stopAt     int
	tracing    bool
}

func main() {
	options := readArguments()

	if err := emulateFile(options); err != nil {
		fmt.Println(fmt.Errorf("emulation failed: %w", err))
		os.Exit(1)
	}
}

func readArguments() optionFlags {
	flags := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	options := optionFlags{}

	flags.BoolVar(&options.debug, "d", false, "start built-in webserver for debug mode")
	flags.StringVar(&options.debugAddress, "a", "127.0.0.1:8080", "listening address for the debug server to use")
	flags.IntVar(&options.entrypoint, "e", -1, "entrypoint to start the CPU")
	flags.BoolVar(&options.noGui, "c", false, "console mode, disable GUI")
	flags.IntVar(&options.stopAt, "s", -1, "stop execution at address")
	flags.BoolVar(&options.tracing, "t", false, "print CPU tracing")

	err := flags.Parse(os.Args[1:])
	args := flags.Args()
	if err != nil || len(args) == 0 {
		printBanner()
		fmt.Printf("usage: nesgoemu [options] <file to emulate>\n\n")
		flags.PrintDefaults()
		os.Exit(1)
	}
	options.input = args[0]

	return options
}

func printBanner() {
	fmt.Println("[-----------------------------]")
	fmt.Println("[ nesgoemu - NES ROM emulator ]")
	fmt.Printf("[-----------------------------]\n\n")
	fmt.Printf("version: %s\n\n", buildinfo.Version(version, commit, date))
}

func emulateFile(options optionFlags) error {
	file, err := os.Open(options.input)
	if err != nil {
		return fmt.Errorf("opening file '%s': %w", options.input, err)
	}

	cart, err := cartridge.LoadFile(file)
	if err != nil {
		return fmt.Errorf("reading file: %w", err)
	}
	_ = file.Close()

	opts := []nes.Option{
		nes.WithEmulator(),
		nes.WithCartridge(cart),
	}

	if options.debug {
		opts = append(opts, nes.WithDebug(options.debugAddress))
	}
	if options.tracing {
		opts = append(opts, nes.WithTracing())
	}
	if options.entrypoint >= 0 {
		opts = append(opts, nes.WithEntrypoint(options.entrypoint))
	}
	if options.stopAt >= 0 {
		opts = append(opts, nes.WithStopAt(options.stopAt))
	}
	if options.noGui {
		opts = append(opts, nes.WithDisabledGUI())
	}

	nes.Start(nil, opts...)
	return nil
}
