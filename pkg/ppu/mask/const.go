package mask

const (
	// TODO remove duplicate from ppu once compiler supports parsing subdirectories

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
