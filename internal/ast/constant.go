package ast

import (
	"fmt"
	"strconv"
)

// Constant is a constant declaration.
type Constant struct {
	Name  string
	Value int64
}

// NewConstant returns a constant specification.
func NewConstant(expr *Identifier, val *Value) (Node, error) {
	i, err := strconv.ParseInt(val.Value, 0, 64)
	if err != nil {
		return nil, fmt.Errorf("parsing constant '%s': %w", val.Value, err)
	}

	return &Constant{
		Name:  expr.Name,
		Value: i,
	}, nil
}

// String implement the fmt.Stringer interface.
func (c Constant) String() string {
	return fmt.Sprintf("const, %s, %d", c.Name, c.Value)
}
