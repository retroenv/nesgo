// Package gamegenie implements NES Game Genie code and decode support.
package gamegenie

import (
	"fmt"
	"strings"

	"github.com/retroenv/retrogolib/nes/addressing"
)

// Patch defines a patch to apply to a NES ROM.
type Patch struct {
	Address    uint16
	Data       byte
	Compare    byte
	HasCompare bool
}

// Decode decodes a game genie code into a patch.
func Decode(code string) (Patch, error) {
	length := len(code)
	if length != 6 && length != 8 {
		return Patch{}, fmt.Errorf("invalid code length %d", length)
	}

	code = strings.ToUpper(code)
	n := make([]uint16, length)
	for i := range code {
		c := code[i]
		value, ok := translationCharToValue[c]
		if !ok {
			return Patch{}, fmt.Errorf("invalid character %v", c)
		}
		n[i] = uint16(value)
	}

	address := addressing.CodeBaseAddress + ((n[3] & 7) << 12) |
		((n[5] & 7) << 8) | ((n[4] & 8) << 8) |
		((n[2] & 7) << 4) | ((n[1] & 8) << 4) |
		(n[4] & 7) | (n[3] & 8)

	data := ((n[1] & 7) << 4) | ((n[0] & 8) << 4) | (n[0] & 7) | (n[length-1] & 8)

	patch := Patch{
		Address: address,
		Data:    byte(data),
	}

	if length == 8 {
		patch.HasCompare = true
		patch.Compare = byte(((n[7] & 7) << 4) | ((n[6] & 8) << 4) | (n[6] & 7) | (n[5] & 8))
	}

	return patch, nil
}

// Encode encodes a NES ROM patch into a game genie code.
func Encode(patch Patch) (string, error) {
	if patch.Address < addressing.CodeBaseAddress {
		return "", fmt.Errorf("address $%04X is not supported", patch.Address)
	}

	n := []uint16{
		uint16(((patch.Data >> 4) & 8) | (patch.Data & 7)),
		((patch.Address >> 4) & 8) | ((uint16(patch.Data) >> 4) & 7),
		8 | ((patch.Address >> 4) & 7),
		(patch.Address & 8) | ((patch.Address >> 12) & 7),
		((patch.Address >> 8) & 8) | (patch.Address & 7),
	}

	if patch.HasCompare {
		n = append(n,
			(uint16(patch.Compare)&8)|((patch.Address>>8)&7),
			uint16(((patch.Compare>>4)&8)|(patch.Compare&7)),
			uint16((patch.Data&8)|((patch.Compare>>4)&7)),
		)
	} else {
		n = append(n, (uint16(patch.Data)&8)|((patch.Address>>8)&7))
	}

	buf := strings.Builder{}
	for _, value := range n {
		character := translationValueToChar[value]
		buf.WriteByte(character)
	}
	return buf.String(), nil
}
