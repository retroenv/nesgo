package cpu

import "github.com/retroenv/retrogolib/arch/cpu/m6502"

// LinkInstructionFuncs links cpu instruction emulation functions to the CPU instance.
// Defining it directly in the instruction instances like ParamFunc: (CPU*).Adc does not work
// due to an initialization loop of Opcodes refers to adc refers to CPU.Adc refers to
// instructionHook refers to Opcodes.
// nolint: funlen
func LinkInstructionFuncs(c *CPU) {
	m6502.Adc.ParamFunc = c.Adc
	m6502.And.ParamFunc = c.And
	m6502.Asl.ParamFunc = c.Asl
	m6502.Bcc.ParamFunc = c.BccInternal
	m6502.Bcs.ParamFunc = c.BcsInternal
	m6502.Beq.ParamFunc = c.BeqInternal
	m6502.Bit.ParamFunc = c.Bit
	m6502.Bmi.ParamFunc = c.BmiInternal
	m6502.Bne.ParamFunc = c.BneInternal
	m6502.Bpl.ParamFunc = c.BplInternal
	m6502.Brk.NoParamFunc = c.Brk
	m6502.Bvc.ParamFunc = c.BvcInternal
	m6502.Bvs.ParamFunc = c.BvsInternal
	m6502.Clc.NoParamFunc = c.Clc
	m6502.Cld.NoParamFunc = c.Cld
	m6502.Cli.NoParamFunc = c.Cli
	m6502.Clv.NoParamFunc = c.Clv
	m6502.Cmp.ParamFunc = c.Cmp
	m6502.Cpx.ParamFunc = c.Cpx
	m6502.Cpy.ParamFunc = c.Cpy
	m6502.Dec.ParamFunc = c.Dec
	m6502.Dex.NoParamFunc = c.Dex
	m6502.Dey.NoParamFunc = c.Dey
	m6502.Eor.ParamFunc = c.Eor
	m6502.Inc.ParamFunc = c.Inc
	m6502.Inx.NoParamFunc = c.Inx
	m6502.Iny.NoParamFunc = c.Iny
	m6502.Jmp.ParamFunc = c.Jmp
	m6502.Jsr.ParamFunc = c.Jsr
	m6502.Lda.ParamFunc = c.Lda
	m6502.Ldx.ParamFunc = c.Ldx
	m6502.Ldy.ParamFunc = c.Ldy
	m6502.Lsr.ParamFunc = c.Lsr
	m6502.Nop.NoParamFunc = c.Nop
	m6502.Ora.ParamFunc = c.Ora
	m6502.Pha.NoParamFunc = c.Pha
	m6502.Php.NoParamFunc = c.Php
	m6502.Pla.NoParamFunc = c.Pla
	m6502.Plp.NoParamFunc = c.Plp
	m6502.Rol.ParamFunc = c.Rol
	m6502.Ror.ParamFunc = c.Ror
	m6502.Rti.NoParamFunc = c.Rti
	m6502.Rts.NoParamFunc = c.Rts
	m6502.Sbc.ParamFunc = c.Sbc
	m6502.Sec.NoParamFunc = c.Sec
	m6502.Sed.NoParamFunc = c.Sed
	m6502.Sei.NoParamFunc = c.Sei
	m6502.Sta.ParamFunc = c.Sta
	m6502.Stx.ParamFunc = c.Stx
	m6502.Sty.ParamFunc = c.Sty
	m6502.Tax.NoParamFunc = c.Tax
	m6502.Tay.NoParamFunc = c.Tay
	m6502.Tsx.NoParamFunc = c.Tsx
	m6502.Txa.NoParamFunc = c.Txa
	m6502.Txs.NoParamFunc = c.Txs
	m6502.Tya.NoParamFunc = c.Tya

	linkUnofficialInstructionFuncs(c)
}
