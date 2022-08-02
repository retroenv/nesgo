package cartridge

// MirrorMode defines a mirror mode.
type MirrorMode int

// mirror modes.
const (
	// MirrorHorizontal is a vertical arrangement of the nametables,
	// resulting in horizontal mirroring, which makes a 32x60 tilemap.
	MirrorHorizontal MirrorMode = iota
	// MirrorVertical is a horizontal arrangement of the nametables,
	// resulting in vertical mirroring, which makes a 64x30 tilemap.
	MirrorVertical
	// Single-screen mirroring is only available with certain mappers,
	// resulting in two 32x30 tilemaps.
	MirrorSingle0
	MirrorSingle1
	// Mirror4 offers 4 unique nametables can be addressed through the PPU bus,
	// creating a 64x60 tilemap, allowing for more flexible screen layouts.
	Mirror4
	MirrorDiagonal
	MirrorLShaped
	Mirror3Vertical
	Mirror3Horizontal
	Mirror3Diagonal
)

// mirrorLookup maps the mirror mode nametables to nametable indexes.
var mirrorLookup = map[MirrorMode][4]uint16{
	MirrorHorizontal:  {0, 0, 1, 1},
	MirrorVertical:    {0, 1, 0, 1},
	MirrorSingle0:     {0, 0, 0, 0},
	MirrorSingle1:     {1, 1, 1, 1},
	Mirror4:           {0, 1, 2, 3},
	MirrorDiagonal:    {0, 1, 1, 0},
	MirrorLShaped:     {0, 1, 1, 1},
	Mirror3Vertical:   {0, 2, 1, 2},
	Mirror3Horizontal: {0, 1, 2, 2},
	Mirror3Diagonal:   {0, 1, 1, 2},
}

// NametableIndexes returns the nametable indexes of the mirror mode.
func (m MirrorMode) NametableIndexes() [4]uint16 {
	indexes := mirrorLookup[m]
	return indexes
}
