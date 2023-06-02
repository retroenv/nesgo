package mapperdb

import (
	"testing"

	"github.com/retroenv/nesgo/pkg/bus"
	"github.com/retroenv/nesgo/pkg/mapper/mapperbase"
	"github.com/retroenv/nesgo/pkg/ppu/nametable"
	"github.com/retroenv/retrogolib/arch/nes/cartridge"
	"github.com/retroenv/retrogolib/assert"
)

func TestMapperMMC1(t *testing.T) {
	chr := make([]byte, 0x1000*3) // 4K banks
	prg := make([]byte, 0x4000*3) // 16K banks

	base := mapperbase.New(&bus.Bus{
		Cartridge: &cartridge.Cartridge{
			CHR: chr,
			PRG: prg,
		},
		NameTable: nametable.New(cartridge.MirrorHorizontal),
	})
	m := NewMMC1(base)

	chr[0x0000] = 0x01
	chr[0x2000] = 0x02
	prg[0x0000] = 0x03
	prg[0x8000] = 0x04

	m.Write(0x8000, 0)
	m.Write(0x8000, 1) // select bank 2
	m.Write(0x8000, 0)
	m.Write(0x8000, 0)
	m.Write(0xE000, 0) // set prg bank

	m.Write(0x8000, 0)
	m.Write(0x8000, 1) // select bank 2
	m.Write(0x8000, 0)
	m.Write(0x8000, 0)
	m.Write(0xC000, 0) // set chr 1 bank

	m.Write(0x8000, 1) // mirror mode 1
	m.Write(0x8000, 0)
	m.Write(0x8000, 0)
	m.Write(0x8000, 1) // prg mode 2
	m.Write(0x9000, 1) // chr mode 1, set control

	assert.Equal(t, 0x01, m.Read(0x0000))
	assert.Equal(t, 0x02, m.Read(0x1000))
	assert.Equal(t, 0x03, m.Read(0x8000))
	assert.Equal(t, 0x04, m.Read(0xC000))

	mode := m.MirrorMode()
	assert.Equal(t, cartridge.MirrorSingle1, mode)
}
