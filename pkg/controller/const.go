package controller

const (
	JOYPAD1 = 0x4016
	JOYPAD2 = 0x4017
)

// AddressToName maps address constants from address to name.
var AddressToName = map[uint16]string{
	JOYPAD1: "JOYPAD1",
	JOYPAD2: "JOYPAD2",
}
