package cartridge

// MirrorMode defines a mirror mode.
type MirrorMode int

// mirror modes.
const (
	MirrorHorizontal MirrorMode = iota
	MirrorVertical
	MirrorSingle0
	MirrorSingle1
	MirrorFour
)

// mirrorLookup maps the mirror mode nametables to nametable indexes.
var mirrorLookup = map[MirrorMode][4]uint16{
	MirrorHorizontal: {0, 0, 1, 1},
	MirrorVertical:   {0, 1, 0, 1},
	MirrorSingle0:    {0, 0, 0, 0},
	MirrorSingle1:    {1, 1, 1, 1},
	MirrorFour:       {0, 1, 2, 3},
}

// NametableIndexes returns the nametable indexes of the mirror mode.
func (m MirrorMode) NametableIndexes() [4]uint16 {
	indexes := mirrorLookup[m]
	return indexes
}
