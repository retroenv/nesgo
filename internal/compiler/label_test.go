package compiler

import (
	"testing"

	"github.com/retroenv/nesgo/internal/ast"
	"github.com/retroenv/retrogolib/assert"
)

func TestFixLabelNameCollisions(t *testing.T) {
	fun := &Function{
		Labels: map[string]*ast.Label{
			"test": {Name: "test"},
		},
	}

	label := &ast.Label{Name: "test"}
	jmp1 := &ast.Branching{Instruction: "jmp", DestinationName: "test", Destination: label}
	jmp2 := &ast.Branching{Instruction: "bne", DestinationName: "test", Destination: label}

	nodes := []ast.Node{label, jmp1, jmp2}
	fixed := fixLabelNameCollisions(fun, nodes)

	assert.True(t, len(nodes) == len(fixed))

	// make sure that the original objects have not been modified
	assert.Equal(t, "test", label.Name)
	assert.Equal(t, "test", jmp1.DestinationName)
	assert.Equal(t, "test", jmp2.DestinationName)

	newLabel, ok := fixed[0].(*ast.Label)
	assert.True(t, ok)
	newJmp1, ok := fixed[1].(*ast.Branching)
	assert.True(t, ok)
	newJmp2, ok := fixed[2].(*ast.Branching)
	assert.True(t, ok)

	assert.False(t, newLabel.Name == label.Name)
	assert.Equal(t, newLabel.Name, newJmp1.DestinationName)
	assert.True(t, newLabel == newJmp1.Destination)
	assert.Equal(t, newLabel.Name, newJmp2.DestinationName)
	assert.True(t, newLabel == newJmp2.Destination)
}
