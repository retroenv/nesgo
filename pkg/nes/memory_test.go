package nes

import (
	"testing"

	"github.com/retroenv/nesgo/internal/assert"
)

func TestMemoryImmediate(t *testing.T) {
	m := newMemory()

	i := NewUint8(0)
	m.writeMemoryAddressModes(1, i)
	assert.Equal(t, 1, *i)

	assert.Equal(t, 1, m.readMemoryAddressModes(true, i))
	assert.Equal(t, 1, m.readMemoryAddressModes(true, 1))
}

func TestMemoryAbsoluteInt(t *testing.T) {
	m := newMemory()

	m.writeMemoryAddressModes(1, 2)
	assert.Equal(t, 1, m.readMemory(2))
	assert.Equal(t, 1, m.readMemoryAddressModes(false, 2))

	m.writeMemoryAddressModes(1, Absolute(3))
	assert.Equal(t, 1, m.readMemory(2))
	assert.Equal(t, 1, m.readMemoryAddressModes(false, Absolute(3)))
}
