//go:build !nesgo
// +build !nesgo

package ppu

type ram interface {
	Reset()

	ReadMemory(address uint16) byte
	WriteMemory(address uint16, value byte)
}
