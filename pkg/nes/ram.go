//go:build !nesgo
// +build !nesgo

package nes

import (
	"sync"
)

// RAM implements the Random-access memory.
type RAM struct {
	mu     sync.RWMutex // protects data
	offset uint16
	data   []byte
}

func newRAM(offset uint16) *RAM {
	r := &RAM{
		offset: offset,
	}
	r.reset()
	return r
}

func (r *RAM) reset() {
	r.mu.Lock()
	r.data = make([]byte, 0x2000)
	r.mu.Unlock()
}

func (r *RAM) readMemory(address uint16) byte {
	r.mu.RLock()
	b := r.data[address-r.offset]
	r.mu.RUnlock()
	return b
}

func (r *RAM) writeMemory(address uint16, value byte) {
	r.mu.Lock()
	r.data[address-r.offset] = value
	r.mu.Unlock()
}
