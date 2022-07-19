//go:build !nesgo
// +build !nesgo

package ppu

func (p *PPU) setOamAddress(value byte) {
	p.oamAddress = value
}

func (p *PPU) readOamData() byte {
	// TODO handle special case of reading during rendering
	data := p.oamData[p.oamAddress]
	return data
}

func (p *PPU) writeOamData(value byte) {
	p.oamData[p.oamAddress] = value
	p.oamAddress++

	// TODO handle scroll glitch
}

func (p *PPU) writeOamDMA(value byte) {
	address := uint16(value) << 8

	for i := 0; i < 256; i++ {
		data := p.ram.ReadMemory(address)
		p.oamData[p.oamAddress] = data
		p.oamAddress++
		address++
	}

	// 1 wait state cycle while waiting for writes to complete,
	// +1 if on an odd CPU cycle, then 256 alternating read/write cycles
	stall := uint16(1 + 256 + 256)
	if p.bus.CPU.Cycles()%2 == 1 {
		stall++
	}
	p.bus.CPU.StallCycles(stall)
}
