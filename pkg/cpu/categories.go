package cpu

// BranchingInstructions contains all branching instructions.
var BranchingInstructions = map[string]struct{}{
	bcc.Name: {},
	bcs.Name: {},
	beq.Name: {},
	bmi.Name: {},
	bne.Name: {},
	bpl.Name: {},
	bvc.Name: {},
	bvs.Name: {},
	jmp.Name: {},
	jsr.Name: {},
}

// NotExecutingFollowingOpcodeInstructions contains all instructions that jump
// to a different address and do not return to execute the following opcode.
var NotExecutingFollowingOpcodeInstructions = map[string]struct{}{
	jmp.Name: {},
	rti.Name: {},
	rts.Name: {},
}

// MemoryReadInstructions contains all instructions that can read from an
// absolute memory address.
var MemoryReadInstructions = map[string]struct{}{
	and.Name: {},
	bit.Name: {},
	cmp.Name: {},
	cpx.Name: {},
	cpy.Name: {},
	lda.Name: {},
	ldx.Name: {},
	ldy.Name: {},
}

// MemoryWriteInstructions contains all instructions that can write to a
// absolute memory address.
var MemoryWriteInstructions = map[string]struct{}{
	sta.Name: {},
	stx.Name: {},
	sty.Name: {},
}
