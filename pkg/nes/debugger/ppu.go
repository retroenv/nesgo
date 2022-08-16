//go:build !nesgo

package debugger

import (
	"fmt"
	"net/http"
	"strings"
)

func (d *Debugger) ppuPalette(w http.ResponseWriter, r *http.Request) {
	palette := d.bus.PPU.Palette()
	data := palette.Data()

	buf := strings.Builder{}
	fmt.Fprintf(&buf, "background color: %s\n", bytesToHex(data[0:1]))
	fmt.Fprintf(&buf, "background palette 0: %s\n", bytesToHex(data[1:4]))
	fmt.Fprintf(&buf, "background palette 1: %s\n", bytesToHex(data[4:7]))
	fmt.Fprintf(&buf, "background palette 2: %s\n", bytesToHex(data[7:10]))
	fmt.Fprintf(&buf, "sprite palette 0: %s\n", bytesToHex(data[10:13]))
	fmt.Fprintf(&buf, "sprite palette 1: %s\n", bytesToHex(data[13:16]))
	fmt.Fprintf(&buf, "sprite palette 2: %s\n", bytesToHex(data[16:19]))
	fmt.Fprintf(&buf, "sprite palette 3: %s\n", bytesToHex(data[19:22]))
	_, _ = w.Write([]byte(buf.String()))
}

func (d *Debugger) ppuNameTables(w http.ResponseWriter, r *http.Request) {
	tables := d.bus.NameTable.Data()

	buf := strings.Builder{}

	for table := 0; table < 4; table++ {
		fmt.Fprintf(&buf, "nametable %d\n", table)

		data := tables[table]
		for row := 0; row < 30; row++ {
			address := row * 32
			fmt.Fprintf(&buf, "%s\n", bytesToHex(data[address:address+32]))
		}
	}
	_, _ = w.Write([]byte(buf.String()))
}
