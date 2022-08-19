package cpu

import "github.com/retroenv/nesgo/pkg/bus"

// State returns the current state of the CPU.
func (c *CPU) State() bus.CPUState {
	c.mu.RLock()
	defer c.mu.RUnlock()

	state := bus.CPUState{
		A:      c.A,
		X:      c.X,
		Y:      c.Y,
		PC:     c.PC,
		SP:     c.SP,
		Cycles: c.cycles,
		Flags: bus.CPUFlags{
			C: c.Flags.C,
			Z: c.Flags.Z,
			I: c.Flags.I,
			D: c.Flags.D,
			B: c.Flags.B,
			V: c.Flags.V,
			N: c.Flags.N,
		},
		Interrupts: bus.CPUInterrupts{
			NMITriggered: c.triggerNmi,
			NMIRunning:   c.nmiRunning,
			IrqTriggered: c.triggerIrq,
			IrqRunning:   c.irqRunning,
		},
	}
	return state
}
