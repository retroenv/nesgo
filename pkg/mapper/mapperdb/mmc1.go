package mapperdb

/*
Boards: SKROM, SLROM, SNROM, others
PRG ROM capacity: 256K (512K)
PRG ROM window: 16K + 16K fixed or 32K
PRG RAM capacity: 32K
PRG RAM window:8K
CHR capacity: 128K
CHR window: 4K + 4K or 8K
*/

import (
	"github.com/retroenv/nesgo/pkg/bus"
	"github.com/retroenv/nesgo/pkg/cartridge"
)

type mapperMMC1 struct {
	Base

	ram []byte

	shiftCount    byte
	shiftRegister byte

	control  byte
	chrBank0 int
	chrBank1 int
	prgBank  int
}

// NewMMC1 returns a new mapper instance.
func NewMMC1(base Base) bus.Mapper {
	m := &mapperMMC1{
		Base: base,
		ram:  make([]byte, 0x8000), // 32K
	}
	m.SetName("MMC1")
	m.SetChrWindowSize(0x1000) // 4K
	m.SetPrgRAM(m.ram)
	m.Initialize()

	m.AddWriteHook(0x8000, 0x7FFF, m.writeShiftBit)

	m.SetPrgWindow(1, -1)

	// TODO support mmc1 variants

	return m
}

func (m *mapperMMC1) resetShift() {
	m.shiftCount = 0
	m.shiftRegister = 0
	m.applyControl()
	m.control = 0x0C
}

func (m *mapperMMC1) writeShiftBit(address uint16, value uint8) {
	if value&0x80 != 0 {
		m.resetShift()
		return
	}

	// the shift register gets written from lowest to highest bit
	bit := (value & 1) << m.shiftCount
	m.shiftRegister |= bit
	m.shiftCount++

	if m.shiftCount < 5 {
		return
	}

	switch {
	case address < 0xA000: // $8000-$9FFF
		m.control = m.shiftRegister

	case address < 0xC000: // $A000-$BFFF
		m.chrBank0 = int(m.shiftRegister)

	case address < 0xE000: // $C000-$DFFF
		m.chrBank1 = int(m.shiftRegister)

	case address >= 0xE000: // $E000-$FFFF
		m.prgBank = int(m.shiftRegister) & 0b00001111
	}

	m.resetShift()
}

func (m *mapperMMC1) applyControl() {
	mirror := m.control & 0b00000011
	switch mirror {
	case 0:
		m.SetNameTableMirrorMode(cartridge.MirrorSingle0)
	case 1:
		m.SetNameTableMirrorMode(cartridge.MirrorSingle1)
	case 2:
		m.SetNameTableMirrorMode(cartridge.MirrorVertical)
	case 3:
		m.SetNameTableMirrorMode(cartridge.MirrorHorizontal)
	}

	prgMode := (m.control >> 2) & 0b00000011
	switch prgMode {
	case 0, 1:
		// switch 32 KB at $8000, ignoring low bit of bank number
		// low bit ignored in 32 KB mode but since window size is 16KB the second part can be set to bank + 1
		m.SetPrgWindow(0, m.prgBank)
		m.SetPrgWindow(1, m.prgBank+1)

	case 2:
		// fix first bank at $8000 and switch 16 KB bank at $C000
		m.SetPrgWindow(0, 0)
		m.SetPrgWindow(1, m.prgBank)

	case 3:
		// fix last bank at $C000 and switch 16 KB bank at $8000
		m.SetPrgWindow(0, m.prgBank)
		m.SetPrgWindow(1, -1)
	}

	m.SetChrWindow(0, m.chrBank0)
	chrMode := (m.control >> 4) & 1
	if chrMode == 0 {
		// switch 8 KB at a time
		m.SetChrWindow(1, m.chrBank0+1)
	} else {
		// switch two separate 4 KB banks
		m.SetChrWindow(1, m.chrBank1)
	}
}
