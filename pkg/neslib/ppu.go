package neslib

import . "github.com/retroenv/nesgo/pkg/nes"

// WaitSync waits for vertical sync to start.
func WaitSync() {
	for Bpl() {
		Bit(PPU_STATUS)
	}
	return
}

// StartPPUTransfer starts the PPU transfer to the passed address.
func StartPPUTransfer(address uint16, _ ...inline) {
	LdxAddr(PPU_STATUS)
	Ldx(uint8(address >> 8))
	Stx(PPU_ADDR)
	Ldx(uint8(address))
	Stx(PPU_ADDR)
}

// PPUTransfer transfers a constant to the PPU.
func PPUTransfer(data uint8, _ ...inline) {
	Lda(data)
	Sta(PPU_DATA)
}

// PPUTransferVar transfers a variable content to the PPU.
func PPUTransferVar(data *uint8, _ ...inline) {
	LdaAddr(data)
	Sta(PPU_DATA)
}

// PPUMask sets the PPU mask.
func PPUMask(flags uint8, _ ...inline) {
	Lda(flags)
	Sta(PPU_MASK)
}
