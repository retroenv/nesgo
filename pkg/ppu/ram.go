//go:build !nesgo
// +build !nesgo

package ppu

import "github.com/retroenv/nesgo/pkg/mapper"

type ram interface {
	mapper.Memory

	Reset()
}
