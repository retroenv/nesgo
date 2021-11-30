// Package neslib provides helper functions for writing NES programs in Golang.
// This package needs to be imported using the dot notation.
package neslib

type inline interface{}

// NesGoVariableInit is a placeholder for a variable initialization function
// that nesgo uses to initialize variables on program startup. If the function
// is not called from the code it will be called automatically from the first
// instruction of the reset handler code. The location of this call can be
// customized by placing a call to this function anywhere into the program
// code.
func NesGoVariableInit() {}
