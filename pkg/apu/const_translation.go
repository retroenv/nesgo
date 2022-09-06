//go:build !nesgo

package apu

import . "github.com/retroenv/retrogolib/nes/addressing"

// AddressToName maps address constants from address to name.
var AddressToName = map[uint16]AccessModeConstant{
	SQ1_VOL:       {Constant: "SQ1_VOL", Mode: WriteAccess},
	SQ1_SWEEP:     {Constant: "SQ1_SWEEP", Mode: WriteAccess},
	SQ1_LO:        {Constant: "SQ1_LO", Mode: WriteAccess},
	SQ1_HI:        {Constant: "SQ1_HI", Mode: WriteAccess},
	SQ2_VOL:       {Constant: "SQ2_VOL", Mode: WriteAccess},
	SQ2_SWEEP:     {Constant: "SQ2_SWEEP", Mode: WriteAccess},
	SQ2_LO:        {Constant: "SQ2_LO", Mode: WriteAccess},
	SQ2_HI:        {Constant: "SQ2_HI", Mode: WriteAccess},
	TRI_LINEAR:    {Constant: "TRI_LINEAR", Mode: WriteAccess},
	TRI_LO:        {Constant: "TRI_LO", Mode: WriteAccess},
	TRI_HI:        {Constant: "TRI_HI", Mode: WriteAccess},
	NOISE_VOL:     {Constant: "NOISE_VOL", Mode: WriteAccess},
	NOISE_LO:      {Constant: "NOISE_LO", Mode: WriteAccess},
	NOISE_HI:      {Constant: "NOISE_HI", Mode: WriteAccess},
	APU_DMC_CTRL:  {Constant: "APU_DMC_CTRL", Mode: WriteAccess},
	APU_CHAN_CTRL: {Constant: "APU_CHAN_CTRL", Mode: ReadWriteAccess},
	APU_FRAME:     {Constant: "APU_FRAME", Mode: WriteAccess},
}
