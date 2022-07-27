package mapperdb

/*
Boards: UNROM, UOROM
PRG ROM capacity: 256K/4096K
PRG ROM window:16K + 16K fixed
CHR capacity: 8K
*/

import (
	"github.com/retroenv/nesgo/pkg/bus"
)

type mapperUxRom struct {
	Base
}

// NewMapperUxRom returns a new mapper instance.
func NewMapperUxRom(base Base) bus.Mapper {
	m := &mapperUxRom{
		Base: base,
	}
	m.SetName("UxROM")
	m.Initialize()
	m.AddWriteHook(0x8000, 0xFFFF, m.setPrgWindow)
	return m
}

func (m *mapperUxRom) setPrgWindow(address uint16, value uint8) {
	// Select 16 KB PRG ROM bank for CPU $8000-$BFFF
	banks := m.PrgBankCount()
	bank := int(value) % banks
	m.SetPrgWindow(0, bank)
}
