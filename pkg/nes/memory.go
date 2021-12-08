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

func newMemory() *Memory {
	r := &Memory{
		ram:         newRAM(0),
		ppu:         newPPU(),
		controller1: newController(),
		controller2: newController(),
	}
	return r
}

func (m *Memory) writeMemoryAbsolute(address interface{}, value byte, reg ...interface{}) {
	switch len(reg) {
	case 0:
		m.writeMemoryAbsoluteNoRegister(address, value)

	case 1:
		switch p := reg[0].(type) {
		case *uint8: // X/Y register referenced in normal code
			m.writeMemoryAbsoluteRegister(address, value, p)
		case uint8: // X/Y register referenced in unit test as system.X
			m.writeMemoryAbsoluteRegister(address, value, &p)
		default:
			panic(fmt.Sprintf("unsupported extra parameter type %T for absolute memory write", reg[0]))
		}

	default:
		panic(fmt.Sprintf("unsupported extra parameter count %d for absolute memory write", len(reg)))
	}
}

func (m *Memory) writeMemoryAbsoluteNoRegister(address interface{}, value byte) {
	switch addr := address.(type) {
	case uint8:
		m.writeMemory(uint16(addr), value)
	case *uint8:
		*addr = value
	case uint16:
		m.writeMemory(addr, value)
	case *uint16:
		*addr = uint16(value)
	case int:
		m.writeMemory(uint16(addr), value)
	case Absolute:
		m.writeMemory(uint16(addr), value)
	default:
		panic(fmt.Sprintf("unsupported address type %T for absolute memory write", address))
	}
}

func (m *Memory) writeMemoryAbsoluteRegister(address interface{}, value byte, register *uint8) {
	switch addr := address.(type) {
	case int8:
		m.writeMemory(uint16(addr)+uint16(*register), value)
	case uint8:
		m.writeMemory(uint16(addr)+uint16(*register), value)
	case uint16:
		m.writeMemory(addr+uint16(*register), value)
	case int:
		m.writeMemory(uint16(addr)+uint16(*register), value)
	case Absolute:
		m.writeMemory(uint16(addr)+uint16(*register), value)
	default:
		panic(fmt.Sprintf("unsupported address type %T for absolute memory write with register", address))
	}
}

func (m *Memory) writeMemoryIndirect(address interface{}, value byte, reg ...interface{}) {
	switch len(reg) {
	case 0:
		panic("register parameter missing for indirect memory addressing")

	case 1:
		p, ok := reg[0].(*uint8)
		if !ok {
			panic(fmt.Sprintf("unsupported extra parameter type %T for indirect memory write", reg[0]))
		}
		addr, ok := reg[0].(uint16)
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

	default:
		panic(fmt.Sprintf("unsupported extra parameter count %d for indirect memory write", len(reg)))
	}
}

func (m *Memory) readMemoryAbsolute(address interface{}, reg ...interface{}) byte {
	switch len(reg) {
	case 0:
		return m.readMemoryAbsoluteOffset(address, 0)

	case 1:
		var offset uint8
		switch val := reg[0].(type) {
		case uint8:
			offset = val
		case *uint8:
			offset = *val
		default:
			panic(fmt.Sprintf("unsupported extra parameter type %T for absolute memory read", reg[0]))
		}
		return m.readMemoryAbsoluteOffset(address, uint16(offset))

	default:
		panic(fmt.Sprintf("unsupported extra parameter count %d for absolute memory write", len(reg)))
	}
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

func (m *Memory) readMemoryAddressModes(param interface{}, reg ...interface{}) byte {
	switch val := param.(type) {
	case int:
		if reg == nil && val <= math.MaxUint8 {
			return uint8(val) // immediate, not an address
		}
		return m.readMemoryAbsolute(val, reg...)
	case uint8:
		return val // immediate, not an address
	case *uint8: // variable
		return *val
	case Absolute:
		return m.readMemoryAbsolute(val, reg...)
	}
	panic(fmt.Sprintf("unsupported memory read addressing mode type %T", param))
}

func (m *Memory) writeMemoryAddressModes(param interface{}, value byte, reg ...interface{}) {
	switch val := param.(type) {
	case int:
		m.writeMemoryAbsolute(val, value, reg...)
	case Absolute:
		m.writeMemoryAbsolute(val, value, reg...)
	case Indirect:
		m.writeMemoryIndirect(val, value, reg...)
	case *uint8: // variable
		*val = value
	default:
		panic(fmt.Sprintf("unsupported memory write addressing mode type %T", param))
	}
}
