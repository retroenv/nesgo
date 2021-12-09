package nes

import (
	"testing"

	"github.com/retroenv/nesgo/internal/assert"
)

func TestMemoryImmediate(t *testing.T) {
	t.Parallel()
	m := newMemory()

	i := NewUint8(0)
	m.writeMemoryAddressModes(1, i)
	assert.Equal(t, 1, *i)

	assert.Equal(t, 1, m.readMemoryAddressModes(true, i))
	assert.Equal(t, 1, m.readMemoryAddressModes(true, 1))
}

func TestMemoryAbsoluteInt(t *testing.T) {
	t.Parallel()
	m := newMemory()

	m.writeMemoryAddressModes(1, 2)
	assert.Equal(t, 1, m.readMemory(2))
	assert.Equal(t, 1, m.readMemoryAddressModes(false, 2))

	m.writeMemoryAddressModes(1, Absolute(3))
	assert.Equal(t, 1, m.readMemory(2))
	assert.Equal(t, 1, m.readMemoryAddressModes(false, Absolute(3)))
}

func TestMemoryAbsoluteIndirect(t *testing.T) {
	t.Parallel()
	m := newSystem()

	m.writeMemory(3, 0x00)
	m.writeMemory(4, 0x10)
	*m.memory.x = 1
	m.writeMemoryAddressModes(1, Indirect(2), m.memory.x)
	assert.Equal(t, 1, m.readMemory(0x10))

	m.writeMemory(8, 0x00)
	m.writeMemory(9, 0x18)
	*m.memory.y = 1
	m.writeMemoryAddressModes(1, Indirect(8), m.memory.y)
	assert.Equal(t, 1, m.readMemory(0x19))
}
