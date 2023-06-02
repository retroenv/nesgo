package memory

import (
	"testing"

	. "github.com/retroenv/retrogolib/addressing"
	"github.com/retroenv/retrogolib/assert"
)

func TestMemoryImmediate(t *testing.T) {
	t.Parallel()
	m := New(nil)

	i := new(uint8)
	m.WriteAddressModes(1, i)
	assert.Equal(t, 1, *i)

	assert.Equal(t, 1, m.ReadAddressModes(true, i))
	assert.Equal(t, 1, m.ReadAddressModes(true, 1))
}

func TestMemoryAbsoluteInt(t *testing.T) {
	t.Parallel()
	m := New(nil)

	m.WriteAddressModes(1, 2)
	assert.Equal(t, 1, m.Read(2))
	assert.Equal(t, 1, m.ReadAddressModes(false, 2))

	m.WriteAddressModes(1, Absolute(3))
	assert.Equal(t, 1, m.Read(2))
	assert.Equal(t, 1, m.ReadAddressModes(false, Absolute(3)))
}

func TestMemoryAbsoluteIndirect(t *testing.T) {
	t.Parallel()
	m := New(nil)
	x := new(uint8)
	y := new(uint8)
	m.LinkRegisters(x, y, x, y)

	m.Write(3, 0x00)
	m.Write(4, 0x10)
	*x = 1
	m.WriteAddressModes(1, Indirect(2), x)
	assert.Equal(t, 1, m.Read(0x1000))

	m.Write(8, 0x00)
	m.Write(9, 0x18)
	*y = 1
	m.WriteAddressModes(1, Indirect(8), y)
	assert.Equal(t, 1, m.Read(0x1800))
}

func TestReadWord(t *testing.T) {
	m := New(nil)
	m.Write(0, 1)
	m.Write(1, 2)
	assert.Equal(t, 0x201, m.ReadWord(0))
}

func TestReadWordBug(t *testing.T) {
	m := New(nil)
	m.Write(0x2ff, 1)
	m.Write(0x200, 2)
	assert.Equal(t, 0x201, m.ReadWordBug(0x02FF))
}

func TestWriteWord(t *testing.T) {
	m := New(nil)
	m.WriteWord(0, 0x201)
	assert.Equal(t, 0x201, m.ReadWord(0))
}
