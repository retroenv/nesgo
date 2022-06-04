//go:build !nesgo && !nogui && !noopengl
// +build !nesgo,!nogui,!noopengl

package nes

import (
	"fmt"
	"image"
	"runtime"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/retroenv/nesgo/pkg/controller"
	"github.com/retroenv/nesgo/pkg/ppu"
)

var guiStarter = setupOpenGLGui

var openGLKeyMapping = map[glfw.Key]controller.Button{
	glfw.KeyUp:        controller.ButtonUp,
	glfw.KeyDown:      controller.ButtonDown,
	glfw.KeyLeft:      controller.ButtonLeft,
	glfw.KeyRight:     controller.ButtonRight,
	glfw.KeyZ:         controller.ButtonA,
	glfw.KeyX:         controller.ButtonB,
	glfw.KeyEnter:     controller.ButtonStart,
	glfw.KeyBackspace: controller.ButtonSelect,
}

func setupOpenGLGui(system *System) (guiRender func() (bool, error), guiCleanup func(), err error) {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()

	window, texture, err := setupOpenGL(system)
	if err != nil {
		return nil, nil, err
	}
	render := func() (bool, error) {
		img := system.ppu.Image()
		renderOpenGL(img, window, texture)
		return !window.ShouldClose(), nil
	}
	cleanup := func() {
		gl.DeleteTextures(1, &texture)
		glfw.Terminate()
	}
	return render, cleanup, nil
}

func setupOpenGL(system *System) (*glfw.Window, uint32, error) {
	// setup GLFW
	if err := glfw.Init(); err != nil {
		return nil, 0, fmt.Errorf("initializing GLFW: %w", err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)

	window, err := glfw.CreateWindow(ppu.Width*scaleFactor, ppu.Height*scaleFactor,
		windowTitle, nil, nil)
	if err != nil {
		return nil, 0, fmt.Errorf("creating GLFW window: %w", err)
	}

	window.SetKeyCallback(onGLFWKey(system))
	window.MakeContextCurrent()
	glfw.SwapInterval(1)

	// setup OpenGL
	if err = gl.Init(); err != nil {
		return nil, 0, fmt.Errorf("initializing OpenGL: %w", err)
	}
	gl.Enable(gl.TEXTURE_2D)
	var texture uint32
	gl.GenTextures(1, &texture)
	gl.BindTexture(gl.TEXTURE_2D, texture)
	img := system.ppu.Image()
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, ppu.Width, ppu.Height,
		0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(&img.Pix[0]))

	return window, texture, nil
}

func renderOpenGL(img *image.RGBA, window *glfw.Window, texture uint32) {
	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.TexSubImage2D(gl.TEXTURE_2D, 0, 0, 0, ppu.Width, ppu.Height,
		gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(&img.Pix[0]))

	// disable any filtering to avoid blurring the texture
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)

	// set an orthogonal projection (2D) with the size of the screen
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(0.0, ppu.Width, 0.0, ppu.Height, -1.0, 1.0)
	gl.MatrixMode(gl.MODELVIEW)

	// render a single quad with the size of the screen and with the
	// contents of the emulator frame buffer
	gl.Begin(gl.QUADS)
	gl.TexCoord2d(0.0, 1.0)
	gl.Vertex2d(0.0, 0.0)
	gl.TexCoord2d(1.0, 1.0)
	gl.Vertex2d(ppu.Width, 0.0)
	gl.TexCoord2d(1.0, 0.0)
	gl.Vertex2d(ppu.Width, ppu.Height)
	gl.TexCoord2d(0.0, 0.0)
	gl.Vertex2d(0.0, ppu.Height)
	gl.End()

	window.SwapBuffers()
	glfw.PollEvents()
}

func onGLFWKey(system *System) func(window *glfw.Window, key glfw.Key, _ int, action glfw.Action, _ glfw.ModifierKey) {
	return func(window *glfw.Window, key glfw.Key, _ int, action glfw.Action, _ glfw.ModifierKey) {
		if action == glfw.Press && key == glfw.KeyEscape {
			window.SetShouldClose(true)
		}

		controllerKey, ok := openGLKeyMapping[key]
		if !ok {
			return
		}
		switch action {
		case glfw.Press:
			system.controller1.SetButtonState(controllerKey, true)
		case glfw.Release:
			system.controller1.SetButtonState(controllerKey, false)
		}
	}
}
