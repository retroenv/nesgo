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

	OAM_DMA = 0x4014

	// PPU_CTRL flags
	CTRL_NMI         = 0b10000000 // Execute Non-Maskable Interrupt on VBlank
	CTRL_MASTERSLAVE = 0b01000000 // Master/Slave select
	CTRL_8x8         = 0b00000000 // Use 8x8 Sprites
	CTRL_8x16        = 0b00100000 // Use 8x16 Sprites
	CTRL_BG_0000     = 0b00000000 // Background Pattern Table at 0x0000 in VRAM
	CTRL_BG_1000     = 0b00010000 // Background Pattern Table at 0x1000 in VRAM
	CTRL_SPR_0000    = 0b00000000 // Sprite Pattern Table at 0x0000 in VRAM
	CTRL_SPR_1000    = 0b00001000 // Sprite Pattern Table at 0x1000 in VRAM
	CTRL_INC_1       = 0b00000000 // Increment PPU Address by 1 (Horizontal rendering)
	CTRL_INC_32      = 0b00000100 // Increment PPU Address by 32 (Vertical rendering)
	CTRL_NT_2000     = 0b00000000 // Name Table Address at 0x2000
	CTRL_NT_2400     = 0b00000001 // Name Table Address at 0x2400
	CTRL_NT_2800     = 0b00000010 // Name Table Address at 0x2800
	CTRL_NT_2C00     = 0b00000011 // Name Table Address at 0x2C00

	// PPU_MASK flags
	MASK_COLOR      = 0b00000000 // Display in Color
	MASK_MONO       = 0b00000001 // Display in Monochrome
	MASK_BG_CLIP    = 0b00000010 // Background clipped on left column
	MASK_SPR_CLIP   = 0b00000100 // Sprites clipped on left column
	MASK_BG         = 0b00001000 // Backgrounds Visible
	MASK_SPR        = 0b00010000 // Sprites Visible
	MASK_TINT_RED   = 0b00100000 // Red Background
	MASK_TINT_BLUE  = 0b01000000 // Blue Background
	MASK_TINT_GREEN = 0b10000000 // Green Background
)
