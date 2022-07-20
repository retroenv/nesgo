package nametable

import (
	"testing"

	"github.com/retroenv/nesgo/internal/assert"
	"github.com/retroenv/nesgo/pkg/cartridge"
	"github.com/retroenv/nesgo/pkg/memory"
)

func TestNameTable(t *testing.T) {
	t.Parallel()

	mapper := memory.NewRAM(0, maximumAddress+1)
	n := New(mapper)

	mapper.Write(0x2000, 1)

	value := n.Read(0x2400, cartridge.MirrorHorizontal)
	assert.Equal(t, 1, value)

	value = n.Read(0x2400, cartridge.MirrorVertical)
	assert.Equal(t, 0, value)

	n.Fetch(0x2000)
	value = n.Value()
	assert.Equal(t, 1, value)
}
