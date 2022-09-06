package neslib

import (
	"testing"

	. "github.com/retroenv/nesgo/pkg/nes"
	"github.com/retroenv/retrogolib/assert"
)

func TestDivSigned16(t *testing.T) {
	sys := NewSystem(nil)
	sys.LinkAliases()
	sys.ResetCycles()

	*A = 0b1000_0001 // -127
	DivSigned16()
	assert.Equal(t, 0b1111_1000, *A)  // -8
	assert.Equal(t, 14, sys.Cycles()) // -16

	*A = 0b0010_1010 // 42
	DivSigned16()
	assert.Equal(t, 0b0000_0010, *A) // 2
}

func TestDivSigned8(t *testing.T) {
	sys := NewSystem(nil)
	sys.LinkAliases()
	sys.ResetCycles()

	*A = 0b1000_0001 // -127
	DivSigned8()
	assert.Equal(t, 0b1111_0000, *A)  // -16
	assert.Equal(t, 12, sys.Cycles()) // -16

	*A = 0b0010_1010 // 42
	DivSigned8()
	assert.Equal(t, 0b0000_0101, *A) // 5
}
