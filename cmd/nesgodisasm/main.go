package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/retroenv/nesgo/pkg/cartridge"
	"github.com/retroenv/nesgo/pkg/disasm"
)

func main() {
	input := flag.String("f", "", "nes file to load")
	output := flag.String("o", "", "name of the output .asm file, printed on stdout if no name given")

	flag.Parse()

	if *input == "" {
		flag.Usage()
		os.Exit(1)
	}

	if err := disasmFile(*input, *output); err != nil {
		fmt.Println(fmt.Errorf("disassembling failed: %w", err))
		os.Exit(1)
	}
}

func disasmFile(input, output string) error {
	file, err := os.Open(input)
	if err != nil {
		return fmt.Errorf("opening file '%s': %w", input, err)
	}

	cart, err := cartridge.LoadFile(file)
	if err != nil {
		return fmt.Errorf("reading file: %w", err)
	}
	_ = file.Close()

	dis, err := disasm.New(cart, "ca65")
	if err != nil {
		return fmt.Errorf("initializing disassembler: %w", err)
	}

	var outputFile io.WriteCloser
	if output == "" {
		outputFile = os.Stdout
	} else {
		outputFile, err = os.Create(output)
		if err != nil {
			return fmt.Errorf("creating file '%s': %w", output, err)
		}
	}
	if err := dis.Process(outputFile); err != nil {
		return fmt.Errorf("processing file: %w", err)
	}
	return outputFile.Close()
}
