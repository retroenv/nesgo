//go:build !nesgo
// +build !nesgo

package memory

import (
	"sync"
)

// RAM implements the Random-access memory.
type RAM struct {
	mu     sync.RWMutex // protects data
	offset uint16
	data   []byte
}

// NewRAM returns a new ram.
func NewRAM(offset uint16) *RAM {
	r := &RAM{
		offset: offset,
	}
	r.Reset()
	return r
}

// Reset resets the RAM content.
func (r *RAM) Reset() {
	r.mu.Lock()
	r.data = make([]byte, 0x2000)
	r.mu.Unlock()
}

// Read a byte from a memory address.
func (r *RAM) Read(address uint16) byte {
	r.mu.RLock()
	b := r.data[address-r.offset]
	r.mu.RUnlock()
	return b
}

// Write a byte to a memory address.
func (r *RAM) Write(address uint16, value byte) {
	r.mu.Lock()
	r.data[address-r.offset] = value
	r.mu.Unlock()
}
