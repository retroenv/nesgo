package ast

import (
	"fmt"
)

// NewForStatement returns a for statement resolved as instructions.
func NewForStatement(clause, block Node) (Node, error) {
	blockNodes, ok := block.(*NodeList)
	if !ok {
		return nil, fmt.Errorf("for statement blocks does not support type %T", block)
	}

	label := &Label{Name: "loop"}
	list := &NodeList{}
	initializer, condition, post, err := extractForClauseItems(clause)
	if err != nil {
		return nil, err
	}

	if initializer != nil {
		list.AddNodes(initializer)
	}

	list.AddNodes(label)

	var finalLabel *Label
	if len(blockNodes.Nodes) > 0 {
		blockNodes.Nodes, finalLabel, err = processForBody(blockNodes.Nodes)
		if err != nil {
			return nil, err
		}
		list.AddNodes(blockNodes.Nodes...)
	}

	if post != nil {
		ex, ok := post.(*ExpressionList)
		if ok {
			ins, err := resolveForPostExpressionList(ex)
			if err != nil {
				return nil, err
			}
			list.AddNodes(ins)
		} else {
			list.AddNodes(post)
		}
	}

	conditionNodes, err := processForCondition(condition, label)
	if err != nil {
		return nil, err
	}
	list.AddNodes(conditionNodes...)
	if finalLabel != nil {
		list.AddNodes(finalLabel)
	}

	return list, nil
}

func processForCondition(condition Node, label *Label) ([]Node, error) {
	var nodes []Node

	if condition == nil {
		condition = &Branching{
			Instruction:     JmpInstruction,
			DestinationName: label.Name,
			Destination:     label,
		}
	} else {
		if ex, ok := condition.(*ExpressionList); ok {
			ins, branch, err := resolveForConditionExpressionList(ex)
			if err != nil {
				return nil, err
			}
			nodes = append(nodes, ins)

			condition = &Branching{
				Instruction:     branch,
				DestinationName: label.Name,
				Destination:     label,
			}
		}
	}

	switch n := condition.(type) {
	case nil:
		break

	case *Branching:
		if n.Instruction != JmpInstruction {
			label.Name = n.Instruction + "_loop"
		}
		n.DestinationName = label.Name
		n.Destination = label

	case *Instruction:
		return nil, fmt.Errorf("expected a branching instruction as for "+
			"statement condition but found '%s'", n.Name)

	default:
		return nil, fmt.Errorf("for statement conditions do not support type %T", condition)
	}

	nodes = append(nodes, condition)
	return nodes, nil
}

// processForBody processes a for body block and checks for break
// instructions and returns a fixed up body and optional label for
// the end of the for code to allow breaking to.
func processForBody(body []Node) ([]Node, *Label, error) {
	var finalLabel *Label

	var newBody []Node
	for i, node := range body {
		branch, ok := node.(*Branching)
		if !ok || branch.Instruction != BreakOperator {
			newBody = append(newBody, node)
			continue
		}
		if i == 0 {
			return nil, nil, ErrBreakNotAfterBranching
		}

		previousNode := body[i-1]
		branch, ok = previousNode.(*Branching)
		if !ok {
			return nil, nil, fmt.Errorf("break statement has to be after a "+
				"branching instruction but found type %T", previousNode)
		}
		if finalLabel == nil {
			finalLabel = &Label{Name: "loop_end"}
		}

		branch.Destination = finalLabel
		branch.DestinationName = finalLabel.Name
	}

	return newBody, finalLabel, nil
}

// NewForClause returns a clause list for a for statement.
func NewForClause(init, condition, post Node) (Node, error) {
	return &NodeList{
		Nodes: []Node{init, condition, post},
	}, nil
}

// extractForClauseItems returns the 3 for clause items from the node list
// and returns them, each item can be nil.
func extractForClauseItems(clause Node) (initializer, condition, post Node, err error) {
	if clause != nil {
		clauseList, ok := clause.(*NodeList)
		if !ok {
			return nil, nil, nil,
				fmt.Errorf("for statement clause does not support type %T", clause)
		}

		initializer = clauseList.Nodes[0]
		condition = clauseList.Nodes[1]
		post = clauseList.Nodes[2]
	}
	return
}

// resolveForConditionExpressionList converts the for condition expression
// list to resolved nodes with instructions.
func resolveForConditionExpressionList(list *ExpressionList) (Node, string, error) {
	if len(list.Nodes) != 3 {
		return nil, "", ErrForOnlySimpleConditions
	}

	id, ok := list.Nodes[0].(*Identifier)
	if !ok {
		return nil, "", fmt.Errorf("for condition identifier does not support type %T", list.Nodes[0])
	}
	st, ok := list.Nodes[1].(*Statement)
	if !ok {
		return nil, "", fmt.Errorf("for condition statement does not support type %T", list.Nodes[1])
	}
	val, ok := list.Nodes[2].(*Value)
	if !ok {
		return nil, "", fmt.Errorf("for condition value does not support type %T", list.Nodes[2])
	}

	ins := &Instruction{
		Comment: list.String(),
	}
	switch id.Name {
	case "X":
		ins.Name = "cpx"
	case "Y":
		ins.Name = "cpy"
	default:
		return nil, "", fmt.Errorf("for condition identifier does not support '%s'", id.Name)
	}
	ins.Arguments = Arguments{val}

	var branch string
	switch st.Op {
	case "<":
		branch = "bcc"
	case "!=":
		branch = "bne"
	case ">=":
		branch = "bcs"

	default:
		return nil, "", fmt.Errorf("for condition statement does not support '%s'", st.Op)
	}

	return ins, branch, nil
}

// resolveForConditionExpressionList converts the for post expression
// list to a resolved instruction node.
func resolveForPostExpressionList(list *ExpressionList) (Node, error) {
	if len(list.Nodes) != 2 {
		return nil, ErrForOnlySimplePostExpressions
	}

	id, ok := list.Nodes[0].(*Identifier)
	if !ok {
		return nil, fmt.Errorf("for post expressions does not support type %T", list.Nodes[0])
	}
	st, ok := list.Nodes[1].(*Statement)
	if !ok {
		return nil, fmt.Errorf("for post expressions does not support type %T", list.Nodes[1])
	}

	ins := &Instruction{
		Comment: list.String(),
	}

	switch st.Op {
	case "++":
		ins.Name = "in"
	case "--":
		ins.Name = "de"
	default:
		return nil, fmt.Errorf("for post expressions statement does not support '%s'", st.Op)
	}

	switch id.Name {
	case "X":
		ins.Name += "x"
	case "Y":
		ins.Name += "y"
	default:
		return nil, fmt.Errorf("for post expressions identifier does not support '%s'", id.Name)
	}

	return ins, nil
}
