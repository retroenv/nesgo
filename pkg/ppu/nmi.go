//go:build !nesgo
// +build !nesgo

package ppu

import "github.com/retroenv/nesgo/pkg/bus"

type nmi struct {
	occurred bool
	output   bool
	previous bool
	delay    byte
}

// change triggers internal handling after a change to the flags.
func (n *nmi) change() {
	nmi := n.output && n.occurred
	if nmi && !n.previous {
		// TODO: this fixes some games but the delay shouldn't have to be so
		// long, so the timings are off somewhere
		n.delay = 15
	}
	n.previous = nmi
}

func (n *nmi) checkTrigger(cpu bus.CPU) {
	if n.delay == 0 {
		return
	}

	n.delay--
	if n.delay == 0 && n.output && n.occurred {
		cpu.TriggerNMI()
	}
}
