package mapper

import (
	"fmt"

	"github.com/retroenv/nesgo/pkg/addressing"
	"github.com/retroenv/nesgo/pkg/bus"
)

type mapper0 struct {
	bus    *bus.Bus
	prgMod uint16
}

func (m mapper0) ReadMemory(address uint16) uint8 {
	switch {
	case address < 0x2000:
		return m.bus.Cartridge.CHR[address]
	case address >= addressing.CodeBaseAddress:
		offset := (address - addressing.CodeBaseAddress) % m.prgMod
		return m.bus.Cartridge.PRG[offset]
	default:
		panic(fmt.Sprintf("invalid read from address #%0000x", address))
	}
}

func (m mapper0) WriteMemory(address uint16, value uint8) {
}

func newMapper0(bus *bus.Bus) bus.BasicMemory {
	return &mapper0{
		bus:    bus,
		prgMod: uint16(len(bus.Cartridge.PRG)),
	}
}
