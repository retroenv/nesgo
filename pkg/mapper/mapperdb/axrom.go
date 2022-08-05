package mapperdb

import (
	"github.com/retroenv/nesgo/pkg/bus"
	"github.com/retroenv/nesgo/pkg/cartridge"
	"github.com/retroenv/nesgo/pkg/mapper/mapperbase"
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

// NewAxROM returns a new mapper instance.
func NewAxROM(base Base) bus.Mapper {
	m := &mapperAxROM{
		Base: base,
	}
	m.SetName("AxROM")
	m.SetPrgWindowSize(0x8000) // 32K
	m.Initialize()

	translation := mapperbase.MirrorModeTranslation{
		0: cartridge.MirrorSingle0,
		1: cartridge.MirrorSingle1,
	}
	m.SetMirrorModeTranslation(translation)

	m.AddWriteHook(0x8000, 0xFFFF, m.setPrgWindow)
	return m
}

func (m *mapperAxROM) setPrgWindow(address uint16, value uint8) {
	value &= 0b00000111
	m.SetPrgWindow(0, int(value)) // select 32 KB PRG ROM bank for CPU $8000-$FFFF

	mirrorMode := (value >> 4) & 1
	m.SetNameTableMirrorModeIndex(mirrorMode)
}
