//go:build !nesgo

package cpu

import (
	"fmt"
)

// instructionHook is a hook that is executed before a CPU instruction is executed.
// It allows for accounting of the instruction timing and trace logging.
// Params can be of length 0 to 2.
// At the end of the function the write lock is taken and a unlocker function returned.
func (c *CPU) instructionHook(instruction *Instruction, params ...any) func() {
	if !c.emulator {
		// trigger interrupt checking here as the system is not looping through the instructions in go mode
		c.CheckInterrupts()
	}

	startCycles := c.cycles

	if c.tracing == NoTracing {
		addressing := c.addressModeFromCall(instruction, params...)
		if !instruction.HasAddressing(addressing) {
			panic(fmt.Sprintf("unexpected addressing mode type %T", addressing))
		}

		opcode := instruction.Addressing[addressing].Opcode
		opcodeInfo := Opcodes[opcode]
		c.cycles += uint64(opcodeInfo.Timing)
	} else {
		if err := c.trace(instruction, params...); err != nil {
			panic(err)
		}
		c.cycles += uint64(c.TraceStep.Timing)

		if c.TraceStep.PageCrossed && c.TraceStep.PageCrossCycle {
			c.cycles++
		}
	}

	// this executes the ppu steps before the instruction
	cpuCycles := c.cycles - startCycles
	ppuCycles := cpuCycles * 3
	c.bus.PPU.Step(int(ppuCycles))

	return c.writeLock()
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
