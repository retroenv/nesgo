package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/retroenv/nesgo/pkg/cartridge"
	"github.com/retroenv/nesgo/pkg/nes"
)

type optionFlags struct {
	input      *string
	entrypoint *int
	tracing    *bool
}

func main() {
	flags := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	options := optionFlags{
		input:      flag.String("f", "", "nes file to load"),
		entrypoint: flag.Int("e", -1, "entrypoint to start the CPU"),
		tracing:    flag.Bool("t", false, "print CPU tracing"),
	}
	if err := flags.Parse(os.Args[1:]); err != nil || *options.input == "" {
		fmt.Printf("nesgoemu is a tool for emulating NES programs.\n\n")
		fmt.Printf("usage: nesgoemu [options]\n\n")
		flag.CommandLine.PrintDefaults()
		os.Exit(1)
	}

	if err := emulateFile(options); err != nil {
		fmt.Println(fmt.Errorf("emulation failed: %w", err))
		os.Exit(1)
	}
}

func emulateFile(options optionFlags) error {
	file, err := os.Open(*options.input)
	if err != nil {
		return fmt.Errorf("opening file '%s': %w", *options.input, err)
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

	if *options.tracing {
		opts = append(opts, nes.WithTracing())
	}
	if *options.entrypoint >= 0 {
		opts = append(opts, nes.WithEntrypoint(*options.entrypoint))
	}

	nes.Start(nil, opts...)
	return nil
}
