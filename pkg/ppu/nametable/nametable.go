//go:build !nesgo

// Package nametable handles PPU nametables.
package nametable

import (
	"sync"

	"github.com/retroenv/retrogolib/nes/cartridge"
)

const (
	baseAddress = 0x2000 // $2000 contains the nametables
	count       = 4      // 4 nametables
	size        = 0x0400 // 1024 byte per nametable
	// VramSize is the size of the nametable buffer.
	// It is normally mapped to the 2kB NES internal VRAM, providing 2 nametables with a mirroring configuration
	// controlled by the cartridge, but it can be partly or fully remapped to RAM on the cartridge,
	// allowing up to 4 simultaneous nametables
	VramSize = count * size
)

// NameTable implements PPU nametable support.
// A nametable is a 1024 byte area of memory used by the PPU to lay out backgrounds.
// Each byte in the nametable controls one 8x8 pixel character cell, and each nametable has 30 rows
// of 32 tiles each, for 960 ($3C0) bytes; the rest is used by each nametable's attribute table.
// With each tile being 8x8 pixels, this makes a total of 256x240 pixels in one map,
// the same size as one full screen.
type NameTable struct {
	mu sync.RWMutex

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

// Data returns the nametable data as byte arrays.
func (n *NameTable) Data() [4][]byte {
	nameTableIndexes := n.mirrorMode.NametableIndexes()
	data := [4][]byte{}

	n.mu.RLock()
	for table := 0; table < 4; table++ {
		nameTableIndex := nameTableIndexes[table]
		base := nameTableIndex * size
		b := n.vram[base : base+size]
		data[table] = b
	}
	n.mu.RLock()
	return data
}

// SetVRAM sets the VRAM data buffer. This gets called by the mapper to allow nametable switching.
func (n *NameTable) SetVRAM(vram []byte) {
	n.mu.Lock()
	n.vram = vram
	n.mu.Unlock()
}

// MirrorMode returns the set mirror mode.
func (n *NameTable) MirrorMode() cartridge.MirrorMode {
	n.mu.RLock()
	defer n.mu.RUnlock()

	return n.mirrorMode
}

// SetMirrorMode sets the mirror mode.
func (n *NameTable) SetMirrorMode(mirrorMode cartridge.MirrorMode) {
	n.mu.Lock()
	n.mirrorMode = mirrorMode
	n.mu.Unlock()
}

// Read a value from the nametable address.
func (n *NameTable) Read(address uint16) byte {
	base := n.mirroredNameTableAddressToBase(address)

	n.mu.RLock()
	value := n.vram[base]
	n.mu.RUnlock()
	return value
}

// Write a value to a nametable address.
func (n *NameTable) Write(address uint16, value byte) {
	base := n.mirroredNameTableAddressToBase(address)

	n.mu.Lock()
	n.vram[base] = value
	n.mu.Unlock()
}

// Fetch a byte from the address and store it in the internal value storage for later retrieval.
func (n *NameTable) Fetch(address uint16) {
	value := n.Read(address)
	n.mu.Lock()
	n.value = value
	n.mu.Unlock()
}

// Value returns the earlier fetched value.
func (n *NameTable) Value() byte {
	n.mu.RLock()
	defer n.mu.RUnlock()

	return n.value
}

func (n *NameTable) mirroredNameTableAddressToBase(address uint16) uint16 {
	address = (address - baseAddress) % (count * size)
	table := address / size
	offset := address % size

	nameTableIndexes := n.mirrorMode.NametableIndexes()
	nameTableIndex := nameTableIndexes[table]

	base := nameTableIndex*size + offset
	return base
}
