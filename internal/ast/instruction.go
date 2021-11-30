package ast

import "fmt"

// Instruction is an instruction declaration.
type Instruction struct {
	Name      string
	Arguments Arguments
	Comment   string
}

// newInstruction creates an instruction specification.
func newInstruction(name string) Node {
	return &Instruction{
		Name: name,
	}
}

func (i *Instruction) addArgument(arg interface{}) {
	switch val := arg.(type) {
	case *Identifier:
		i.Arguments = append(i.Arguments, &ArgumentValue{Value: val.Name})

	case *Value:
		i.Arguments = append(i.Arguments, &ArgumentValue{Value: val.Value})

	case *NodeList:
		for _, node := range val.Nodes {
			i.addArgument(node)
		}

	case *ExpressionList:
		i.Arguments = append(i.Arguments, val)
	}
}

// String implement the fmt.Stringer interface.
func (i Instruction) String() string {
	if len(i.Arguments) == 0 {
		return fmt.Sprintf("inst, %s", i.Name)
	}
	return fmt.Sprintf("inst, %s, %s", i.Name, i.Arguments)
}
