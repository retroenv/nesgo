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

func NewRAM(offset uint16) *RAM {
	r := &RAM{
		offset: offset,
	}
	r.Reset()
	return r
}

func (r *RAM) Reset() {
	r.mu.Lock()
	r.data = make([]byte, 0x2000)
	r.mu.Unlock()
}

func (r *RAM) ReadMemory(address uint16) byte {
	r.mu.RLock()
	b := r.data[address-r.offset]
	r.mu.RUnlock()
	return b
}

func (r *RAM) WriteMemory(address uint16, value byte) {
	r.mu.Lock()
	r.data[address-r.offset] = value
	r.mu.Unlock()
}
