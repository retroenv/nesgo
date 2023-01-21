package cpu

import "github.com/retroenv/retrogolib/nes/cpu"

// LinkInstructionFuncs links cpu instruction emulation functions to the CPU instance.
// Defining it directly in the instruction instances like ParamFunc: (CPU*).Adc does not work
// due to an initialization loop of Opcodes refers to adc refers to CPU.Adc refers to
// instructionHook refers to Opcodes.
// nolint: funlen
func LinkInstructionFuncs(c *CPU) {
	cpu.Adc.ParamFunc = c.Adc
	cpu.And.ParamFunc = c.And
	cpu.Asl.ParamFunc = c.Asl
	cpu.Bcc.ParamFunc = c.BccInternal
	cpu.Bcs.ParamFunc = c.BcsInternal
	cpu.Beq.ParamFunc = c.BeqInternal
	cpu.Bit.ParamFunc = c.Bit
	cpu.Bmi.ParamFunc = c.BmiInternal
	cpu.Bne.ParamFunc = c.BneInternal
	cpu.Bpl.ParamFunc = c.BplInternal
	cpu.Brk.NoParamFunc = c.Brk
	cpu.Bvc.ParamFunc = c.BvcInternal
	cpu.Bvs.ParamFunc = c.BvsInternal
	cpu.Clc.NoParamFunc = c.Clc
	cpu.Cld.NoParamFunc = c.Cld
	cpu.Cli.NoParamFunc = c.Cli
	cpu.Clv.NoParamFunc = c.Clv
	cpu.Cmp.ParamFunc = c.Cmp
	cpu.Cpx.ParamFunc = c.Cpx
	cpu.Cpy.ParamFunc = c.Cpy
	cpu.Dec.ParamFunc = c.Dec
	cpu.Dex.NoParamFunc = c.Dex
	cpu.Dey.NoParamFunc = c.Dey
	cpu.Eor.ParamFunc = c.Eor
	cpu.Inc.ParamFunc = c.Inc
	cpu.Inx.NoParamFunc = c.Inx
	cpu.Iny.NoParamFunc = c.Iny
	cpu.Jmp.ParamFunc = c.Jmp
	cpu.Jsr.ParamFunc = c.Jsr
	cpu.Lda.ParamFunc = c.Lda
	cpu.Ldx.ParamFunc = c.Ldx
	cpu.Ldy.ParamFunc = c.Ldy
	cpu.Lsr.ParamFunc = c.Lsr
	cpu.Nop.NoParamFunc = c.Nop
	cpu.Ora.ParamFunc = c.Ora
	cpu.Pha.NoParamFunc = c.Pha
	cpu.Php.NoParamFunc = c.Php
	cpu.Pla.NoParamFunc = c.Pla
	cpu.Plp.NoParamFunc = c.Plp
	cpu.Rol.ParamFunc = c.Rol
	cpu.Ror.ParamFunc = c.Ror
	cpu.Rti.NoParamFunc = c.Rti
	cpu.Rts.NoParamFunc = c.Rts
	cpu.Sbc.ParamFunc = c.Sbc
	cpu.Sec.NoParamFunc = c.Sec
	cpu.Sed.NoParamFunc = c.Sed
	cpu.Sei.NoParamFunc = c.Sei
	cpu.Sta.ParamFunc = c.Sta
	cpu.Stx.ParamFunc = c.Stx
	cpu.Sty.ParamFunc = c.Sty
	cpu.Tax.NoParamFunc = c.Tax
	cpu.Tay.NoParamFunc = c.Tay
	cpu.Tsx.NoParamFunc = c.Tsx
	cpu.Txa.NoParamFunc = c.Txa
	cpu.Txs.NoParamFunc = c.Txs
	cpu.Tya.NoParamFunc = c.Tya

	linkUnofficialInstructionFuncs(c)
}
