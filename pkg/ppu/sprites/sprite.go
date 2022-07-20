package sprites

// Sprite defines a sprite that can be drawn on screen and moved.
type Sprite struct {
	y          byte // y position of top of sprite, sprite data is delayed by one scanline
	index      byte // Tile index number
	attributes byte
	x          byte // x position of left side of sprite.
}

func (s Sprite) field(index byte) byte {
	switch index {
	case 0:
		return s.y
	case 1:
		return s.index
	case 2:
		return s.attributes
	default:
		return s.x
	}
}

func (s *Sprite) setField(index, value byte) {
	switch index {
	case 0:
		s.y = value
	case 1:
		s.index = value
	case 2:
		s.attributes = value
	default:
		s.x = value
	}
}

// Priority returns the sprite priority attribute.
func (s Sprite) Priority() byte {
	priority := (s.attributes >> 5) & 1
	return priority
}

func (s Sprite) flipHorizontally() bool {
	flip := (s.attributes >> 6) & 1
	return flip == 1
}

func (s Sprite) flipVertically() bool {
	flip := (s.attributes >> 7) & 1
	return flip == 1
}
