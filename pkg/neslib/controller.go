package neslib

import . "github.com/retroenv/nesgo/pkg/nes"

// ReadJoypad returns all joypad bits in the A register.
// index defines the joypad, must be set to 0 or 1.
func ReadJoypad(index uint8) {
	Ldy(index)
	Lda(1)
	Sta(JOYPAD1, Y) // set strobe bit
	Lsr()           // now A is 0
	Sta(JOYPAD1, Y) // clear strobe bit
	Ldx(8)          // read 8 bits
	for Bne() {     // repeat while X is 0
		Pha()           // save A (result)
		Lda(JOYPAD1, Y) // load controller state
		Lsr()           // bit 0 -> carry
		Pla()           // restore A (result)
		Rol()           // carry -> bit 0 of result
		Dex()           // X = X - 1
	}
}
