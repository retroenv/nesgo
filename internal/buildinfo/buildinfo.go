// Package buildinfo formats build information that is embedded into the binaries.
package buildinfo

import (
	"fmt"
	"runtime"
	"strings"
)

// BuildVersion builds a version string based on binary release information.
func BuildVersion(version, commit, date string) string {
	buf := strings.Builder{}
	buf.WriteString(version)

	if commit != "" {
		buf.WriteString(fmt.Sprintf(" commit: %s", commit))
	}
	if date != "" {
		buf.WriteString(fmt.Sprintf(" built at: %s", date))
	}
	goVersion := runtime.Version()
	buf.WriteString(fmt.Sprintf(" built with: %s", goVersion))
	return buf.String()
}
