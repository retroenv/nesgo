// Package main implements a Golang for NES Compiler
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/retroenv/nesgo/internal/buildinfo"
	"github.com/retroenv/nesgo/internal/compiler"
	"github.com/retroenv/nesgo/pkg/disasm/ca65"
)

type optionFlags struct {
	input  string
	output string

	quiet bool
}

func main() {
	options := readArguments()

	if !options.quiet {
		printBanner(options)
		fmt.Printf("Compiling %s\n", options.input)
	}

	if err := compileFile(options); err != nil {
		fmt.Println(fmt.Errorf("error: %w", err))
		os.Exit(1)
	}

	if !options.quiet {
		fmt.Printf("Output file %s created successfully\n", options.output)
	}
}

func readArguments() optionFlags {
	flags := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	options := optionFlags{}

	flags.StringVar(&options.output, "o", "", "name of the output .nes file")
	flags.BoolVar(&options.quiet, "q", false, "perform operations quietly")

	err := flags.Parse(os.Args[1:])
	args := flags.Args()
	if err != nil || len(args) == 0 || options.output == "" {
		printBanner(options)
		fmt.Printf("usage: nesgo [options] <file to compile>\n\n")
		flags.PrintDefaults()
		os.Exit(1)
	}

	options.input = args[0]
	return options
}

func printBanner(options optionFlags) {
	if !options.quiet {
		fmt.Println("[---------------------------------]")
		fmt.Println("[ nesgo - Golang for NES Compiler ]")
		fmt.Printf("[---------------------------------]\n\n")
		fmt.Printf("version: %s\n\n", buildinfo.BuildVersion(version, commit, date))
	}
}

func compileFile(options optionFlags) error {
	cfg := &compiler.Config{}
	c, err := compiler.New(cfg)
	if err != nil {
		return fmt.Errorf("creating compiler: %w", err)
	}

	data, err := os.ReadFile(options.input)
	if err != nil {
		return fmt.Errorf("reading file: %w", err)
	}
	if err = c.Parse(options.input, data); err != nil {
		return fmt.Errorf("parsing file '%s': %w", options.input, err)
	}

	asmFile, objectFile, err := c.OutputAsmFile(options.output)
	if err != nil {
		return fmt.Errorf("compiling to file '%s' failed: %w", options.output, err)
	}

	// TODO pass real options
	ca65Config := ca65.Config{
		PRGSize: 0x8000,
		CHRSize: 0x2000,
	}

	if err = ca65.AssembleUsingExternalApp(asmFile, objectFile, options.output, ca65Config); err != nil {
		return fmt.Errorf("creating .nes file '%s' failed: %w", options.output, err)
	}

	return nil
}
