package ines

import (
	"encoding/binary"
	"fmt"
	"io"
)

// Cartridge contains a NES cartridge content.
type Cartridge struct {
	PRG     []byte // PRG-ROM banks
	CHR     []byte // CHR-ROM banks
	Trainer []byte
	Mapper  byte // mapper type
	Mirror  byte // mirroring mode
	Battery byte // battery present
}

// Save the cartridge content in iNES format.
func (c *Cartridge) Save(writer io.Writer) error {
	header := header{
		Magic:  iNESFileMagic,
		NumPRG: byte(len(c.PRG) / 16384),
		NumCHR: byte(len(c.CHR) / 8192),
	}

	header.Control1 |= (c.Battery & 1) << 1

	header.Control1 |= c.Mirror & 1
	header.Control1 |= ((c.Mirror >> 1) & 1) << 3

	header.Control1 |= mergeNibbles(c.Mapper, header.Control1)
	header.Control2 |= mergeNibbles(highNibble(c.Mapper), header.Control2)

	if len(c.Trainer) > 0 {
		header.Control1 |= trainerFlag
	}

	if err := binary.Write(writer, binary.LittleEndian, header); err != nil {
		return fmt.Errorf("writing header: %w", err)
	}

	if len(c.Trainer) > 0 {
		if err := binary.Write(writer, binary.LittleEndian, c.Trainer); err != nil {
			return fmt.Errorf("writing trainer: %w", err)
		}
	}

	if err := binary.Write(writer, binary.LittleEndian, c.PRG); err != nil {
		return fmt.Errorf("writing PRG: %w", err)
	}

	if len(c.CHR) > 0 {
		if err := binary.Write(writer, binary.LittleEndian, c.CHR); err != nil {
			return fmt.Errorf("writing CHR: %w", err)
		}
	}

	return nil
}
