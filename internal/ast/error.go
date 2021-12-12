package ast

import "errors"

// errors.
var (
	ErrIfBranchingEmpty             = errors.New("if statement with branch can not have an empty block, goto or break expected")
	ErrBreakNotAfterBranching       = errors.New("break statement has to be after a branching instruction")
	ErrContinueNotAfterBranching    = errors.New("continue statement has to be after a branching instruction")
	ErrForOnlySimpleConditions      = errors.New("only simple conditions are supported in for loops")
	ErrForOnlySimplePostExpressions = errors.New("only simple statements are supported in for loop post expressions")
	ErrInvalidInitializer           = errors.New("variable initialization has to use constructors like NewUint8()")
	ErrInvalidVariableName          = errors.New("variable can not use reserved name")
)
