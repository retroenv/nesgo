// Package codedatalog implements support for Code/Data Logging in a FCEUX/Mesen emulator compatible format.
package codedatalog

import (
	"errors"
	"io"

	"github.com/retroenv/retrogolib/nes/cartridge"
)

// PrgFlag defines flags for the PRG ROM.
type PrgFlag byte

// ChrFlag defines flags for the CHR ROM.
type ChrFlag byte

const (
	Code              PrgFlag = 0b0000_0001
	Data              PrgFlag = 0b0000_0010
	RomBankMappedLow  PrgFlag = 0b0000_0100
	RomBankMappedHigh PrgFlag = 0b0000_1000
	IndirectCode      PrgFlag = 0b0001_0000
	IndirectData      PrgFlag = 0b0010_0000
	PCMAudio          PrgFlag = 0b0100_0000
	SubEntryPoint     PrgFlag = 0b1000_0000

	DrawnOnScreen        ChrFlag = 0b0000_0001
	ReadProgrammatically ChrFlag = 0b0000_0010
)

// LoadFile loads an .cdl file in FCEUX/Mesen Code/Data Logger format.
func LoadFile(cart *cartridge.Cartridge, reader io.Reader) ([]PrgFlag, error) {
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	if len(data) < len(cart.PRG) {
		return nil, errors.New("invalid cdl file size for PRG ROM")
	}

	prgFlags := make([]PrgFlag, len(cart.PRG))
	for i := 0; i < len(cart.PRG); i++ {
		prgFlags[i] = PrgFlag(data[i])
	}

	return prgFlags, nil
}
