package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/retroenv/nesgo/pkg/cartridge"
	"github.com/retroenv/nesgo/pkg/disasm"
)

func main() {
	input := flag.String("f", "", "nes file to load")

	flag.Parse()

	if *input == "" {
		flag.Usage()
		os.Exit(1)
	}

	if err := disasmFile(*input); err != nil {
		fmt.Println(fmt.Errorf("disassembling failed: %w", err))
		os.Exit(1)
	}
}

func disasmFile(input string) error {
	file, err := os.Open(input)
	if err != nil {
		return fmt.Errorf("opening file '%s': %w", input, err)
	}

	cart, err := cartridge.LoadFile(file)
	if err != nil {
		return fmt.Errorf("reading file: %w", err)
	}
	_ = file.Close()

	dis := disasm.New(cart)
	return dis.Process()
}
