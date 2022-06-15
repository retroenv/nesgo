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
