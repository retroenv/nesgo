package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/retroenv/nesgo/pkg/cartridge"
	"github.com/retroenv/nesgo/pkg/disasm"
	"github.com/retroenv/nesgo/pkg/disasm/ca65"
)

var (
	input  = flag.String("f", "", "nes file to load")
	output = flag.String("o", "", "name of the output .asm file, printed on console if no name given")
	verify = flag.Bool("v", false, "verify using ca65 that the generated output matches the input")
)

func main() {
	flag.Parse()

	if *input == "" {
		fmt.Printf("nesgodisasm is a tool for deassembling NES programs.\n\n")
		fmt.Printf("usage: nesgodisasm [options]\n\n")
		flag.CommandLine.PrintDefaults()
		os.Exit(1)
	}

	if err := disasmFile(); err != nil {
		fmt.Println(fmt.Errorf("disassembling failed: %w", err))
		os.Exit(1)
	}
}

func disasmFile() error {
	file, err := os.Open(*input)
	if err != nil {
		return fmt.Errorf("opening file '%s': %w", *input, err)
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
	if *output == "" {
		outputFile = os.Stdout
	} else {
		outputFile, err = os.Create(*output)
		if err != nil {
			return fmt.Errorf("creating file '%s': %w", *output, err)
		}
	}
	if err = dis.Process(outputFile); err != nil {
		return fmt.Errorf("processing file: %w", err)
	}
	if err = outputFile.Close(); err != nil {
		return fmt.Errorf("closing file: %w", err)
	}

	if *verify {
		return verifyOutput()
	}
	return nil
}

func verifyOutput() error {
	if *output == "" {
		return errors.New("can not verify console output")
	}

	filePart := filepath.Ext(*output)
	objectFile, err := os.CreateTemp("", filePart+".*.o")
	if err != nil {
		return err
	}
	defer func() {
		_ = os.Remove(objectFile.Name())
	}()

	outputFile, err := os.CreateTemp("", filePart+".*.nes")
	if err != nil {
		return err
	}
	defer func() {
		_ = os.Remove(outputFile.Name())
	}()

	if err = ca65.AssembleUsingExternalApp(*output, objectFile.Name(), outputFile.Name()); err != nil {
		return fmt.Errorf("creating .nes file failed: %w", err)
	}

	source, err := os.ReadFile(*input)
	if err != nil {
		return fmt.Errorf("reading file for comparison: %w", err)
	}

	destination, err := os.ReadFile(outputFile.Name())
	if err != nil {
		return fmt.Errorf("reading file for comparison: %w", err)
	}

	return checkEqual(source, destination)
}

func checkEqual(bs1, bs2 []byte) error {
	if len(bs1) != len(bs2) {
		return fmt.Errorf("mismatched lengths, %d != %d", len(bs1), len(bs2))
	}

	var byteDiffs uint64
	for i := range bs1 {
		if bs1[i] != bs2[i] {
			byteDiffs++
		}
	}
	if byteDiffs == 0 {
		return nil
	}
	return fmt.Errorf("%d offset mismatches", byteDiffs)
}
