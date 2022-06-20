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

	// TODO: use a temp file
	cmd = exec.Command(linker, "-C", "pkg/disasm/ca65/cfg/nrom.cfg", "-o", outputFile, objectFile)
	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("linking file: %s: %w", string(out), err)
	}

	return nil
}
