package compiler

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/retroenv/nesgo/internal/ast"
)

const (
	VarInitFunctionName     = "NesGoVariableInit"
	VarInitFunctionFullName = "github.com/retroenv/nesgo/pkg/neslib.NesGoVariableInit"
)

// Function is a function declaration.
type Function struct {
	Definition *ast.FunctionDefinition
	Body       *ast.NodeList

	Package *Package              // package that this function belongs to
	Labels  map[string]*ast.Label // map of all labels in this function
}

// resolveFunctionNodes parses all nodes of a function and resolves
// references to variables or constants, as well as assign statements.
func (c *Compiler) resolveFunctionNodes(f *Function) error {
	caller := f.Definition.Name
	newNodes := make([]ast.Node, 0, len(f.Body.Nodes))

	for _, node := range f.Body.Nodes {
		switch n := node.(type) {
		case *ast.Call:
			if err := c.resolveCall(f, n, caller); err != nil {
				return err
			}
			newNodes = append(newNodes, node)

		case *ast.Instruction:
			if err := c.resolveInstruction(f.Package, f, caller, n); err != nil {
				return fmt.Errorf("parsing instruction: %w", err)
			}
			newNodes = append(newNodes, node)

		case *ast.Statement:
			nodes, err := c.resolveStatement(f.Package, caller, n)
			if err != nil {
				return fmt.Errorf("handling statement: %w", err)
			}
			newNodes = append(newNodes, nodes...)

		default:
			newNodes = append(newNodes, node)
		}
	}

	f.Body.Nodes = newNodes
	addFunctionReturn(f)
	return nil
}

func (c *Compiler) resolveCall(f *Function, n *ast.Call, caller string) error {
	fullName, fun, err := f.Package.findFunction(c.packages, caller, n.Function)
	if err != nil {
		if caller == "main" && n.Function == "Start" {
			return c.handleStartCall(n)
		}
		return err
	}

	for _, param := range n.Parameter {
		identifier, ok := param.(*ast.Identifier)
		if !ok {
			continue
		}

		if variable, err := f.Package.findVariable(c.packages, caller, identifier.Name); err == nil {
			c.addVariable(variable)
		}
	}

	n.Function = fullName

	if _, ok := c.functionsAdded[fullName]; !ok {
		c.functionsToParse[fullName] = fun
	}
	return nil
}

func (c *Compiler) handleStartCall(n *ast.Call) error {
	arg := n.Parameter[0]
	identifier, ok := arg.(*ast.Identifier)
	if !ok {
		return fmt.Errorf("type %T is not supported as Start call parameter", arg)
	}
	c.resetHandler = identifier.Name

	if len(n.Parameter) > 2 {
		arg = n.Parameter[2]
		identifier, ok = arg.(*ast.Identifier)
		if !ok {
			return fmt.Errorf("type %T is not supported as Start call parameter", arg)
		}
		c.irqHandler = identifier.Name
	}

	if len(n.Parameter) > 1 {
		arg = n.Parameter[1]
		identifier, ok = arg.(*ast.Identifier)
		if !ok {
			return fmt.Errorf("type %T is not supported as Start call parameter", arg)
		}
		c.nmiHandler = identifier.Name
	}

	return nil
}

// resolveInstruction replaces all constant references of instruction
// arguments with the value of the constant.
func (c *Compiler) resolveInstruction(p *Package, f *Function, caller string,
	ins *ast.Instruction) error {
	for i, arg := range ins.Arguments {
		node, ok := arg.(*ast.ArgumentValue)
		if !ok {
			if _, ok = arg.(*ast.ExpressionList); ok {
				continue
			}
			return fmt.Errorf("wrong argument type %T for instruction", arg)
		}

		_, err := strconv.ParseUint(node.Value, 0, cpuRegisterSize)
		if err == nil {
			continue // skip if valid number
		}
		if _, ok = ast.CPURegisters[node.Value]; ok {
			continue // skip if references cpu register
		}

		if idx, ok := f.Definition.ParamIndex[node.Value]; ok {
			ins.Arguments[i] = &ast.ArgumentParam{Index: idx}
			continue // skip if references function parameter
		}

		if con, err := p.findConstant(c.packages, caller, node.Value); err == nil {
			ins.Arguments[i] = &ast.ArgumentValue{Value: fmt.Sprint(con.Value)}
			ins.Comment = node.Value
			continue
		}
		if variable, err := p.findVariable(c.packages, caller, node.Value); err == nil {
			ins.Arguments[i] = &ast.ArgumentValue{Value: variable.Name}
			c.addVariable(variable)
			continue
		}
		return fmt.Errorf("unknown argument '%s' for instruction '%s'", node.Value, ins.Name)
	}
	return nil
}

