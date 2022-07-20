//go:build !nesgo
// +build !nesgo

package ppu

func (p *PPU) backgroundPixel() byte {
	if !p.mask.RenderBackground {
		return 0
	}

	data := p.fetchTileData() >> ((7 - p.fineX) * 4)
	return byte(data & 0x0F)
}

func (p *PPU) fetchTileData() uint32 {
	return uint32(p.tileData >> 32)
}

func (p *PPU) storeTileData() {
	var data uint32
	for i := 0; i < 8; i++ {
		a := p.attributeTableByte
		p1 := (p.lowTileByte & 0x80) >> 7
		p2 := (p.highTileByte & 0x80) >> 6
		p.lowTileByte <<= 1
		p.highTileByte <<= 1
		data <<= 4
		data |= uint32(a | p1 | p2)
	}
	p.tileData |= uint64(data)
}

func (p *PPU) tileAddress() uint16 {
	table := p.control.BackgroundPatternTable
	tile := p.nameTable.Value()
	address := 0x1000*table + uint16(tile)*16 + p.addressing.FineY()
	return address
}

func (p *PPU) fetchLowTileByte() {
	address := p.tileAddress()
	p.lowTileByte = p.bus.Memory.Read(address)
}

func (p *PPU) fetchHighTileByte() {
	address := p.tileAddress()
	p.highTileByte = p.bus.Memory.Read(address + 8)
}
