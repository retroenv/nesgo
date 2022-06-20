package ca65

import (
	"errors"
	"fmt"
	"io"
	"sort"

	"github.com/retroenv/nesgo/pkg/disasm/program"
)

var header = `.byte "NES", $1a ; Magic string that always begins an iNES header
`

var headerByte = ".byte $%02x        ; %s\n"

var headerRemainder = `.byte %00000001  ; Vertical mirroring, no save RAM, no mapper
.byte %00000000  ; No special-case flags set, no mapper
.byte $00        ; No PRG-RAM present
.byte $00        ; NTSC format

`

var footer = `
.segment "VECTORS"
.addr %s, %s, %s

.segment "CHARS"
.res 8192
.segment "STARTUP"
`

// FileWriter writes the assembly file content.
type FileWriter struct {
}

// Write writes the assembly file content including header, footer, code and data.
func (f FileWriter) Write(app *program.Program, writer io.Writer) error {
	if err := f.writeSegment(writer, "HEADER"); err != nil {
		return err
	}
	if _, err := fmt.Fprint(writer, header); err != nil {
		return err
	}
	if _, err := fmt.Fprintf(writer, headerByte, len(app.PRG)/16384, "Number of 16KB PRG-ROM banks"); err != nil {
		return err
	}
	if _, err := fmt.Fprintf(writer, headerByte, len(app.CHR)/8192, "Number of 8KB CHR-ROM banks"); err != nil {
		return err
	}
	if _, err := fmt.Fprint(writer, headerRemainder); err != nil {
		return err
	}

	if len(app.Constants) > 0 {
		if err := f.writeConstants(app, writer); err != nil {
			return err
		}
	}

	if err := f.writeCode(app, writer); err != nil {
		return err
	}

	if _, err := fmt.Fprintf(writer, footer, app.Handlers.NMI, app.Handlers.Reset, app.Handlers.IRQ); err != nil {
		return err
	}
	return nil
}

func (f FileWriter) writeSegment(writer io.Writer, name string) error {
	_, err := fmt.Fprintf(writer, ".segment \"%s\"\n", name)
	return err
}

func (f FileWriter) writeConstants(app *program.Program, writer io.Writer) error {
	names := make([]string, 0, len(app.Constants))
	for constant := range app.Constants {
		names = append(names, constant)
	}
	sort.Strings(names)

	for _, constant := range names {
		address := app.Constants[constant]
		if _, err := fmt.Fprintf(writer, "%s = $%04X\n", constant, address); err != nil {
			return err
		}
	}
	_, err := fmt.Fprint(writer, "\n")
	return err
}

func (f FileWriter) writeCode(app *program.Program, writer io.Writer) error {
	if err := f.writeSegment(writer, "CODE"); err != nil {
		return err
	}

	lastNonZeroByte, err := getLastNonZeroByte(app)
	if err != nil {
		return err
	}

	for i := 0; i < lastNonZeroByte; i++ {
		res := app.PRG[i]
		if res.CodeOutput == "" {
			if res.HasData {
				// TODO bundle data outputs
				if _, err := fmt.Fprintf(writer, ".byte $%02x\n", res.Data); err != nil {
					return err
				}
			}
			continue
		}

		if res.Label != "" {
			if res.IsCallTarget {
				if _, err := fmt.Fprintln(writer); err != nil {
					return err
				}
			}
			if _, err := fmt.Fprintf(writer, "%s:\n", res.Label); err != nil {
				return err
			}
		}
		if _, err := fmt.Fprintf(writer, "  %s\n", res.CodeOutput); err != nil {
			return err
		}
	}
	return nil
}

// getLastNonZeroByte searches for the last byte in PRG that is not zero
func getLastNonZeroByte(app *program.Program) (int, error) {
	start := len(app.PRG) - 1 - 6 // skip irq pointers

	for i := start; i > 0; i-- {
		res := app.PRG[i]
		if res.CodeOutput == "" && res.Data == 0 {
			continue
		}
		return i + 1, nil
	}
	return 0, errors.New("could not find last zero byte")
}
