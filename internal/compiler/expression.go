package compiler

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/retroenv/nesgo/internal/ast"
)

func evaluateExpressionList(functionContext *Function, call *ast.Call,
	packages map[string]*Package, list *ast.ExpressionList) (string, error) {
	rpn, err := expressionListToRPN(list)
	if err != nil {
		return "", fmt.Errorf("parsing expression list: %w", err)
	}

	stack := make([]interface{}, 0)

	for i := 0; i < len(rpn) || len(stack) > 1; i++ {
		item := rpn[i]
		switch n := item.(type) {
		case *ast.Type:
			stack = append(stack, item)

		case *ast.Value:
			val, err := strconv.ParseInt(n.Value, 0, 16)
			if err != nil {
				return "", fmt.Errorf("parsing constant '%s': %w", n.Value, err)
			}
			stack = append(stack, int(val))

		case *ast.Statement:
			if n.Op == "(" {
				continue
			}

			stack, err = evaluateStatement(stack, n)
			if err != nil {
				return "", fmt.Errorf("evaluating statement: %w", err)
			}

		case *ast.Identifier:
			stack, err = evaluateIdentifier(stack, n.Name, functionContext, call, packages)
			if err != nil {
				return "", fmt.Errorf("evaluating identifier: %w", err)
			}

		default:
			return "", fmt.Errorf("can not evaluate type %T", n)
		}
	}

	item := stack[0]
	i, ok := item.(int)
	if !ok {
		return "", fmt.Errorf("unexpected result type %T", item)
	}
	s := strconv.Itoa(i)
	return s, nil
}

func popStack(stack []interface{}) (interface{}, []interface{}) {
	return stack[len(stack)-1], stack[:len(stack)-1]
}

func evaluateStatement(stack []interface{}, statement *ast.Statement) ([]interface{}, error) {
	var a, b interface{}
	b, stack = popStack(stack)
	a, stack = popStack(stack)

	if statement.Op == "cast" {
		val, err := evaluateCast(a, b)
		if err != nil {
			return nil, fmt.Errorf("evaluating cast: %w", err)
		}
		stack = append(stack, val)
	} else {
		val, err := evaluateOperator(a, b, statement)
		if err != nil {
			return nil, fmt.Errorf("evaluating operator: %w", err)
		}
		stack = append(stack, val)
	}
	return stack, nil
}

func evaluateIdentifier(stack []interface{}, identifier string, functionContext *Function,
	call *ast.Call, packages map[string]*Package) ([]interface{}, error) {
	if call != nil {
		return evaluateCallIdentifier(stack, identifier, functionContext, call, packages)
	}
	return evaluateConstant(stack, identifier, functionContext, packages)
}

func evaluateCallIdentifier(stack []interface{}, identifier string,
	functionContext *Function, call *ast.Call, packages map[string]*Package) ([]interface{}, error) {
	idx, ok := functionContext.Definition.ParamIndex[identifier]
	if !ok {
		return nil, fmt.Errorf("call identifier '%s' is not a parameter", identifier)
	}

	param := call.Parameter[idx]
	switch p := param.(type) {
	case *ast.Value:
		val, err := strconv.ParseInt(p.Value, 0, 16)
		if err != nil {
			return nil, fmt.Errorf("parsing constant '%s': %w", p.Value, err)
		}

		stack = append(stack, int(val))
		return stack, nil

	case *ast.Identifier:
		return evaluateConstant(stack, p.Name, functionContext, packages)

	default:
		return nil, fmt.Errorf("unexpected call identifier param type %T", param)
	}
}

func evaluateConstant(stack []interface{}, identifier string,
	functionContext *Function, packages map[string]*Package) ([]interface{}, error) {
	p := functionContext.Package
	con, err := p.findConstant(packages, functionContext.Definition.Name, identifier)
	if err != nil {
		return nil, fmt.Errorf("constant '%s' for evaluation not found", identifier)
	}

	stack = append(stack, int(con.Value))
	return stack, nil
}

func evaluateOperator(p1, p2 interface{}, statement *ast.Statement) (int, error) {
	a, ok := p1.(int)
	if !ok {
		return 0, fmt.Errorf("unexpected operator param type %T", p1)
	}
	b, ok := p2.(int)
	if !ok {
		return 0, fmt.Errorf("unexpected operator param type %T", p2)
	}

	switch statement.Op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "|":
		return a | b, nil
	case "^":
		return a ^ b, nil
	case ">>":
		return a >> b, nil
	case "<<":
		return a >> b, nil
	default:
		return 0, fmt.Errorf("can not evaluate operator '%s'", statement.Op)
	}
}

func evaluateCast(p1, p2 interface{}) (int, error) {
	typ, ok := p1.(*ast.Type)
	if !ok {
		return 0, fmt.Errorf("unexpected cast param type %T", p1)
	}
	i, ok := p2.(int)
	if !ok {
		return 0, fmt.Errorf("unexpected cast param type %T", p2)
	}

	switch typ.Name {
	case "int8":
		val := int8(i)
		return int(val), nil
	case "uint8":
		val := uint8(i)
		return int(val), nil
	default:
		return 0, fmt.Errorf("unsupported cast type '%s'", typ.Name)
	}
}

// convert an expression list to Reverse Polish Notation.
func expressionListToRPN(list *ast.ExpressionList) ([]interface{}, error) {
	var rpn []interface{}
	var ops []*ast.Statement

	var leftParen, rightParen int

	for _, item := range list.Nodes {
		switch n := item.(type) {
		case *ast.Statement:
			if n.Op == "(" {
				leftParen++
			}
			if n.Op != ")" {
				ops = append(ops, n)
				continue
			}

			rightParen++
			for len(ops) != 0 {
				op := ops[len(ops)-1]
				ops = ops[:len(ops)-1]
				if op.Op != "(" {
					rpn = append(rpn, op)
				}
			}

		default:
			rpn = append(rpn, n)
			if len(ops) == 0 {
				continue
			}

			op := ops[len(ops)-1]
			ops = ops[:len(ops)-1]

			rpn = append(rpn, op)
		}
	}

	if leftParen != rightParen {
		return nil, errors.New("invalid number of parenthesis")
	}

	return rpn, nil
}
