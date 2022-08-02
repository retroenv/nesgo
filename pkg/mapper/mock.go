package mapper

import (
	"github.com/retroenv/nesgo/pkg/bus"
	"github.com/retroenv/nesgo/pkg/cartridge"
	"github.com/retroenv/nesgo/pkg/memory"
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

// Name returns the name of the mapper.
func (m *MockMapper) Name() string {
	return ""
}

// MirrorMode returns the set mirror mode.
func (m *MockMapper) MirrorMode() cartridge.MirrorMode {
	return cartridge.MirrorHorizontal
}
