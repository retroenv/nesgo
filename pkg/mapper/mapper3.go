package mapper

/*
Name: CNROM
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
	m.setDefaultBankSizes()
	m.setBanks()
	m.setWindows()
	return m
}

// Write a byte to a CHR or PRG memory address.
func (m *mapper3) Write(address uint16, value uint8) {
	switch {
	case address >= 0x8000:
		// Select 8 KB CHR ROM bank for PPU $0000-$1FFF
		bank := int(value) % len(m.chrBanks)
		m.chrWindows[0] = bank

	default:
		m.Base.Write(address, value)
	}
}
