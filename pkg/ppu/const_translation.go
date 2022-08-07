//go:build !nesgo

package ppu

import . "github.com/retroenv/nesgo/pkg/addressing"

// AddressToName maps address constants from address to name.
var AddressToName = map[uint16]AccessModeConstant{
	PPU_CTRL:   {Constant: "PPU_CTRL", Mode: WriteAccess},
	PPU_MASK:   {Constant: "PPU_MASK", Mode: WriteAccess},
	PPU_STATUS: {Constant: "PPU_STATUS", Mode: ReadAccess},
	OAM_ADDR:   {Constant: "OAM_ADDR", Mode: WriteAccess},
	OAM_DATA:   {Constant: "OAM_DATA", Mode: ReadWriteAccess},
	PPU_SCROLL: {Constant: "PPU_SCROLL", Mode: WriteAccess},
	PPU_ADDR:   {Constant: "PPU_ADDR", Mode: WriteAccess},
	PPU_DATA:   {Constant: "PPU_DATA", Mode: ReadWriteAccess},

	PALETTE_START: {Constant: "PALETTE_START", Mode: ReadWriteAccess},

	OAM_DMA: {Constant: "OAM_DMA", Mode: WriteAccess},
}
