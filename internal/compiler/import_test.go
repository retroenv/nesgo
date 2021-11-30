package compiler

import (
	"testing"
)

func TestCurrentPackage(t *testing.T) {
	if _, _, err := currentPackage(); err != nil {
		t.Error(err)
		return
	}
}
