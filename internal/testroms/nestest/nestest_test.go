package nestest

import (
	"bufio"
	"bytes"
	"os"
	"testing"

	"github.com/retroenv/nesgo/internal/assert"
	"github.com/retroenv/nesgo/pkg/cartridge"
	"github.com/retroenv/nesgo/pkg/nes"
)

func TestNestest(t *testing.T) {
	file, err := os.Open("nestest.nes")
	assert.NoError(t, err)

	cart, err := cartridge.LoadFile(file)
	assert.NoError(t, err)
	assert.NoError(t, file.Close())

	var b bytes.Buffer
	trace := bufio.NewWriter(&b)

	opts := []nes.Option{
		nes.WithEmulator(),
		nes.WithCartridge(cart),
		nes.WithEntrypoint(0xc000),
		nes.WithStopAt(0x0001),
		nes.WithDisabledGUI(),
		nes.WithTracing(),
		nes.WithTracingTarget(trace),
	}
	nes.Start(nil, opts...)

	assert.NoError(t, trace.Flush())

	file, err = os.Open("nestest_no_ppu.log")
	assert.NoError(t, err)

	nestest := bufio.NewScanner(file)
	emulator := bufio.NewScanner(bufio.NewReader(&b))
	for nestest.Scan() {
		expected := nestest.Text()
		assert.True(t, emulator.Scan())
		got := emulator.Text()
		assert.Equal(t, expected, got)
	}

	assert.NoError(t, file.Close())
}
