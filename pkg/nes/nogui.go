//go:build !nesgo

package nes

import (
	"time"

	"github.com/retroenv/nesgo/pkg/ppu"
	"github.com/retroenv/retrogolib/gui"
)

func setupNoGui(_ gui.Backend) (guiRender func() (bool, error), guiCleanup func(), err error) {
	render := func() (bool, error) {
		time.Sleep(time.Second / ppu.FPS)
		return true, nil
	}
	cleanup := func() {}

	return render, cleanup, nil
}
