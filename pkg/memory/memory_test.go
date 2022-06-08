package memory

import (
	"testing"

	"github.com/retroenv/nesgo/internal/assert"
	. "github.com/retroenv/nesgo/pkg/addressing"
)

func TestMemoryImmediate(t *testing.T) {
	t.Parallel()
	m := New(nil, nil, nil, nil, nil)

	i := new(uint8)
	m.WriteMemoryAddressModes(1, i)
	assert.Equal(t, 1, *i)

	assert.Equal(t, 1, m.ReadMemoryAddressModes(true, i))
	assert.Equal(t, 1, m.ReadMemoryAddressModes(true, 1))
}

func TestMemoryAbsoluteInt(t *testing.T) {
	t.Parallel()
	m := New(nil, nil, nil, nil, nil)

	m.WriteMemoryAddressModes(1, 2)
	assert.Equal(t, 1, m.ReadMemory(2))
	assert.Equal(t, 1, m.ReadMemoryAddressModes(false, 2))

	m.WriteMemoryAddressModes(1, Absolute(3))
	assert.Equal(t, 1, m.ReadMemory(2))
	assert.Equal(t, 1, m.ReadMemoryAddressModes(false, Absolute(3)))
}

func TestMemoryAbsoluteIndirect(t *testing.T) {
	t.Parallel()
	m := New(nil, nil, nil, nil, nil)
	x := new(uint8)
	y := new(uint8)
	m.LinkRegisters(x, y, x, y)

	m.WriteMemory(3, 0x00)
	m.WriteMemory(4, 0x10)
	*x = 1
	m.WriteMemoryAddressModes(1, Indirect(2), x)
	assert.Equal(t, 1, m.ReadMemory(0x1000))

	m.WriteMemory(8, 0x00)
	m.WriteMemory(9, 0x18)
	*y = 1
	m.WriteMemoryAddressModes(1, Indirect(8), y)
	assert.Equal(t, 1, m.ReadMemory(0x1800))
}

func TestReadMemory16(t *testing.T) {
	m := New(nil, nil, nil, nil, nil)
	m.WriteMemory(0, 1)
	m.WriteMemory(1, 2)
	assert.Equal(t, 0x201, m.ReadMemory16(0))
}

func TestWriteMemory16(t *testing.T) {
	m := New(nil, nil, nil, nil, nil)
	m.WriteMemory16(0, 0x201)
	assert.Equal(t, 0x201, m.ReadMemory16(0))
}
