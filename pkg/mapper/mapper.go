//go:build !nesgo
// +build !nesgo

// Package mapper provides hardware mapper support.
// It maps CHR and PRG chips into the NES address space.
package mapper

import (
	"fmt"

	"github.com/retroenv/nesgo/pkg/cartridge"
)

// Mapper offers a normal memory access interface to access the mapper
// functionality.
type Mapper interface {
	ReadMemory(address uint16) uint8
	WriteMemory(address uint16, value uint8)
}

// New creates a new mapper for the mapper defined by the cartridge.
func New(cart *cartridge.Cartridge) (Mapper, error) {
	switch cart.Mapper {
	case 0:
		return newMapper0(cart), nil
	default:
		return nil, fmt.Errorf("mapper %d is not supported", cart.Mapper)
	}
}
