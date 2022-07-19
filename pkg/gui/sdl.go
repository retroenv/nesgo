//go:build !nesgo && !nogui && sdl
// +build !nesgo,!nogui,sdl

package gui

import (
	"fmt"

	"github.com/retroenv/nesgo/pkg/bus"
	"github.com/retroenv/nesgo/pkg/controller"
	"github.com/retroenv/nesgo/pkg/nes"
	"github.com/retroenv/nesgo/pkg/ppu"
	"github.com/veandco/go-sdl2/sdl"
)

func init() {
	nes.GuiStarter = setupSDLGui
}

var sdlKeyMapping = map[sdl.Keycode]controller.Button{
	sdl.K_UP:        controller.Up,
	sdl.K_DOWN:      controller.Down,
	sdl.K_LEFT:      controller.Left,
	sdl.K_RIGHT:     controller.Right,
	sdl.K_z:         controller.A,
	sdl.K_x:         controller.B,
	sdl.K_RETURN:    controller.Start,
	sdl.K_BACKSPACE: controller.Select,
}

func setupSDLGui(bus *bus.Bus) (guiRender func() (bool, error), guiCleanup func(), err error) {
	window, renderer, tex, err := setupSDL()
	if err != nil {
		return nil, nil, err
	}
	render := func() (bool, error) {
		return renderSDL(bus, renderer, tex)
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

	tex, err := renderer.CreateTexture(uint32(sdl.PIXELFORMAT_ABGR8888),
		sdl.TEXTUREACCESS_STREAMING, int32(ppu.Width), int32(ppu.Height))
	if err != nil {
		return nil, nil, nil, fmt.Errorf("creating SDL texture: %w", err)
	}

	return window, renderer, tex, nil
}

func renderSDL(bus *bus.Bus, renderer *sdl.Renderer, tex *sdl.Texture) (bool, error) {
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
			onSDLKey(bus, et)
		}
	}

	if err := tex.Update(nil, bus.PPU.Image().Pix, ppu.Width); err != nil {
		return false, err
	}
	if err := renderer.Copy(tex, nil, nil); err != nil {
		return false, err
	}
	renderer.Present()
	return running, nil
}

func onSDLKey(bus *bus.Bus, event *sdl.KeyboardEvent) {
	controllerKey, ok := sdlKeyMapping[event.Keysym.Sym]
	if !ok {
		return
	}
	switch event.Type {
	case sdl.KEYDOWN:
		bus.Controller1.SetButtonState(controllerKey, true)
	case sdl.KEYUP:
		bus.Controller1.SetButtonState(controllerKey, false)
	}
}
