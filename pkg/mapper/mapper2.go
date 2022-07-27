package mapper

/*
Boards: UNROM, UOROM
PRG ROM capacity: 256K/4096K
PRG ROM window:16K + 16K fixed
CHR capacity: 8K
*/

import (
	"github.com/retroenv/nesgo/pkg/bus"
)

type mapper2 struct {
	*Base
}

func newMapper2(bus *bus.Bus) bus.Mapper {
	m := &mapper2{
		Base: newBase(bus),
	}
	m.name = "UxROM"
	m.initialize()
	m.addWriteHook(0x8000, 0xFFFF, m.setPrgWindow)
	return m
}

func (m *mapper2) setPrgWindow(address uint16, value uint8) {
	// Select 16 KB PRG ROM bank for CPU $8000-$BFFF
	bank := int(value) % len(m.prgBanks)
	m.prgWindows[0] = bank
}
