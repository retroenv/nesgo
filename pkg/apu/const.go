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
