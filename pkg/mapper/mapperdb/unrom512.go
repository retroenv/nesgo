package mapperdb

import (
	"github.com/retroenv/nesgo/pkg/bus"
	"github.com/retroenv/nesgo/pkg/cartridge"
	"github.com/retroenv/nesgo/pkg/mapper/mapperbase"
)

/*
Boards: UNROM-512-8, UNROM-512-16, UNROM-512-32, INL-D-RAM, UNROM-512-F
PRG ROM capacity: 256K/512K
PRG ROM window: 16K + 16K fixed
CHR capacity: 32K
CHR window: 8K
*/

type mapperUNROM512 struct {
	Base
}

// NewUNROM512 returns a new mapper instance.
func NewUNROM512(base Base) bus.Mapper {
	m := &mapperUNROM512{
		Base: base,
	}
	m.SetName("UNROM 512")
	m.SetChrRAM(make([]byte, 0x8000)) // 32K
	m.Initialize()

	m.AddWriteHook(0x8000, 0xFFFF, m.setBanks)

	translation := mapperbase.MirrorModeTranslation{
		0: cartridge.MirrorHorizontal,
		1: cartridge.MirrorVertical,
		2: cartridge.MirrorSingle0,
		3: cartridge.Mirror4,
	}
	m.SetMirrorModeTranslation(translation)

	cart := m.Cartridge()
	m.SetNameTableMirrorModeIndex(uint8(cart.Mirror))

	m.SetPrgWindow(1, -1)
	return m
}

func (m *mapperUNROM512) setBanks(address uint16, value uint8) {
	prgBank := value & 0b00011111

	m.SetPrgWindow(0, int(prgBank)) // select 16 KB PRG ROM bank at $8000

	chrBank := int(value>>5) & 0b00000011
	m.SetChrWindow(0, chrBank)

	screen := int(value>>7) & 1
	if screen == 0 {
		m.SetNameTableMirrorMode(cartridge.MirrorSingle0)
	} else {
		m.SetNameTableMirrorMode(cartridge.MirrorSingle1)
	}
}
