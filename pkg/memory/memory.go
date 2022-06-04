//go:build !nesgo
// +build !nesgo

// Package memory provides Memory functionality.
package memory

import (
	"fmt"
	"math"

	"github.com/retroenv/nesgo/pkg/addressing"
	"github.com/retroenv/nesgo/pkg/apu"
	"github.com/retroenv/nesgo/pkg/controller"
	"github.com/retroenv/nesgo/pkg/ines"
	"github.com/retroenv/nesgo/pkg/ppu"
)

// Memory implements the memory controller.
type Memory struct {
	ram         *RAM
	ppu         *ppu.PPU
	controller1 *controller.Controller
	controller2 *controller.Controller
	cartridge   *ines.Cartridge

	// point to X/Y for comparison of indirect register
	// parameters in unit tests.
	x, globalX *uint8
	y, globalY *uint8
}

// New returns a new memory instance, embedded it has
// new instances for PPU and the Controllers.
func New(cartridge *ines.Cartridge, ppu *ppu.PPU, controller1, controller2 *controller.Controller) *Memory {
	r := &Memory{
		ram:         NewRAM(0),
		ppu:         ppu,
		controller1: controller1,
		controller2: controller2,
		cartridge:   cartridge,
	}
	return r
}

// LinkRegisters points the internal x/y registers for unit test usage
// to the actual processor registers.
func (m *Memory) LinkRegisters(x *uint8, y *uint8, globalX *uint8, globalY *uint8) {
	m.x = x
	m.globalX = globalX
	m.y = y
	m.globalY = globalY
}

// WriteMemoryAddressModes writes to memory using different address modes:
// Absolute: the absolut memory address is used to write the value
// Absolute, X: the absolut memory address with offset from X is used
// Absolute, Y: the absolut memory address with offset from Y is used
// (Indirect, X): the absolut memory address to read the value from is
//                read from (indirect address + X)
// (Indirect), Y: the pointer to the memory address is read from the
//                indirect parameter and adjusted after reading it
//                by adding Y. The value is read from this pointer
func (m *Memory) WriteMemoryAddressModes(value byte, params ...interface{}) {
	param := params[0]
	var register interface{}
	if len(params) > 1 {
		register = params[1]
	}

	switch address := param.(type) {
	case int:
		m.writeMemoryAbsolute(address, value, register)
	case *uint8: // variable
		*address = value
	case addressing.Absolute:
		m.writeMemoryAbsolute(address, value, register)
	case addressing.Indirect:
		m.writeMemoryIndirect(address, value, register)
	default:
		panic(fmt.Sprintf("unsupported memory write addressing mode type %T", param))
	}
}

func (m *Memory) writeMemoryIndirect(address addressing.Indirect, value byte, register interface{}) {
	pointer := m.indirectMemoryPointer(address, register)
	m.WriteMemory(pointer, value)
}

func (m *Memory) writeMemoryAbsolute(address interface{}, value byte, register interface{}) {
	if register == nil {
		m.writeMemoryAbsoluteOffset(address, value, 0)
		return
	}

	var offset uint16
	switch val := register.(type) {
	case *uint8: // X/Y register referenced in normal code
		offset = uint16(*val)
	case uint8: // X/Y register referenced in unit test as system.X
		offset = uint16(val)
	default:
		panic(fmt.Sprintf("unsupported extra parameter type %T for absolute memory write", register))
	}

	m.writeMemoryAbsoluteOffset(address, value, offset)
}

func (m *Memory) writeMemoryAbsoluteOffset(address interface{}, value byte, offset uint16) {
	switch addr := address.(type) {
	case int8:
		m.WriteMemory(uint16(addr)+offset, value)
	case uint8:
		m.WriteMemory(uint16(addr)+offset, value)
	case *uint8:
		*addr = value
	case uint16:
		m.WriteMemory(addr+offset, value)
	case *uint16:
		*addr = uint16(value)
	case int:
		m.WriteMemory(uint16(addr)+offset, value)
	case addressing.Absolute:
		m.WriteMemory(uint16(addr)+offset, value)
	default:
		panic(fmt.Sprintf("unsupported address type %T for absolute memory write with register", address))
	}
}

// WriteMemory writes a byte to a memory address.
func (m *Memory) WriteMemory(address uint16, value byte) {
	switch {
	case address < 0x2000:
		m.ram.WriteMemory(address&0x07FF, value)
	case address < 0x4000:
		m.ppu.WriteRegister(address, value)
	case address == controller.JOYPAD1:
		m.controller1.SetStrobeMode(value)
		m.controller2.SetStrobeMode(value)
	case address == apu.DMC_FREQ:
		return // TODO
	case address == apu.APU_CHAN_CTRL:
		return // TODO
	case address == apu.APU_FRAME:
		return // TODO
	default:
		panic(fmt.Sprintf("unhandled memory write at address: 0x%04X", address))
	}
}

