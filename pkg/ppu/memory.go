//go:build !nesgo
// +build !nesgo

package ppu

import "fmt"

// Read from a PPU memory address.
func (p *PPU) Read(address uint16) uint8 {
	if address < 0x2000 {
		return p.bus.Mapper.Read(address)
	}

	base := mirroredAddressToBase(address)
	switch base {
	case PPU_CTRL:
		return p.control.value
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

// Write to a PPU memory address.
func (p *PPU) Write(address uint16, value uint8) {
	if address < 0x2000 {
		p.bus.Mapper.Write(address, value)
		return
	}

	base := mirroredAddressToBase(address)
	switch base {
	case PPU_CTRL:
		p.setControl(value)
	case PPU_MASK:
		p.mask.Set(value)
	case OAM_ADDR:
		p.sprites.SetAddress(value)
	case OAM_DATA:
		p.sprites.Write(value)
	case PPU_SCROLL:
		p.setScroll(value)
	case PPU_ADDR:
		p.addressing.SetAddress(value)
	case PPU_DATA:
		p.writeData(value)
	case OAM_DMA:
		p.sprites.WriteDMA(value)

	default:
		panic(fmt.Sprintf("unhandled ppu write at address: 0x%04X", address))
	}
}
