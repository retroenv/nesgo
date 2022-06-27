package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/retroenv/nesgo/internal/buildinfo"
	"github.com/retroenv/nesgo/pkg/cartridge"
	_ "github.com/retroenv/nesgo/pkg/gui"
	"github.com/retroenv/nesgo/pkg/nes"
)

var (
	version = "dev"
	commit  = ""
	date    = ""
)

type optionFlags struct {
	input string

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

	flags.BoolVar(&options.noGui, "c", false, "console mode, disable GUI")
	flags.IntVar(&options.entrypoint, "e", -1, "entrypoint to start the CPU")
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
	fmt.Printf("version: %s\n", buildinfo.BuildVersion(version, commit, date))
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
