//go:build !nesgo
// +build !nesgo

// Package ppu provides PPU (Picture Processing Unit) functionality.
package ppu

import (
	"fmt"
	"image"
	"time"

	"github.com/retroenv/nesgo/pkg/mapper"
)

const (
	Width  = 256
	Height = 240
	FPS    = 60
)

type control struct {
	BaseNameTable          uint16
	VRAMIncrement          uint8 // 0: add 1, going across; 1: add 32, going down
	SpritePatternTable     uint16
	BackgroundPatternTable uint16
	SpriteSize             uint8 // 0: 8x8 pixels; 1: 8x16 pixels
	MasterSlave            uint8
	NmiOutput              bool
}

type mask struct {
	Grayscale            bool
	RenderBackgroundLeft bool
	RenderSpritesLeft    bool
	RenderBackground     bool
	RenderSprites        bool
	EnhanceRed           bool
	EnhanceGreen         bool
	EnhanceBlue          bool
}

type state struct {
	Control byte
	Mask    byte
}

type cpu interface {
	Cycles() uint64
	StallCycles(cycles uint16)
}

// PPU represents the Picture Processing Unit.
type PPU struct {
	cpu     cpu
	ram     ram
	mapper  mapper.Memory
	control control
	mask    mask
	state   state

	addressLatch bool
	vramAddress  register
	tempAddress  register

	fineX uint16

	oamData    [256]byte
	oamAddress byte

	image *image.RGBA
}

// New returns a new PPU.
func New(ram ram, mapper mapper.Memory, cpu cpu) *PPU {
	p := &PPU{
		ram:    ram,
		mapper: mapper,
		cpu:    cpu,
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

	p.setControl(0x00)
	p.setMask(0x00)
	p.setOamAddress(0x00)
}

// Image returns the rendered image to display.
func (p *PPU) Image() *image.RGBA {
	return p.image
}

// ReadMemory reads from a PPU memory address.
func (p *PPU) ReadMemory(address uint16) uint8 {
	switch address {
	case PPU_CTRL:
		return p.state.Control

	case PPU_MASK:
		return p.state.Mask

	case PPU_STATUS:
		p.addressLatch = false
		b := p.ram.ReadMemory(address)
		p.clearVBlank()
		return b

	default:
		if address < 0x2000 {
			return p.mapper.ReadMemory(address)
		}

		panic(fmt.Sprintf("unhandled ppu read at address: 0x%04X", address))
	}
}

// WriteMemory writes to a PPU memory address.
func (p *PPU) WriteMemory(address uint16, value uint8) {
	switch address {
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
		if address < 0x2000 {
			p.mapper.WriteMemory(address, value)
			return
		}

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

// StartRender starts the rendering process.
func (p *PPU) StartRender() {
	p.setVBlank()
	time.Sleep(time.Second / FPS)
}

// FinishRender finishes the rendering process.
func (p *PPU) FinishRender() {
	status := p.ram.ReadMemory(PPU_STATUS)
	status &= 0xbf
	p.ram.WriteMemory(PPU_STATUS, status)
	p.clearVBlank()
}

// RenderScreen renders the screen into the internal image.
func (p *PPU) RenderScreen() {
	if p.mask.RenderBackground {
		p.renderBackground()
	}
}

func (p *PPU) renderBackground() {
	idx := int(p.ram.ReadMemory(PALETTE_START))
	idx %= len(colors)
	c := colors[idx]

	for y := 0; y < Height; y++ {
		for x := 0; x < Width; x++ {
			p.image.SetRGBA(x, y, c)
		}
	}
}

func (p *PPU) setControl(value byte) {
	p.state.Control = value

	p.control.BaseNameTable = (uint16(value&CTRL_NT_2C00) << 10) + 0x2000

	increment := (value & CTRL_INC_32) >> 2
	if increment == 0 {
		p.control.VRAMIncrement = 1
	} else {
		p.control.VRAMIncrement = 32
	}

	p.control.SpritePatternTable = uint16(value&CTRL_SPR_1000) << 9
	p.control.BackgroundPatternTable = uint16(value&CTRL_BG_1000) << 8
	p.control.SpriteSize = value & CTRL_8x16 >> 5
	p.control.MasterSlave = value & CTRL_MASTERSLAVE >> 6
	p.control.NmiOutput = value&CTRL_NMI != 0

	p.tempAddress.NameTableX = uint16(value & CTRL_NT_2400)
	p.tempAddress.NameTableY = uint16(value&CTRL_NT_2800) >> 1
}

func (p *PPU) setMask(value byte) {
	p.state.Mask = value

	p.mask.Grayscale = value&MASK_MONO != 0
	p.mask.RenderBackgroundLeft = value&MASK_BG_CLIP != 0
	p.mask.RenderSpritesLeft = value&MASK_SPR_CLIP != 0
	p.mask.RenderBackground = value&MASK_BG != 0
	p.mask.RenderSprites = value&MASK_SPR != 0
	p.mask.EnhanceRed = value&MASK_TINT_RED != 0
	p.mask.EnhanceGreen = value&MASK_TINT_GREEN != 0
	p.mask.EnhanceBlue = value&MASK_TINT_BLUE != 0
}

func (p *PPU) setOamAddress(value byte) {
	p.oamAddress = value
}

func (p *PPU) writeOamData(value byte) {
	p.oamData[p.oamAddress] = value
	p.oamAddress++

	// TODO handle scroll glitch
}

func (p *PPU) writeOamDMA(value byte) {
	address := uint16(value) << 8

	for i := 0; i < 256; i++ {
		data := p.ram.ReadMemory(address)
		p.oamData[p.oamAddress] = data
		p.oamAddress++
		address++
	}

	// 1 wait state cycle while waiting for writes to complete,
	// +1 if on an odd CPU cycle, then 256 alternating read/write cycles
	stall := uint16(1 + 256 + 256)
	if p.cpu.Cycles()%2 == 1 {
		stall++
	}
	p.cpu.StallCycles(stall)
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
