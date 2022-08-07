//go:build !nesgo

// Package tiles handles PPU tiles support.
package tiles

import (
	"github.com/retroenv/nesgo/pkg/bus"
)

type addressing interface {
	Address() uint16
	FineY() uint16
}

// Tiles implements PPU tiles support.
type Tiles struct {
	addressing addressing
	memory     bus.BasicMemory
	nameTable  bus.NameTable

	attribute              byte
	backgroundPatternTable uint16
	lowByte                byte
	highByte               byte
	data                   uint64
}

// New returns a new tiles manager.
func New(addressing addressing, memory bus.BasicMemory, nameTable bus.NameTable) *Tiles {
	return &Tiles{
		addressing: addressing,
		memory:     memory,
		nameTable:  nameTable,
	}
}

// FetchCycle runs a fetch cycle for tile data. Based on the current cycle, different data is fetched.
func (t *Tiles) FetchCycle(cycle int) {
	t.data <<= 4

	switch cycle % 8 {
	case 0:
		t.storeTileData()

	case 1:
		t.nameTable.Fetch(t.addressing.Address())

	case 3:
		t.fetchAttributeTableByte()

	case 5:
		address := t.tileAddress()
		t.lowByte = t.memory.Read(address)

	case 7:
		address := t.tileAddress()
		t.highByte = t.memory.Read(address + 8)
	}
}

// SetBackgroundPatternTable sets the temp register nametable from the passed PPU control byte.
func (t *Tiles) SetBackgroundPatternTable(backgroundPatternTable uint16) {
	t.backgroundPatternTable = backgroundPatternTable
}

// BackgroundPixel returns the background pixel for the given X coordinate.
func (t *Tiles) BackgroundPixel(fineX uint16) byte {
	data := uint32(t.data >> 32)
	shift := (7 - fineX) * 4
	data >>= shift
	data &= 0x0F
	return byte(data)
}

func (t *Tiles) fetchAttributeTableByte() {
	address := t.addressing.Address()
	shift := ((address >> 4) & 4) | (address & 2)
	address = 0x23C0 | (address & 0x0C00) | ((address >> 4) & 0x38) | ((address >> 2) & 0x07)

	value := t.memory.Read(address)
	t.attribute = ((value >> shift) & 3) << 2
}

func (t *Tiles) storeTileData() {
	var data uint32
	for i := 0; i < 8; i++ {
		a := t.attribute
		p1 := (t.lowByte & 0x80) >> 7
		p2 := (t.highByte & 0x80) >> 6
		t.lowByte <<= 1
		t.highByte <<= 1
		data <<= 4
		data |= uint32(a | p1 | p2)
	}
	t.data |= uint64(data)
}

func (t *Tiles) tileAddress() uint16 {
	table := t.backgroundPatternTable
	tile := t.nameTable.Value()
	address := 0x1000*table + uint16(tile)*16 + t.addressing.FineY()
	return address
}
