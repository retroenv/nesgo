package mapper

/*
Name: UxROM
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
	m.setDefaultBankSizes()
	m.setBanks()
	m.setWindows()
	return m
}

// Write a byte to a CHR or PRG memory address.
func (m *mapper2) Write(address uint16, value uint8) {
	switch {
	case address >= 0x8000:
		// Select 16 KB PRG ROM bank for CPU $8000-$BFFF
		bank := int(value) % len(m.prgBanks)
		m.prgWindows[0] = bank

	default:
		m.Base.Write(address, value)
	}
}
