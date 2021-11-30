// Package nes provides helper functions for writing NES programs in Golang.
// This package needs to be imported using the dot notation.
package nes

const (
	// PPU_CTRL = PPU Control 1.
	PPU_CTRL = 0x2000
	// PPU_MASK = PPU Control 2
	PPU_MASK = 0x2001
	// PPU_STATUS = PPU Status.
	PPU_STATUS = 0x2002
	// OAM_ADDR = OAM Address.
	OAM_ADDR = 0x2003
	// OAM_DATA = OAM Data.
	OAM_DATA = 0x2004
	// PPU_SCROLL = Background Scroll Position \newline (write X then Y).
	PPU_SCROLL = 0x2005
	// PPU_ADDR = PPU Address \newline (write upper then lower).
	PPU_ADDR = 0x2006
	// PPU_DATA = PPU Data.
	PPU_DATA = 0x2007

	// PALETTE_START = Universal background color.
	PALETTE_START = 0x3f00

	// PPU_OAM_DMA = Sprite Page DMA Transfer.
	PPU_OAM_DMA     = 0x4014
	DMC_FREQ        = 0x4010
	APU_STATUS      = 0x4015
	APU_NOISE_VOL   = 0x400C
	APU_NOISE_FREQ  = 0x400E
	APU_NOISE_TIMER = 0x400F
	APU_DMC_CTRL    = 0x4010
	APU_CHAN_CTRL   = 0x4015
	APU_FRAME       = 0x4017

	JOYPAD1 = 0x4016
	JOYPAD2 = 0x4017

	OAM_DMA = 0x4014
	OAM_RAM = 0x0200

	// PPU_CTRL flags
	CTRL_NMI      = 0b10000000 // Execute Non-Maskable Interrupt on VBlank
	CTRL_8x8      = 0b00000000 // Use 8x8 Sprites
	CTRL_8x16     = 0b00100000 // Use 8x16 Sprites
	CTRL_BG_0000  = 0b00000000 // Background Pattern Table at 0x0000 in VRAM
	CTRL_BG_1000  = 0b00010000 // Background Pattern Table at 0x1000 in VRAM
	CTRL_SPR_0000 = 0b00000000 // Sprite Pattern Table at 0x0000 in VRAM
	CTRL_SPR_1000 = 0b00001000 // Sprite Pattern Table at 0x1000 in VRAM
	CTRL_INC_1    = 0b00000000 // Increment PPU Address by 1 (Horizontal rendering)
	CTRL_INC_32   = 0b00000100 // Increment PPU Address by 32 (Vertical rendering)
	CTRL_NT_2000  = 0b00000000 // Name Table Address at 0x2000
	CTRL_NT_2400  = 0b00000001 // Name Table Address at 0x2400
	CTRL_NT_2800  = 0b00000010 // Name Table Address at 0x2800
	CTRL_NT_2C00  = 0b00000011 // Name Table Address at 0x2C00

	// PPU_MASK flags
	MASK_TINT_RED   = 0b00100000 // Red Background
	MASK_TINT_BLUE  = 0b01000000 // Blue Background
	MASK_TINT_GREEN = 0b10000000 // Green Background
	MASK_SPR        = 0b00010000 // Sprites Visible
	MASK_BG         = 0b00001000 // Backgrounds Visible
	MASK_SPR_CLIP   = 0b00000100 // Sprites clipped on left column
	MASK_BG_CLIP    = 0b00000010 // Background clipped on left column
	MASK_COLOR      = 0b00000000 // Display in Color
	MASK_MONO       = 0b00000001 // Display in Monochrome

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
