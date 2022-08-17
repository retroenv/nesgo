package mapperbase

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

// ChrBankCount returns the amount of CHR banks.
func (b *Base) ChrBankCount() int {
	return len(b.chrBanks)
}

// SetChrWindow sets a CHR window to a specific bank.
func (b *Base) SetChrWindow(window, bank int) {
	if bank < 0 {
		bank = len(b.chrBanks) + bank
	}
	bank %= len(b.chrBanks)

	b.mu.Lock()
	b.chrWindows[window] = bank
	b.mu.Unlock()
}

// SetChrWindowSize sets the CHR window size.
func (b *Base) SetChrWindowSize(size int) {
	b.chrWindowSize = size
}

// SetChrRAM enables the usage of CHR RAM and sets the RAM buffer.
func (b *Base) SetChrRAM(ram []byte) {
	b.chrRAM = ram
}
