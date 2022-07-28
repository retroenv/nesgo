package mapper

import "github.com/retroenv/nesgo/pkg/ppu/nametable"

type bank struct {
	data   []byte
	length int
}

// setDefaultBankSizes sets the default CHR and PRG sizes based on the set window size.
func (b *Base) setDefaultBankSizes() {
	b.setDefaultChrBankSizes()
	b.setDefaultPrgBankSizes()
}

// setDefaultPrgBankSizes creates the banks with default lengths set based on the PRG and PRG window size.
// Some mapper and modes can have a mix of sizes, for example One 16KB bank and two 8KB banks for MMC5.
// In case of mixed sizing the bank sizes need to be initialized manually.
func (b *Base) setDefaultPrgBankSizes() {
	prgSize := len(b.bus.Cartridge.PRG)
	banks := prgSize / b.prgWindowSize
	b.prgBanks = make([]bank, banks)

	for i := 0; i < banks; i++ {
		bank := &b.prgBanks[i]
		bank.length = b.prgWindowSize
	}
}

// setDefaultChrBankSizes creates the banks with default lengths set based on the CHR and CHR window size.
// Some mapper and modes can have a mix of sizes, for example two 2KB banks and four 1KB banks for MMC6.
// In case of mixed sizing the bank sizes need to be initialized manually.
func (b *Base) setDefaultChrBankSizes() {
	chrSize := len(b.bus.Cartridge.CHR)
	if chrSize == 0 {
		chrSize = len(b.chrRAM)
	}

	banks := chrSize / b.chrWindowSize
	b.chrBanks = make([]bank, banks)

	for i := 0; i < banks; i++ {
		bank := &b.chrBanks[i]
		bank.length = b.chrWindowSize
	}
}

// setBanks sets the bank data based on each bank's length. This needs to be called after the bank lengths
// have been set.
func (b *Base) setBanks() {
	b.setChrBanks()
	b.setPrgBanks()
	b.createNameTableBanks()
}

// setPrgBanks sets the bank data based on each bank's length. This needs to be called after the bank lengths
// have been set, either by calling setDefaultPrgBankSizes() or setting it manually.
func (b *Base) setPrgBanks() {
	prg := b.bus.Cartridge.PRG
	startOffset := 0

	for i := 0; i < len(b.prgBanks); i++ {
		bank := &b.prgBanks[i]
		endOffset := startOffset + bank.length
		bank.data = prg[startOffset:endOffset]
		startOffset += bank.length
	}
}

// setChrBanks sets the bank data based on each bank's length. This needs to be called after the bank lengths
// have been set, either by calling setDefaultChrBankSizes() or setting it manually.
func (b *Base) setChrBanks() {
	chr := b.bus.Cartridge.CHR
	if len(chr) == 0 {
		chr = b.chrRAM
	}

	startOffset := 0

	for i := 0; i < len(b.chrBanks); i++ {
		bank := &b.chrBanks[i]
		endOffset := startOffset + bank.length
		bank.data = chr[startOffset:endOffset]
		startOffset += bank.length
	}
}

// createNameTableBanks creates the VRAM banks.
func (b *Base) createNameTableBanks() {
	b.nameTableBanks = make([]bank, b.nameTableCount)

	for i := 0; i < b.nameTableCount; i++ {
		bank := &b.nameTableBanks[i]
		bank.length = nametable.VramSize
		bank.data = make([]byte, bank.length)
	}

	b.SetNameTableWindow(0)
}

// setWindows sets the CHR and PRG windows to banks based on a static window size.
func (b *Base) setWindows() {
	windows := chrMemSize / b.chrWindowSize
	b.chrWindows = make([]int, windows)
	bank := 0
	for i := 0; i < windows; i++ {
		b.chrWindows[i] = bank

		if bank+1 < len(b.chrBanks) {
			bank++
		} else {
			bank = 0
		}
	}

	windows = prgMemSize / b.prgWindowSize
	b.prgWindows = make([]int, windows)
	bank = 0
	for i := 0; i < windows; i++ {
		b.prgWindows[i] = bank

		if bank+1 < len(b.prgBanks) {
			bank++
		} else {
			bank = 0
		}
	}
}

// ChrBankCount returns the amount of CHR banks.
func (b *Base) ChrBankCount() int {
	return len(b.chrBanks)
}

// PrgBankCount returns the amount of PRG banks.
func (b *Base) PrgBankCount() int {
	return len(b.prgBanks)
}

// SetChrWindow sets a CHR window to a specific bank.
func (b *Base) SetChrWindow(window, bank int) {
	b.chrWindows[window] = bank
}

// SetPrgWindow sets a PRG window to a specific bank.
func (b *Base) SetPrgWindow(window, bank int) {
	b.prgWindows[window] = bank
}

// SetChrWindowSize sets the CHR window size.
func (b *Base) SetChrWindowSize(size int) {
	b.chrWindowSize = size
}

// SetPrgWindowSize sets the PRG window size.
func (b *Base) SetPrgWindowSize(size int) {
	b.prgWindowSize = size
}

// SetChrRAM enables the usage of CHR RAM and sets the RAM buffer.
func (b *Base) SetChrRAM(ram []byte) {
	b.chrRAM = ram
}

// SetNameTableCount sets amount of nametables.
func (b *Base) SetNameTableCount(count int) {
	b.nameTableCount = count
}

// SetNameTableWindow sets the nametable window to a specific bank.
func (b *Base) SetNameTableWindow(bank int) {
	bank %= len(b.nameTableBanks)
	nameTable := &b.nameTableBanks[bank]
	b.bus.NameTable.SetVRAM(nameTable.data)
}

// NameTable returns the nametable buffer of a specific bank. Used in tests.
func (b *Base) NameTable(bank int) []byte {
	bank %= len(b.nameTableBanks)
	nameTable := &b.nameTableBanks[bank]
	return nameTable.data
}
