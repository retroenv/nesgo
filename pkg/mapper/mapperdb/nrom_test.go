package mapperdb

import (
	"testing"

	"github.com/retroenv/nesgo/pkg/bus"
	"github.com/retroenv/nesgo/pkg/mapper/mapperbase"
	"github.com/retroenv/nesgo/pkg/ppu/nametable"
	"github.com/retroenv/retrogolib/arch/nes/cartridge"
	"github.com/retroenv/retrogolib/assert"
)

func TestMapperNROMPrg16k(t *testing.T) {
	chr := make([]byte, 0x2000)
	prg := make([]byte, 0x4000)

	base := mapperbase.New(&bus.Bus{
		Cartridge: &cartridge.Cartridge{
			CHR: chr,
			PRG: prg,
		},
		NameTable: nametable.New(cartridge.MirrorHorizontal),
	})
	m := NewNROM(base)

	chr[0x0001] = 0x02 // bank 0
	assert.Equal(t, 0x02, m.Read(0x0001))

	prg[0x0010] = 0x03 // bank 0
	assert.Equal(t, 0x03, m.Read(0x8010))
	assert.Equal(t, 0x03, m.Read(0xC010))
}

func TestMapperNROMPrg32k(t *testing.T) {
	chr := make([]byte, 0x2000)
	prg := make([]byte, 0x8000)

	base := mapperbase.New(&bus.Bus{
		Cartridge: &cartridge.Cartridge{
			CHR: chr,
			PRG: prg,
		},
		NameTable: nametable.New(cartridge.MirrorHorizontal),
	})
	m := NewNROM(base)

	chr[0x0001] = 0x02 // bank 0
	assert.Equal(t, 0x02, m.Read(0x0001))

	prg[0x0010] = 0x03 // bank 0
	prg[0x4010] = 0x04 // bank 1
	assert.Equal(t, 0x03, m.Read(0x8010))
	assert.Equal(t, 0x04, m.Read(0xC010))
}
