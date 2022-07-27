package mapperdb

/*
Boards: CNROM "and similar"
PRG ROM capacity: 16K or 32K
CHR capacity: 32K (2M oversize version)
CHR window:8K
*/

import (
	"github.com/retroenv/nesgo/pkg/bus"
)

type mapper3 struct {
	Base
}

// NewMapperCNROM returns a new mapper instance.
func NewMapperCNROM(base Base) bus.Mapper {
	m := &mapper3{
		Base: base,
	}
	m.SetName("CNROM")
	m.Initialize()
	m.AddWriteHook(0x8000, 0xFFFF, m.setChrWindow)
	return m
}

func (m *mapper3) setChrWindow(address uint16, value uint8) {
	// Select 8 KB CHR ROM bank for PPU $0000-$1FFF
	banks := m.ChrBankCount()
	bank := int(value) % banks
	m.SetChrWindow(0, bank)
}
