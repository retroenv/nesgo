package cpu

type memory interface {
	ReadMemory(address uint16) byte
	ReadMemoryAbsolute(address interface{}, register interface{}) byte
	ReadMemoryAddressModes(immediate bool, params ...interface{}) byte
	WriteMemory(address uint16, value byte)
	WriteMemoryAddressModes(value byte, params ...interface{})
}
