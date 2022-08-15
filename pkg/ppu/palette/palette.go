//go:build !nesgo

// Package palette handles PPU palette support.
package palette

import "sync"

const size = 32

// Palette implements PPU palette support.
type Palette struct {
	mu   sync.RWMutex
	data [size]byte // contains color indexes
}

// New returns a new palette manager.
func New() *Palette {
	return &Palette{}
}

// Read a value from the palette address.
func (p *Palette) Read(address uint16) byte {
	base := mirroredPaletteAddressToBase(address)
	p.mu.RLock()
	value := p.data[base]
	p.mu.RUnlock()
	return value
}

// Write a value to a palette address.
func (p *Palette) Write(address uint16, value byte) {
	base := mirroredPaletteAddressToBase(address)
	p.mu.Lock()
	p.data[base] = value
	p.mu.Unlock()
}

// Data returns the palette data as byte array.
func (p *Palette) Data() [32]byte {
	p.mu.RLock()
	data := p.data
	p.mu.RUnlock()
	return data
}

func mirroredPaletteAddressToBase(address uint16) uint16 {
	// $3F20-$3FFF are mirrors of $3F00-$3F1F
	address %= size

	// $3F10/$3F14/$3F18/$3F1C are mirrors of $3F00/$3F04/$3F08/$3F0C
	if address >= 0x10 && address%4 == 0 {
		address -= 0x10
	}
	return address
}
