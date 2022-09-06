package mapperdb

import (
	"testing"

	"github.com/retroenv/nesgo/pkg/bus"
	"github.com/retroenv/nesgo/pkg/mapper/mapperbase"
	"github.com/retroenv/nesgo/pkg/ppu/nametable"
	"github.com/retroenv/retrogolib/assert"
	"github.com/retroenv/retrogolib/nes/cartridge"
)

func TestMapperGTROM(t *testing.T) {
	prg := make([]byte, 0x8000*2)

	nameTable := nametable.New(cartridge.Mirror4)
	base := mapperbase.New(&bus.Bus{
		Cartridge: &cartridge.Cartridge{
			PRG: prg,
		},
		NameTable: nameTable,
	})
	m := NewGTROM(base)

	chr := make([]byte, 0x4000)
	base.SetChrRAM(chr)
	base.Initialize()

	prg[0x7010] = 0x03 // bank 0
	prg[0xF010] = 0x04 // bank 1
	assert.Equal(t, 0x03, m.Read(0xF010))

	m.Write(0x5000, 1) // select bank 1

	assert.Equal(t, 0x04, m.Read(0xF010))

	chr[0x1010] = 0x03 // bank 0
	chr[0x3010] = 0x04 // bank 1
	assert.Equal(t, 0x03, m.Read(0x1010))

	m.Write(0x5000, 1<<4) // select bank 1
	assert.Equal(t, 0x04, m.Read(0x1010))

	data := base.NameTable(0)
	data[0x0100] = 0x05 // bank 0
	data = base.NameTable(1)
	data[0x0100] = 0x06 // bank 1

	assert.Equal(t, 0x05, nameTable.Read(0x2100))
	m.Write(0x5000, 1<<5) // select bank 1
	assert.Equal(t, 0x06, nameTable.Read(0x2100))
}
