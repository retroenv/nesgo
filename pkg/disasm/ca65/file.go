package ca65

import (
	"errors"
	"fmt"
	"io"
	"sort"

	"github.com/retroenv/nesgo/pkg/cartridge"
	"github.com/retroenv/nesgo/pkg/disasm/program"
)

var header = `.byte "NES", $1a ; Magic string that always begins an iNES header
`

var headerByte = ".byte $%02x        ; %s\n"

var headerRemainder = `.byte $00        ; No PRG-RAM present
.byte $00        ; NTSC format

`

var footer = `
.segment "VECTORS"
.addr %s, %s, %s
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

	control1, control2 := cartridge.ControlBytes(app.Battery, app.Mirror, app.Mapper, len(app.Trainer) > 0)
	if _, err := fmt.Fprintf(writer, headerByte, control1, "control bits"); err != nil {
		return err
	}
	if _, err := fmt.Fprintf(writer, headerByte, control2, "control bits"); err != nil {
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

	if err := f.writeCHR(app, writer); err != nil {
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

func (f FileWriter) writeCHR(app *program.Program, writer io.Writer) error {
	if err := f.writeSegment(writer, "TILES"); err != nil {
		return err
	}

	lastNonZeroByte := getLastNonZeroCHRByte(app)
	for i := 0; i < lastNonZeroByte; i++ {
		b := app.CHR[i]
		// TODO bundle data outputs
		if _, err := fmt.Fprintf(writer, ".byte $%02x\n", b); err != nil {
			return err
		}
	}

	if _, err := fmt.Fprintln(writer); err != nil {
		return err
	}
	return nil
}
func (f FileWriter) writeCode(app *program.Program, writer io.Writer) error {
	if err := f.writeSegment(writer, "CODE"); err != nil {
		return err
	}

	lastNonZeroByte, err := getLastNonZeroPRGByte(app)
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

	if _, err := fmt.Fprintln(writer); err != nil {
		return err
	}
	return nil
}

// getLastNonZeroPRGByte searches for the last byte in PRG that is not zero
func getLastNonZeroPRGByte(app *program.Program) (int, error) {
	start := len(app.PRG) - 1 - 6 // skip irq pointers

	for i := start; i >= 0; i-- {
		res := app.PRG[i]
		if res.CodeOutput == "" && res.Data == 0 {
			continue
		}
		return i + 1, nil
	}
	return 0, errors.New("could not find last zero byte")
}

// getLastNonZeroCHRByte searches for the last byte in CHR that is not zero
func getLastNonZeroCHRByte(app *program.Program) int {
	for i := len(app.CHR) - 1; i >= 0; i-- {
		b := app.CHR[i]
		if b == 0 {
			continue
		}
		return i + 1
	}
	return 0
}
