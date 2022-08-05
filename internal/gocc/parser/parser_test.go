package parser

var _ = []byte(`package main

import . "github.com/retroenv/nesgo/pkg/nes"

const PPU_CTRL = 0x2000

type inline any

const (
  A = 0x1000
  B = 0b1000 // comment
)

var i int8

var (
  C int8
  D int8 // comment
)


type inter any

// comment1

/* comment2 */

func test() {
	i = 1
	Inx()
}

func empty() {
}

func TestInline(... inline) {
	test()
x:
}

func main() {
	Init() // comment1
	Lda(1) /* comment2 */
	Sta(PPU_ADDR, 2)

a:
	if Bne() {
		goto a
	}
	return
}
`)

var _ = `package, main
import, ., github.com/retroenv/nesgo/pkg/nes
const, PPU_CTRL, 0x2000
const, A, 0x1000
const, B, 0b1000
var, i, int8
var, C, int8
var, D, int8
func, test
  op, =, i,1
  inst, inx, 
func, empty

func, inline, TestInline
  call, test
  label, x
func, main
  call, Init
  inst, lda, 1
  inst, sta, PPU_ADDR,2
  label, a
  inst, bne, a
  inst, rts, 
`
