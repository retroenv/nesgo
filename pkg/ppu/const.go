package ppu

import (
	"github.com/retroenv/nesgo/pkg/ppu/control"
	"github.com/retroenv/nesgo/pkg/ppu/mask"
)

const (
	PPU_CTRL   = 0x2000
	PPU_MASK   = 0x2001
	PPU_STATUS = 0x2002
	OAM_ADDR   = 0x2003
	OAM_DATA   = 0x2004
	PPU_SCROLL = 0x2005
	PPU_ADDR   = 0x2006
	PPU_DATA   = 0x2007

	PALETTE_START = 0x3f00

	OAM_DMA = 0x4014

	// PPU_CTRL flags
	CTRL_NMI         = control.CTRL_NMI         // Execute Non-Maskable Interrupt on VBlank
	CTRL_MASTERSLAVE = control.CTRL_MASTERSLAVE // Master/Slave select
	CTRL_8x8         = control.CTRL_8x8         // Use 8x8 Sprites
	CTRL_8x16        = control.CTRL_8x16        // Use 8x16 Sprites
	CTRL_BG_0000     = control.CTRL_BG_0000     // Background Pattern Table at 0x0000 in VRAM
	CTRL_BG_1000     = control.CTRL_BG_1000     // Background Pattern Table at 0x1000 in VRAM
	CTRL_SPR_0000    = control.CTRL_SPR_0000    // Sprite Pattern Table at 0x0000 in VRAM
	CTRL_SPR_1000    = control.CTRL_SPR_1000    // Sprite Pattern Table at 0x1000 in VRAM
	CTRL_INC_1       = control.CTRL_INC_1       // Increment PPU Address by 1 (Horizontal rendering)
	CTRL_INC_32      = control.CTRL_INC_32      // Increment PPU Address by 32 (Vertical rendering)
	CTRL_NT_2000     = control.CTRL_NT_2000     // Name Table Address at 0x2000
	CTRL_NT_2400     = control.CTRL_NT_2400     // Name Table Address at 0x2400
	CTRL_NT_2800     = control.CTRL_NT_2800     // Name Table Address at 0x2800
	CTRL_NT_2C00     = control.CTRL_NT_2C00     // Name Table Address at 0x2C00

	// PPU_MASK flags
	MASK_COLOR      = mask.MASK_COLOR      // Display in Color
	MASK_MONO       = mask.MASK_MONO       // Display in Monochrome
	MASK_BG_CLIP    = mask.MASK_BG_CLIP    // Background clipped on left column
	MASK_SPR_CLIP   = mask.MASK_SPR_CLIP   // Sprites clipped on left column
	MASK_BG         = mask.MASK_BG         // Backgrounds Visible
	MASK_SPR        = mask.MASK_SPR        // Sprites Visible
	MASK_TINT_RED   = mask.MASK_TINT_RED   // Red Background
	MASK_TINT_BLUE  = mask.MASK_TINT_BLUE  // Blue Background
	MASK_TINT_GREEN = mask.MASK_TINT_GREEN // Green Background
)
