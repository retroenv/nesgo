package mapperdb

/*
Boards: GTROM
PRG ROM capacity: 512K
PRG ROM window: 32K
CHR capacity: 16K
CHR window: 8K

32K CHR RAM used as two 8K CHR RAM and two 8K nametables
*/

import (
	"github.com/retroenv/nesgo/pkg/bus"
)

type mapperGTROM struct {
	Base
}

// NewMapperGTROM returns a new mapper instance.
func NewMapperGTROM(base Base) bus.Mapper {
	m := &mapperGTROM{
		Base: base,
	}
	m.SetName("Cheapocabra (GTROM)")
	m.SetPrgWindowSize(0x8000) // 32K
	m.Base.SetChrRAM(make([]byte, 0x4000))
	m.Initialize()

	m.AddReadHook(0x5000, 0x5FFF, m.getControl)
	m.AddReadHook(0x7000, 0x7FFF, m.getControl)
	m.AddWriteHook(0x5000, 0x5FFF, m.setBanks)
	m.AddWriteHook(0x7000, 0x7FFF, m.setBanks)

	// TODO support nametable RAM

	return m
}

func (m *mapperGTROM) getControl(address uint16) uint8 {
	return 0 // TODO should return open bus value
}

func (m *mapperGTROM) setBanks(address uint16, value uint8) {
	prgBank := value & 0b00001111

	// Select 32 KB PRG ROM bank for CPU $8000-$FFFF
	banks := m.PrgBankCount()
	bank := int(prgBank) % banks
	m.SetPrgWindow(0, bank)

	chrBank := int(value>>4) & 1
	m.SetChrWindow(0, chrBank)

	nameTableBank := int(value>>5) & 1
	_ = nameTableBank // TODO support nametable switching
}

// SetChrRAM sets the CHR RAM and reinitializes the mapper base structures. Used in unit tests.
func (m *mapperGTROM) SetChrRAM(ram []byte) {
	m.Base.SetChrRAM(ram)
	m.Initialize()
}
