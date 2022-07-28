package mapperdb_test

import (
	"testing"

	"github.com/retroenv/nesgo/internal/assert"
	"github.com/retroenv/nesgo/pkg/bus"
	"github.com/retroenv/nesgo/pkg/cartridge"
	"github.com/retroenv/nesgo/pkg/mapper"
	"github.com/retroenv/nesgo/pkg/ppu/nametable"
)

func TestMapperUxROM(t *testing.T) {
	prg := make([]byte, 0xC000)

	b := &bus.Bus{
		Cartridge: &cartridge.Cartridge{
			Mapper: 2,
			CHR:    make([]byte, 0x2000),
			PRG:    prg,
		},
		NameTable: nametable.New(cartridge.MirrorHorizontal),
	}

	m, err := mapper.New(b)
	assert.NoError(t, err)

	prg[0x0010] = 0x03 // bank 0
	prg[0x4010] = 0x04 // bank 1
	prg[0x8010] = 0x05 // bank 2
	assert.Equal(t, 0x03, m.Read(0x8010))
	assert.Equal(t, 0x04, m.Read(0xC010))

	m.Write(0x8000, 2) // select bank 2
	assert.Equal(t, 0x05, m.Read(0x8010))
}
