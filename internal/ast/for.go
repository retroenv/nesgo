package ast

import (
	"fmt"
)

type forStatement struct {
	condition Node

	initializer    Node
	bodyStart      *Label
	body           []Node
	postStart      *Label
	post           Node
	ConditionStart *Label
	conditionNodes []Node
	loopEnd        *Label
}

// NewForClause returns a clause list for a for statement.
func NewForClause(init, condition, post Node) (Node, error) {
	return &NodeList{
		Nodes: []Node{init, condition, post},
	}, nil
}

// NewForStatement returns a for statement resolved as instructions.
func NewForStatement(clause, block Node) (Node, error) {
	blockNodes, ok := block.(*NodeList)
	if !ok {
		return nil, fmt.Errorf("for statement blocks does not support type %T", block)
	}

	f := forStatement{
		bodyStart: &Label{Name: "loop"},
	}
	if err := f.extractForClauseItems(clause); err != nil {
		return nil, err
	}
	if err := f.processForCondition(); err != nil {
		return nil, err
	}
	if err := f.processForBody(blockNodes.Nodes); err != nil {
		return nil, err
	}

	if f.post != nil {
		ex, ok := f.post.(*ExpressionList)
		if ok {
			ins, err := resolveForPostExpressionList(ex)
			if err != nil {
				return nil, err
			}
			f.post = ins
		}
	}

	list := f.nodes()
	return list, nil
}

func (f *forStatement) processForCondition() error {
	if f.condition == nil {
		f.condition = &Branching{
			Instruction:     JmpInstruction,
			DestinationName: f.bodyStart.Name,
			Destination:     f.bodyStart,
		}
	} else {
		if ex, ok := f.condition.(*ExpressionList); ok {
			ins, branch, err := resolveForConditionExpressionList(ex)
			if err != nil {
				return err
			}
			f.conditionNodes = append(f.conditionNodes, ins)

			f.condition = &Branching{
				Instruction:     branch,
				DestinationName: f.bodyStart.Name,
				Destination:     f.bodyStart,
			}
		}
	}

	switch n := f.condition.(type) {
	case nil:
		break

	case *Branching:
		if n.Instruction != JmpInstruction {
			f.bodyStart.Name = n.Instruction + "_loop"
		}
		n.DestinationName = f.bodyStart.Name
		n.Destination = f.bodyStart

	case *Instruction:
		return fmt.Errorf("expected a branching instruction as for statement condition but found '%s'", n.Name)

	default:
		return fmt.Errorf("for statement conditions do not support type %T", f.condition)
	}

	f.conditionNodes = append(f.conditionNodes, f.condition)
	return nil
}

// processForBody processes a for body block and checks for break
// instructions and returns a fixed up body and optional label for
// the end of the for condition and code to allow jumping to.
func (f *forStatement) processForBody(body []Node) error {
	if f.ConditionStart != nil {
		f.body = []Node{&Branching{
			Instruction:     JmpInstruction,
			DestinationName: f.ConditionStart.Name,
			Destination:     f.ConditionStart,
		}}
	}

	for i, node := range body {
		branch, ok := node.(*Branching)
		if !ok {
			f.body = append(f.body, node)
			continue
		}

		switch branch.Instruction {
		case breakStatement:
			if i == 0 {
				return ErrBreakNotAfterBranching // TODO: support
			}

			previousNode := body[i-1]
			branch, ok = previousNode.(*Branching)
			if !ok {
				return fmt.Errorf("break statement has to be after a branching instruction but found type %T", previousNode)
			}

			if f.loopEnd == nil {
				f.loopEnd = &Label{Name: "loop_end"}
			}
			branch.Destination = f.loopEnd
			branch.DestinationName = f.loopEnd.Name

		case continueStatement:
			if i == 0 {
				return ErrContinueNotAfterBranching // TODO: support
			}

			previousNode := body[i-1]
			branch, ok = previousNode.(*Branching)
			if !ok {
				return fmt.Errorf("break statement has to be after a branching instruction but found type %T", previousNode)
			}
			if f.postStart == nil {
				f.postStart = &Label{Name: "loop_post"}
			}
			branch.Destination = f.postStart
			branch.DestinationName = f.postStart.Name

		default:
			f.body = append(f.body, node)
		}
	}
	return nil
}

// extractForClauseItems extracts the 3 for clause items from the node list.
func (f *forStatement) extractForClauseItems(clause Node) error {
	if clause == nil {
		return nil
	}

	clauseList, ok := clause.(*NodeList)
	if !ok {
		return fmt.Errorf("for statement clause does not support type %T", clause)
	}

	f.initializer = clauseList.Nodes[0]
	f.condition = clauseList.Nodes[1]
	f.post = clauseList.Nodes[2]
	return nil
}

func (f *forStatement) nodes() Node {
	list := &NodeList{}
	list.AddNodes(f.initializer, f.bodyStart)
	list.AddNodes(f.body...)
	list.AddNodes(f.post)
	if f.postStart != nil {
		list.AddNodes(f.postStart)
	}
	if f.ConditionStart != nil {
		list.AddNodes(f.ConditionStart)
	}
	list.AddNodes(f.conditionNodes...)
	if f.loopEnd != nil {
		list.AddNodes(f.loopEnd)
	}
	return list
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
