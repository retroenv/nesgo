//go:build !nesgo
// +build !nesgo

package ppu

func (p *PPU) fetchAttributeTableByte() {
	address := p.addressing.Address()
	shift := ((address >> 4) & 4) | (address & 2)
	address = 0x23C0 | (address & 0x0C00) | ((address >> 4) & 0x38) | ((address >> 2) & 0x07)

	value := p.bus.Memory.Read(address)
	p.attributeTableByte = ((value >> shift) & 3) << 2
}
