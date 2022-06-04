//go:build !nesgo
// +build !nesgo

// Package gui implements multiple emulator GUIs.
package gui

import (
	"github.com/retroenv/nesgo/pkg/system"
)

const (
	scaleFactor = 2.0
	windowTitle = "nesgo"
)

// RunRenderer starts the chosen GUI renderer.
func RunRenderer(sys *system.System) error {
	render, cleanup, err := guiStarter(sys)
	if err != nil {
		return err
	}
	defer cleanup()

	go func() {
		sys.ResetHandler()
		for { // forever loop in case reset handler returns
		}
	}()

	running := true
	for running {
		sys.PPU.StartRender()

		sys.PPU.RenderScreen()

		running, err = render()
		if err != nil {
			return err
		}

		sys.PPU.FinishRender()
	}
	return nil
}
