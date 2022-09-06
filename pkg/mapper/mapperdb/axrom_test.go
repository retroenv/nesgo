package mapperdb

import (
	"testing"

	"github.com/retroenv/nesgo/pkg/bus"
	"github.com/retroenv/nesgo/pkg/mapper/mapperbase"
	"github.com/retroenv/nesgo/pkg/ppu/nametable"
	"github.com/retroenv/retrogolib/assert"
	"github.com/retroenv/retrogolib/nes/cartridge"
)

func TestMapperAxROM(t *testing.T) {
	prg := make([]byte, 0x8000*2)

	base := mapperbase.New(&bus.Bus{
		Cartridge: &cartridge.Cartridge{
			CHR: make([]byte, 0x2000),
			PRG: prg,
		},
		NameTable: nametable.New(cartridge.MirrorHorizontal),
	})
	m := NewAxROM(base)

	prg[0x0010] = 0x03 // bank 0
	prg[0x8010] = 0x04 // bank 1
	assert.Equal(t, 0x03, m.Read(0x8010))

	m.Write(0x8000, 1) // select bank 1
	assert.Equal(t, 0x04, m.Read(0x8010))
}
