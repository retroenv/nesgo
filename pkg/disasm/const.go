package disasm

import (
	"github.com/retroenv/nesgo/pkg/apu"
	"github.com/retroenv/nesgo/pkg/controller"
	"github.com/retroenv/nesgo/pkg/ppu"
)

// buildConstMap builds the map of all known NES constants from all
// modules that maps an address to a constant name.
func buildConstMap() map[uint16]string {
	m := map[uint16]string{}
	mergeMaps(m, apu.AddressToName)
	mergeMaps(m, controller.AddressToName)
	mergeMaps(m, ppu.AddressToName)
	return m
}

func mergeMaps(destination, source map[uint16]string) {
	for k, v := range source {
		destination[k] = v
	}
}
