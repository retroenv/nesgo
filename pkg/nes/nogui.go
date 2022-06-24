//go:build !nesgo
// +build !nesgo

package nes

import (
	"time"

	"github.com/retroenv/nesgo/pkg/ppu"
	"github.com/retroenv/nesgo/pkg/system"
)

func setupNoGui(_ *system.System) (guiRender func() (bool, error), guiCleanup func(), err error) {
	render := func() (bool, error) {
		time.Sleep(time.Second / ppu.FPS)
		return true, nil
	}
	cleanup := func() {}

	return render, cleanup, nil
}
