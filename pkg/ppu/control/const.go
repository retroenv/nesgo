package control

const (
	// TODO remove duplicate from ppu once compiler supports parsing subdirectories

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
)
