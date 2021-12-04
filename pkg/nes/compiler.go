package nes

// Inline can be used to declare the function to be inlined by the compiler.
// This should be used as last parameter in a function, as variadic parameter
// so that any caller does not need to pass any extra argument, for example:
// func Name(_ ...Inline)
type Inline interface{}

// VariableInit is a placeholder for a variable initialization function
// that nesgo uses to initialize variables on program startup. If the function
// is not called from the code it will be called automatically from the first
// instruction of the reset handler code. The location of this call can be
// customized by placing a call to this function anywhere into the program
// code.
func VariableInit() {}

// ZeroPage indicates that the parameter for the instruction is addressing
// the zero page.
type ZeroPage uint8

// Absolute indicates that the parameter for the instruction is an
// absolute address.
type Absolute uint16

// Indirect indicates that the parameter for the instruction is using
// indirect addressing using an address and X or Y register.
type Indirect uint8
