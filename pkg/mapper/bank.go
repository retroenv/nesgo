package mapper

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
	banks := prgSize / b.prgWindow
	b.prgBanks = make([]bank, banks)

	for i := 0; i < banks; i++ {
		bank := &b.prgBanks[i]
		bank.length = b.prgWindow
	}
}

// setDefaultChrBankSizes creates the banks with default lengths set based on the CHR and CHR window size.
// Some mapper and modes can have a mix of sizes, for example two 2KB banks and four 1KB banks for MMC6.
// In case of mixed sizing the bank sizes need to be initialized manually.
func (b *Base) setDefaultChrBankSizes() {
	chrSize := len(b.bus.Cartridge.CHR)
	banks := chrSize / b.chrWindow
	b.chrBanks = make([]bank, banks)

	for i := 0; i < banks; i++ {
		bank := &b.chrBanks[i]
		bank.length = b.chrWindow
	}
}

// setBanks sets the bank data based on each bank's length. This needs to be called after the bank lengths
// have been set.
func (b *Base) setBanks() {
	b.setChrBanks()
	b.setPrgBanks()
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
	startOffset := 0

	for i := 0; i < len(b.chrBanks); i++ {
		bank := &b.chrBanks[i]
		endOffset := startOffset + bank.length
		bank.data = chr[startOffset:endOffset]
		startOffset += bank.length
	}
}

// setWindows sets the CHR and PRG windows to banks based on a static window size.
func (b *Base) setWindows() {
	windows := chrMemSize / b.chrWindow
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

	windows = prgMemSize / b.prgWindow
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
