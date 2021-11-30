package ast

import (
	"fmt"
	"strings"
)

// Statement is a statement declaration.
type Statement struct {
	Op        string
	Arguments []string
}

// String implement the fmt.Stringer interface.
func (s Statement) String() string {
	return fmt.Sprintf("op, %s, %s", s.Op, strings.Join(s.Arguments, ","))
}

// NewAssignStatement returns an assignment statement.
func NewAssignStatement(id *Identifier, val interface{}) (Node, error) {
	s := &Statement{
		Op: "=",
	}

	switch n := val.(type) {
	case *Identifier:
		s.Arguments = []string{id.Name, n.Name}
	case *Value:
		s.Arguments = []string{id.Name, n.Value}
	default:
		return nil, fmt.Errorf("type %T is not supported for assign statements", val)
	}

	return s, nil
}

// NewReturnStatement returns a return statement.
func NewReturnStatement() (Node, error) {
	return newInstruction("rts"), nil
}
