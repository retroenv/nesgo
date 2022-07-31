// Package mapperdb contains all mapper implementations.
package mapperdb

import (
	"github.com/retroenv/nesgo/pkg/bus"
	"github.com/retroenv/nesgo/pkg/cartridge"
)

// Hook defines a hook type that can be configured after creation.
type Hook interface {
	SetProxyOnly(proxy bool)
}

// Base defines the base mapper interface that contains helper functions for shared functionality.
type Base interface {
	bus.Mapper

	ChrBankCount() int
	SetChrRAM(ram []byte)
	SetChrWindow(window, bank int)
	SetChrWindowSize(size int)

	PrgBankCount() int
	SetPrgWindow(window, bank int)
	SetPrgWindowSize(size int)

	NameTable(bank int) []byte
	SetNameTableCount(count int)
	SetNameTableMirrorMode(mirrorMode cartridge.MirrorMode)
	SetNameTableWindow(bank int)

	AddReadHook(startAddress, endAddress uint16, hookFunc func(address uint16) uint8) Hook
	AddWriteHook(startAddress, endAddress uint16, hookFunc func(address uint16, value uint8)) Hook
	Initialize()
	SetName(name string)
}
