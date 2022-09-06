package nametable

import (
	"testing"

	"github.com/retroenv/retrogolib/assert"
	"github.com/retroenv/retrogolib/nes/cartridge"
)

func TestNameTable(t *testing.T) {
	t.Parallel()

	n := New(cartridge.MirrorHorizontal)
	n.SetVRAM(make([]byte, VramSize))

	n.vram[0] = 1

	value := n.Read(0x2400)
	assert.Equal(t, 1, value)

	n.mirrorMode = cartridge.MirrorVertical
	value = n.Read(0x2400)
	assert.Equal(t, 0, value)

	n.Fetch(0x2000)
	value = n.Value()
	assert.Equal(t, 1, value)
}
