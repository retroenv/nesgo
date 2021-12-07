//go:build !nesgo
// +build !nesgo

package nes

import (
	"sync"
)

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
