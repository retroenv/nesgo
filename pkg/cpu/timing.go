//go:build !nesgo
// +build !nesgo

package cpu

import (
	"fmt"
	"time"
)

// TODO disable timing in unit tests
func instructionHook(instruction *Instruction, params ...interface{}) {
	// TODO get addressing mode based on passed params
	// TODO account for exact cycles

	if len(params) != 100 { // TODO add tracing
		fmt.Println(instruction.Name)
	}

	time.Sleep(time.Microsecond)
}
