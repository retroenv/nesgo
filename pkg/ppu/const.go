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
	CTRL_NMI         = 0b1000_0000 // Execute Non-Maskable Interrupt on VBlank
	CTRL_MASTERSLAVE = 0b0100_0000 // Master/Slave select
	CTRL_8x8         = 0b0000_0000 // Use 8x8 Sprites
	CTRL_8x16        = 0b0010_0000 // Use 8x16 Sprites
	CTRL_BG_0000     = 0b0000_0000 // Background Pattern Table at 0x0000 in VRAM
	CTRL_BG_1000     = 0b0001_0000 // Background Pattern Table at 0x1000 in VRAM
	CTRL_SPR_0000    = 0b0000_0000 // Sprite Pattern Table at 0x0000 in VRAM
	CTRL_SPR_1000    = 0b0000_1000 // Sprite Pattern Table at 0x1000 in VRAM
	CTRL_INC_1       = 0b0000_0000 // Increment PPU Address by 1 (Horizontal rendering)
	CTRL_INC_32      = 0b0000_0100 // Increment PPU Address by 32 (Vertical rendering)
	CTRL_NT_2000     = 0b0000_0000 // Name Table Address at 0x2000
	CTRL_NT_2400     = 0b0000_0001 // Name Table Address at 0x2400
	CTRL_NT_2800     = 0b0000_0010 // Name Table Address at 0x2800
	CTRL_NT_2C00     = 0b0000_0011 // Name Table Address at 0x2C00

	// PPU_MASK flags
	MASK_COLOR      = 0b0000_0000 // Display in Color
	MASK_MONO       = 0b0000_0001 // Display in Monochrome
	MASK_BG_CLIP    = 0b0000_0010 // Background clipped on left column
	MASK_SPR_CLIP   = 0b0000_0100 // Sprites clipped on left column
	MASK_BG         = 0b0000_1000 // Backgrounds Visible
	MASK_SPR        = 0b0001_0000 // Sprites Visible
	MASK_TINT_RED   = 0b0010_0000 // Red Background
	MASK_TINT_BLUE  = 0b0100_0000 // Blue Background
	MASK_TINT_GREEN = 0b1000_0000 // Green Background
)
