package ast

import (
	"errors"
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
				_, _ = fmt.Fprint(b, ", ")
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
	Inline           bool
	Name             string
	Params           []*Variable
	ParamInitializer map[string]*Instruction
	ParamIndex       map[string]int // maps parameter name to index
}

// NewFunction returns a function declaration.
func NewFunction(def *FunctionDefinition, body any) (Node, error) {
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

	if len(f.Definition.Params) > 0 && !f.Definition.Inline {
		if err := f.processParams(); err != nil {
			return nil, err
		}
	}

	return f, nil
}

// NewFunctionHeader returns a function header.
func NewFunctionHeader(id *Identifier, signature any) (any, error) {
	f := &FunctionDefinition{
		Name:             id.Name,
		ParamInitializer: map[string]*Instruction{},
		ParamIndex:       map[string]int{},
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

	return f, nil
}

func (f *Function) processParams() error {
	if f.Body == nil {
		return errors.New("missing function body")
	}
	body := f.Body.Nodes[:0] // filter initializer nodes without allocation

	// look for instructions that use the function parameters
	for _, n := range f.Body.Nodes {
		ins, ok := n.(*Instruction)
		if !ok || len(ins.Arguments) == 0 {
			body = append(body, n)
			continue // not an instruction or no arguments
		}

		arg, ok := ins.Arguments[0].(*ArgumentValue)
		if !ok {
			body = append(body, n)
			continue // wrong type
		}
		if _, ok = f.Definition.ParamIndex[arg.Value]; !ok {
			body = append(body, n)
			continue // not a param of the function
		}

		if _, ok = f.Definition.ParamInitializer[arg.Value]; ok {
			return fmt.Errorf("parameter %s can only be referenced once", arg.Value)
		}
		f.Definition.ParamInitializer[arg.Value] = ins
	}

	f.Body.Nodes = body
	return nil
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
func NewUntypedParamListEntry(name string, definition any) any {
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
