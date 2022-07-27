package mapperdb_test

import (
	"testing"

	"github.com/retroenv/nesgo/internal/assert"
	"github.com/retroenv/nesgo/pkg/bus"
	"github.com/retroenv/nesgo/pkg/cartridge"
	"github.com/retroenv/nesgo/pkg/mapper"
	"github.com/retroenv/nesgo/pkg/mapper/mapperdb"
)

func TestMapperGTROM(t *testing.T) {
	prg := make([]byte, 0x8000*2)

	b := &bus.Bus{
		Cartridge: &cartridge.Cartridge{
			PRG: prg,
		},
	}

	base := mapper.NewBase(b)
	m := mapperdb.NewMapperGTROM(base)

	type SetChrRAMMapper interface {
		SetChrRAM(ram []byte)
	}

	gtromMapper, ok := m.(SetChrRAMMapper)
	assert.True(t, ok)
	chr := make([]byte, 0x4000)
	gtromMapper.SetChrRAM(chr)

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
}
