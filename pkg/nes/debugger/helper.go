//go:build !nesgo

package debugger

import (
	"fmt"
	"strings"
)

func bytesToHex(data []byte) string {
	parts := make([]string, 0, len(data))
	for _, b := range data {
		parts = append(parts, fmt.Sprintf("%02X", b))
	}
	s := strings.Join(parts, ",")
	return s
}
