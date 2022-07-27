package mapper

type writeHookFunc func(address uint16, value uint8)

type writeHook struct {
	startAddress uint16
	endAddress   uint16
	hook         writeHookFunc
}

func (b *Base) addWriteHook(startAddress, endAddress uint16, hookFunc writeHookFunc) {
	hook := writeHook{
		startAddress: startAddress,
		endAddress:   endAddress,
		hook:         hookFunc,
	}
	b.writeHooks = append(b.writeHooks, hook)
}
