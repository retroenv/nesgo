package mapperdb_test

import (
	"testing"

	"github.com/retroenv/nesgo/internal/assert"
	"github.com/retroenv/nesgo/pkg/bus"
	"github.com/retroenv/nesgo/pkg/cartridge"
	"github.com/retroenv/nesgo/pkg/mapper"
	"github.com/retroenv/nesgo/pkg/ppu/nametable"
)

func TestMapperCNROM(t *testing.T) {
	chr := make([]byte, 0x6000)

	b := &bus.Bus{
		Cartridge: &cartridge.Cartridge{
			Mapper: 3,
			CHR:    chr,
			PRG:    make([]byte, 0x4000),
		},
		NameTable: nametable.New(cartridge.MirrorHorizontal),
	}

	m, err := mapper.New(b)
	assert.NoError(t, err)

	chr[0x0010] = 0x03 // bank 0
	chr[0x2010] = 0x04 // bank 1
	chr[0x4010] = 0x05 // bank 2

	assert.Equal(t, 0x03, m.Read(0x0010))

	m.Write(0x8000, 1) // select bank 1
	assert.Equal(t, 0x04, m.Read(0x0010))

	m.Write(0x8000, 2) // select bank 2
	assert.Equal(t, 0x05, m.Read(0x0010))
}
