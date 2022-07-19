//go:build !nesgo
// +build !nesgo

package ppu

type status struct {
	OpenBus        byte // 0001 1111
	SpriteOverflow byte // 0010 0000
	SpriteZeroHit  byte // 0100 0000
	VerticalBlank  byte // 1000 0000
}

func (p *PPU) getStatus() byte {
	p.addressLatch = false
	p.clearVBlank()

	value := p.status.value()
	return value
}

func (s status) value() byte {
	value := s.OpenBus // TODO implement support for open bus value reading
	value |= s.SpriteOverflow << 5
	value |= s.SpriteZeroHit << 6
	value |= s.VerticalBlank << 7
	return value
}
