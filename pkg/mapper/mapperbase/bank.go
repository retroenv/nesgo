package mapperbase

type bank struct {
	data   []byte
	length int
}

// setDefaultBankSizes sets the default CHR and PRG sizes based on the set window size.
func (b *Base) setDefaultBankSizes() {
	b.setDefaultChrBankSizes()
	b.setDefaultPrgBankSizes()
}

// setBanks sets the bank data based on each bank's length. This needs to be called after the bank lengths
// have been set.
func (b *Base) setBanks() {
	b.setChrBanks()
	b.setPrgBanks()
	b.createNameTableBanks()
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
