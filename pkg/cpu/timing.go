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

	c.cycles += uint64(c.TraceStep.Timing)
	if c.TraceStep.PageCrossed && c.TraceStep.PageCrossCycle {
		c.cycles++
	}

	// TODO slow down emulation and add option to disable it
	time.Sleep(time.Microsecond)
}

// AccountBranchingPageCrossCycle accounts for a branch page crossing extra CPU cycle.
func (c *CPU) AccountBranchingPageCrossCycle(ins *Instruction) {
	if _, ok := BranchingInstructions[ins.Name]; !ok {
		return
	}
	if ins.Name != jmp.Name && ins.Name != jsr.Name {
		c.cycles++
	}
}
