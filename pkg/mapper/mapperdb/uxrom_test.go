package mapperdb

import (
	"testing"

	"github.com/retroenv/nesgo/pkg/bus"
	"github.com/retroenv/nesgo/pkg/mapper/mapperbase"
	"github.com/retroenv/nesgo/pkg/ppu/nametable"
	"github.com/retroenv/retrogolib/arch/nes/cartridge"
	"github.com/retroenv/retrogolib/assert"
)

func TestMapperUxROMOr(t *testing.T) {
	prg := make([]byte, 0xC000)

	base := mapperbase.New(&bus.Bus{
		Cartridge: &cartridge.Cartridge{
			CHR: make([]byte, 0x2000),
			PRG: prg,
		},
		NameTable: nametable.New(cartridge.MirrorHorizontal),
	})
	m := NewUxROMOr(base)

	prg[0x0010] = 0x03 // bank 0
	prg[0x4010] = 0x04 // bank 1
	prg[0x8010] = 0x05 // bank 2
	assert.Equal(t, 0x03, m.Read(0x8010))
	assert.Equal(t, 0x05, m.Read(0xC010))

	m.Write(0x8000, 1) // select bank 1
	assert.Equal(t, 0x04, m.Read(0x8010))
}

func TestMapperUxROMAnd(t *testing.T) {
	prg := make([]byte, 0xC000)

	base := mapperbase.New(&bus.Bus{
		Cartridge: &cartridge.Cartridge{
			CHR: make([]byte, 0x2000),
			PRG: prg,
		},
		NameTable: nametable.New(cartridge.MirrorHorizontal),
	})
	m := NewUxROMAnd(base)

	prg[0x0010] = 0x03 // bank 0
	prg[0x4010] = 0x04 // bank 1
	prg[0x8010] = 0x05 // bank 2
	assert.Equal(t, 0x03, m.Read(0x8010))
	assert.Equal(t, 0x04, m.Read(0xC010))

	m.Write(0x8000, 2) // select bank 2
	assert.Equal(t, 0x05, m.Read(0xC010))
}

func TestMapperUN1ROM(t *testing.T) {
	prg := make([]byte, 0xC000)

	base := mapperbase.New(&bus.Bus{
		Cartridge: &cartridge.Cartridge{
			CHR: make([]byte, 0x2000),
			PRG: prg,
		},
		NameTable: nametable.New(cartridge.MirrorHorizontal),
	})
	m := NewUN1ROM(base)

	prg[0x0010] = 0x03 // bank 0
	prg[0x4010] = 0x04 // bank 1
	prg[0x8010] = 0x05 // bank 2
	assert.Equal(t, 0x03, m.Read(0x8010))
	assert.Equal(t, 0x05, m.Read(0xC010))

	m.Write(0x8000, 1<<2) // select bank 1
	assert.Equal(t, 0x04, m.Read(0x8010))
}
