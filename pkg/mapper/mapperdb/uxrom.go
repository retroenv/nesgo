package mapperdb

import "github.com/retroenv/nesgo/pkg/bus"

/*
Boards: UNROM, UOROM
PRG ROM capacity: 256K/4096K
PRG ROM window: 16K + 16K fixed
CHR capacity: 8K
*/

type mapperUxROM struct {
	Base

	valueShift  int
	windowIndex int
}

// NewUxROMOr returns a new mapper instance with OR logic (74HC32) configuration.
func NewUxROMOr(base Base) bus.Mapper {
	m := newMapperUxROM(base)
	m.SetName("UxROM")

	// $8000-$BFFF: 16 KB switchable PRG ROM bank
	// $C000-$FFFF: 16 KB PRG ROM bank, fixed to the last bank
	m.SetPrgWindow(1, -1) // $C000-$FFFF: 16 KB PRG ROM bank, fixed to the last bank
	return m
}

// NewUN1ROM returns a new mapper instance with OR logic (74HC32) configuration and a value shifter of 2.
func NewUN1ROM(base Base) bus.Mapper {
	m := newMapperUxROM(base)
	m.SetName("UN1ROM")

	// $8000-$BFFF: 16 KB switchable PRG ROM bank
	// $C000-$FFFF: 16 KB PRG ROM bank, fixed to the last bank
	m.SetPrgWindow(1, -1) // $C000-$FFFF: 16 KB PRG ROM bank, fixed to the last bank
	m.valueShift = 2      // very similar to UxROM, but the register is shifted by two bits
	return m
}

// NewUxROMAnd returns a new mapper instance with AND logic (74HC08) configuration.
func NewUxROMAnd(base Base) bus.Mapper {
	m := newMapperUxROM(base)
	m.SetName("UxROM")

	// $8000-$BFFF: 16 KB PRG ROM bank, fixed to the first bank
	// $C000-$FFFF: 16 KB switchable PRG ROM bank
	m.windowIndex = 1
	return m
}

func newMapperUxROM(base Base) *mapperUxROM {
	m := &mapperUxROM{
		Base: base,
	}
	m.Initialize()

	m.AddWriteHook(0x8000, 0xFFFF, m.setPrgWindow)
	return m
}

func (m *mapperUxROM) setPrgWindow(address uint16, value uint8) {
	value >>= m.valueShift
	value &= 0b00000111                       // UNROM uses bits 2-0; UOROM/UN1ROM uses bits 3-0
	m.SetPrgWindow(m.windowIndex, int(value)) // select 16 KB PRG ROM bank
}
