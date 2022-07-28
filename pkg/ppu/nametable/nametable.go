//go:build !nesgo
// +build !nesgo

// Package nametable handles PPU nametables.
package nametable

import (
	"github.com/retroenv/nesgo/pkg/cartridge"
)

const (
	baseAddress   = 0x2000 // $2000 contains the nametables
	count         = 4      // 4 nametables
	nameTableSize = 0x0400 // 1024 byte per nametable
	// VramSize is the size of the nametable buffer.
	// It is normally mapped to the 2kB NES internal VRAM, providing 2 nametables with a mirroring configuration
	// controlled by the cartridge, but it can be partly or fully remapped to RAM on the cartridge,
	// allowing up to 4 simultaneous nametables
	VramSize = count * nameTableSize
)

// NameTable implements PPU nametable support.
// A nametable is a 1024 byte area of memory used by the PPU to lay out backgrounds.
// Each byte in the nametable controls one 8x8 pixel character cell, and each nametable has 30 rows
// of 32 tiles each, for 960 ($3C0) bytes; the rest is used by each nametable's attribute table.
// With each tile being 8x8 pixels, this makes a total of 256x240 pixels in one map,
// the same size as one full screen.
type NameTable struct {
	mirrorMode cartridge.MirrorMode

	value byte
	vram  []byte
}

// New returns a new nametable manager.
func New(mirrorMode cartridge.MirrorMode) *NameTable {
	return &NameTable{
		mirrorMode: mirrorMode,
	}
}

// SetVRAM sets the VRAM data buffer. This gets called by the mapper to allow nametable switching.
func (n *NameTable) SetVRAM(vram []byte) {
	n.vram = vram
}

// Read a value from the nametable address.
func (n NameTable) Read(address uint16) byte {
	base := n.mirroredNameTableAddressToBase(address)
	value := n.vram[base]
	return value
}

// Write a value to a nametable address.
func (n *NameTable) Write(address uint16, value byte) {
	base := n.mirroredNameTableAddressToBase(address)
	n.vram[base] = value
}

// Fetch a byte from the address and store it in the internal value storage for later retrieval.
func (n *NameTable) Fetch(address uint16) {
	n.value = n.Read(address)
}

// Value returns the earlier fetched value.
func (n NameTable) Value() byte {
	return n.value
}

func (n *NameTable) mirroredNameTableAddressToBase(address uint16) uint16 {
	address = (address - baseAddress) % (count * nameTableSize)
	table := address / nameTableSize
	offset := address % nameTableSize

	nameTableIndexes := n.mirrorMode.NametableIndexes()
	nameTableIndex := nameTableIndexes[table]

	base := nameTableIndex*nameTableSize + offset
	return base
}
