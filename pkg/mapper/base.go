package mapper

import (
	"fmt"

	"github.com/retroenv/nesgo/pkg/addressing"
	"github.com/retroenv/nesgo/pkg/bus"
)

const (
	chrMemSize           = 0x2000 // 8K
	defaultChrWindowSize = 0x2000 // 8K
	defaultPrgWindowSize = 0x4000 // 16K
	prgMemSize           = 0x8000 // 32K
)

// bankMapper maps an address to a bank number and offset into that bank.
type bankMapper func(address uint16) (int, uint16)

// Base provides common functionality for most mappers.
type Base struct {
	bus  *bus.Bus
	name string // optional

	chrRAM []byte

	nameTableCount int
	nameTableBanks []bank

	chrWindowSize int
	prgWindowSize int

	chrBanks []bank
	prgBanks []bank

	chrWindows []int
	prgWindows []int

	chrBankMapper bankMapper
	prgBankMapper bankMapper

	readHooks  []readHook
	writeHooks []writeHook
}

// NewBase creates a new mapper base.
func NewBase(bus *bus.Bus) *Base {
	return &Base{
		bus: bus,

		chrWindowSize: defaultChrWindowSize,
		prgWindowSize: defaultPrgWindowSize,

		nameTableCount: 1,
	}
}

// Name returns the name of the mapper.
func (b *Base) Name() string {
	return b.name
}

// SetName sets the name of the mapper.
func (b *Base) SetName(name string) {
	b.name = name
}

// Read a byte from a CHR or PRG memory address.
func (b *Base) Read(address uint16) uint8 {
	var value byte

	for _, hook := range b.readHooks {
		if address >= hook.startAddress && address <= hook.endAddress {
			value = hook.hook(address)
			return value
		}
	}

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

	if len(b.chrRAM) > 0 && address < 0x2000 {
		bankNr, offset := b.chrBankMapper(address)
		bank := &b.chrBanks[bankNr]
		bank.data[offset] = value
		return
	}

	panic(fmt.Sprintf("invalid write to address #%0000x", address))
}

// Initialize the mapper base with default settings.
func (b *Base) Initialize() {
	b.chrBankMapper = b.defaultChrBankMapper
	b.prgBankMapper = b.defaultPrgBankMapper

	b.setDefaultBankSizes()
	b.setBanks()
	b.setWindows()
}

func (b *Base) defaultChrBankMapper(address uint16) (int, uint16) {
	offset := address % uint16(b.chrWindowSize)
	windowNr := address / uint16(b.chrWindowSize)
	bankNr := b.chrWindows[windowNr]
	return bankNr, offset
}

func (b *Base) defaultPrgBankMapper(address uint16) (int, uint16) {
	address -= addressing.CodeBaseAddress
	offset := address % uint16(b.prgWindowSize)
	windowNr := address / uint16(b.prgWindowSize)
	bankNr := b.prgWindows[windowNr]
	return bankNr, offset
}
