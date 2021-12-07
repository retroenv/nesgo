//go:build !nesgo && !nogui && sdl
// +build !nesgo,!nogui,sdl

package nes

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

var guiStarter = setupSDLGui

var sdlKeyMapping = map[sdl.Keycode]button{
	sdl.K_UP:        buttonUp,
	sdl.K_DOWN:      buttonDown,
	sdl.K_LEFT:      buttonLeft,
	sdl.K_RIGHT:     buttonRight,
	sdl.K_z:         buttonA,
	sdl.K_x:         buttonB,
	sdl.K_RETURN:    buttonStart,
	sdl.K_BACKSPACE: buttonSelect,
}

func setupSDLGui(ppu *PPU) (guiRender func() (bool, error), guiCleanup func(), err error) {
	window, renderer, tex, err := setupSDL()
	if err != nil {
		return nil, nil, err
	}
	render := func() (bool, error) {
		return renderSDL(ppu, renderer, tex)
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
		sdl.WINDOWPOS_CENTERED, width*scaleFactor, height*scaleFactor,
		sdl.WINDOW_SHOWN)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("creating SDL window: %w", err)
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("creating SDL renderer: %w", err)
	}

	tex, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888,
		sdl.TEXTUREACCESS_STREAMING, int32(width), int32(height))
	if err != nil {
		return nil, nil, nil, fmt.Errorf("creating SDL texture: %w", err)
	}

	return window, renderer, tex, nil
}

func renderSDL(ppu *PPU, renderer *sdl.Renderer, tex *sdl.Texture) (bool, error) {
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
			onSDLKey(et)
		}
	}

	if err := tex.Update(nil, ppu.image.Pix, width); err != nil {
		return false, err
	}
	if err := renderer.Copy(tex, nil, nil); err != nil {
		return false, err
	}
	renderer.Present()
	return running, nil
}

func onSDLKey(event *sdl.KeyboardEvent) {
	controllerKey, ok := sdlKeyMapping[event.Keysym.Sym]
	if !ok {
		return
	}
	switch event.Type {
	case sdl.KEYDOWN:
		system.memory.controller1.setButtonState(controllerKey, true)
	case sdl.KEYUP:
		system.memory.controller1.setButtonState(controllerKey, false)
	}
}
