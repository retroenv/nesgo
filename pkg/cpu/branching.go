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
