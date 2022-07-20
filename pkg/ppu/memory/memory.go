//go:build !nesgo
// +build !nesgo

// Package memory provides PPU memory access.
package memory

import (
	"fmt"

	"github.com/retroenv/nesgo/pkg/bus"
	"github.com/retroenv/nesgo/pkg/ppu/nametable"
	"github.com/retroenv/nesgo/pkg/ppu/palette"
)

// Memory implements PPU memory support.
type Memory struct {
	mapper    bus.Mapper
	nametable *nametable.NameTable
	palette   *palette.Palette
}

// New returns a new memory manager.
func New(mapper bus.Mapper, nametable *nametable.NameTable, palette *palette.Palette) *Memory {
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

	case address >= 0x3F00:
		return m.palette.Read(address)

	default:
		panic(fmt.Sprintf("unhandled ppu read at address: 0x%04X", address))
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

	case address >= 0x3F00:
		m.palette.Write(address, value)

	default:
		panic(fmt.Sprintf("unhandled ppu write at address: 0x%04X", address))
	}
}
