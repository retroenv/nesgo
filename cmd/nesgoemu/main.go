package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/retroenv/nesgo/pkg/cartridge"
	"github.com/retroenv/nesgo/pkg/nes"
)

var (
	input      = flag.String("f", "", "nes file to load")
	entrypoint = flag.Int("e", -1, "entrypoint to start the CPU")
	tracing    = flag.Bool("t", false, "print CPU tracing")
)

func main() {
	flag.Parse()

	if *input == "" {
		fmt.Printf("nesgoemu is a tool for emulating NES programs.\n\n")
		fmt.Printf("usage: nesgoemu [options]\n\n")
		flag.CommandLine.PrintDefaults()
		os.Exit(1)
	}

	if err := emulateFile(*input); err != nil {
		fmt.Println(fmt.Errorf("emulation failed: %w", err))
		os.Exit(1)
	}
}

func emulateFile(input string) error {
	file, err := os.Open(input)
	if err != nil {
		return fmt.Errorf("opening file '%s': %w", input, err)
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

	if *tracing {
		opts = append(opts, nes.WithTracing())
	}
	if *entrypoint >= 0 {
		opts = append(opts, nes.WithEntrypoint(*entrypoint))
	}

	nes.Start(nil, opts...)
	return nil
}
