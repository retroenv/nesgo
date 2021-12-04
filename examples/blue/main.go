package main

import (
	. "github.com/retroenv/nesgo/pkg/nes"
	. "github.com/retroenv/nesgo/pkg/neslib"
)

var backgroundColor = NewUint8(COLOR_MEDIUM_BLUE)

func main() {
	Start(resetHandler)
}

func resetHandler() {
	Init()

	WaitSync()     // wait for VSYNC
	ClearRAM()     // clear RAM
	WaitSync()     // wait for VSYNC (and PPU warmup)
	VariableInit() // initialize variables after RAM has been cleared

	StartPPUTransfer(PALETTE_START)
	PPUTransfer(backgroundColor)
	PPUMask(MASK_BG_CLIP | MASK_SPR_CLIP | MASK_BG | MASK_SPR)
}
