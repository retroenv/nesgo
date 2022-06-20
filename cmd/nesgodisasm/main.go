package main

import (
	"bytes"
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

type optionFlags struct {
	input  *string
	output *string
	verify *bool
}

func main() {
	flags := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	options := optionFlags{
		input:  flags.String("f", "", "nes file to load"),
		output: flags.String("o", "", "name of the output .asm file, printed on console if no name given"),
		verify: flags.Bool("v", false, "verify using ca65 that the generated output matches the input"),
	}
	fmt.Printf("[ nesgodisasm - NES program disassembler ]\n\n")
	if err := flags.Parse(os.Args[1:]); err != nil || *options.input == "" {
		fmt.Printf("usage: nesgodisasm [options]\n\n")
		flags.PrintDefaults()
		os.Exit(1)
	}

	if err := disasmFile(options); err != nil {
		fmt.Println(fmt.Errorf("disassembling failed: %w", err))
		os.Exit(1)
	}
}

func disasmFile(options optionFlags) error {
	file, err := os.Open(*options.input)
	if err != nil {
		return fmt.Errorf("opening file '%s': %w", *options.input, err)
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
	if *options.output == "" {
		outputFile = os.Stdout
	} else {
		outputFile, err = os.Create(*options.output)
		if err != nil {
			return fmt.Errorf("creating file '%s': %w", *options.output, err)
		}
	}
	if err = dis.Process(outputFile); err != nil {
		return fmt.Errorf("processing file: %w", err)
	}
	if err = outputFile.Close(); err != nil {
		return fmt.Errorf("closing file: %w", err)
	}

	if *options.verify {
		if err = verifyOutput(options); err != nil {
			return err
		}
		fmt.Println("Output file matched input file.")
	}
	return nil
}

func verifyOutput(options optionFlags) error {
	if *options.output == "" {
		return errors.New("can not verify console output")
	}

	filePart := filepath.Ext(*options.output)
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

	if err = ca65.AssembleUsingExternalApp(*options.output, objectFile.Name(), outputFile.Name()); err != nil {
		return fmt.Errorf("creating .nes file failed: %w", err)
	}

	source, err := os.ReadFile(*options.input)
	if err != nil {
		return fmt.Errorf("reading file for comparison: %w", err)
	}

	destination, err := os.ReadFile(outputFile.Name())
	if err != nil {
		return fmt.Errorf("reading file for comparison: %w", err)
	}

	if err := checkBufferEqual(source, destination); err != nil {
		if detailsErr := compareCartridgeDetails(source, destination); detailsErr != nil {
			return fmt.Errorf("comparing cartridge details: %w", detailsErr)
		}
		return err
	}
	return nil
}

func checkBufferEqual(input, output []byte) error {
	if len(input) != len(output) {
		return fmt.Errorf("mismatched lengths, %d != %d", len(input), len(output))
	}

	var diffs uint64
	firstDiff := -1
	for i := range input {
		if input[i] == output[i] {
			continue
		}
		diffs++
		if firstDiff == -1 {
			firstDiff = i
		}
	}
	if diffs == 0 {
		return nil
	}
	return fmt.Errorf("%d offset mismatches, first at offset %d", diffs, firstDiff)
}

func compareCartridgeDetails(input, output []byte) error {
	inputReader := bytes.NewReader(input)
	outputReader := bytes.NewReader(output)

	cart1, err := cartridge.LoadFile(inputReader)
	if err != nil {
		return err
	}
	cart2, err := cartridge.LoadFile(outputReader)
	if err != nil {
		return err
	}

	if err := checkBufferEqual(cart1.PRG, cart2.PRG); err != nil {
		fmt.Printf("PRG difference: %s\n", err)
	}
	if err := checkBufferEqual(cart1.CHR, cart2.CHR); err != nil {
		fmt.Printf("CHR difference: %s\n", err)
	}
	if err := checkBufferEqual(cart1.Trainer, cart2.Trainer); err != nil {
		fmt.Printf("Trainer difference: %s\n", err)
	}
	if cart1.Mapper != cart2.Mapper {
		fmt.Println("Mapper header does not match")
	}
	if cart1.Mirror != cart2.Mirror {
		fmt.Println("Mirror header does not match")
	}
	if cart1.Battery != cart2.Battery {
		fmt.Println("Battery header does not match")
	}
	return nil
}
