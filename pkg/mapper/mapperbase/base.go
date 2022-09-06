// Package mapperbase provides common functionality for most mappers.
package mapperbase

import (
	"fmt"
	"sync"

	"github.com/retroenv/nesgo/pkg/addressing"
	"github.com/retroenv/nesgo/pkg/bus"
	"github.com/retroenv/retrogolib/nes/cartridge"
)

const (
	chrMemSize           = 0x2000 // 8K
	defaultChrWindowSize = 0x2000 // 8K
	defaultPrgWindowSize = 0x4000 // 16K
	prgMemSize           = 0x8000 // 32K

	prgRAMStart = 0x6000
	prgRAMEnd   = 0x7FFF
)

// bankMapper maps an address to a bank number and offset into that bank.
type bankMapper func(address uint16) (int, uint16)

// Base provides common functionality for most mappers.
type Base struct {
	mu   sync.RWMutex
	bus  *bus.Bus
	name string // optional

	chrRAM []byte
	prgRAM []byte

	mirrorModeTranslation MirrorModeTranslation
	nameTableCount        int
	nameTableBanks        []bank

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

// New creates a new mapper base.
func New(bus *bus.Bus) *Base {
	return &Base{
		bus: bus,

		chrWindowSize: defaultChrWindowSize,
		prgWindowSize: defaultPrgWindowSize,

		nameTableCount: 1,
	}
}

// State returns the current state of the mapper.
func (b *Base) State() bus.MapperState {
	b.mu.RLock()
	defer b.mu.RUnlock()

	state := bus.MapperState{
		ID:         b.bus.Cartridge.Mapper,
		Name:       b.name,
		ChrWindows: b.chrWindows,
		PrgWindows: b.prgWindows,
	}

	return state
}

// SetName sets the name of the mapper.
func (b *Base) SetName(name string) {
	b.mu.Lock()
	b.name = name
	b.mu.Unlock()
}

// Read a byte from a CHR or PRG memory address.
func (b *Base) Read(address uint16) uint8 {
	var value byte

	for _, hook := range b.readHooks {
		if address >= hook.startAddress && address <= hook.endAddress {
			value = hook.hookFunc(address)
			if !hook.onlyProxy {
				return value
			}
		}
	}

	switch {
	case address < 0x2000:
		bankNr, offset := b.chrBankMapper(address)
		b.mu.RLock()
		bank := &b.chrBanks[bankNr]
		b.mu.RUnlock()
		value = bank.data[offset]

	case address >= prgRAMStart && address <= prgRAMEnd && len(b.prgRAM) > 0:
		offset := address - prgRAMStart
		value = b.prgRAM[offset]

	case address >= addressing.CodeBaseAddress:
		bankNr, offset := b.prgBankMapper(address)
		b.mu.RLock()
		bank := &b.prgBanks[bankNr]
		b.mu.RUnlock()
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
			hook.hookFunc(address, value)
			if !hook.onlyProxy {
				return
			}
		}
	}

	switch {
	case address < 0x2000 && len(b.chrRAM) > 0:
		bankNr, offset := b.chrBankMapper(address)
		bank := &b.chrBanks[bankNr]
		bank.data[offset] = value

	case address >= prgRAMStart && address <= prgRAMEnd && len(b.prgRAM) > 0:
		offset := address - prgRAMStart
		b.prgRAM[offset] = value

	default:
		panic(fmt.Sprintf("invalid write to address #%0000x", address))
	}
}

// Initialize the mapper base with default settings.
func (b *Base) Initialize() {
	b.chrBankMapper = b.defaultChrBankMapper
	b.prgBankMapper = b.defaultPrgBankMapper

	b.setDefaultBankSizes()
	b.setBanks()
	b.setWindows()
}

// Cartridge returns the current cartridge.
func (b *Base) Cartridge() *cartridge.Cartridge {
	return b.bus.Cartridge
}

func (b *Base) defaultChrBankMapper(address uint16) (int, uint16) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	offset := address % uint16(b.chrWindowSize)
	windowNr := address / uint16(b.chrWindowSize)
	bankNr := b.chrWindows[windowNr]
	return bankNr, offset
}

func (b *Base) defaultPrgBankMapper(address uint16) (int, uint16) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	address -= addressing.CodeBaseAddress
	offset := address % uint16(b.prgWindowSize)
	windowNr := address / uint16(b.prgWindowSize)
	bankNr := b.prgWindows[windowNr]
	return bankNr, offset
}
