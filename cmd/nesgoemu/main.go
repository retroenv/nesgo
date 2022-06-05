package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/retroenv/nesgo/pkg/cartridge"
	"github.com/retroenv/nesgo/pkg/nes"
)

func main() {
	input := flag.String("f", "", "nes file to load")
	tracing := flag.Bool("t", false, "print CPU tracing")

	flag.Parse()

	if *input == "" {
		flag.Usage()
		os.Exit(1)
	}

	if err := emulateFile(*input, *tracing); err != nil {
		fmt.Println(fmt.Errorf("emulation failed: %w", err))
		os.Exit(1)
	}
}

func emulateFile(input string, tracing bool) error {
	file, err := os.Open(input)
	if err != nil {
		return fmt.Errorf("opening file '%s': %w", input, err)
	}

	defer func() {
		_ = file.Close()
	}()

	cart, err := cartridge.LoadFile(file)
	if err != nil {
		return fmt.Errorf("reading file: %w", err)
	}

	nes.StartEmulator(cart, tracing)
	return nil
}
