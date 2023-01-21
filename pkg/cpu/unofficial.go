//go:build !nesgo

// This file contains support for unofficial CPU instructions.
// https://www.nesdev.org/wiki/Programming_with_unofficial_opcodes

package cpu

import "github.com/retroenv/retrogolib/nes/cpu"

func linkUnofficialInstructionFuncs(c *CPU) {
	cpu.Dcp.ParamFunc = c.Dcp
	cpu.Isc.ParamFunc = c.Isc
	cpu.Lax.ParamFunc = c.Lax
	cpu.NopUnofficial.ParamFunc = c.NopUnofficial
	cpu.Rla.ParamFunc = c.Rla
	cpu.Rra.ParamFunc = c.Rra
	cpu.Sax.ParamFunc = c.Sax
	cpu.SbcUnofficial.ParamFunc = c.SbcUnofficial
	cpu.Slo.ParamFunc = c.Slo
	cpu.Sre.ParamFunc = c.Sre
}
