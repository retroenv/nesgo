package main

import (
	"fmt"

	. "github.com/retroenv/nesgo/pkg/nes"
)

func main() {
	Start(resetHandler)
}

func resetHandler() {
	fmt.Println(*X) // print the content of the X CPU register on start, 0
	Ldx(0x34)       // set X to 0x34
	fmt.Println(*X) // print the content of the X CPU register, now 0x34
}
