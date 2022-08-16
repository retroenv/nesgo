//go:build !nesgo

package debugger

import (
	"encoding/json"
	"fmt"
	"strings"
)

// hexArray implements a byte array alias that JSON marshalls to a hex array output.
type hexArray []byte

func (a hexArray) MarshalJSON() ([]byte, error) {
	parts := make([]string, len(a))
	for i, b := range a {
		parts[i] = fmt.Sprintf("%02X", b)
	}

	return json.Marshal(parts)
}

// hexArrayCombined implements a byte array alias that JSON marshalls to a hex string output.
type hexArrayCombined []byte

func (a hexArrayCombined) MarshalJSON() ([]byte, error) {
	buf := strings.Builder{}

	for _, b := range a {
		s := fmt.Sprintf("%02X", b)
		buf.WriteString(s)
	}

	return json.Marshal(buf.String())
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
