package nestest

import (
	"bufio"
	"bytes"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/retroenv/nesgo/internal/assert"
	"github.com/retroenv/nesgo/pkg/cartridge"
	"github.com/retroenv/nesgo/pkg/nes"
)

func TestNestest(t *testing.T) {
	_, filename, _, ok := runtime.Caller(0)
	assert.True(t, ok)
	testDir := filepath.Dir(filename)

	file, err := os.Open(path.Join(testDir, "nestest.nes"))
	assert.NoError(t, err)

	cart, err := cartridge.LoadFile(file)
	assert.NoError(t, err)
	assert.NoError(t, file.Close())

	var b bytes.Buffer
	trace := bufio.NewWriter(&b)

	opts := nes.NewOptions(
		nes.WithEmulator(),
		nes.WithCartridge(cart),
		nes.WithEntrypoint(0xc000),
		nes.WithTracing(),
		nes.WithTracingTarget(trace),
	)
	sys := nes.InitializeSystem(opts)
	nes.RunEmulatorUntil(sys, 0x0001)

	assert.NoError(t, trace.Flush())

	file, err = os.Open(path.Join(testDir, "nestest_noclock.log"))
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
