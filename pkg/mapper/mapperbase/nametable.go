package mapperbase

import (
	"fmt"

	"github.com/retroenv/nesgo/pkg/cartridge"
	"github.com/retroenv/nesgo/pkg/ppu/nametable"
)

// MirrorModeTranslation maps a 8bit index to a nametable mirror mode.
type MirrorModeTranslation map[uint8]cartridge.MirrorMode

// createNameTableBanks creates the VRAM banks.
func (b *Base) createNameTableBanks() {
	b.nameTableBanks = make([]bank, b.nameTableCount)

	for i := 0; i < b.nameTableCount; i++ {
		bank := &b.nameTableBanks[i]
		bank.length = nametable.VramSize
		bank.data = make([]byte, bank.length)
	}

	b.SetNameTableWindow(0)
}

// SetNameTableCount sets amount of nametables.
func (b *Base) SetNameTableCount(count int) {
	b.nameTableCount = count
}

// SetNameTableWindow sets the nametable window to a specific bank.
func (b *Base) SetNameTableWindow(bank int) {
	bank %= len(b.nameTableBanks)
	nameTable := &b.nameTableBanks[bank]
	b.bus.NameTable.SetVRAM(nameTable.data)
}

// NameTable returns the nametable buffer of a specific bank. Used in tests.
func (b *Base) NameTable(bank int) []byte {
	bank %= len(b.nameTableBanks)
	nameTable := &b.nameTableBanks[bank]
	return nameTable.data
}

// SetNameTableMirrorMode sets the nametable mirror mode.
func (b *Base) SetNameTableMirrorMode(mirrorMode cartridge.MirrorMode) {
	b.bus.NameTable.SetMirrorMode(mirrorMode)
}

// MirrorMode returns the set mirror mode.
func (b *Base) MirrorMode() cartridge.MirrorMode {
	return b.bus.NameTable.MirrorMode()
}

// SetMirrorModeTranslation set the mirror mode translation map.
func (b *Base) SetMirrorModeTranslation(translation MirrorModeTranslation) {
	b.mirrorModeTranslation = translation
}

// SetNameTableMirrorModeIndex sets the nametable mirror mode based on the value of the mapper based
// translation map from index to mirror mode.
func (b *Base) SetNameTableMirrorModeIndex(index uint8) {
	mode, ok := b.mirrorModeTranslation[index]
	if !ok {
		panic(fmt.Sprintf("invalid nametable mirror mode index %d", index))
	}
	b.bus.NameTable.SetMirrorMode(mode)
}
