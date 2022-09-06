package compiler

import (
	"testing"

	"github.com/retroenv/retrogolib/assert"
)

var fileIgnoredTestCases = []struct {
	name           string
	input          string
	expectedResult bool
	expectedError  string
}{
	{
		"old single tag",
		`// +build !nesgo`,
		true,
		"",
	},
	{
		"old multiple tags",
		`// +build !nesgo,!nogui`,
		true,
		"",
	},
	{
		"new single tag",
		`//go:build !nesgo`,
		true,
		"",
	},
	{
		"new multiple tags",
		`//go:build !nesgo && !nogui`,
		true,
		"",
	},
	{
		"no build tag match",
		`//go:build macos`,
		false,
		"",
	},
	{
		"no build header",
		`package nes`,
		false,
		"",
	},
}

func TestIsFileIgnored(t *testing.T) {
	for _, test := range fileIgnoredTestCases {
		result, err := isFileIgnored([]byte(test.input + "\n"))
		if test.expectedError == "" {
			assert.NoError(t, err, test.name)
		} else {
			assert.Error(t, err, test.expectedError, test.name)
			return
		}
		assert.Equal(t, test.expectedResult, result)
	}
}
