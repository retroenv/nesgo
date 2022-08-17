package mapperbase

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

// PrgBankCount returns the amount of PRG banks.
func (b *Base) PrgBankCount() int {
	return len(b.prgBanks)
}

// SetPrgWindow sets a PRG window to a specific bank.
func (b *Base) SetPrgWindow(window, bank int) {
	if bank < 0 {
		bank = len(b.prgBanks) + bank
	}
	bank %= len(b.prgBanks)

	b.mu.Lock()
	b.prgWindows[window] = bank
	b.mu.Unlock()
}

// SetPrgWindowSize sets the PRG window size.
func (b *Base) SetPrgWindowSize(size int) {
	b.prgWindowSize = size
}

// SetPrgRAM enables the usage of PRG RAM and sets the RAM buffer.
func (b *Base) SetPrgRAM(ram []byte) {
	b.prgRAM = ram
}
