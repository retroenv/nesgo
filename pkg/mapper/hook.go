package mapper

type readHook struct {
	startAddress uint16
	endAddress   uint16
	hook         func(address uint16) uint8
}

type writeHook struct {
	startAddress uint16
	endAddress   uint16
	hook         func(address uint16, value uint8)
}

// AddReadHook adds an address range read hook that gets called when a read from given range is made.
func (b *Base) AddReadHook(startAddress, endAddress uint16, hookFunc func(address uint16) uint8) {
	hook := readHook{
		startAddress: startAddress,
		endAddress:   endAddress,
		hook:         hookFunc,
	}
	b.readHooks = append(b.readHooks, hook)
}

// AddWriteHook adds an address range write hook that gets called when a write into the given range is made.
func (b *Base) AddWriteHook(startAddress, endAddress uint16, hookFunc func(address uint16, value uint8)) {
	hook := writeHook{
		startAddress: startAddress,
		endAddress:   endAddress,
		hook:         hookFunc,
	}
	b.writeHooks = append(b.writeHooks, hook)
}
