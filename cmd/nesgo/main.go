package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/retroenv/nesgo/internal/compiler"
	"github.com/retroenv/nesgo/pkg/disasm/ca65"
)

type optionFlags struct {
	input  *string
	output *string
	quiet  *bool
}

func main() {
	flags := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	options := optionFlags{
		input:  flags.String("f", "", "go file to parse"),
		output: flags.String("o", "", "name of the output .nes file"),
		quiet:  flags.Bool("q", false, "perform operations quietly"),
	}
	if err := flags.Parse(os.Args[1:]); err != nil || *options.input == "" || *options.output == "" {
		fmt.Println("[ nesgo - Golang for NES Compiler ]")
		fmt.Printf("usage: nesgo [options]\n\n")
		flags.PrintDefaults()
		os.Exit(1)
	}

	if !*options.quiet {
		fmt.Println("[ nesgo - Golang to NES Compiler ]")
		fmt.Printf("Compiling %s\n", *options.input)
	}

	if err := compileFile(options); err != nil {
		fmt.Println(fmt.Errorf("error: %w", err))
		os.Exit(1)
	}

	if !*options.quiet {
		fmt.Printf("Output file %s created successfully\n", *options.output)
	}
}

func compileFile(options optionFlags) error {
	cfg := &compiler.Config{}
	c, err := compiler.New(cfg)
	if err != nil {
		return fmt.Errorf("creating compiler: %w", err)
	}

	data, err := os.ReadFile(*options.input)
	if err != nil {
		return fmt.Errorf("reading file: %w", err)
	}
	if err = c.Parse(*options.input, data); err != nil {
		return fmt.Errorf("parsing file '%s': %w", *options.input, err)
	}

	asmFile, objectFile, err := c.OutputAsmFile(*options.output)
	if err != nil {
		return fmt.Errorf("compiling to file '%s' failed: %w", *options.output, err)
	}

	if err = ca65.AssembleUsingExternalApp(asmFile, objectFile, *options.output); err != nil {
		return fmt.Errorf("creating .nes file '%s' failed: %w", *options.output, err)
	}

	return nil
}
