package mapperdb

/*
Boards: NROM, HROM*, RROM, RTROM, SROM, STROM
PRG ROM capacity: 16K or 32K
CHR capacity: 8K
*/

import (
	"github.com/retroenv/nesgo/pkg/bus"
)

type mapperNROM struct {
	Base
}

// NewNROM returns a new mapper instance.
func NewNROM(base Base) bus.Mapper {
	m := &mapperNROM{
		Base: base,
	}
	m.SetName("NROM")
	m.Initialize()
	return m
}
