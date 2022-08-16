//go:build !nesgo

package debugger

import (
	"encoding/json"
	"net/http"
)

type cpuFlags struct {
	C hexByte `json:"c"`
	Z hexByte `json:"z"`
	I hexByte `json:"i"`
	D hexByte `json:"d"`
	B hexByte `json:"b"`
	V hexByte `json:"v"`
	N hexByte `json:"n"`
}

type cpuState struct {
	A      hexByte  `json:"a"`
	X      hexByte  `json:"x"`
	Y      hexByte  `json:"y"`
	PC     hexWord  `json:"pc"`
	SP     hexByte  `json:"sp"`
	Cycles hexQword `json:"cycles"`
	Flags  cpuFlags `json:"flags"`
}

func (d *Debugger) cpuState(w http.ResponseWriter, r *http.Request) {
	state := d.bus.CPU.State()

	res := cpuState{
		A:      hexByte(state.A),
		X:      hexByte(state.X),
		Y:      hexByte(state.Y),
		PC:     hexWord(state.PC),
		SP:     hexByte(state.SP),
		Cycles: hexQword(state.Cycles),
		Flags: cpuFlags{
			C: hexByte(state.Flags.C),
			Z: hexByte(state.Flags.Z),
			I: hexByte(state.Flags.I),
			D: hexByte(state.Flags.D),
			B: hexByte(state.Flags.B),
			V: hexByte(state.Flags.V),
			N: hexByte(state.Flags.N),
		},
	}

	_ = json.NewEncoder(w).Encode(res)
}

func (d *Debugger) cpuPause(w http.ResponseWriter, r *http.Request) {
	// TODO implement
}
