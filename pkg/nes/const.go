// Package nes provides helper functions for writing NES programs in Golang.
// This package needs to be imported using the dot notation.
package nes

import (
	"github.com/retroenv/nesgo/pkg/apu"
	"github.com/retroenv/nesgo/pkg/controller"
	"github.com/retroenv/nesgo/pkg/ppu"
)

const (
	// PPU_CTRL = PPU Control 1.
	PPU_CTRL = ppu.PPU_CTRL
	// PPU_MASK = PPU Control 2
	PPU_MASK = ppu.PPU_MASK
	// PPU_STATUS = PPU Status.
	PPU_STATUS = ppu.PPU_STATUS
	// OAM_ADDR = OAM Address.
	OAM_ADDR = ppu.OAM_ADDR
	// OAM_DATA = OAM Data.
	OAM_DATA = ppu.OAM_DATA
	// PPU_SCROLL = Background Scroll Position (write X then Y).
	PPU_SCROLL = ppu.PPU_SCROLL
	// PPU_ADDR = PPU Address (write upper then lower).
	PPU_ADDR = ppu.PPU_ADDR
	// PPU_DATA = PPU Data.
	PPU_DATA = ppu.PPU_DATA

	// PALETTE_START = Universal background color.
	PALETTE_START = ppu.PALETTE_START

	// PPU_OAM_DMA = Sprite Page DMA Transfer.
	PPU_OAM_DMA = ppu.OAM_DMA

	APU_SQ1_VOL    = apu.SQ1_VOL
	APU_SQ1_SWEEP  = apu.SQ1_SWEEP
	APU_SQ1_LO     = apu.SQ1_LO
	APU_SQ1_HI     = apu.SQ1_HI
	APU_SQ2_VOL    = apu.SQ2_VOL
	APU_SQ2_SWEEP  = apu.SQ2_SWEEP
	APU_SQ2_LO     = apu.SQ2_LO
	APU_SQ2_HI     = apu.SQ2_HI
	APU_TRI_LINEAR = apu.TRI_LINEAR
	APU_TRI_LO     = apu.TRI_LO
	APU_TRI_HI     = apu.TRI_HI
	APU_NOISE_VOL  = apu.NOISE_VOL
	APU_NOISE_LO   = apu.NOISE_LO
	APU_NOISE_HI   = apu.NOISE_HI
	APU_DMC_CTRL   = apu.APU_DMC_CTRL
	APU_CHAN_CTRL  = apu.APU_CHAN_CTRL
	APU_FRAME      = apu.APU_FRAME

	JOYPAD1 = controller.JOYPAD1
	JOYPAD2 = controller.JOYPAD2

	// PPU_CTRL flags
	CTRL_NMI         = ppu.CTRL_NMI         // Execute Non-Maskable Interrupt on VBlank
	CTRL_MASTERSLAVE = ppu.CTRL_MASTERSLAVE // Master/Slave select
	CTRL_8x8         = ppu.CTRL_8x8         // Use 8x8 Sprites
	CTRL_8x16        = ppu.CTRL_8x16        // Use 8x16 Sprites
	CTRL_BG_0000     = ppu.CTRL_BG_0000     // Background Pattern Table at 0x0000 in VRAM
	CTRL_BG_1000     = ppu.CTRL_BG_1000     // Background Pattern Table at 0x1000 in VRAM
	CTRL_SPR_0000    = ppu.CTRL_SPR_0000    // Sprite Pattern Table at 0x0000 in VRAM
	CTRL_SPR_1000    = ppu.CTRL_SPR_1000    // Sprite Pattern Table at 0x1000 in VRAM
	CTRL_INC_1       = ppu.CTRL_INC_1       // Increment PPU Address by 1 (Horizontal rendering)
	CTRL_INC_32      = ppu.CTRL_INC_32      // Increment PPU Address by 32 (Vertical rendering)
	CTRL_NT_2000     = ppu.CTRL_NT_2000     // Name Table Address at 0x2000
	CTRL_NT_2400     = ppu.CTRL_NT_2400     // Name Table Address at 0x2400
	CTRL_NT_2800     = ppu.CTRL_NT_2800     // Name Table Address at 0x2800
	CTRL_NT_2C00     = ppu.CTRL_NT_2C00     // Name Table Address at 0x2C00

	// PPU_MASK flags
	MASK_TINT_RED   = ppu.MASK_TINT_RED   // Red Background
	MASK_TINT_BLUE  = ppu.MASK_TINT_BLUE  // Blue Background
	MASK_TINT_GREEN = ppu.MASK_TINT_GREEN // Green Background
	MASK_SPR        = ppu.MASK_SPR        // Sprites Visible
	MASK_BG         = ppu.MASK_BG         // Backgrounds Visible
	MASK_SPR_CLIP   = ppu.MASK_SPR_CLIP   // Sprites clipped on left column
	MASK_BG_CLIP    = ppu.MASK_BG_CLIP    // Background clipped on left column
	MASK_COLOR      = ppu.MASK_COLOR      // Display in Color
	MASK_MONO       = ppu.MASK_MONO       // Display in Monochrome

	// read flags
	F_BLANK   = 0b10000000 // VBlank Active
	F_SPRITE0 = 0b01000000 // VBlank hit Sprite 0
	F_SCAN8   = 0b00100000 // More than 8 sprites on current scanline
	F_WIGNORE = 0b00010000 // VRAM Writes currently ignored.
)

