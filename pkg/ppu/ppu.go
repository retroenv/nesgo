//go:build !nesgo
// +build !nesgo

// Package ppu provides PPU (Picture Processing Unit) functionality.
package ppu

import (
	"fmt"
	"image"

	"github.com/retroenv/nesgo/pkg/bus"
	"github.com/retroenv/nesgo/pkg/memory"
)

const (
	Width  = 256
	Height = 240
	FPS    = 60
)

// PPU represents the Picture Processing Unit.
type PPU struct {
	bus *bus.Bus
	ram ram

	control control
	mask    mask
	status  status

	addressLatch bool
	vramAddress  register
	tempAddress  register

	fineX uint16

	oamData    [256]byte
	oamAddress byte

	image *image.RGBA
}

// New returns a new PPU.
func New(bus *bus.Bus) *PPU {
	p := &PPU{
		bus: bus,
		ram: memory.NewRAM(0x2000),
	}
	p.reset()
	return p
}

func (p *PPU) reset() {
	p.vramAddress = register{}
	p.tempAddress = register{}
	p.addressLatch = false
	p.fineX = 0

	p.image = image.NewRGBA(image.Rect(0, 0, Width, Height))
	p.oamData = [256]byte{}
	p.ram.Reset()

	p.status = status{}
	p.setControl(0x00)
	p.setMask(0x00)
	p.setOamAddress(0x00)
}

// ReadMemory reads from a PPU memory address.
func (p *PPU) ReadMemory(address uint16) uint8 {
	if address < 0x2000 {
		return p.bus.Mapper.ReadMemory(address)
	}
	if address > 0x3FFF {
		panic(fmt.Sprintf("unhandled ppu read at address: 0x%04X", address))
	}

	base := mirroredAddressToBase(address)
	switch base {
	case PPU_CTRL:
		return p.control.value
	case PPU_MASK:
		return p.mask.value
	case PPU_STATUS:
		return p.getStatus()
	case OAM_DATA:
		return p.readOamData()
	case PPU_DATA:
		return p.readData()

	default:
		panic(fmt.Sprintf("unhandled ppu read at address: 0x%04X", address))
	}
}

// WriteMemory writes to a PPU memory address.
func (p *PPU) WriteMemory(address uint16, value uint8) {
	if address < 0x2000 {
		p.bus.Mapper.WriteMemory(address, value)
		return
	}
	if address > 0x3FFF {
		panic(fmt.Sprintf("unhandled ppu write at address: 0x%04X", address))
	}

	base := mirroredAddressToBase(address)
	switch base {
	case PPU_CTRL:
		p.setControl(value)
	case PPU_MASK:
		p.setMask(value)
	case OAM_ADDR:
		p.setOamAddress(value)
	case OAM_DATA:
		p.writeOamData(value)
	case PPU_SCROLL:
		p.setScroll(uint16(value))
	case PPU_ADDR:
		p.setAddress(uint16(value))
	case PPU_DATA:
		p.writeData(value)
	case OAM_DMA:
		p.writeOamDMA(value)

	default:
		panic(fmt.Sprintf("unhandled ppu write at address: 0x%04X", address))
	}
}

func (p *PPU) setVBlank() {
	status := p.ram.ReadMemory(PPU_STATUS)
	status |= 0x80
	p.ram.WriteMemory(PPU_STATUS, status)
	// TODO handle NMI
}

func (p *PPU) clearVBlank() {
	status := p.ram.ReadMemory(PPU_STATUS)
	status &= 0x7f
	p.ram.WriteMemory(PPU_STATUS, status)
}

func (p *PPU) readData() byte {
	address := p.vramAddress.address()
	address &= 0x3FFF // valid addresses are $0000-$3FFF; higher addresses will be mirrored down

	data := p.ram.ReadMemory(address)
	// TODO handle special case of reading during rendering
	p.vramAddress.increment(p.control.VRAMIncrement)
	return data
}

func (p *PPU) writeData(value byte) {
	address := p.vramAddress.address()
	address &= 0x3FFF // valid addresses are $0000-$3FFF; higher addresses will be mirrored down

	p.ram.WriteMemory(address, value)
	p.vramAddress.increment(p.control.VRAMIncrement)
}

func (p *PPU) setScroll(value uint16) {
	if p.addressLatch {
		p.tempAddress.FineY = value & 0x07
		p.tempAddress.CoarseY = value >> 3
	} else {
		p.fineX = value & 0x07
		p.tempAddress.CoarseX = value >> 3
	}

	p.addressLatch = !p.addressLatch
}

func (p *PPU) setAddress(value uint16) {
	if p.addressLatch {
		address := p.tempAddress.address() & 0xFF00
		address |= value
		p.tempAddress.set(address)
		p.vramAddress = p.tempAddress
	} else {
		address := p.tempAddress.address() & 0x00FF
		address |= value << 8
		p.tempAddress.set(address)
	}

	p.addressLatch = !p.addressLatch
}

// mirroredAddressToBase converts the mirrored addresses to the base address.
// PPU registers are mirrored in every 8 bytes from $2008 through $3FFF.
func mirroredAddressToBase(address uint16) uint16 {
	base := 0x2000 + address&0b00000111
	return base
}
