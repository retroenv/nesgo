//go:build !nesgo

package nes

import (
	"time"

	"github.com/retroenv/nesgo/pkg/bus"
	"github.com/retroenv/nesgo/pkg/ppu"
)

func setupNoGui(_ *bus.Bus) (guiRender func() (bool, error), guiCleanup func(), err error) {
	render := func() (bool, error) {
		time.Sleep(time.Second / ppu.FPS)
		return true, nil
	}
	cleanup := func() {}

	return render, cleanup, nil
}
