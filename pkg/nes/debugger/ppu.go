//go:build !nesgo

package debugger

import (
	"encoding/json"
	"net/http"

	"github.com/retroenv/retrogolib/nes/cartridge"
)

type ppuPaletteBackground struct {
	Color    hexArray `json:"color"`
	Palette0 hexArray `json:"palette0"`
	Palette1 hexArray `json:"palette1"`
	Palette2 hexArray `json:"palette2"`
}

type ppuPaletteSprite struct {
	Palette0 hexArray `json:"palette0"`
	Palette1 hexArray `json:"palette1"`
	Palette2 hexArray `json:"palette2"`
	Palette3 hexArray `json:"palette3"`
}

type ppuPalette struct {
	Background ppuPaletteBackground `json:"background"`
	Sprite     ppuPaletteSprite     `json:"sprite"`
}

func (d *Debugger) ppuPalette(w http.ResponseWriter, r *http.Request) {
	palette := d.bus.PPU.Palette()
	data := palette.Data()

	res := ppuPalette{
		Background: ppuPaletteBackground{
			Color:    data[0:1],
			Palette0: data[1:4],
			Palette1: data[4:7],
			Palette2: data[7:10],
		},
		Sprite: ppuPaletteSprite{
			Palette0: data[10:13],
			Palette1: data[13:16],
			Palette2: data[16:19],
			Palette3: data[19:22],
		},
	}

	_ = json.NewEncoder(w).Encode(res)
}

type ppuNameTables struct {
	NameTable0 []hexArrayCombined `json:"nametable0"`
	NameTable1 []hexArrayCombined `json:"nametable1"`
	NameTable2 []hexArrayCombined `json:"nametable2"`
	NameTable3 []hexArrayCombined `json:"nametable3"`
}

func (d *Debugger) ppuNameTables(w http.ResponseWriter, r *http.Request) {
	tables := d.bus.NameTable.Data()

	tableLen := 30 * 32
	res := ppuNameTables{
		NameTable0: bytesToSliceArrayCombined(tables[0][:tableLen], 30, 32),
		NameTable1: bytesToSliceArrayCombined(tables[1][:tableLen], 30, 32),
		NameTable2: bytesToSliceArrayCombined(tables[2][:tableLen], 30, 32),
		NameTable3: bytesToSliceArrayCombined(tables[3][:tableLen], 30, 32),
	}

	_ = json.NewEncoder(w).Encode(res)
}

type ppuMirrorMode struct {
	MirrorMode cartridge.MirrorMode `json:"mirrorMode"`
}

func (d *Debugger) ppuMirrorMode(w http.ResponseWriter, r *http.Request) {
	mirrorMode := d.bus.NameTable.MirrorMode()

	res := ppuMirrorMode{
		MirrorMode: mirrorMode,
	}

	_ = json.NewEncoder(w).Encode(res)
}
