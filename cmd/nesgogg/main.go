// Package main implements a NES Game Genie decoder/encoder
package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/retroenv/nesgo/internal/buildinfo"
	"github.com/retroenv/nesgo/pkg/gamegenie"
)

type optionFlags struct {
	code    string
	address string
	value   string
	compare string
}

func main() {
	options := readArguments()

	if options.code != "" {
		if err := decode(options); err != nil {
			fmt.Println(fmt.Errorf("decoding failed: %w", err))
			os.Exit(1)
		}
		return
	}

	if err := encode(options); err != nil {
		fmt.Println(fmt.Errorf("encoding failed: %w", err))
		os.Exit(1)
	}
}

func readArguments() optionFlags {
	flags := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	options := optionFlags{}

	flags.StringVar(&options.address, "a", "", "address to patch in decimal or hex with 0x prefix")
	flags.StringVar(&options.value, "v", "", "value to write in decimal or hex with 0x prefix")
	flags.StringVar(&options.compare, "c", "", "compare value in decimal or hex with 0x prefix")

	err := flags.Parse(os.Args[1:])
	// nolint:ifshort
	args := flags.Args()

	if err != nil || (len(args) == 0 && options.address == "") {
		printBanner()
		fmt.Printf("usage: nesgogg [options] <code>\n\n")
		flags.PrintDefaults()
		os.Exit(1)
	}
	if len(args) > 0 {
		options.code = args[0]
	}

	return options
}

func printBanner() {
	fmt.Println("[------------------------------------------]")
	fmt.Println("[ nesgogg - NES Game Genie decoder/encoder ]")
	fmt.Printf("[------------------------------------------]\n\n")
	fmt.Printf("version: %s\n\n", buildinfo.BuildVersion(version, commit, date))
}

func decode(options optionFlags) error {
	patch, err := gamegenie.Decode(options.code)
	if err != nil {
		return fmt.Errorf("decoding code: %w", err)
	}

	fmt.Printf("address: 0x%04X\n", patch.Address)
	fmt.Printf("value: 0x%02X\n", patch.Data)

	if patch.HasCompare {
		fmt.Printf("compare: 0x%02X\n", patch.Compare)
	}

	return nil
}

func encode(options optionFlags) error {
	address, err := strconv.ParseUint(options.address, 0, 16)
	if err != nil {
		return fmt.Errorf("parsing address: %w", err)
	}

	value, err := strconv.ParseUint(options.value, 0, 8)
	if err != nil {
		return fmt.Errorf("parsing value: %w", err)
	}

	patch := gamegenie.Patch{
		Address: uint16(address),
		Data:    byte(value),
	}

	if options.compare != "" {
		patch.HasCompare = true
		compare, err := strconv.ParseUint(options.compare, 0, 8)
		if err != nil {
			return fmt.Errorf("parsing compare: %w", err)
		}
		patch.Compare = byte(compare)
	}

	code, err := gamegenie.Encode(patch)
	if err != nil {
		return fmt.Errorf("encoding code: %w", err)
	}
	fmt.Printf("Game Genie code: %s\n", code)
	return nil
}
