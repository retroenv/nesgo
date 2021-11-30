package compiler

import (
	"fmt"
	"math/rand"

	"github.com/retroenv/nesgo/internal/ast"
)

// collectAndLinkAllLabels collects all labels and stores them in the functions
// label map. All branching instructions will be linked to their label pointer
// if it does not exist already.
func collectAndLinkAllLabels(fun *Function) error {
	for _, item := range fun.Body.Nodes {
		if label, ok := item.(*ast.Label); ok {
			fun.Labels[label.Name] = label
		}
	}

	for _, item := range fun.Body.Nodes {
		branch, ok := item.(*ast.Branching)
		if !ok {
			continue
		}
		label, ok := fun.Labels[branch.DestinationName]
		if !ok {
			return fmt.Errorf("branching destination label '%s' not found", branch.DestinationName)
		}
		branch.Destination = label
	}
	return nil
}

// fixLabelNameCollisions makes sure that the labels to inline have a unique
// name in the function context. A copy of the nodes is returned to allow
// modification at the caller .
func fixLabelNameCollisions(fun *Function, nodes []ast.Node) []ast.Node {
	newNodes := make([]ast.Node, 0, len(nodes))

	newLabels := map[string]*ast.Label{}
	for _, node := range nodes {
		switch n := node.(type) {
		case *ast.Instruction:
			cp := &ast.Instruction{
				Name:      n.Name,
				Arguments: make(ast.Arguments, len(n.Arguments)),
				Comment:   n.Comment,
			}
			copy(cp.Arguments, n.Arguments)
			newNodes = append(newNodes, cp)

		case *ast.Branching:
			cp := &ast.Branching{
				Instruction:     n.Instruction,
				DestinationName: n.DestinationName,
			}
			newNodes = append(newNodes, cp)

		case *ast.Label:
			cp := &ast.Label{
				Name: n.Name,
			}
			newNodes = append(newNodes, cp)
			newLabels[n.Name] = cp

		default:
			newNodes = append(newNodes, n)
		}
	}

	if len(newLabels) > 0 {
		relinkBranches(fun, newNodes, newLabels)
	}
	return newNodes
}

// relinkBranches links all branching instructions to labels that in case
// of a name collision got renamed.
func relinkBranches(fun *Function, nodes []ast.Node, newLabels map[string]*ast.Label) {
	for _, node := range nodes {
		branch, ok := node.(*ast.Branching)
		if !ok {
			continue
		}

		label := newLabels[branch.DestinationName]
		// point to copy
		branch.Destination = label
		// update in case a previous branching to the multiple times
		// referenced label got a new name already
		branch.DestinationName = label.Name

		// generate a new label name in case it is in use at the
		// function scope and has not been randomized yet.
		funLabel, ok := fun.Labels[branch.DestinationName]
		if ok && funLabel.Name == branch.DestinationName {
			label.Name = randomLabelName(fun, label.Name)
			branch.DestinationName = label.Name
		}
	}

	// add new labels to the function context
	for _, label := range newLabels {
		fun.Labels[label.Name] = label
	}
}

// randomLabelName returns a random label name that is not used in
// the given function.
func randomLabelName(fun *Function, base string) string {
	for {
		i := rand.Int31() % 0xfff
		label := fmt.Sprintf("%s_%d", base, i)
		if _, ok := fun.Labels[label]; !ok {
			return label
		}
	}
}
