//go:build !nesgo

// Package sprites handles PPU sprites handling.
package sprites

import (
	"github.com/retroenv/nesgo/pkg/bus"
)

const (
	defaultSpriteSize  = 8  // zero value of control bit = 8x8 pixels
	maxSprites         = 64 // maximum amount of sprites
	maxSpritesOnScreen = 8  // maximum amount of sprites on screen
	oamMemorySize      = 256
	spriteStructSize   = 4 // size of sprite structure
)

type renderState interface {
	Cycle() int
	ScanLine() int
}

type status interface {
	SetSpriteOverflow(value bool)
}

// Sprites implements PPU sprites support.
type Sprites struct {
	cpu         bus.CPU
	mapper      bus.Mapper
	memory      bus.Memory
	renderState renderState
	status      status

	spriteSize         int // 8: 8x8 pixels; 16: 8x16 pixels
	spritePatternTable uint16

	sprites            [maxSprites]Sprite
	patterns           [maxSpritesOnScreen]uint32
	visibleSprites     [maxSpritesOnScreen]int // value is index of sprite in sprites
	visibleSpriteCount int                     // number of sprites on the screen
	address            byte                    // next address in sprites buffer to access
}

// New returns a new sprites manager.
func New(cpu bus.CPU, mapper bus.Mapper, memory bus.Memory, renderState renderState, status status) *Sprites {
	return &Sprites{
		cpu:         cpu,
		mapper:      mapper,
		memory:      memory,
		renderState: renderState,
		status:      status,

		spriteSize: defaultSpriteSize,
	}
}

// SetAddress sets the address for the next read or write operation.
func (s *Sprites) SetAddress(address byte) {
	s.address = address
}

// SetSpriteSize sets the sprite height in pixels, 8 or 16.
func (s *Sprites) SetSpriteSize(size int) {
	s.spriteSize = size
}

// SetSpritePatternTable sets the address of the sprite pattern table.
func (s *Sprites) SetSpritePatternTable(table uint16) {
	s.spritePatternTable = table * 0x1000
}

// Read a sprite field, based on the previously set address.
func (s *Sprites) Read() byte {
	// TODO handle special case of reading during rendering
	index := s.address / spriteStructSize
	field := s.address % spriteStructSize

	sprite := &s.sprites[index]
	return sprite.field(field)
}

// Write to a sprite field, based on the previously set address.
func (s *Sprites) Write(value byte) {
	index := s.address / spriteStructSize
	field := s.address % spriteStructSize

	sprite := &s.sprites[index]
	sprite.setField(field, value)

	s.address++

	// TODO handle scroll glitch
}

// WriteDMA writes all sprite fields using Direct memory access mode.
func (s *Sprites) WriteDMA(value byte) {
	address := uint16(value) << 8

	for i := 0; i < oamMemorySize; i++ {
		data := s.memory.Read(address)
		s.Write(data)
		address++
	}

	// 1 wait state cycle while waiting for writes to complete,
	// +1 if on an odd CPU cycle, then 256 alternating read/write cycles
	stall := uint16(1 + 256 + 256)
	if s.cpu.Cycles()%2 == 1 {
		stall++
	}
	s.cpu.StallCycles(stall)
}

// Render executes a sprites render cycle.
func (s *Sprites) Render() {
	if s.renderState.Cycle() != 257 {
		return
	}

	if s.renderState.ScanLine() < 240 {
		s.evaluate()
	} else {
		s.visibleSpriteCount = 0
	}
}

// Pixel returns the rendered sprite pixel for the current render state position.
// The returned values are:
// 1. priority of the sprite of which a pixel is drawn
// 2. flag whether the sprite is sprite with index 0
// 3. the color pattern of the sprite pixel
func (s *Sprites) Pixel() (bool, bool, byte) {
	cycle := s.renderState.Cycle() - 1

	for i := 0; i < s.visibleSpriteCount; i++ {
		index := s.visibleSprites[i]
		sprite := &s.sprites[index]

		offset := cycle - int(sprite.x)
		if offset < 0 || offset > 7 {
			continue
		}
		offset = 7 - offset

		color := byte((s.patterns[i] >> byte(offset*4)) & 0x0F)
		if color%4 == 0 {
			continue
		}

		zeroHit := i == 0
		priority := sprite.priority()
		return priority, zeroHit, color
	}
	return false, false, 0
}

// Sprite evaluation does not cause sprite 0 hit. This is handled by sprite rendering instead.
func (s *Sprites) evaluate() {
	s.visibleSpriteCount = 0

	for i := 0; i < maxSprites; i++ {
		sprite := &s.sprites[i]
		row := s.renderState.ScanLine() - int(sprite.y)
		if row < 0 || row >= s.spriteSize {
			continue
		}

		if s.visibleSpriteCount < maxSpritesOnScreen {
			s.patterns[s.visibleSpriteCount] = s.fetchSpritePattern(sprite, row)
			s.visibleSprites[s.visibleSpriteCount] = i
		}
		s.visibleSpriteCount++
	}

	if s.visibleSpriteCount > maxSpritesOnScreen {
		s.visibleSpriteCount = maxSpritesOnScreen
		s.status.SetSpriteOverflow(true)
	}
}

func (s *Sprites) fetchSpritePattern(sprite *Sprite, row int) uint32 {
	tile := sprite.index
	var address uint16

	if s.spriteSize == 8 {
		if sprite.flipVertically() {
			row = 7 - row
		}
		address = s.spritePatternTable + uint16(tile)*16 + uint16(row)
	} else {
		if sprite.flipVertically() {
			row = 15 - row
		}
		table := tile & 1
		tile &= 0xFE
		if row > 7 {
			tile++
			row -= 8
		}
		address = 0x1000*uint16(table) + uint16(tile)*16 + uint16(row)
	}

	a := (sprite.attributes & 3) << 2
	lowTileByte := s.mapper.Read(address)
	highTileByte := s.mapper.Read(address + 8)

	var data uint32
	for i := 0; i < maxSpritesOnScreen; i++ {
		var p1, p2 byte
		if sprite.flipHorizontally() {
			p1 = (lowTileByte & 1) << 0
			p2 = (highTileByte & 1) << 1
			lowTileByte >>= 1
			highTileByte >>= 1
		} else {
			p1 = (lowTileByte & 0x80) >> 7
			p2 = (highTileByte & 0x80) >> 6
			lowTileByte <<= 1
			highTileByte <<= 1
		}
		data <<= 4
		data |= uint32(a | p1 | p2)
	}
	return data
}
