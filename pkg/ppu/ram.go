//go:build !nesgo
// +build !nesgo

package ppu

import (
	"github.com/retroenv/nesgo/pkg/bus"
)

type ram interface {
	bus.Memory

	Reset()
}
