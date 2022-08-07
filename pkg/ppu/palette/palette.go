//go:build !nesgo

// Package palette handles PPU palette support.
package palette

// Palette implements PPU palette support.
type Palette struct {
	data [32]byte // contains color indexes
}

// New returns a new palette manager.
func New() *Palette {
	return &Palette{}
}

// Read a value from the palette address.
func (p *Palette) Read(address uint16) byte {
	base := mirroredPaletteAddressToBase(address)
	value := p.data[base]
	return value
}

// Write a value to a palette address.
func (p *Palette) Write(address uint16, value byte) {
	base := mirroredPaletteAddressToBase(address)
	p.data[base] = value
}

func mirroredPaletteAddressToBase(address uint16) uint16 {
	// $3F20-$3FFF are mirrors of $3F00-$3F1F
	address %= 0x20

	// $3F10/$3F14/$3F18/$3F1C are mirrors of $3F00/$3F04/$3F08/$3F0C
	if address >= 0x10 && address%4 == 0 {
		address -= 0x10
	}
	return address
}
