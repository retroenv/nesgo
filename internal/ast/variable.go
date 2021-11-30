package ast

import "fmt"

var reservedNames = map[string]struct{}{
	"x": {},
	"X": {},
	"y": {},
	"Y": {},
}

// Variable is a variable declaration.
type Variable struct {
	Name  string
	Type  string
	Value string
}

// NewVariable creates a variable specification.
func NewVariable(expr Node, t *Type, value interface{}) (Node, error) {
	v := &Variable{
		Type: t.Name,
	}

	switch val := value.(type) {
	case nil:

	case *Identifier:
		v.Value = val.Name

	case *Value:
		if !t.InitializerUsed {
			return nil, ErrInvalidInitializer
		}
		v.Value = val.Value

	default:
		return nil, fmt.Errorf("unexpected intitializer type %T", value)
	}

	switch e := expr.(type) {
	case *Identifier:
		if _, ok := reservedNames[e.Name]; ok {
			return nil, ErrInvalidVariableName
		}
		v.Name = e.Name
		return v, nil

	case *NodeList:
		vars := &NodeList{}
		for _, node := range e.Nodes {
			id, ok := node.(*Identifier)
			if !ok {
				return nil, fmt.Errorf("unexpected node list expression type %T", value)
			}

			newVar := &Variable{
				Name:  id.Name,
				Type:  t.Name,
				Value: v.Value,
			}
			vars.Nodes = append(vars.Nodes, newVar)
		}
		return vars, nil

	default:
		return nil, fmt.Errorf("unexpected expression type %T", value)
	}
}

// String implement the fmt.Stringer interface.
func (v Variable) String() string {
	if v.Value == "" {
		return fmt.Sprintf("var, %s, %s", v.Name, v.Type)
	}
	return fmt.Sprintf("var, %s, %s, %s", v.Name, v.Type, v.Value)
}
