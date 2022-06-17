package ca65

import (
	"fmt"
	"io"

	"github.com/retroenv/nesgo/pkg/disasm/program"
)

var header = `.segment "HEADER"
.byte "NES", $1a ; Magic string that always begins an iNES header
`

var headerByte = ".byte $%02x        ; %s\n"

var headerRemainder = `.byte %00000001  ; Vertical mirroring, no save RAM, no mapper
.byte %00000000  ; No special-case flags set, no mapper
.byte $00        ; No PRG-RAM present
.byte $00        ; NTSC format

.segment "CODE"
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

	// TODO output constants

	for i := 0; i < len(app.PRG); i++ {
		res := app.PRG[i]
		if res.Output == "" {
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
		if _, err := fmt.Fprintf(writer, "  %s\n", res.Output); err != nil {
			return err
		}
	}

	if _, err := fmt.Fprintf(writer, footer, app.Handlers.NMI, app.Handlers.Reset, app.Handlers.IRQ); err != nil {
		return err
	}
	return nil
}
