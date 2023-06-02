// Package mapperdb contains all mapper implementations.
package mapperdb

import (
	"github.com/retroenv/nesgo/pkg/bus"
	"github.com/retroenv/nesgo/pkg/mapper/mapperbase"
	"github.com/retroenv/retrogolib/arch/nes/cartridge"
)

// Base defines the base mapper interface that contains helper functions for shared functionality.
type Base interface {
	bus.Mapper

	ChrBankCount() int
	SetChrRAM(ram []byte)
	SetChrWindow(window, bank int)
	SetChrWindowSize(size int)

	PrgBankCount() int
	SetPrgRAM(ram []byte)
	SetPrgWindow(window, bank int)
	SetPrgWindowSize(size int)

	NameTable(bank int) []byte
	SetMirrorModeTranslation(translation mapperbase.MirrorModeTranslation)
	SetNameTableCount(count int)
	SetNameTableMirrorMode(mirrorMode cartridge.MirrorMode)
	SetNameTableMirrorModeIndex(index uint8)
	SetNameTableWindow(bank int)

	AddReadHook(startAddress, endAddress uint16, hookFunc func(address uint16) uint8) mapperbase.Hook
	AddWriteHook(startAddress, endAddress uint16, hookFunc func(address uint16, value uint8)) mapperbase.Hook
	Cartridge() *cartridge.Cartridge
	Initialize()
	SetName(name string)
}
