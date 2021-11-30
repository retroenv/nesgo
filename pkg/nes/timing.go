//go:build !nesgo
// +build !nesgo

package nes

import "time"

// TODO disable timing in unit tests
// account for exact cycles
func timeInstructionExecution() {
	time.Sleep(time.Microsecond)
}