const (
	COLOR_DARK_GRAY     = 0x00
	COLOR_MEDIUM_GRAY   = 0x10
	COLOR_LIGHT_GRAY    = 0x20
	COLOR_LIGHTEST_GRAY = 0x30

	COLOR_DARK_BLUE     = 0x01
	COLOR_MEDIUM_BLUE   = 0x11
	COLOR_LIGHT_BLUE    = 0x21
	COLOR_LIGHTEST_BLUE = 0x31

	COLOR_DARK_INDIGO     = 0x02
	COLOR_MEDIUM_INDIGO   = 0x12
	COLOR_LIGHT_INDIGO    = 0x22
	COLOR_LIGHTEST_INDIGO = 0x32

	COLOR_DARK_VIOLET     = 0x03
	COLOR_MEDIUM_VIOLET   = 0x13
	COLOR_LIGHT_VIOLET    = 0x23
	COLOR_LIGHTEST_VIOLET = 0x33

	COLOR_DARK_PURPLE     = 0x04
	COLOR_MEDIUM_PURPLE   = 0x14
	COLOR_LIGHT_PURPLE    = 0x24
	COLOR_LIGHTEST_PURPLE = 0x24

	COLOR_DARK_REDVIOLET     = 0x05
	COLOR_MEDIUM_REDVIOLET   = 0x15
	COLOR_LIGHT_REDVIOLET    = 0x25
	COLOR_LIGHTEST_REDVIOLET = 0x35

	COLOR_DARK_RED     = 0x06
	COLOR_MEDIUM_RED   = 0x16
	COLOR_LIGHT_RED    = 0x26
	COLOR_LIGHTEST_RED = 0x36

	COLOR_DARK_ORANGE     = 0x07
	COLOR_MEDIUM_ORANGE   = 0x17
	COLOR_LIGHT_ORANGE    = 0x27
	COLOR_LIGHTEST_ORANGE = 0x37

	COLOR_DARK_YELLOW     = 0x08
	COLOR_MEDIUM_YELLOW   = 0x18
	COLOR_LIGHT_YELLOW    = 0x28
	COLOR_LIGHTEST_YELLOW = 0x38

	COLOR_DARK_CHARTREUSE     = 0x09
	COLOR_MEDIUM_CHARTREUSE   = 0x19
	COLOR_LIGHT_CHARTREUSE    = 0x29
	COLOR_LIGHTEST_CHARTREUSE = 0x39

	COLOR_DARK_GREEN     = 0x0a
	COLOR_MEDIUM_GREEN   = 0x1a
	COLOR_LIGHT_GREEN    = 0x2a
	COLOR_LIGHTEST_GREEN = 0x3a

	COLOR_DARK_CYAN     = 0x0b
	COLOR_MEDIUM_CYAN   = 0x1b
	COLOR_LIGHT_CYAN    = 0x2b
	COLOR_LIGHTEST_CYAN = 0x3b

	COLOR_DARK_TURQUOISE     = 0x0c
	COLOR_MEDIUM_TURQUOISE   = 0x1c
	COLOR_LIGHT_TURQUOISE    = 0x1c
	COLOR_LIGHTEST_TURQUOISE = 0x1c

	COLOR_BLACK        = 0x0f
	COLOR_SILVER       = 0x10
	COLOR_LIME         = 0x2a
	COLOR_DARKEST_GRAY = 0x2d
	COLOR_MEDIUM_GRAY2 = 0x3d
)
