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
	000: mapperdb.NewMapperNROM,
	// TODO 001: mapperdb.NewMapperMMC1,
	002: mapperdb.NewMapperUxROM,
	003: mapperdb.NewMapperCNROM,
	// TODO 030: mapperdb.NewMapperUNROM512,
	111: mapperdb.NewMapperGTROM,
}

// New creates a new mapper for the mapper defined by the cartridge.
func New(bus *bus.Bus) (bus.Mapper, error) {
	mapperNumber := bus.Cartridge.Mapper
	initializer, ok := mappers[mapperNumber]
	if !ok {
		return nil, fmt.Errorf("mapper %d is not supported", mapperNumber)
	}

	base := NewBase(bus)
	mapper := initializer(base)
	return mapper, nil
}
