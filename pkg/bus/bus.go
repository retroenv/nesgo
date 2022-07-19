// Package bus provides a system Bus connecting all main system parts.
package bus

import (
	"image"

	"github.com/retroenv/nesgo/pkg/cartridge"
	"github.com/retroenv/nesgo/pkg/controller"
)

// Bus contains all NES sub system components.
// Since many components access other components, this structure
// allows an easy access and reduces the import dependencies and
// initialization order issues.
type Bus struct {
	Cartridge   *cartridge.Cartridge
	Controller1 Controller
	Controller2 Controller
	CPU         CPU
	Mapper      BasicMemory
	Memory      Memory
	PPU         PPU
}

// Controller represents a hardware controller.
type Controller interface {
	Read() uint8
	SetButtonState(key controller.Button, pressed bool)
	SetStrobeMode(mode uint8)
}

// CPU represents the Central Processing Unit.
type CPU interface {
	Cycles() uint64
	StallCycles(cycles uint16)
}

// BasicMemory represents a basic memory access interface.
type BasicMemory interface {
	Read(address uint16) uint8
	Write(address uint16, value uint8)
}

// Memory represents an advanced memory access interface.
type Memory interface {
	BasicMemory

	ReadAbsolute(address interface{}, register interface{}) byte
	ReadAddressModes(immediate bool, params ...interface{}) byte
	ReadWord(address uint16) uint16
	ReadWordBug(address uint16) uint16
	WriteAddressModes(value byte, params ...interface{})
	WriteWord(address, value uint16)

	LinkRegisters(x *uint8, y *uint8, globalX *uint8, globalY *uint8)
}

// PPU represents the Picture Processing Unit.
type PPU interface {
	BasicMemory

	Image() *image.RGBA

	FinishRender()
	RenderScreen()
	StartRender()
}
