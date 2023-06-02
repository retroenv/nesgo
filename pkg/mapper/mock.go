package mapper

import (
	"github.com/retroenv/nesgo/pkg/bus"
	"github.com/retroenv/nesgo/pkg/memory"
	"github.com/retroenv/retrogolib/arch/nes/cartridge"
)

// MockMapper implements a mock mapper for use in tests.
type MockMapper struct {
	*memory.Memory
}

// NewMockMapper returns a new mock mapper.
func NewMockMapper(bus *bus.Bus) bus.Mapper {
	return &MockMapper{
		Memory: memory.New(bus),
	}
}

// State returns the current state of the mapper.
func (m *MockMapper) State() bus.MapperState {
	return bus.MapperState{}
}

// MirrorMode returns the set mirror mode.
func (m *MockMapper) MirrorMode() cartridge.MirrorMode {
	return cartridge.MirrorHorizontal
}
