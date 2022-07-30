package mapperdb

/*
Boards: HVC-UN1ROM
PRG ROM capacity: 128K
PRG ROM window: 16K + 16K fixed
CHR capacity: 8K
*/

import (
	"github.com/retroenv/nesgo/pkg/bus"
)

type mapperUN1ROM struct {
	Base
}

// NewMapperUN1ROM returns a new mapper instance.
func NewMapperUN1ROM(base Base) bus.Mapper {
	m := &mapperUN1ROM{
		Base: base,
	}
	m.SetName("UxROM")
	m.Initialize()

	m.AddWriteHook(0x8000, 0xFFFF, m.setPrgWindow)

	m.SetPrgWindow(1, -1) // $C000-$FFFF: 16 KB PRG ROM bank, fixed to the last bank
	return m
}

func (m *mapperUN1ROM) setPrgWindow(address uint16, value uint8) {
	value >>= 2 // it is very similar to UxROM, but the register is shifted by two bits
	value &= 0b00000111
	m.SetPrgWindow(0, int(value)) // select 16 KB PRG ROM bank for CPU $8000-$BFFF
}
