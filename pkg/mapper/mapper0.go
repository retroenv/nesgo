package mapper

/*
Name: NROM
Boards: NROM, HROM*, RROM, RTROM, SROM, STROM
PRG ROM capacity: 16K or 32K
CHR capacity: 8K
*/

import (
	"github.com/retroenv/nesgo/pkg/bus"
)

type mapper0 struct {
	*Base
}

func newMapper0(bus *bus.Bus) bus.Mapper {
	m := &mapper0{
		Base: newBase(bus),
	}
	m.setDefaultBankSizes()
	m.setBanks()
	m.setWindows()
	return m
}
