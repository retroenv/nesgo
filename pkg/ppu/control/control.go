//go:build !nesgo

// Package control contains the PPU control manager.
package control

type addressing interface {
	SetTempNameTables(nameTableX, nameTableY byte)
}

type nmi interface {
	SetEnabled(enabled bool)
}

type sprites interface {
	SetSpritePatternTable(address uint16)
	SetSpriteSize(size int)
}

type tiles interface {
	SetBackgroundPatternTable(backgroundPatternTable uint16)
}

// Control implements a PPU control manager.
type Control struct {
	addressing addressing
	nmi        nmi
	sprites    sprites
	tiles      tiles

	value byte // cached value since the fields are never modified directly

	BaseNameTable          uint16
	VRAMIncrement          uint8 // 0: add 1, going across; 1: add 32, going down
	SpritePatternTable     uint16
	BackgroundPatternTable uint16
	SpriteSize             uint8 // 0: 8x8 pixels; 1: 8x16 pixels
	MasterSlave            uint8
}

// New returns a new mask manager.
func New(addressing addressing, nmi nmi, sprites sprites, tiles tiles) *Control {
	return &Control{
		addressing: addressing,
		nmi:        nmi,
		sprites:    sprites,
		tiles:      tiles,
	}
}

// Set and extract the mask fields from given byte value.
func (c *Control) Set(value byte) {
	c.value = value

	c.BaseNameTable = (uint16(value&CTRL_NT_2C00) << 10) + 0x2000

	increment := (value & CTRL_INC_32) >> 2
	if increment == 0 {
		c.VRAMIncrement = 1
	} else {
		c.VRAMIncrement = 32
	}

	c.SpritePatternTable = uint16(value&CTRL_SPR_1000) << 9
	c.sprites.SetSpritePatternTable(c.SpritePatternTable)

	c.BackgroundPatternTable = uint16(value&CTRL_BG_1000) << 8
	c.tiles.SetBackgroundPatternTable(c.BackgroundPatternTable)

	c.SpriteSize = value & CTRL_8x16 >> 5
	if c.SpriteSize == 0 {
		c.sprites.SetSpriteSize(8)
	} else {
		c.sprites.SetSpriteSize(16)
	}

	c.MasterSlave = value & CTRL_MASTERSLAVE >> 6

	c.nmi.SetEnabled(value&CTRL_NMI != 0)

	nameTableX := value & CTRL_NT_2400
	nameTableY := value & CTRL_NT_2800 >> 1
	c.addressing.SetTempNameTables(nameTableX, nameTableY)
}

// Value returns the control fields encoded as byte.
func (c *Control) Value() byte {
	return c.value
}
