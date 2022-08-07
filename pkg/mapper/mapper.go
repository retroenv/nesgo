//go:build !nesgo

// Package mapper provides hardware mapper support.
// It maps CHR and PRG chips into the NES address space.
package mapper

import (
	"fmt"

	"github.com/retroenv/nesgo/pkg/bus"
	"github.com/retroenv/nesgo/pkg/mapper/mapperbase"
	"github.com/retroenv/nesgo/pkg/mapper/mapperdb"
)

type mapperInitializer func(base mapperdb.Base) bus.Mapper

var mappers = map[byte]mapperInitializer{
	0:   mapperdb.NewNROM,
	1:   mapperdb.NewMMC1,
	2:   mapperdb.NewUxROMOr,
	3:   mapperdb.NewCNROM,
	7:   mapperdb.NewAxROM,
	30:  mapperdb.NewUNROM512,
	94:  mapperdb.NewUN1ROM,
	111: mapperdb.NewGTROM,
	180: mapperdb.NewUxROMAnd,
}

// New creates a new mapper for the mapper defined by the cartridge.
func New(bus *bus.Bus) (bus.Mapper, error) {
	mapperNumber := bus.Cartridge.Mapper
	initializer, ok := mappers[mapperNumber]
	if !ok {
		return nil, fmt.Errorf("mapper %d is not supported", mapperNumber)
	}

	base := mapperbase.New(bus)
	mapper := initializer(base)
	return mapper, nil
}
