package neslib

import . "github.com/retroenv/nesgo/pkg/nes"

// Init is the NES setup start routine.
func Init(_ ...Inline) {
	Sei()
	Cld()
	Ldx(0xff)
	Txs()
	Inx()
	Stx(PPU_MASK)      // disable rendering
	Stx(APU_DMC_CTRL)  // disable DMC interrupts
	Stx(PPU_CTRL)      // disable NMI interrupts
	Bit(PPU_STATUS)    // clear VBL flag
	Bit(APU_CHAN_CTRL) // ack DMC IRQ bit 7
	Lda(0x40)
	Sta(APU_FRAME) // disable APU Frame IRQ
	Lda(0x0F)
	Sta(APU_CHAN_CTRL) // disable DMC, enable/init other channels.
}

// ClearRAM clears all RAM, except for last 2 bytes of CPU stack.
func ClearRAM() {
	Lda(0) // A = 0
	Tax()  // X = 0

	for Bne() { // loop 256 times
		Sta(0, X) // clear 0x0-0xff

		Cpx(0xfe)   // last 2 bytes of stack?
		if !Bcs() { // don't clear it
			Sta(0x100, X) // clear 0x100-0x1fd
		}

		Sta(0x200, X) // clear 0x200-0x2ff
		Sta(0x300, X) // clear 0x300-0x3ff
		Sta(0x400, X) // clear 0x400-0x4ff
		Sta(0x500, X) // clear 0x500-0x5ff
		Sta(0x600, X) // clear 0x600-0x6ff
		Sta(0x700, X) // clear 0x700-0x7ff
		Inx()         // X = X + 1
	}
}
