//go:build !nesgo

// This file contains support for unofficial CPU instructions.
// https://www.nesdev.org/wiki/Programming_with_unofficial_opcodes

package cpu

import (
	"github.com/retroenv/retrogolib/arch/cpu/m6502"
)

func linkUnofficialInstructionFuncs(c *CPU) {
	m6502.Dcp.ParamFunc = c.Dcp
	m6502.Isc.ParamFunc = c.Isc
	m6502.Lax.ParamFunc = c.Lax
	m6502.NopUnofficial.ParamFunc = c.NopUnofficial
	m6502.Rla.ParamFunc = c.Rla
	m6502.Rra.ParamFunc = c.Rra
	m6502.Sax.ParamFunc = c.Sax
	m6502.SbcUnofficial.ParamFunc = c.SbcUnofficial
	m6502.Slo.ParamFunc = c.Slo
	m6502.Sre.ParamFunc = c.Sre
}
