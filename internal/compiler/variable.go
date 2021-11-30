package compiler

import (
	"errors"
	"fmt"
	"strings"

	"github.com/retroenv/nesgo/internal/ast"
)

func (p *Package) addVariables(variables []*ast.Variable) error {
	for _, v := range variables {
		name := v.Name
		if _, ok := p.variables[name]; ok {
			return fmt.Errorf("variable '%s' is defined multiple times", name)
		}
		p.variables[name] = v
	}
	return nil
}

func (p *Package) findVariable(packages map[string]*Package,
	caller, variable string) (*ast.Variable, error) {
	sl := strings.Split(variable, ".")
	if len(sl) > 1 {
		// TODO support non dot imported functions
		return nil, errors.New("non dot imports from external packages are not support yet")
	}

	if v, ok := p.variables[variable]; ok {
		return v, nil
	}

	imports := p.functionFile[caller].Imports
	for _, imp := range imports {
		impPack := packages[imp.Path]
		if v, ok := impPack.variables[variable]; ok {
			return v, nil
		}
	}

	return nil, fmt.Errorf("variable '%s' can not be found", variable)
}

func (c *Compiler) addVariable(variable *ast.Variable) {
	c.variables[variable.Name] = variable
	if variable.Value != "" {
		c.variablesInitialized[variable.Name] = variable
	}
}

func (c *Compiler) createVariableInitializations(mainPackage *Package) error {
	resetHandler := mainPackage.functions[c.resetHandler]

	fun, userCalled := c.functionsAdded[VarInitFunctionFullName]
	if !userCalled {
		call := &ast.Call{Function: VarInitFunctionName}
		resetHandler.Body.Nodes = append([]ast.Node{call}, resetHandler.Body.Nodes...)

		fun = &Function{
			Definition: &ast.FunctionDefinition{
				Inline: false,
				Name:   VarInitFunctionName,
			},
			Body: &ast.NodeList{},
		}
	}

	for _, variable := range c.variablesInitialized {
		value := variable.Value
		if con, err := mainPackage.findConstant(c.packages, c.resetHandler, value); err == nil {
			value = fmt.Sprint(con.Value)
		}

		load := &ast.Instruction{
			Name: "lda",
			Arguments: ast.Arguments{
				&ast.ArgumentValue{Value: value}},
		}
		store := &ast.Instruction{
			Name: "sta",
			Arguments: ast.Arguments{
				&ast.ArgumentValue{Value: variable.Name}},
		}
		fun.Body.AddNodes(load, store)
	}

	ret, _ := ast.NewReturnStatement()
	fun.Body.AddNodes(ret)

	if !userCalled {
		return c.addFunction(VarInitFunctionName, fun)
	}
	return nil
}
