//go:build !nesgo
// +build !nesgo

package nes

import (
	"fmt"
	"math"
)

// Memory implements the memory controller.
type Memory struct {
	ram         *RAM
	ppu         *PPU
	controller1 *controller
	controller2 *controller
}

// newMemory returns a new memory instance, embedded it has
// new instances for PPU and the Controllers.
func newMemory() *Memory {
	r := &Memory{
		ram:         newRAM(0),
		ppu:         newPPU(),
		controller1: newController(),
		controller2: newController(),
	}
	return r
}

// writeMemoryAddressModes writes to memory using different address modes:
// Absolute: the absolut memory address is used to write the value
// Absolute, X: the absolut memory address with offset from X is used
// Absolute, Y: the absolut memory address with offset from Y is used
// (Indirect, X): TODO: add
// (Indirect), Y: TODO: add
func (m *Memory) writeMemoryAddressModes(value byte, params ...interface{}) {
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
	case Absolute:
		m.writeMemoryAbsolute(address, value, register)
	case Indirect:
		m.writeMemoryIndirect(address, value, register)
	default:
		panic(fmt.Sprintf("unsupported memory write addressing mode type %T", param))
	}
}

func (m *Memory) writeMemoryIndirect(address interface{}, value byte, register interface{}) {
	if register == nil {
		panic("register parameter missing for indirect memory addressing")
	}

	p, ok := register.(*uint8)
	if !ok {
		panic(fmt.Sprintf("unsupported extra parameter type %T for indirect memory write", register))
	}
	addr, ok := address.(uint16)
	if !ok {
		panic(fmt.Sprintf("unsupported address parameter type %T for indirect memory write", address))
	}
	switch {
	case p == X:
		m.writeMemory(addr+uint16(*p), value)
	case p == Y:
		ptr := m.readPointer(addr)
		ptr += uint16(*p)
		m.writeMemory(ptr, value)
	default:
		panic("only X and Y registers are supported for indirect addressing")
	}
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
		m.writeMemory(uint16(addr)+offset, value)
	case uint8:
		m.writeMemory(uint16(addr)+offset, value)
	case *uint8:
		*addr = value
	case uint16:
		m.writeMemory(addr+offset, value)
	case *uint16:
		*addr = uint16(value)
	case int:
		m.writeMemory(uint16(addr)+offset, value)
	case Absolute:
		m.writeMemory(uint16(addr)+offset, value)
	default:
		panic(fmt.Sprintf("unsupported address type %T for absolute memory write with register", address))
	}
}

func (m *Memory) writeMemory(address uint16, value byte) {
	switch {
	case address < 0x2000:
		m.ram.writeMemory(address&0x07FF, value)
	case address < 0x4000:
		m.ppu.writeRegister(address, value)
	case address == JOYPAD1:
		m.controller1.setStrobeMode(value)
		m.controller2.setStrobeMode(value)
	case address == DMC_FREQ:
		return // TODO
	case address == APU_CHAN_CTRL:
		return // TODO
	case address == APU_FRAME:
		return // TODO
	default:
		panic(fmt.Sprintf("unhandled memory write at address: 0x%04X", address))
	}
}

// readMemoryAddressModes reads memory using different address modes:
// Immediate: if immediate is true and the passed first param fits into
//            a byte, it's immediate value is returned
// Absolute: the absolut memory address is used to read the value
// Absolute, X: the absolut memory address with offset from X is used
// Absolute, Y: the absolut memory address with offset from Y is used
// (Indirect, X): TODO: add
// (Indirect), Y: TODO: add
func (m *Memory) readMemoryAddressModes(immediate bool, params ...interface{}) byte {
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
		return m.readMemoryAbsolute(address, register)
	case uint8:
		return address // immediate, not an address
	case *uint8: // variable
		return *address
	case Absolute:
		return m.readMemoryAbsolute(address, register)
	default:
		panic(fmt.Sprintf("unsupported memory read addressing mode type %T", param))
	}
}

func (m *Memory) readMemoryAbsolute(address interface{}, register interface{}) byte {
	if register == nil {
		return m.readMemoryAbsoluteOffset(address, 0)
	}

	var offset uint16
	switch val := register.(type) {
	case uint8:
		offset = uint16(val)
	case *uint8:
		offset = uint16(*val)
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
		return m.readMemory(addr + offset)
	case int:
		return m.readMemory(uint16(addr) + offset)
	case Absolute:
		return m.readMemory(uint16(addr) + offset)
	default:
		panic(fmt.Sprintf("unsupported address type %T for absolute memory write", address))
	}
}

func (m *Memory) readPointer(address uint16) uint16 {
	b1 := m.readMemory(address)
	b2 := m.readMemory(address + 1)
	ptr := uint16(b1)<<8 + uint16(b2)
	return ptr
}

func (m *Memory) readMemory(address uint16) byte {
	switch {
	case address < 0x2000:
		return m.ram.readMemory(address & 0x07FF)
	case address < 0x4000:
		return m.ppu.readRegister(address)
	case address == JOYPAD1:
		return m.controller1.read()
	case address == JOYPAD2:
		return m.controller2.read()
	case address == APU_CHAN_CTRL:
		return 0 // TODO
	default:
		panic(fmt.Sprintf("unhandled memory read at address: 0x%04X", address))
	}
}
