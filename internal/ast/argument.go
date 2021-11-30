package ast

import (
	"fmt"
	"strings"
)

// ArgumentValue is an instruction argument value.
type ArgumentValue struct {
	Value string
}

// String implement the fmt.Stringer interface.
func (a ArgumentValue) String() string {
	return a.Value
}

// ArgumentParam is an instruction argument parameter reference.
type ArgumentParam struct {
	Index int
}

// String implement the fmt.Stringer interface.
func (a ArgumentParam) String() string {
	return fmt.Sprintf("param[%d]", a.Index)
}

// Arguments defines a list of arguments.
type Arguments []Node

// String implement the fmt.Stringer interface.
func (a Arguments) String() string {
	args := make([]string, 0, len(a))
	for _, arg := range a {
		args = append(args, arg.String())
	}
	s := strings.Join(args, ",")
	return s
}
