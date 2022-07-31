package mapper

import "github.com/retroenv/nesgo/pkg/mapper/mapperdb"

type hook struct {
	startAddress uint16
	endAddress   uint16

	onlyProxy bool // whether to continue mapper memory function execution after hook call
}

type readHook struct {
	hook

	hookFunc func(address uint16) uint8
}

type writeHook struct {
	hook

	hookFunc func(address uint16, value uint8)
}

func (h *hook) SetProxyOnly(proxy bool) {
	h.onlyProxy = proxy
}

// AddReadHook adds an address range read hook that gets called when a read from given range is made.
func (b *Base) AddReadHook(startAddress, endAddress uint16, hookFunc func(address uint16) uint8) mapperdb.Hook {
	hook := readHook{
		hook: hook{
			startAddress: startAddress,
			endAddress:   endAddress,
		},
		hookFunc: hookFunc,
	}
	b.readHooks = append(b.readHooks, hook)
	return &hook.hook
}

// AddWriteHook adds an address range write hook that gets called when a write into the given range is made.
func (b *Base) AddWriteHook(startAddress, endAddress uint16, hookFunc func(address uint16, value uint8)) mapperdb.Hook {
	hook := writeHook{
		hook: hook{
			startAddress: startAddress,
			endAddress:   endAddress,
		},
		hookFunc: hookFunc,
	}
	b.writeHooks = append(b.writeHooks, hook)
	return &hook.hook
}
