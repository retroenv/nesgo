package cpu

import (
	"testing"

	"github.com/retroenv/nesgo/internal/assert"
)

// TestVerifyOpcodes verifies that all opcode and addressing mode info match.
func TestVerifyOpcodes(t *testing.T) {
	t.Parallel()

	for b, op := range Opcodes {
		info := op.Instruction.Addressing[op.Addressing]
		assert.Equal(t, b, info.Opcode)
	}
}
