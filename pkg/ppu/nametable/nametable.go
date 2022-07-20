//go:build !nesgo
// +build !nesgo

// Package nametable handles PPU nametables.
package nametable

import (
	"github.com/retroenv/nesgo/pkg/bus"
	"github.com/retroenv/nesgo/pkg/cartridge"
)

const (
	baseAddress    = 0x2000 // $2000  contains the nametables
	count          = 4      // 4 nametables
	maximumAddress = 0x2FFF // $2FFF end of nametable 4
	size           = 0x0400 // 1024 byte per nametable
)

// NameTable implements PPU nametable support.
// A nametable is a 1024 byte area of memory used by the PPU to lay out backgrounds.
// Each byte in the nametable controls one 8x8 pixel character cell, and each nametable has 30 rows
// of 32 tiles each, for 960 ($3C0) bytes; the rest is used by each nametable's attribute table.
// With each tile being 8x8 pixels, this makes a total of 256x240 pixels in one map,
// the same size as one full screen.
type NameTable struct {
	mapper bus.Mapper
	value  byte
}

// New returns a new nametable manager.
func New(mapper bus.Mapper) *NameTable {
	return &NameTable{
		mapper: mapper,
	}
}

// Read a value from the nametable address.
func (n NameTable) Read(address uint16, mirrorMode cartridge.MirrorMode) byte {
	base := mirroredNameTableAddressToBase(address, mirrorMode)
	value := n.mapper.Read(base)
	return value
}

// Write a value to a nametable address.
func (n *NameTable) Write(address uint16, value byte, mirrorMode cartridge.MirrorMode) {
	base := mirroredNameTableAddressToBase(address, mirrorMode)
	n.mapper.Write(base, value)
}

// Fetch a byte from the address and store it in the internal value storage for later retrieval.
func (n *NameTable) Fetch(address uint16) {
	address &= maximumAddress
	n.value = n.mapper.Read(address)
}

// Value returns the earlier fetched value.
func (n NameTable) Value() byte {
	return n.value
}

func mirroredNameTableAddressToBase(address uint16, mirrorMode cartridge.MirrorMode) uint16 {
	address = (address - baseAddress) % (count * size)
	table := address / size
	offset := address % size

	nameTableIndexes := mirrorMode.NametableIndexes()
	nameTableIndex := nameTableIndexes[table]

	base := baseAddress + nameTableIndex*size + offset
	return base
}
