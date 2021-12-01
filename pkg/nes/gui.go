//go:build !nesgo
// +build !nesgo

package nes

const (
	scaleFactor = 2.0
	windowTitle = "nesgo"
)

func runRenderer() error {
	render, cleanup, err := guiStarter()
	if err != nil {
		return err
	}
	defer cleanup()

	go func() {
		resetHandler()
		for { // forever loop in case reset handler returns
		}
	}()

	running := true
	for running {
		ppu.startRender()

		ppu.renderScreen()

		running, err = render()
		if err != nil {
			return err
		}

		ppu.finishRender()
	}
	return nil
}
