package cartridge

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

var iNESFileMagic = [4]byte{'N', 'E', 'S', 0x1a}

const (
	trainerFlag = 1 << 3
)

type header struct {
	Magic       [4]byte // iNES magic number
	NumPRG      byte    // number of PRG-ROM banks (16KB each)
	NumCHR      byte    // number of CHR-ROM banks (8KB each)
	Control1    byte    // control bits
	Control2    byte    // control bits
	NumRAM      byte    // PRG-RAM size (x 8KB)
	VideoFormat byte    // 0 NTSC, 1 PAL
	Reserved    [6]byte // unused padding
}

// LoadFile loads an .nes file in iNES format.
func LoadFile(reader io.Reader) (*Cartridge, error) {
	var header header
	if err := binary.Read(reader, binary.LittleEndian, &header); err != nil {
		return nil, fmt.Errorf("reading header: %w", err)
	}

	if header.Magic != iNESFileMagic {
		return nil, errors.New("invalid file header magic")
	}

	mapper := mergeNibbles(highNibble(header.Control2), highNibble(header.Control1))

	mirror1 := header.Control1 & 1
	mirror2 := (header.Control1 >> 3) & 1
	mirror := mirror1 | mirror2<<1

	battery := (header.Control1 >> 1) & 1

	var trainer []byte
	if header.Control1&trainerFlag != 0 { // check if trainer is present
		trainer = make([]byte, 512)
		if _, err := io.ReadFull(reader, trainer); err != nil {
			return nil, fmt.Errorf("reading trainer: %w", err)
		}
	}

	prg := make([]byte, int(header.NumPRG)*16384)
	if _, err := io.ReadFull(reader, prg); err != nil {
		return nil, fmt.Errorf("reading PRG: %w", err)
	}

	var chr []byte
	if header.NumCHR > 0 {
		chr = make([]byte, int(header.NumCHR)*8192)
		if _, err := io.ReadFull(reader, chr); err != nil {
			return nil, fmt.Errorf("reading CHR: %w", err)
		}
	}

	return &Cartridge{
		PRG:         prg,
		CHR:         chr,
		RAM:         header.NumRAM,
		Trainer:     trainer,
		Mapper:      mapper,
		Mirror:      MirrorMode(mirror),
		Battery:     battery,
		VideoFormat: header.VideoFormat,
	}, nil
}
