//go:build !nesgo && !nogui && sdl
// +build !nesgo,!nogui,sdl

package gui

import (
	"fmt"

	"github.com/retroenv/nesgo/pkg/controller"
	"github.com/retroenv/nesgo/pkg/ppu"
	"github.com/retroenv/nesgo/pkg/system"
	"github.com/veandco/go-sdl2/sdl"
)

var guiStarter = setupSDLGui

var sdlKeyMapping = map[sdl.Keycode]controller.Button{
	sdl.K_UP:        controller.ButtonUp,
	sdl.K_DOWN:      controller.ButtonDown,
	sdl.K_LEFT:      controller.ButtonLeft,
	sdl.K_RIGHT:     controller.ButtonRight,
	sdl.K_z:         controller.ButtonA,
	sdl.K_x:         controller.ButtonB,
	sdl.K_RETURN:    controller.ButtonStart,
	sdl.K_BACKSPACE: controller.ButtonSelect,
}

func setupSDLGui(sys *system.System) (guiRender func() (bool, error), guiCleanup func(), err error) {
	window, renderer, tex, err := setupSDL()
	if err != nil {
		return nil, nil, err
	}
	render := func() (bool, error) {
		return renderSDL(sys, renderer, tex)
	}
	cleanup := func() {
		_ = tex.Destroy()
		_ = renderer.Destroy()
		_ = window.Destroy()
		sdl.Quit()
	}
	return render, cleanup, nil
}

func setupSDL() (*sdl.Window, *sdl.Renderer, *sdl.Texture, error) {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return nil, nil, nil, fmt.Errorf("initializing SDL: %w", err)
	}

	window, err := sdl.CreateWindow(windowTitle, sdl.WINDOWPOS_CENTERED,
		sdl.WINDOWPOS_CENTERED, ppu.Width*scaleFactor, ppu.Height*scaleFactor,
		sdl.WINDOW_SHOWN)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("creating SDL window: %w", err)
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("creating SDL renderer: %w", err)
	}

	tex, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888,
		sdl.TEXTUREACCESS_STREAMING, int32(ppu.Width), int32(ppu.Height))
	if err != nil {
		return nil, nil, nil, fmt.Errorf("creating SDL texture: %w", err)
	}

	return window, renderer, tex, nil
}

func renderSDL(sys *system.System, renderer *sdl.Renderer, tex *sdl.Texture) (bool, error) {
	running := true
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch et := event.(type) {
		case *sdl.QuitEvent:
			running = false
			break

		case *sdl.KeyboardEvent:
			if et.Type == sdl.KEYDOWN && et.Keysym.Sym == sdl.K_ESCAPE {
				running = false
				break
			}
			onSDLKey(sys, et)
		}
	}

	if err := tex.Update(nil, sys.PPU.Image().Pix, ppu.Width); err != nil {
		return false, err
	}
	if err := renderer.Copy(tex, nil, nil); err != nil {
		return false, err
	}
	renderer.Present()
	return running, nil
}

func onSDLKey(sys *system.System, event *sdl.KeyboardEvent) {
	controllerKey, ok := sdlKeyMapping[event.Keysym.Sym]
	if !ok {
		return
	}
	switch event.Type {
	case sdl.KEYDOWN:
		sys.Controller1.SetButtonState(controllerKey, true)
	case sdl.KEYUP:
		sys.Controller1.SetButtonState(controllerKey, false)
	}
}
