//go:build !nesgo
// +build !nesgo

package ppu

// register represents an internal PPU register that has the address decoded to fields.
type register struct {
	CoarseX    uint16
	CoarseY    uint16
	NameTableX uint16
	NameTableY uint16
	FineY      uint16
	Unused     uint16
}

// address returns the final address calculated from the internal fields.
func (r register) address() uint16 {
	address := r.CoarseX
	address |= r.CoarseY << 5
	address |= r.NameTableX << 10
	address |= r.NameTableY << 11
	address |= r.FineY << 12
	address |= r.Unused << 15
	return address
}

// set the internal decoded fields from an address.
func (r *register) set(address uint16) {
	r.CoarseX = address & 0b00011111
	r.CoarseY = (address >> 5) & 0b00011111
	r.NameTableX = (address >> 10) & 1
	r.NameTableY = (address >> 11) & 1
	r.FineY = (address >> 12) & 0b00000111
	r.Unused = (address >> 15) & 1
}

// increment the address by the given pixel count.
func (r *register) increment(value byte) {
	r.set(r.address() + uint16(value))
}
