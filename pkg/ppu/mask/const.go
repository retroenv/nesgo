package mask

const (
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
