//go:build !nesgo

package addressing

// register represents an internal PPU register that has the address decoded to fields.
// It is also known as loopy register. The data could be stored more efficient as a
// single uint16 but the decomposed structure is chosen for ease of development and readability of code.
type register struct {
	CoarseX    uint16 // 0000 0000 0001 1111
	CoarseY    uint16 // 0000 0011 1110 0000
	NameTableX uint16 // 0000 0100 0000 0000
	NameTableY uint16 // 0000 1000 0000 0000
	FineY      uint16 // 0111 0000 0000 0000
	Unused     uint16 // 1000 0000 0000 0000
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
	r.CoarseX = address & 0b0001_1111
	r.CoarseY = (address >> 5) & 0b0001_1111
	r.NameTableX = (address >> 10) & 1
	r.NameTableY = (address >> 11) & 1
	r.FineY = (address >> 12) & 0b0000_0111
	r.Unused = (address >> 15) & 1
}

// increment the address by the given pixel count.
func (r *register) increment(value byte) {
	r.set(r.address() + uint16(value))
}

// incrementX increments coarse x and wraps the nameTables horizontally.
func (r *register) incrementX() {
	if r.CoarseX < 31 {
		r.CoarseX++
		return
	}

	r.CoarseX = 0
	r.NameTableX ^= 1 // switch horizontal nameTable
}

// incrementY increments fine Y, overflowing to coarse Y and wraps the nameTables vertically.
func (r *register) incrementY() {
	if r.FineY < 7 {
		r.FineY++
		return
	}

	r.FineY = 0

	switch r.CoarseY {
	case 29:
		// row 29 is the last row of tiles in a nameTable
		r.CoarseY = 0
		r.NameTableY ^= 1 // switch vertical nameTable

	case 31:
		// coarse Y can be set out of bounds (> 29), which will cause the PPU to read the attribute data stored
		// there as tile data. if coarse Y is incremented from 31, it will wrap to 0, but the nameTable will not switch.
		r.CoarseY = 0

	default:
		r.CoarseY++
	}
}
