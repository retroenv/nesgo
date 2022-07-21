//go:build !nesgo
// +build !nesgo

package ppu

func (p *PPU) setVerticalBlank() {
	p.screen.FinishRendering()
	p.nmi.occurred = true
	p.nmi.change()
}

func (p *PPU) clearVerticalBlank() {
	p.nmi.occurred = false
	p.nmi.change()
}
