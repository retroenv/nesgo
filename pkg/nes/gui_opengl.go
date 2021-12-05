//go:build !nesgo && !nogui && !noopengl
// +build !nesgo,!nogui,!noopengl

package nes

import (
	"fmt"
	"runtime"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

var guiStarter = setupOpenGLGui

var openGLKeyMapping = map[glfw.Key]button{
	glfw.KeyUp:        buttonUp,
	glfw.KeyDown:      buttonDown,
	glfw.KeyLeft:      buttonLeft,
	glfw.KeyRight:     buttonRight,
	glfw.KeyZ:         buttonA,
	glfw.KeyX:         buttonB,
	glfw.KeyEnter:     buttonStart,
	glfw.KeyBackspace: buttonSelect,
}

func setupOpenGLGui() (guiRender func() (bool, error), guiCleanup func(), err error) {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()

	window, texture, err := setupOpenGL()
	if err != nil {
		return nil, nil, err
	}
	render := func() (bool, error) {
		renderOpenGL(window, texture)
		return !window.ShouldClose(), nil
	}
	cleanup := func() {
		gl.DeleteTextures(1, &texture)
		glfw.Terminate()
	}
	return render, cleanup, nil
}

func setupOpenGL() (*glfw.Window, uint32, error) {
	// setup GLFW
	if err := glfw.Init(); err != nil {
		return nil, 0, fmt.Errorf("initializing GLFW: %w", err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)

	window, err := glfw.CreateWindow(width*scaleFactor, height*scaleFactor,
		windowTitle, nil, nil)
	if err != nil {
		return nil, 0, fmt.Errorf("creating GLFW window: %w", err)
	}

	window.SetKeyCallback(onGLFWKey)
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
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, width, height,
		0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(&ppu.image.Pix[0]))

	return window, texture, nil
}

func renderOpenGL(window *glfw.Window, texture uint32) {
	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.TexSubImage2D(gl.TEXTURE_2D, 0, 0, 0, width, height,
		gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(&ppu.image.Pix[0]))

	// disable any filtering to avoid blurring the texture
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)

	// set an orthogonal projection (2D) with the size of
	// the Game Boy Screen
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(0.0, width, 0.0, height, -1.0, 1.0)
	gl.MatrixMode(gl.MODELVIEW)

	// render a single quad with the size of the Game Boy
	// screen and with the contents of the emulator
	// frame buffer (already in the texture)
	gl.Begin(gl.QUADS)
	gl.TexCoord2d(0.0, 1.0)
	gl.Vertex2d(0.0, 0.0)
	gl.TexCoord2d(1.0, 1.0)
	gl.Vertex2d(width, 0.0)
	gl.TexCoord2d(1.0, 0.0)
	gl.Vertex2d(width, height)
	gl.TexCoord2d(0.0, 0.0)
	gl.Vertex2d(0.0, height)
	gl.End()

	window.SwapBuffers()
	glfw.PollEvents()
}

func onGLFWKey(window *glfw.Window, key glfw.Key, _ int, action glfw.Action, _ glfw.ModifierKey) {
	if action == glfw.Press && key == glfw.KeyEscape {
		window.SetShouldClose(true)
	}

	controllerKey, ok := openGLKeyMapping[key]
	if !ok {
		return
	}
	switch action {
	case glfw.Press:
		controller1.setButtonState(controllerKey, true)
	case glfw.Release:
		controller1.setButtonState(controllerKey, false)
	}
}
