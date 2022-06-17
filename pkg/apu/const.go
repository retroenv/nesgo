// Package apu provides APU (Audio Processing Unit) functionality.
package apu

const (
	SQ1_VOL       = 0x4000
	SQ1_SWEEP     = 0x4001
	SQ1_LO        = 0x4002
	SQ1_HI        = 0x4003
	SQ2_VOL       = 0x4004
	SQ2_SWEEP     = 0x4005
	SQ2_LO        = 0x4006
	SQ2_HI        = 0x4007
	TRI_LINEAR    = 0x4008
	TRI_LO        = 0x400A
	TRI_HI        = 0x400B
	NOISE_VOL     = 0x400C
	NOISE_LO      = 0x400E
	NOISE_HI      = 0x400F
	APU_DMC_CTRL  = 0x4010
	APU_CHAN_CTRL = 0x4015
	APU_FRAME     = 0x4017
)

// AddressToName maps address constants from address to name.
var AddressToName = map[uint16]string{
	SQ1_VOL:       "SQ1_VOL",
	SQ1_SWEEP:     "SQ1_SWEEP",
	SQ1_LO:        "SQ1_LO",
	SQ1_HI:        "SQ1_HI",
	SQ2_VOL:       "SQ2_VOL",
	SQ2_SWEEP:     "SQ2_SWEEP",
	SQ2_LO:        "SQ2_LO",
	SQ2_HI:        "SQ2_HI",
	TRI_LINEAR:    "TRI_LINEAR",
	TRI_LO:        "TRI_LO",
	TRI_HI:        "TRI_HI",
	NOISE_VOL:     "NOISE_VOL",
	NOISE_LO:      "NOISE_LO",
	NOISE_HI:      "NOISE_HI",
	APU_DMC_CTRL:  "APU_DMC_CTRL",
	APU_CHAN_CTRL: "APU_CHAN_CTRL",
	APU_FRAME:     "APU_FRAME",
}
