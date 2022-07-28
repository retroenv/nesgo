package bus

// NameTable represents a name table interface.
type NameTable interface {
	BasicMemory

	Fetch(address uint16)
	SetVRAM(vram []byte)
	Value() byte
}
