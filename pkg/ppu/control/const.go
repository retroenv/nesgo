package control

const (
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
)
