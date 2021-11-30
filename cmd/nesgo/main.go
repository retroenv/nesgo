package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"

	"github.com/retroenv/nesgo/internal/compiler"
)

const (
	assembler = "ca65"
	linker    = "ld65"
)

func main() {
	input := flag.String("f", "", "go file to parse")
	output := flag.String("o", "", "name of the output .nes file")

	flag.Parse()

	if *input == "" || *output == "" {
		flag.Usage()
		os.Exit(1)
	}

	fmt.Println("[ goNES - Golang to NES Compiler ]")
	fmt.Printf("Compiling %s\n", *input)

	if err := compileFile(*input, *output); err != nil {
		fmt.Println(fmt.Errorf("error: %w", err))
		os.Exit(1)
	}

	fmt.Printf("Output file %s created successfully\n", *output)
}

func compileFile(input, output string) error {
	cfg := &compiler.Config{}
	c, err := compiler.New(cfg)
	if err != nil {
		return fmt.Errorf("creating compiler: %w", err)
	}

	if err = c.Parse(input); err != nil {
		return fmt.Errorf("parsing file '%s': %w", input, err)
	}

	asmFile, objectFile, err := c.OutputAsmFile(output)
	if err != nil {
		return fmt.Errorf("compiling to file '%s' failed: %w", output, err)
	}

	if err = createNESFile(asmFile, objectFile, output); err != nil {
		return fmt.Errorf("creating .nes file '%s' failed: %w", output, err)
	}

	return nil
}

func createNESFile(asmFile, objectFile, outputFile string) error {
	if _, err := exec.LookPath(assembler); err != nil {
		return fmt.Errorf("%s is not installed", assembler)
	}
	if _, err := exec.LookPath(linker); err != nil {
		return fmt.Errorf("%s is not installed", linker)
	}

	cmd := exec.Command(assembler, asmFile)
	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("assembling file: %s: %w", string(out), err)
	}

	cmd = exec.Command(linker, objectFile, "-t", "nes", "-o", outputFile)
	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("assembling file: %s: %w", string(out), err)
	}

	return nil
}
