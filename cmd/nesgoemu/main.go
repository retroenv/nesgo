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

	flag.Parse()

	if *input == "" {
		flag.Usage()
		os.Exit(1)
	}

	if err := emulateFile(*input); err != nil {
		fmt.Println(fmt.Errorf("error: %w", err))
		os.Exit(1)
	}
}

func emulateFile(input string) error {
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

	nes.StartEmulator(cart)
	return nil
}
