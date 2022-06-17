package disasm

import (
	"fmt"

	. "github.com/retroenv/nesgo/pkg/addressing"
	"github.com/retroenv/nesgo/pkg/apu"
	"github.com/retroenv/nesgo/pkg/controller"
	"github.com/retroenv/nesgo/pkg/ppu"
)

type constTranslation struct {
	Read  string
	Write string
}

// buildConstMap builds the map of all known NES constants from all
// modules that maps an address to a constant name.
func buildConstMap() (map[uint16]constTranslation, error) {
	m := map[uint16]constTranslation{}
	if err := mergeConstantsMaps(m, apu.AddressToName); err != nil {
		return nil, fmt.Errorf("processing apu constants: %w", err)
	}
	if err := mergeConstantsMaps(m, controller.AddressToName); err != nil {
		return nil, fmt.Errorf("processing controller constants: %w", err)
	}
	if err := mergeConstantsMaps(m, ppu.AddressToName); err != nil {
		return nil, fmt.Errorf("processing ppu constants: %w", err)
	}
	return m, nil
}

func mergeConstantsMaps(destination map[uint16]constTranslation, source map[uint16]AccessModeConstant) error {
	for address, constantInfo := range source {
		translation := destination[address]

		if constantInfo.Mode&ReadAccess != 0 {
			if translation.Read != "" {
				return fmt.Errorf("constant with address 0x%04X and read mode is defined twice", address)
			}
			translation.Read = constantInfo.Constant
		}

		if constantInfo.Mode&WriteAccess != 0 {
			if translation.Write != "" {
				return fmt.Errorf("constant with address 0x%04X and write mode is defined twice", address)
			}
			translation.Write = constantInfo.Constant
		}

		destination[address] = translation
	}
	return nil
}
