//go:build !nesgo
// +build !nesgo

package nes

import (
	"fmt"
	"sync"
)

var ram *RAM

// RAM implements the Random-access memory.
type RAM struct {
	data []byte
	mu   sync.RWMutex
}

func newRAM() *RAM {
	r := &RAM{}
	r.reset()
	return r
}

func (r *RAM) reset() {
	r.mu.Lock()
	r.data = make([]byte, 0x4000)
	r.mu.Unlock()
}

func (r *RAM) readMemory(address uint16) byte {
	r.mu.RLock()
	b := r.data[address]
	r.mu.RUnlock()
	return b
}

func (r *RAM) writeMemory(address uint16, value byte) {
	r.mu.Lock()
	r.data[address] = value
	r.mu.Unlock()
}

func writeMemoryAbsolute(address interface{}, value byte, reg ...interface{}) {
	switch len(reg) {
	case 0:
		writeMemoryAbsoluteNoRegister(address, value)

	case 1:
		p, ok := reg[0].(*uint8)
		if !ok {
			panic(fmt.Sprintf("unsupported extra parameter type %T for absolute memory write", reg[0]))
		}
		writeMemoryAbsoluteRegister(address, value, p)

	default:
		panic(fmt.Sprintf("unsupported extra parameter count %d for absolute memory write", len(reg)))
	}
}

func writeMemoryAbsoluteNoRegister(address interface{}, value byte) {
	switch addr := address.(type) {
	case uint8:
		writeMemory(uint16(addr), value)
	case *uint8:
		*addr = value
	case uint16:
		writeMemory(addr, value)
	case *uint16:
		*addr = uint16(value)
	case int:
		writeMemory(uint16(addr), value)
	default:
		panic(fmt.Sprintf("unsupported address type %T for absolute memory write", address))
	}
}

func writeMemoryAbsoluteRegister(address interface{}, value byte, register *uint8) {
	switch addr := address.(type) {
	case int8:
		writeMemory(uint16(addr)+uint16(*register), value)
	case uint8:
		writeMemory(uint16(addr)+uint16(*register), value)
	case uint16:
		writeMemory(addr+uint16(*register), value)
	case int:
		writeMemory(uint16(addr)+uint16(*register), value)
	default:
		panic(fmt.Sprintf("unsupported address type %T for absolute memory write with register", address))
	}
}

func writeMemoryIndirect(address uint16, value byte, reg ...interface{}) {
	switch len(reg) {
	case 0:
		panic("register parameter missing for indirect memory addressing")

	case 1:
		p, ok := reg[0].(*uint8)
		if !ok {
			panic(fmt.Sprintf("unsupported extra parameter type %T for indirect memory write", reg[0]))
		}
		switch {
		case p == X:
			writeMemory(address+uint16(*p), value)
		case p == Y:
			ptr := readPointer(address)
			ptr += uint16(*p)
			writeMemory(ptr, value)
		default:
			panic("only X and Y registers are supported for indirect addressing")
		}

	default:
		panic(fmt.Sprintf("unsupported extra parameter count %d for indirect memory write", len(reg)))
	}
}

func readMemoryAbsolute(address interface{}, reg ...interface{}) byte {
	switch len(reg) {
	case 0:
		return readMemoryAbsoluteNoRegister(address)

	case 1:
		panic(fmt.Sprintf("unsupported extra parameter type %T for absolute memory read", reg[0]))

	default:
		panic(fmt.Sprintf("unsupported extra parameter count %d for absolute memory write", len(reg)))
	}
}

func readMemoryAbsoluteNoRegister(address interface{}) byte {
	switch addr := address.(type) {
	case *uint8:
		return *addr
	case uint16:
		return readMemory(addr)
	case int:
		return readMemory(uint16(addr))
	default:
		panic(fmt.Sprintf("unsupported address type %T for absolute memory write", address))
	}
}

func readPointer(address uint16) uint16 {
	b1 := readMemory(address)
	b2 := readMemory(address + 1)
	ptr := uint16(b1)<<8 + uint16(b2)
	return ptr
}

func readMemory(address uint16) byte {
	switch {
	case address < 0x2000:
		b := ram.readMemory(address & 0x07FF)
		return b
	case address < 0x4000:
		return ppu.readRegister(address)
	case address == APU_CHAN_CTRL:
		return 0 // TODO
	default:
		panic(fmt.Sprintf("unhandled memory read at address: 0x%04X", address))
	}
}

func writeMemory(address uint16, value byte) {
	switch {
	case address < 0x2000:
		ram.writeMemory(address&0x07FF, value)
	case address < 0x4000:
		ppu.writeRegister(address, value)
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
