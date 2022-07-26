// Package cartridge provides .nes ROM loading and saving.
package cartridge

import (
	"encoding/binary"
	"fmt"
	"io"
)

// Cartridge contains a NES cartridge content.
type Cartridge struct {
	PRG     []byte // PRG-ROM banks
	CHR     []byte // CHR-ROM banks
	RAM     byte   // PRG-RAM banks
	Trainer []byte

	Mapper      byte       // mapper type
	Mirror      MirrorMode // mirroring mode
	Battery     byte       // battery present
	VideoFormat byte       // 0 NTSC, 1 PAL
}

// New returns a new cartridge.
func New() *Cartridge {
	return &Cartridge{
		PRG:     make([]byte, 0x8000),
		CHR:     make([]byte, 0x2000),
		Mapper:  0,
		Mirror:  MirrorVertical,
		Battery: 0,
	}
}

// Save the cartridge content in iNES format.
func (c *Cartridge) Save(writer io.Writer) error {
	header := header{
		Magic:       iNESFileMagic,
		NumPRG:      byte(len(c.PRG) / 16384),
		NumCHR:      byte(len(c.CHR) / 8192),
		NumRAM:      c.RAM,
		VideoFormat: c.VideoFormat,
	}

	header.Control1, header.Control2 = ControlBytes(c.Battery, byte(c.Mirror), c.Mapper, len(c.Trainer) > 0)

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
