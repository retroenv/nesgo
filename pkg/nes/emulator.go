//go:build !nesgo
// +build !nesgo

package nes

var (
	ram         *RAM
	cpu         *cPU
	controller1 controller
	controller2 controller
)

func init() {
	ram = newRAM()
	ppu = newPPU()
	reset()
}

func reset() {
	cpu = &cPU{}
	A = &cpu.A
	X = &cpu.X
	Y = &cpu.Y
	SP = 0xFD
	setFlags(0x24)
	ram.reset()
	ppu.reset()
	controller1.reset()
	controller2.reset()
}
