package ast

import (
	"fmt"
	"strings"
)

// Function is a function declaration.
type Function struct {
	Definition *FunctionDefinition
	Body       *NodeList
}

// String implement the fmt.Stringer interface.
func (f Function) String() string {
	b := &strings.Builder{}
	_, _ = fmt.Fprint(b, "func, ")

	if f.Definition.Inline {
		_, _ = fmt.Fprint(b, "inline, ")
	}

	if len(f.Definition.Params) > 0 {
		_, _ = fmt.Fprint(b, "(")

		for i := 0; i < len(f.Definition.Params); i++ {
			p := f.Definition.Params[i]
			_, _ = fmt.Fprint(b, p.Name)
			if i+1 < len(f.Definition.Params) {
				_, _ = fmt.Fprint(b, ",")
			}
			// TODO print type
		}

		_, _ = fmt.Fprint(b, "), ")
	}

	_, _ = fmt.Fprint(b, f.Definition.Name)
	if f.Body != nil {
		_, _ = fmt.Fprintf(b, "\n%v", f.Body)
	}
	s := b.String()
	return s
}

// FunctionDefinition is a function definition.
type FunctionDefinition struct {
	Inline     bool
	Name       string
	Params     []*Variable
	ParamIndex map[string]int
}

// NewFunction returns a function declaration.
func NewFunction(def *FunctionDefinition, body interface{}) (Node, error) {
	f := &Function{
		Definition: def,
		Body:       &NodeList{},
	}
	if body != nil {
		l, ok := body.(*NodeList)
		if !ok {
			l = &NodeList{}
			l.AddNodes(body.(Node))
		}
		f.Body = l
	}
	return f, nil
}

// NewFunctionHeader returns a function header.
func NewFunctionHeader(id *Identifier, signature interface{}) (interface{}, error) {
	f := &FunctionDefinition{
		Name:       id.Name,
		ParamIndex: map[string]int{},
	}

	switch s := signature.(type) {
	case *Variable:
		f.Params = append(f.Params, s)
		f.ParamIndex[s.Name] = 0

	case *Inline:
		f.Inline = true

	case *NodeList:
		for i, node := range s.Nodes {
			switch n := node.(type) {
			case *Inline:
				f.Inline = true
			case *Variable:
				f.Params = append(f.Params, n)
				f.ParamIndex[n.Name] = i
			default:
				return nil, fmt.Errorf("type %T is not supported as function parameter", node)
			}
		}
	}

	if len(f.Params) > 0 && !f.Inline {
		return nil, fmt.Errorf("functions with parameters without inlining are currently not supported")
	}

	return f, nil
}

// Inline is an inline declaration.
type Inline struct{}

// NewInline returns an inline declaration.
func NewInline() (Node, error) {
	return &Inline{}, nil
}

// String implement the fmt.Stringer interface.
func (i Inline) String() string {
	return ""
}

// NewUntypedParamListEntry handles a function parameter list entry without
// a type specifier.
func NewUntypedParamListEntry(name string, definition interface{}) interface{} {
	list, ok := definition.(*NodeList)
	if !ok {
		return nil
	}
	n, ok := list.Nodes[0].(*Variable)
	if !ok {
		return nil
	}
	v := &Variable{
		Name: name,
		Type: n.Type,
	}
	return v
}
