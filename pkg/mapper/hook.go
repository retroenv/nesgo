package mapper

type writeHook struct {
	startAddress uint16
	endAddress   uint16
	hook         func(address uint16, value uint8)
}

// AddWriteHook adds an address write hook that gets called when a write into the given address range is made.
func (b *Base) AddWriteHook(startAddress, endAddress uint16, hookFunc func(address uint16, value uint8)) {
	hook := writeHook{
		startAddress: startAddress,
		endAddress:   endAddress,
		hook:         hookFunc,
	}
	b.writeHooks = append(b.writeHooks, hook)
}
