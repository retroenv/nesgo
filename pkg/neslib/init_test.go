package neslib

import (
	"testing"

	"github.com/retroenv/nesgo/internal/assert"
	. "github.com/retroenv/nesgo/pkg/nes"
)

func TestClearRAM(t *testing.T) {
	sys := NewSystem(nil)
	sys.LinkAliases()

	sys.Bus.Memory.Write(0x7FF, 1)

	value := sys.Bus.Memory.Read(0x7FF)
	assert.Equal(t, 1, value)

	ClearRAM()

	value = sys.Bus.Memory.Read(0x7FF)
	assert.Equal(t, 0, value)
}
