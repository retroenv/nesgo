package neslib

import . "github.com/retroenv/nesgo/pkg/nes"

// NextRandom returns a random number in A register.
func NextRandom() {
	Lsr(A)
	if Bcc() {
		goto NoEor
	}
	Eor(0xd4)
NoEor:
}

// PrevRandom returns a random number in A register.
func PrevRandom() {
	Asl(A)
	if Bcc() {
		goto NoEor
	}
	Eor(0xa9)
NoEor:
}
