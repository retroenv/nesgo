//go:build !nesgo

// Package nmi contains the PPU NMI manager.
package nmi

import "github.com/retroenv/nesgo/pkg/bus"

// Nmi implements a PPU NMI manager.
type Nmi struct {
	enabled  bool
	occurred bool
}

// New returns a new mask manager.
func New() *Nmi {
	return &Nmi{}
}

// Occurred returns whether an NMI occurred.
func (n *Nmi) Occurred() bool {
	return n.occurred
}

// SetOccurred sets the occurred flag.
func (n *Nmi) SetOccurred(occurred bool) {
	n.occurred = occurred
}

// Enabled returns whether NMI triggering is enabled.
func (n *Nmi) Enabled() bool {
	return n.enabled
}

// SetEnabled sets whether NMI is enabled.
func (n *Nmi) SetEnabled(enabled bool) {
	n.enabled = enabled
}

// Trigger an NMI if it occurred and is enabled.
func (n *Nmi) Trigger(cpu bus.CPU) {
	if n.enabled && n.occurred {
		cpu.TriggerNMI()
		n.occurred = false
	}
}
