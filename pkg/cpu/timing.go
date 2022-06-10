//go:build !nesgo
// +build !nesgo

package cpu

import (
	"time"
)

// instructionHook is a hook that is executed before a CPU instruction is executed.
// It allows for accounting of the instruction timing and trace logging.
// TODO add option to disable timing in unit tests
func (c *CPU) instructionHook(instruction *Instruction, params ...interface{}) {
	if c.tracing != NoTracing {
		c.trace(instruction, params...)
	}

	// TODO account for exact cycles
	time.Sleep(time.Microsecond)
}
