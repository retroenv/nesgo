package ca65

import (
	"fmt"
	"os/exec"
)

const (
	assembler = "ca65"
	linker    = "ld65"
)

// AssembleUsingExternalApp calls the external assembler and linker to generate an .nes
// image from the given asm file.
func AssembleUsingExternalApp(asmFile, objectFile, outputFile string) error {
	if _, err := exec.LookPath(assembler); err != nil {
		return fmt.Errorf("%s is not installed", assembler)
	}
	if _, err := exec.LookPath(linker); err != nil {
		return fmt.Errorf("%s is not installed", linker)
	}

	cmd := exec.Command(assembler, asmFile, "-o", objectFile)
	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("assembling file: %s: %w", string(out), err)
	}

	cmd = exec.Command(linker, objectFile, "-t", "nes", "-o", outputFile)
	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("linking file: %s: %w", string(out), err)
	}

	return nil
}
