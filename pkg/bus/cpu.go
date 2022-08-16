package bus

// CPUFlags contains the CPU flags.
type CPUFlags struct {
	C uint8
	Z uint8
	I uint8
	D uint8
	B uint8
	V uint8
	N uint8
}

// CPUState contains the current state of the CPU.
type CPUState struct {
	A      uint8
	X      uint8
	Y      uint8
	PC     uint16
	SP     uint8
	Cycles uint64
	Flags  CPUFlags
}

// CPU represents the Central Processing Unit.
type CPU interface {
	Cycles() uint64
	StallCycles(cycles uint16)
	State() CPUState
	TriggerIrq()
	TriggerNMI()
}
