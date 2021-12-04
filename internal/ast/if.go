package ast

import (
	"fmt"
)

// NewIfStatement returns an if statement, resolved as instructions.
func NewIfStatement(not bool, branch *Branching, block Node) (Node, error) {
	list, ok := block.(*NodeList)
	if !ok {
		return nil, fmt.Errorf("expression type %T is not supported for if statement blocks", block)
	}
	branch.Not = not

	switch len(list.Nodes) {
	case 0:
		return nil, ErrIfBranchingEmpty

	case 1:
		n := list.Nodes[0]
		if b, ok := n.(*Branching); ok {
			return handleIfBlockBranching(branch, b)
		}
	}

	labelIfNot := &Label{Name: "if_not_" + branch.Instruction}
	resolved := &NodeList{
		Nodes: []Node{
			branch,
		},
	}

	if not {
		branch.Destination = labelIfNot
		branch.DestinationName = labelIfNot.Name
	} else {
		labelIf := &Label{Name: "if_" + branch.Instruction}
		branch.Destination = labelIf
		branch.DestinationName = labelIf.Name
		jmp := &Branching{
			Instruction:     JmpInstruction,
			DestinationName: labelIfNot.Name,
			Destination:     labelIfNot,
		}

		resolved.AddNodes(jmp, labelIf)
	}
	resolved.AddNodes(list.Nodes...)
	resolved.AddNodes(labelIfNot)
	return resolved, nil
}

func handleIfBlockBranching(branch *Branching, block *Branching) (Node, error) {
	switch block.Instruction {
	case GotoInstruction:
		branch.DestinationName = block.DestinationName
		return branch, nil

	case breakOperator:
		return NewNodeList(branch, block)

	default:
		return nil, fmt.Errorf("if block does not supported branching "+
			"instruction '%s'", block.Instruction)
	}
}
