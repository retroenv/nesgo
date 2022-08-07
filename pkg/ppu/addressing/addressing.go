//go:build !nesgo

// Package addressing handles PPU addressing of X/Y coordinates and nametables.
package addressing

// Addressing handles PPU addressing of X/Y coordinates and nametables.
type Addressing struct {
	latch bool // address latch toggle for high/low byte
	vram  register
	temp  register
}

// New returns a new addressing manager.
func New() *Addressing {
	return &Addressing{}
}

// SetAddress sets the address using the temp address register and the latch as switch
// to differentiate between the high and low bytes.
func (a *Addressing) SetAddress(value byte) {
	if a.latch {
		address := a.temp.address() & 0xFF00
		address |= uint16(value)
		a.temp.set(address)
		a.vram = a.temp
	} else {
		address := a.temp.address() & 0x00FF
		address |= uint16(value) << 8
		a.temp.set(address)
	}

	a.latch = !a.latch
}

// SetScroll sets the scroll values, X or Y depending on the latch toggle.
func (a *Addressing) SetScroll(value byte) {
	if a.latch {
		a.temp.FineY = uint16(value) & 0x07
		a.temp.CoarseY = uint16(value) >> 3
	} else {
		a.temp.CoarseX = uint16(value) >> 3
	}

	a.latch = !a.latch
}

// SetTempNameTables sets the temp register nametable from the passed PPU control byte.
func (a *Addressing) SetTempNameTables(nameTableX, nameTableY byte) {
	a.temp.NameTableX = uint16(nameTableX)
	a.temp.NameTableY = uint16(nameTableY)
}

// ClearLatch clears the address latch toggle.
func (a *Addressing) ClearLatch() {
	a.latch = false
}

// Latch returns the address latch toggle.
func (a *Addressing) Latch() bool {
	return a.latch
}

// Address returns the current vram address.
func (a *Addressing) Address() uint16 {
	return a.vram.address()
}

// FineY returns FineY of the vram.
func (a *Addressing) FineY() uint16 {
	return a.vram.FineY
}

// Increment the vram address by the given pixel count.
func (a *Addressing) Increment(value byte) {
	a.vram.increment(value)
}

// IncrementX increments coarse x and wraps the nameTables horizontally.
func (a *Addressing) IncrementX() {
	a.vram.incrementX()
}

// IncrementY increments fine Y, overflowing to coarse Y and wraps the nameTables vertically.
func (a *Addressing) IncrementY() {
	a.vram.incrementY()
}

// CopyX copies the temp X coordinates from temp to vram register.
func (a *Addressing) CopyX() {
	a.vram.NameTableX = a.temp.NameTableX
	a.vram.CoarseX = a.temp.CoarseX
}

// CopyY copies the temp Y coordinates from temp to vram register.
func (a *Addressing) CopyY() {
	a.vram.NameTableY = a.temp.NameTableY
	a.vram.CoarseY = a.temp.CoarseY
	a.vram.FineY = a.temp.FineY
}