func (c *Compiler) resolveStatement(p *Package, caller string,
	st *ast.Statement) ([]ast.Node, error) {
	if st.Op != "=" {
		return []ast.Node{st}, nil
	}
	if len(st.Arguments) != 2 {
		return nil, errors.New("invalid assign statement argument count")
	}

	s := st.Arguments[1]
	con, err := p.findConstant(c.packages, caller, s)
	if err != nil {
		return nil, err
	}

	s = st.Arguments[0]
	variable, err := p.findVariable(c.packages, caller, s)
	if err != nil {
		return nil, err
	}
	c.addVariable(variable)

	load := &ast.Instruction{
		Name:      "lda",
		Arguments: ast.Arguments{&ast.ArgumentValue{Value: fmt.Sprint(con.Value)}},
	}
	store := &ast.Instruction{
		Name:      "sta",
		Arguments: ast.Arguments{&ast.ArgumentValue{Value: st.Arguments[0]}},
	}

	return []ast.Node{
		load, store,
	}, nil
}

// addFunction adds a function to the internal map and will return
// an error in case the name already exists.
func (c *Compiler) addFunction(fullName string, f *Function) error {
	if _, exists := c.functionsAdded[fullName]; exists {
		return fmt.Errorf("function '%s' is defined multiple times", fullName)
	}

	c.functionsAdded[fullName] = f
	c.functions = append(c.functions, f)
	delete(c.functionsToParse, fullName)
	return nil
}

// inlineFunctions checks all functions for calls to a function
// that is marked as inline and replaces the calls by inlining
// the call destination function code.
func (c *Compiler) inlineFunctions() error {
	nonInline := make([]*Function, 0, len(c.functions))

	var err error
	for _, fun := range c.functions {
		if fun.Definition.Inline {
			continue
		}
		nonInline = append(nonInline, fun)

		var body []ast.Node

		for _, node := range fun.Body.Nodes {
			switch call := node.(type) {
			case *ast.Call:
				body, err = c.inlineFunctionCall(fun, call, body)
				if err != nil {
					return err
				}

			default:
				body = append(body, call)
			}
		}

		fun.Body.Nodes = body
	}

	c.functions = nonInline
	return nil
}

// inlineFunctionCall inlines a call to a function if the function
// is marked as inline.
func (c *Compiler) inlineFunctionCall(functionContext *Function,
	call *ast.Call, body []ast.Node) ([]ast.Node, error) {
	f := c.functionsAdded[call.Function]
	if !f.Definition.Inline {
		body = append(body, call)
		return body, nil
	}

	nodes := fixLabelNameCollisions(functionContext, f.Body.Nodes)

	for _, node := range nodes {
		ins, ok := node.(*ast.Instruction)
		if !ok {
			continue
		}

		for i, arg := range ins.Arguments {
			if _, ok := arg.(*ast.ArgumentValue); ok {
				continue
			}
			if list, ok := arg.(*ast.ExpressionList); ok {
				val, err := evaluateExpressionList(f, call, c.packages, list)
				if err != nil {
					return nil, err
				}
				ins.Comment = formatNodeListComment(list.Nodes)
				ins.Arguments[i] = &ast.ArgumentValue{Value: val}
				continue
			}

			param, ok := arg.(*ast.ArgumentParam)
			if !ok {
				continue
			}

			val, err := getArgument(functionContext, c.packages, call.Parameter[param.Index])
			if err != nil {
				return nil, err
			}
			ins.Arguments[i] = &ast.ArgumentValue{Value: val}
		}
		body = append(body, node)
	}
	return body, nil
}

// getArgument returns the evaluated function call argument to use it for
// inlining the function.
func getArgument(functionContext *Function, packages map[string]*Package,
	item interface{}) (string, error) {
	switch n := item.(type) {
	case *ast.Value:
		return n.Value, nil

	case *ast.Identifier:
		p := functionContext.Package
		caller := functionContext.Definition.Name
		con, err := p.findConstant(packages, caller, n.Name)
		if err == nil {
			return fmt.Sprint(con.Value), nil
		}
		variable, err := p.findVariable(packages, caller, n.Name)
		if err == nil {
			return variable.Name, nil
		}

		return "", fmt.Errorf("identifier '%s' not found", n.Name)

	case *ast.ExpressionList:
		return evaluateExpressionList(functionContext, nil, packages, n)

	default:
		return "", fmt.Errorf("type %T is not supported as inlining call parameter", n)
	}
}

// addFunctionReturn adds a return at the end of functions unless the
// function is inlined or there is already a branching instruction as
// last node.
func addFunctionReturn(f *Function) {
	if f.Definition.Inline || f.Definition.Name == VarInitFunctionName {
		return
	}

	if len(f.Body.Nodes) > 0 {
		last := f.Body.Nodes[len(f.Body.Nodes)-1]
		switch n := last.(type) {
		case *ast.Branching:
			return

		case *ast.Instruction:
			if n.Name == ast.ReturnInstruction {
				return
			}
		}
	}

	f.Body.Nodes = append(f.Body.Nodes, ast.Instruction{
		Name: ast.ReturnInstruction,
		Comment: "automatically added by nesgo",
	})
}
