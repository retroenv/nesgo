//go:build !nesgo
// +build !nesgo

package ppu

import "fmt"

// Read from a PPU memory register address.
func (p *PPU) Read(address uint16) uint8 {
	base := mirroredRegisterAddressToBase(address)

	switch base {
	case PPU_CTRL:
		return p.control.Value()

	case PPU_MASK:
		return p.mask.Value()

	case PPU_STATUS:
		return p.getStatus()

	case OAM_DATA:
		return p.sprites.Read()

	case PPU_DATA:
		return p.readData()

	default:
		panic(fmt.Sprintf("unhandled ppu read at address: 0x%04X", address))
	}
}

// Write to a PPU memory register address.
func (p *PPU) Write(address uint16, value uint8) {
	base := mirroredRegisterAddressToBase(address)

	switch base {
	case PPU_CTRL:
		p.control.Set(value)

	case PPU_MASK:
		p.mask.Set(value)

	case OAM_ADDR:
		p.sprites.SetAddress(value)

	case OAM_DATA:
		p.sprites.Write(value)

	case PPU_SCROLL:
		if !p.addressing.Latch() {
			p.fineX = uint16(value) & 0x07
		}
		p.addressing.SetScroll(value)

	case PPU_ADDR:
		p.addressing.SetAddress(value)

	case PPU_DATA:
		address := p.addressing.Address()
		p.memory.Write(address, value)
		p.addressing.Increment(p.control.VRAMIncrement)

	case OAM_DMA:
		p.sprites.WriteDMA(value)

	default:
		panic(fmt.Sprintf("unhandled ppu write at address: 0x%04X", address))
	}
}

// mirroredRegisterAddressToBase converts the mirrored addresses to the base address.
// PPU registers are mirrored in every 8 bytes from $2008 through $3FFF.
func mirroredRegisterAddressToBase(address uint16) uint16 {
	base := 0x2000 + address&0b00000111
	return base
}
