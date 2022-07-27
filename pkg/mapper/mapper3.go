package mapper

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
	*Base
}

func newMapper3(bus *bus.Bus) bus.Mapper {
	m := &mapper3{
		Base: newBase(bus),
	}
	m.name = "CNROM"
	m.initialize()
	m.addWriteHook(0x8000, 0xFFFF, m.setChrWindow)
	return m
}

func (m *mapper3) setChrWindow(address uint16, value uint8) {
	// Select 8 KB CHR ROM bank for PPU $0000-$1FFF
	bank := int(value) % len(m.chrBanks)
	m.chrWindows[0] = bank
}
