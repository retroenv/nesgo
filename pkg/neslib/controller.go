package neslib

import . "github.com/retroenv/nesgo/pkg/nes"

// ReadJoypad0 return either joypad bits in A register.
// Y must be set to 0 or 1.
func ReadJoypad0() {
	Ldy(0)
	Lda(1)
	Sta(JOYPAD1, Y) // set strobe bit
	Lsr()           // now A is 0
	Sta(JOYPAD1, Y) // clear strobe bit
	Ldx(8)          // read 8 bits
	for Bne() {     // repeat while X is 0
		Pha()               // save A (result)
		LdaAddr(JOYPAD1, Y) // load controller state
		Lsr()               // bit 0 -> carry
		Pla()               // restore A (result)
		Rol()               // carry -> bit 0 of result
		Dex()               // X = X - 1
	}
	return
}
