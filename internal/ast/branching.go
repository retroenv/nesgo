package ast

import (
	"fmt"
	"strings"

	"github.com/retroenv/retrogolib/arch/cpu/m6502"
)

const (
	breakStatement             = "break"
	continueStatement          = "continue"
	GotoInstruction            = "goto"
	JmpInstruction             = "jmp"
	NotOperator                = "!"
	ReturnInstruction          = "rts"
	ReturnInterruptInstruction = "rti"
)

// Branching is a branching declaration.
type Branching struct {
	Instruction     string
	DestinationName string
	Destination     *Label
	Not             bool
}

// NewBranching returns a goto instruction.
func NewBranching(instruction string, destination string) (Node, error) {
	return &Branching{
		Instruction:     instruction,
		DestinationName: destination,
	}, nil
}

// String implement the fmt.Stringer interface.
func (b Branching) String() string {
	if b.DestinationName == "" {
		return fmt.Sprintf("inst, %s", b.Instruction)
	}
	return fmt.Sprintf("inst, %s, %s", b.Instruction, b.DestinationName)
}

// Label is a label declaration.
type Label struct {
	Name string
}

// NewLabel returns a label definition.
func NewLabel(id *Identifier, instruction any) (Node, error) {
	l := &Label{
		Name: id.Name,
	}
	if instruction == nil {
		return l, nil
	}
	return NewNodeList(l, instruction)
}

// String implement the fmt.Stringer interface.
func (l Label) String() string {
	return fmt.Sprintf("label, %s", l.Name)
}

// NewCall handles a function call that could represent an alias
// for a CPU instruction.
func NewCall(expr *Identifier, arg any) (Node, error) {
	name := strings.ToLower(expr.Name)
	if strings.HasPrefix(name, "fmt.") {
		return nil, nil // nolint: nilnil
	}

	if _, ok := m6502.BranchingInstructions[name]; ok {
		var destination string
		if ins, ok := arg.(*Instruction); ok {
			destination = ins.Name
		}
		return NewBranching(name, destination)
	}

	if _, isInst := m6502.Instructions[name]; isInst {
		i, err := newInstruction(name, arg)
		if err != nil {
			return nil, err
		}
		return i, nil
	}

	return newCall(expr.Name, arg)
}

// Call is a call declaration.
type Call struct {
	Function  string
	Parameter []any
}

// String implement the fmt.Stringer interface.
func (c Call) String() string {
	b := &strings.Builder{}
	_, _ = fmt.Fprintf(b, "call, %s", c.Function)

	for _, p := range c.Parameter {
		_, _ = fmt.Fprintf(b, ", %s", p)
	}

	s := b.String()
	return s
}

// newCall returns a call instruction.
func newCall(name string, arg any) (Node, error) {
	c := &Call{
		Function: name,
	}
	if arg == nil {
		return c, nil
	}

	switch a := arg.(type) {
	case *Identifier:
		c.Parameter = append(c.Parameter, a)
	case *ExpressionList:
		c.Parameter = append(c.Parameter, a)
	case *Value:
		c.Parameter = append(c.Parameter, a)
	case *NodeList:
		for _, node := range a.Nodes {
			switch n := node.(type) {
			case *Value:
				c.Parameter = append(c.Parameter, n)
			case *Identifier:
				c.Parameter = append(c.Parameter, n)
			case *Call:
				c.Parameter = append(c.Parameter, n)
			default:
				return nil, fmt.Errorf("type %T is not supported as call parameter in node list", node)
			}
		}
	default:
		return nil, fmt.Errorf("type %T is not supported as call parameter", arg)
	}

	return c, nil
}
