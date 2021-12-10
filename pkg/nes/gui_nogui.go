//go:build !nesgo && nogui
// +build !nesgo,nogui

package nes

var guiStarter = setupNoGui

func setupNoGui(_ *System) (guiRender func() (bool, error), guiCleanup func(), err error) {
	render := func() (bool, error) {
		return true, nil
	}
	cleanup := func() {}

	// reference to avoid unused linter warning
	_ = scaleFactor
	_ = windowTitle

	return render, cleanup, nil
}
