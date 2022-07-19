//go:build !nesgo
// +build !nesgo

// Package mapper provides hardware mapper support.
// It maps CHR and PRG chips into the NES address space.
package mapper

import (
	"fmt"

	"github.com/retroenv/nesgo/pkg/bus"
)

type mapperInitializer func(*bus.Bus) bus.BasicMemory

var mappers = map[byte]mapperInitializer{
	0: newMapper0,
}

// New creates a new mapper for the mapper defined by the cartridge.
func New(bus *bus.Bus) (bus.BasicMemory, error) {
	mapperNumber := bus.Cartridge.Mapper
	initializer, ok := mappers[mapperNumber]
	if !ok {
		return nil, fmt.Errorf("mapper %d is not supported", mapperNumber)
	}

	mapper := initializer(bus)
	return mapper, nil
}
