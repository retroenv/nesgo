package ast

import (
	"fmt"
	"strconv"
	"strings"
)

// Constant is a constant declaration.
type Constant struct {
	Name  string
	Value int64

	// if set, marks the value replaced from the resolved alias
	AliasName string
	// package will be first the import name, then replaced by the full package path
	AliasPackage string
}

// NewConstant returns a constant specification.
func NewConstant(expr *Identifier, arg interface{}) (Node, error) {
	constant := &Constant{
		Name: expr.Name,
	}

	switch val := arg.(type) {
	case *Identifier:
		sl := strings.Split(val.Name, ".")
		constant.AliasPackage = sl[0]
		constant.AliasName = sl[1]

	case *Value:
		i, err := strconv.ParseInt(val.Value, 0, 64)
		if err != nil {
			return nil, fmt.Errorf("parsing constant '%s': %w", val.Value, err)
		}
		constant.Value = i

	default:
		return nil, fmt.Errorf("type %T is not supported as constant argument", arg)
	}
	return constant, nil
}

// String implement the fmt.Stringer interface.
func (c Constant) String() string {
	return fmt.Sprintf("const, %s, %d", c.Name, c.Value)
}
