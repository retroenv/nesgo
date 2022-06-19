package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/retroenv/nesgo/internal/compiler"
	"github.com/retroenv/nesgo/pkg/disasm/ca65"
)

var (
	input  = flag.String("f", "", "go file to parse")
	output = flag.String("o", "", "name of the output .nes file")
	quiet  = flag.Bool("q", false, "perform operations quietly")
)

func main() {
	flag.Parse()

	if *input == "" || *output == "" {
		fmt.Printf("nesgo is a tool for compiling Go programs to a NES file.\n\n")
		fmt.Printf("usage: nesgo [options]\n\n")
		flag.CommandLine.PrintDefaults()
		os.Exit(1)
	}

	if !*quiet {
		fmt.Println("[ nesgo - Golang to NES Compiler ]")
		fmt.Printf("Compiling %s\n", *input)
	}

	if err := compileFile(); err != nil {
		fmt.Println(fmt.Errorf("error: %w", err))
		os.Exit(1)
	}

	if !*quiet {
		fmt.Printf("Output file %s created successfully\n", *output)
	}
}

func compileFile() error {
	cfg := &compiler.Config{}
	c, err := compiler.New(cfg)
	if err != nil {
		return fmt.Errorf("creating compiler: %w", err)
	}

	data, err := os.ReadFile(*input)
	if err != nil {
		return fmt.Errorf("reading file: %w", err)
	}
	if err = c.Parse(*input, data); err != nil {
		return fmt.Errorf("parsing file '%s': %w", *input, err)
	}

	asmFile, objectFile, err := c.OutputAsmFile(*output)
	if err != nil {
		return fmt.Errorf("compiling to file '%s' failed: %w", *output, err)
	}

	if err = ca65.AssembleUsingExternalApp(asmFile, objectFile, *output); err != nil {
		return fmt.Errorf("creating .nes file '%s' failed: %w", *output, err)
	}

	return nil
}
