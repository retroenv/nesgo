package mapper

import (
	"fmt"

	"github.com/retroenv/nesgo/pkg/addressing"
	"github.com/retroenv/nesgo/pkg/bus"
)

const (
	chrMemSize       = 0x2000 // 8K
	defaultChrWindow = 0x2000 // 8K
	defaultPrgWindow = 0x4000 // 16K
	prgMemSize       = 0x8000 // 32K
)

// bankMapper maps an address to a bank number and offset into that bank.
type bankMapper func(address uint16) (int, uint16)

// Base provides common functionality for most mappers.
type Base struct {
	bus  *bus.Bus
	name string // optional

	chrWindow int
	prgWindow int

	chrBanks []bank
	prgBanks []bank

	chrWindows []int
	prgWindows []int

	chrBankMapper bankMapper
	prgBankMapper bankMapper

	writeHooks []writeHook
}

// newBase creates a new mapper base.
func newBase(bus *bus.Bus) *Base {
	return &Base{
		bus: bus,

		chrWindow: defaultChrWindow,
		prgWindow: defaultPrgWindow,
	}
}

// Name returns the name of the mapper.
func (b *Base) Name() string {
	return b.name
}

// Read a byte from a CHR or PRG memory address.
func (b *Base) Read(address uint16) uint8 {
	var value byte

	switch {
	case address < 0x2000:
		bankNr, offset := b.chrBankMapper(address)
		bank := &b.chrBanks[bankNr]
		value = bank.data[offset]

	case address >= addressing.CodeBaseAddress:
		bankNr, offset := b.prgBankMapper(address)
		bank := &b.prgBanks[bankNr]
		value = bank.data[offset]

	default:
		panic(fmt.Sprintf("invalid read from address #%0000x", address))
	}
	return value
}

// Write a byte to a CHR or PRG memory address.
func (b *Base) Write(address uint16, value uint8) {
	for _, hook := range b.writeHooks {
		if address >= hook.startAddress && address <= hook.endAddress {
			hook.hook(address, value)
			return
		}
	}

	panic(fmt.Sprintf("invalid write to address #%0000x", address))
}

// initialize the mapper base with default settings.
func (b *Base) initialize() {
	b.chrBankMapper = b.defaultChrBankMapper
	b.prgBankMapper = b.defaultPrgBankMapper

	b.setDefaultBankSizes()
	b.setBanks()
	b.setWindows()
}

func (b *Base) defaultChrBankMapper(address uint16) (int, uint16) {
	offset := address % uint16(b.chrWindow)
	windowNr := address / uint16(b.chrWindow)
	bankNr := b.chrWindows[windowNr]
	return bankNr, offset
}

func (b *Base) defaultPrgBankMapper(address uint16) (int, uint16) {
	address -= addressing.CodeBaseAddress
	offset := address % uint16(b.prgWindow)
	windowNr := address / uint16(b.prgWindow)
	bankNr := b.prgWindows[windowNr]
	return bankNr, offset
}
