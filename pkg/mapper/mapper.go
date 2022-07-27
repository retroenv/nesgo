//go:build !nesgo
// +build !nesgo

// Package mapper provides hardware mapper support.
// It maps CHR and PRG chips into the NES address space.
package mapper

import (
	"fmt"

	"github.com/retroenv/nesgo/pkg/bus"
	"github.com/retroenv/nesgo/pkg/mapper/mapperdb"
)

type mapperInitializer func(base mapperdb.Base) bus.Mapper

var mappers = map[byte]mapperInitializer{
	0: mapperdb.NewMapperNROM,
	2: mapperdb.NewMapperUxRom,
	3: mapperdb.NewMapperCNROM,
}

// New creates a new mapper for the mapper defined by the cartridge.
func New(bus *bus.Bus) (bus.Mapper, error) {
	mapperNumber := bus.Cartridge.Mapper
	initializer, ok := mappers[mapperNumber]
	if !ok {
		return nil, fmt.Errorf("mapper %d is not supported", mapperNumber)
	}

	base := newBase(bus)
	mapper := initializer(base)
	return mapper, nil
}
