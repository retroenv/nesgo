//go:build !nesgo

package debugger

import (
	"encoding/json"
	"fmt"
	"strings"
)

// hexArray implements a byte array alias that JSON marshals to a hex array.
type hexArray []byte

func (h hexArray) MarshalJSON() ([]byte, error) {
	parts := make([]string, len(h))
	for i, b := range h {
		parts[i] = fmt.Sprintf("%02X", b)
	}

	return json.Marshal(parts)
}

// hexArrayCombined implements a byte array alias that JSON marshals to a hex string.
type hexArrayCombined []byte

func (h hexArrayCombined) MarshalJSON() ([]byte, error) {
	buf := strings.Builder{}

	for _, b := range h {
		s := fmt.Sprintf("%02X", b)
		buf.WriteString(s)
	}

	return json.Marshal(buf.String())
}

// hexByte implements byte alias that JSON marshals to a hex string.
type hexByte uint8

func (h hexByte) MarshalJSON() ([]byte, error) {
	s := fmt.Sprintf("%02X", h)
	return json.Marshal(s)
}

// hexWord implements word alias that JSON marshals to a hex string.
type hexWord uint16

func (h hexWord) MarshalJSON() ([]byte, error) {
	s := fmt.Sprintf("%04X", h)
	return json.Marshal(s)
}

// hexDword implements qword alias that JSON marshals to a hex string.
type hexQword uint64

func (h hexQword) MarshalJSON() ([]byte, error) {
	s := fmt.Sprintf("%08X", h)
	return json.Marshal(s)
}

// nolint: unparam
func bytesToSliceArrayCombined(data []byte, rows, width int) []hexArrayCombined {
	var result []hexArrayCombined

	for row := 0; row < rows; row++ {
		offset := row * width
		result = append(result, data[offset:offset+width])
	}

	return result
}
