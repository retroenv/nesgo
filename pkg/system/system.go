// Package system provides an initializer for the NES system.
package system

import (
	"github.com/retroenv/nesgo/pkg/bus"
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

	Bus *bus.Bus

	NmiHandler   func()
	IrqHandler   func()
	ResetHandler func()
}

// New creates a new NES system.
func New(cart *cartridge.Cartridge) *System {
	systemBus := &bus.Bus{
		Cartridge:   cart,
		Controller1: controller.New(),
		Controller2: controller.New(),
	}

	var err error
	systemBus.Mapper, err = mapper.New(systemBus)
	if err != nil {
		panic(err)
	}

	sys := &System{
		Bus:    systemBus,
		Memory: memory.New(systemBus),
	}

	sys.CPU = cpu.New(sys.Memory, &sys.IrqHandler)
	systemBus.CPU = sys.CPU
	systemBus.PPU = ppu.New(systemBus)
	return sys
}
