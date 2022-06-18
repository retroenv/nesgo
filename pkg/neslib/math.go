package neslib

import . "github.com/retroenv/nesgo/pkg/nes"

// DivSigned16 divides the signed number in A by 16.
func DivSigned16(_ ...Inline) {
	Lsr()
	Lsr()
	Lsr()
	Lsr()
	Clc()
	Adc(0x78)
	Eor(0x78)
}

// DivSigned8 divides the signed number in A by 8.
func DivSigned8(_ ...Inline) {
	Lsr()
	Lsr()
	Lsr()
	Clc()
	Adc(0x70)
	Eor(0x70)
}
