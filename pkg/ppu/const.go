package ppu

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
)
