package ast

import (
	"fmt"
	"strings"
)

// NodeList contains a list of nodes.
type NodeList struct {
	Nodes []Node
}

// NewNodeList returns a statement list.
func NewNodeList(nodes ...interface{}) (Node, error) {
	top := &NodeList{}

	for _, node := range nodes {
		if node == nil {
			continue
		}
		n, ok := node.(Node)
		if !ok {
			return nil, fmt.Errorf("unexpected parameter type %T", node)
		}
		top.AddNodes(n)
	}

	return top, nil
}

// AddNodes adds a node to the node list if the node is not nil or its
// content is empty.
func (t *NodeList) AddNodes(nodes ...Node) {
	for _, node := range nodes {
		if node == nil {
			continue
		}

		switch n := node.(type) {
		case *NodeList:
			t.Nodes = append(t.Nodes, n.Nodes...)
		default:
			t.Nodes = append(t.Nodes, node)
		}
	}
}

// String implement the fmt.Stringer interface.
func (t NodeList) String() string {
	b := &strings.Builder{}
	for _, n := range t.Nodes {
		_, _ = fmt.Fprintf(b, "%s\n", n)
	}
	s := b.String()
	s = strings.TrimSuffix(s, "\n")
	return s
}

// ExpressionList contains a list of operands and expressions.
type ExpressionList struct {
	NodeList
}

// NewExpressionList returns a expression list.
func NewExpressionList(operand1 interface{}, expression string,
	operands ...interface{}) (interface{}, error) {
	list, ok := operand1.(*ExpressionList)
	if !ok {
		list = &ExpressionList{}
		node, ok := operand1.(Node)
		if !ok {
			return nil, fmt.Errorf("unexpected parameter type %T", operand1)
		}
		list.AddNodes(node)
	}

	s := &Statement{
		Op: expression,
	}
	list.AddNodes(s)

	for _, n := range operands {
		switch val := n.(type) {
		case *Identifier:
			list.AddNodes(val)
		case *Value:
			list.AddNodes(val)
		case *ExpressionList:
			if expression == "cast" {
				list.AddNodes(&Statement{Op: "("})
				list.AddNodes(val.Nodes...)
				list.AddNodes(&Statement{Op: ")"})
			} else {
				list.AddNodes(val.Nodes...)
			}
		default:
			return nil, fmt.Errorf("unexpected parameter type %T", n)
		}
	}

	return list, nil
}
