//go:build !nesgo

package controller

import . "github.com/retroenv/retrogolib/addressing"

// AddressToName maps address constants from address to name.
var AddressToName = map[uint16]AccessModeConstant{
	JOYPAD1: {Constant: "JOYPAD1", Mode: ReadWriteAccess},
	JOYPAD2: {Constant: "JOYPAD2", Mode: ReadAccess},
}
