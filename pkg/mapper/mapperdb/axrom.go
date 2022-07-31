package mapperdb

import (
	"github.com/retroenv/nesgo/pkg/bus"
	"github.com/retroenv/nesgo/pkg/cartridge"
)

/*
Boards: AMROM, ANROM, AN1ROM, AOROM, others
PRG ROM capacity: 256K
PRG ROM window: 32K
CHR capacity: 8K
*/

type mapperAxROM struct {
	Base
}

// NewMapperAxROM returns a new mapper instance.
func NewMapperAxROM(base Base) bus.Mapper {
	m := &mapperAxROM{
		Base: base,
	}
	m.SetName("AxROM")
	m.SetPrgWindowSize(0x8000) // 32K
	m.Initialize()

	m.AddWriteHook(0x8000, 0xFFFF, m.setPrgWindow)
	return m
}

func (m *mapperAxROM) setPrgWindow(address uint16, value uint8) {
	value &= 0b00000111
	m.SetPrgWindow(0, int(value)) // select 32 KB PRG ROM bank for CPU $8000-$FFFF

	if (value>>4)&1 == 0 {
		m.SetNameTableMirrorMode(cartridge.MirrorSingle0)
	} else {
		m.SetNameTableMirrorMode(cartridge.MirrorSingle1)
	}
}
