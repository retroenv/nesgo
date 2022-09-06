package cpu

import (
	"testing"

	"github.com/retroenv/retrogolib/assert"
)

// TestVerifyOpcodes verifies that all opcode and addressing mode info match.
func TestVerifyOpcodes(t *testing.T) {
	t.Parallel()

	for b, op := range Opcodes {
		ins := op.Instruction
		if ins == nil {
			continue
		}
		if ins.Unofficial && ins.Name == NopInstruction {
			// unofficial nop has multiple opcodes for the
			// same addressing mode
			continue
		}

		info := ins.Addressing[op.Addressing]
		assert.Equal(t, b, info.Opcode)
	}
}
