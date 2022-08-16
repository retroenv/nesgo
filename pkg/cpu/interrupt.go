package cpu

// CheckInterrupts checks for triggered interrupts and executes them.
func (c *CPU) CheckInterrupts() {
	if c.triggerNmi {
		c.nmi()
	}
	if c.triggerIrq {
		c.irq()
	}
}

func (c *CPU) nmi() {
	c.triggerNmi = false
	c.executeInterrupt(c.nmiHandler, c.nmiAddress)
}

func (c *CPU) irq() {
	c.triggerIrq = false
	c.executeInterrupt(c.irqHandler, c.irqAddress)
}

func (c *CPU) executeInterrupt(goFun *func(), funAddress uint16) {
	c.Push16(c.PC)
	c.phpInternal()

	if *goFun != nil {
		c.Flags.I = 1
		c.cycles += 7
		f := *goFun
		f()
		return
	}

	if funAddress != 0 {
		c.Flags.I = 1
		c.cycles += 7
		c.PC = funAddress
	}
}
