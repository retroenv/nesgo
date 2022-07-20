package neslib

import (
	"testing"

	"github.com/retroenv/nesgo/internal/assert"
	. "github.com/retroenv/nesgo/pkg/nes"
)

func TestDivSigned16(t *testing.T) {
	sys := NewSystem(nil)
	sys.LinkAliases()
	sys.ResetCycles()

	*A = 0b10000001 // -127
	DivSigned16()
	assert.Equal(t, 0b11111000, *A)   // -8
	assert.Equal(t, 14, sys.Cycles()) // -16

	*A = 0b00101010 // 42
	DivSigned16()
	assert.Equal(t, 0b00000010, *A) // 2
}

func TestDivSigned8(t *testing.T) {
	sys := NewSystem(nil)
	sys.LinkAliases()
	sys.ResetCycles()

	*A = 0b10000001 // -127
	DivSigned8()
	assert.Equal(t, 0b11110000, *A)   // -16
	assert.Equal(t, 12, sys.Cycles()) // -16

	*A = 0b00101010 // 42
	DivSigned8()
	assert.Equal(t, 0b00000101, *A) // 5
}
