package bus

// BasicMemory represents a basic memory access interface.
type BasicMemory interface {
	Read(address uint16) uint8
	Write(address uint16, value uint8)
}

// Memory represents an advanced memory access interface.
type Memory interface {
	BasicMemory

	ReadAbsolute(address any, register any) byte
	ReadAddressModes(immediate bool, params ...any) byte
	ReadWord(address uint16) uint16
	ReadWordBug(address uint16) uint16
	WriteAddressModes(value byte, params ...any)
	WriteWord(address, value uint16)

	LinkRegisters(x *uint8, y *uint8, globalX *uint8, globalY *uint8)
}
