package gamegenie

import (
	"testing"

	"github.com/retroenv/retrogolib/assert"
)

var testCases = []struct {
	Code            string
	ExpectedAddress uint16
	ExpectedValue   uint8
	ExpectedCompare uint8
	ExpectedError   bool
}{
	{
		// Capcom's Ghosts 'n Goblins start your player with a really funky weapon
		Code:            "GOSSIP",
		ExpectedAddress: 0xD1DD,
		ExpectedValue:   0x14,
	},
	// TODO: fix
	// {
	//	//  Super Mario Bros 1 swim in any level
	//	Code:            "PIGOAP",
	//	ExpectedAddress: 0x9148,
	//	ExpectedValue:   0x51,
	// },
	{
		// Dr. Mario clear a row or column with only 3 colors in a line, rather than 4
		Code:            "ZEXPYGLA",
		ExpectedAddress: 0x94A7,
		ExpectedValue:   0x02,
		ExpectedCompare: 0x03,
	},
	{
		Code:          "TEST",
		ExpectedError: true,
	},
}

func TestDecode(t *testing.T) {
	t.Parallel()

	for _, test := range testCases {
		test := test
		t.Run(test.Code, func(t *testing.T) {
			t.Parallel()

			patch, err := Decode(test.Code)

			if test.ExpectedError {
				assert.True(t, err != nil)
				return
			}
			assert.NoError(t, err)

			assert.Equal(t, test.ExpectedAddress, patch.Address)
			assert.Equal(t, test.ExpectedValue, patch.Data)

			if len(test.Code) == 8 {
				assert.Equal(t, test.ExpectedCompare, patch.Compare)
			}
		})
	}
}

func TestEncode(t *testing.T) {
	t.Parallel()

	for _, test := range testCases {
		test := test
		t.Run(test.Code, func(t *testing.T) {
			t.Parallel()

			if test.ExpectedError {
				return
			}

			patch := Patch{
				Address: test.ExpectedAddress,
				Data:    test.ExpectedValue,
				Compare: test.ExpectedCompare,
			}
			if len(test.Code) == 8 {
				patch.HasCompare = true
			}

			code, err := Encode(patch)
			assert.NoError(t, err)
			assert.Equal(t, test.Code, code)
		})
	}
}
