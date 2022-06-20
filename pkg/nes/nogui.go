//go:build !nesgo
// +build !nesgo

package nes

import "github.com/retroenv/nesgo/pkg/system"

func setupNoGui(_ *system.System) (guiRender func() (bool, error), guiCleanup func(), err error) {
	render := func() (bool, error) {
		return true, nil
	}
	cleanup := func() {}

	return render, cleanup, nil
}
