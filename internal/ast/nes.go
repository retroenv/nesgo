package ast

// CPUBranchingInstructions ...
var CPUBranchingInstructions = map[string]struct{}{
	"bcc": {},
	"bcs": {},
	"beq": {},
	"bmi": {},
	"bne": {},
	"bpl": {},
	"bvc": {},
	"bvs": {},
	"jmp": {},
}

// CPURegisters ...
var CPURegisters = map[string]struct{}{
	"A": {},
	"X": {},
	"Y": {},
}
