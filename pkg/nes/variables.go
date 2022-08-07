//go:build !nesgo

package nes

// NewInt8 creates a new int8 variable that can be used with
// instructions like Sta(). This function returns a pointer to
// allow the emulator Go mode to differentiate between a
// passed address or variable.
func NewInt8(value int8) *int8 {
	i := new(int8)
	*i = value
	return i
}

// NewUint8 creates a new uint8 variable that can be used with
// instructions like Sta(). This function returns a pointer to
// allow the emulator Go mode to differentiate between a
// passed address or variable.
func NewUint8(value uint8) *uint8 {
	i := new(uint8)
	*i = value
	return i
}

// NewUint16 creates a new uint16 variable that can be used with
// instructions like Sta(). This function returns a pointer to
// allow the emulator Go mode to differentiate between a
// passed address or variable.
func NewUint16(value uint16) *uint16 {
	i := new(uint16)
	*i = value
	return i
}
