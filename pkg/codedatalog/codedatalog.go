// Package codedatalog implements support for Code/Data Logging in a FCEUX/Mesen emulator compatible format.
package codedatalog

import (
	"errors"
	"io"

	"github.com/retroenv/nesgo/pkg/cartridge"
)

// PrgFlag defines flags for the PRG ROM.
type PrgFlag byte

// ChrFlag defines flags for the CHR ROM.
type ChrFlag byte

const (
	Code              PrgFlag = 0b00000001
	Data              PrgFlag = 0b00000010
	RomBankMappedLow  PrgFlag = 0b00000100
	RomBankMappedHigh PrgFlag = 0b00001000
	IndirectCode      PrgFlag = 0b00010000
	IndirectData      PrgFlag = 0b00100000
	PCMAudio          PrgFlag = 0b01000000
	SubEntryPoint     PrgFlag = 0b10000000

	DrawnOnScreen        ChrFlag = 0b00000001
	ReadProgrammatically ChrFlag = 0b00000010
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
