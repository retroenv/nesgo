//go:build !nesgo
// +build !nesgo

package nes

func init() {
	guiStarter = setupNoGui
}

func setupNoGui() (guiRender func() (bool, error), guiCleanup func(), err error) {
	render := func() (bool, error) {
		return true, nil
	}
	cleanup := func() {}
	return render, cleanup, nil
}