// ReadMemoryAddressModes reads memory using different address modes:
// Immediate: if immediate is true and the passed first param fits into
//            a byte, it's immediate value is returned
// Absolute: the absolut memory address is used to read the value
// Absolute, X: the absolut memory address with offset from X is used
// Absolute, Y: the absolut memory address with offset from Y is used
// (Indirect, X): the absolut memory address to write the value to is
//                read from (indirect address + X)
// (Indirect), Y: the pointer to the memory address is read from the
//                indirect parameter and adjusted after reading it
//                by adding Y. The value is written to this pointer
func (m *Memory) ReadMemoryAddressModes(immediate bool, params ...interface{}) byte {
	param := params[0]
	var register interface{}
	if len(params) > 1 {
		register = params[1]
	}

	switch address := param.(type) {
	case int:
		if immediate && register == nil && address <= math.MaxUint8 {
			return uint8(address) // immediate, not an address
		}
		return m.ReadMemoryAbsolute(address, register)
	case uint8:
		return address // immediate, not an address
	case *uint8: // variable
		return *address
	case addressing.Absolute:
		return m.ReadMemoryAbsolute(address, register)
	case addressing.Indirect:
		return m.readMemoryIndirect(address, register)
	default:
		panic(fmt.Sprintf("unsupported memory read addressing mode type %T", param))
	}
}

// ReadMemoryAbsolute reads a byte from an address using absolute addressing.
func (m *Memory) ReadMemoryAbsolute(address interface{}, register interface{}) byte {
	if register == nil {
		return m.readMemoryAbsoluteOffset(address, 0)
	}

	var offset uint16
	switch val := register.(type) {
	case *uint8: // X/Y register referenced in normal code
		offset = uint16(*val)
	case uint8: // X/Y register referenced in unit test as system.X
		offset = uint16(val)
	default:
		panic(fmt.Sprintf("unsupported extra parameter type %T for absolute memory read", register))
	}
	return m.readMemoryAbsoluteOffset(address, offset)
}

func (m *Memory) readMemoryAbsoluteOffset(address interface{}, offset uint16) byte {
	switch addr := address.(type) {
	case *uint8:
		if offset != 0 {
			panic("memory pointer read with offset is not supported")
		}
		return *addr
	case uint16:
		return m.ReadMemory(addr + offset)
	case int:
		return m.ReadMemory(uint16(addr) + offset)
	case addressing.Absolute:
		return m.ReadMemory(uint16(addr) + offset)
	default:
		panic(fmt.Sprintf("unsupported address type %T for absolute memory write", address))
	}
}

func (m *Memory) readMemoryIndirect(address addressing.Indirect, register interface{}) byte {
	pointer := m.indirectMemoryPointer(address, register)
	return m.ReadMemory(pointer)
}

func (m *Memory) readPointer(address uint16) uint16 {
	b1 := m.ReadMemory(address)
	b2 := m.ReadMemory(address + 1)
	ptr := uint16(b1)<<8 + uint16(b2)
	return ptr
}

// ReadMemory reads a byte from a memory address.
func (m *Memory) ReadMemory(address uint16) byte {
	switch {
	case address < 0x2000:
		return m.ram.ReadMemory(address & 0x07FF)
	case address < 0x4000:
		return m.ppu.ReadRegister(address)
	case address == controller.JOYPAD1:
		return m.controller1.Read()
	case address == controller.JOYPAD2:
		return m.controller2.Read()
	case address == apu.APU_CHAN_CTRL:
		return 0 // TODO
	case address >= 0x8000: // TODO implement mapper
		return m.cartridge.PRG[address-0x8000]
	default:
		panic(fmt.Sprintf("unhandled memory read at address: 0x%04X", address))
	}
}

func (m *Memory) indirectMemoryPointer(address addressing.Indirect, register interface{}) uint16 {
	if register == nil {
		panic("register parameter missing for indirect memory addressing")
	}

	p, ok := register.(*uint8)
	if !ok {
		panic(fmt.Sprintf("unsupported extra parameter type %T for indirect memory addressing", register))
	}

	var pointer uint16
	switch {
	case p == m.globalX, p == m.x:
		pointer = m.readPointer(uint16(address) + uint16(*p))
	case p == m.globalY, p == m.y:
		pointer = m.readPointer(uint16(address))
		pointer += uint16(*p)
	default:
		panic("only X and Y registers are supported for indirect addressing")
	}
	return pointer
}
