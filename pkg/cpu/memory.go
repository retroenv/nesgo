package cpu

type memory interface {
	ReadMemory(address uint16) byte
	ReadMemory16(address uint16) uint16
	ReadMemory16Bug(address uint16) uint16
	ReadMemoryAbsolute(address interface{}, register interface{}) byte
	ReadMemoryAddressModes(immediate bool, params ...interface{}) byte
	WriteMemory(address uint16, value byte)
	WriteMemory16(address, value uint16)
	WriteMemoryAddressModes(value byte, params ...interface{})
}
