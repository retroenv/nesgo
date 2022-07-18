// Package system provides an initializer for the NES system.
package system

import (
	"github.com/retroenv/nesgo/pkg/cartridge"
	"github.com/retroenv/nesgo/pkg/controller"
	"github.com/retroenv/nesgo/pkg/cpu"
	"github.com/retroenv/nesgo/pkg/mapper"
	"github.com/retroenv/nesgo/pkg/memory"
	"github.com/retroenv/nesgo/pkg/ppu"
)

// System implements a NES system.
type System struct {
	*cpu.CPU
	*memory.Memory

	PPU         *ppu.PPU
	Controller1 *controller.Controller
	Controller2 *controller.Controller

	NmiHandler   func()
	IrqHandler   func()
	ResetHandler func()
}

// New creates a new NES system.
func New(cart *cartridge.Cartridge) *System {
	mapp, err := mapper.New(cart)
	if err != nil {
		panic(err)
	}

	sys := &System{
		PPU:         ppu.New(memory.NewRAM(0x2000), mapp, nil), // TODO use bus
		Controller1: controller.New(),
		Controller2: controller.New(),
	}

	sys.Memory = memory.New(cart, sys.PPU, sys.Controller1, sys.Controller2, mapp)
	sys.CPU = cpu.New(sys.Memory, &sys.IrqHandler)
	return sys
}
