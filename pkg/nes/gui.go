//go:build !nesgo
// +build !nesgo

package nes

const (
	scaleFactor = 2.0
	windowTitle = "nesgo"
)

func runRenderer(system *System) error {
	render, cleanup, err := guiStarter(system)
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
		system.ppu.startRender()

		system.ppu.renderScreen()

		running, err = render()
		if err != nil {
			return err
		}

		system.ppu.finishRender()
	}
	return nil
}
