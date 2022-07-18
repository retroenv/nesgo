package mapper

import (
	"fmt"

	"github.com/retroenv/nesgo/pkg/addressing"
	"github.com/retroenv/nesgo/pkg/cartridge"
)

type mapper0 struct {
	cart   *cartridge.Cartridge
	prgMod uint16
}

func (m mapper0) ReadMemory(address uint16) uint8 {
	switch {
	case address < 0x2000:
		return m.cart.CHR[address]
	case address >= addressing.CodeBaseAddress:
		offset := (address - addressing.CodeBaseAddress) % m.prgMod
		return m.cart.PRG[offset]
	default:
		panic(fmt.Sprintf("invalid read from address #%0000x", address))
	}
}

func (m mapper0) WriteMemory(address uint16, value uint8) {
}

func newMapper0(cart *cartridge.Cartridge) mapper0 {
	return mapper0{
		cart:   cart,
		prgMod: uint16(len(cart.PRG)),
	}
}
