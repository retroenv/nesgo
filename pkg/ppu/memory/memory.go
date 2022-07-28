//go:build !nesgo
// +build !nesgo

// Package memory provides PPU memory access.
package memory

import (
	"github.com/retroenv/nesgo/pkg/bus"
)

// Memory implements PPU memory support.
type Memory struct {
	mapper    bus.Mapper
	nametable bus.NameTable
	palette   bus.BasicMemory
}

// New returns a new memory manager.
func New(mapper bus.Mapper, nametable bus.NameTable, palette bus.BasicMemory) *Memory {
	return &Memory{
		mapper:    mapper,
		nametable: nametable,
		palette:   palette,
	}
}

// Read from a PPU memory address.
func (m *Memory) Read(address uint16) uint8 {
	address &= 0x3FFF // valid addresses are $0000-$3FFF; higher addresses will be mirrored down

	switch {
	case address < 0x2000:
		return m.mapper.Read(address)

	case address < 0x3F00:
		return m.nametable.Read(address)

	default: // >= 0x3F00
		return m.palette.Read(address)
	}
}

// Write to a PPU memory address.
func (m *Memory) Write(address uint16, value uint8) {
	address &= 0x3FFF // valid addresses are $0000-$3FFF; higher addresses will be mirrored down

	switch {
	case address < 0x2000:
		m.mapper.Write(address, value)

	case address < 0x3F00:
		m.nametable.Write(address, value)

	default: // >= 0x3F00
		m.palette.Write(address, value)
	}
}
