//go:build !nesgo
// +build !nesgo

package ppu

type control struct {
	value byte

	BaseNameTable          uint16
	VRAMIncrement          uint8 // 0: add 1, going across; 1: add 32, going down
	SpritePatternTable     uint16
	BackgroundPatternTable uint16
	SpriteSize             uint8 // 0: 8x8 pixels; 1: 8x16 pixels
	MasterSlave            uint8
	NmiOutput              bool
}

func (p *PPU) setControl(value byte) {
	p.control.value = value

	p.control.BaseNameTable = (uint16(value&CTRL_NT_2C00) << 10) + 0x2000

	increment := (value & CTRL_INC_32) >> 2
	if increment == 0 {
		p.control.VRAMIncrement = 1
	} else {
		p.control.VRAMIncrement = 32
	}

	p.control.SpritePatternTable = uint16(value&CTRL_SPR_1000) << 9
	p.control.BackgroundPatternTable = uint16(value&CTRL_BG_1000) << 8
	p.control.SpriteSize = value & CTRL_8x16 >> 5
	p.control.MasterSlave = value & CTRL_MASTERSLAVE >> 6
	p.control.NmiOutput = value&CTRL_NMI != 0

	p.tempAddress.NameTableX = uint16(value & CTRL_NT_2400)
	p.tempAddress.NameTableY = uint16(value&CTRL_NT_2800) >> 1
}
