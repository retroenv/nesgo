package main

import (
	"fmt"

	_ "github.com/retroenv/nesgo/pkg/gui"
	. "github.com/retroenv/nesgo/pkg/nes"
	. "github.com/retroenv/nesgo/pkg/neslib"
)

var keyState = NewUint8(0)

func main() {
	Start(resetHandler)
}

func resetHandler() {
	for {
		ReadJoypad(0)
		Cmp(keyState)
		if Bne() {
			fmt.Printf("key state changed: %d\n", *A)
			Sta(keyState)
		}
	}
}
