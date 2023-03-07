//go:build !nesgo

// Package debugger provides a Debugger webserver.
package debugger

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/retroenv/nesgo/pkg/bus"
)

const defaultWebserverTimeout = 5 * time.Second

// Debugger implements a Debugger webserver.
type Debugger struct {
	bus    *bus.Bus
	server *http.Server
}

// New creates a new debugger webserver.
func New(listenAddress string, bus *bus.Bus) *Debugger {
	d := &Debugger{
		bus: bus,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/cpu", d.cpuState)
	mux.HandleFunc("/cpu/pause", d.cpuPause)

	mux.HandleFunc("/mapper", d.mapperState)

	mux.HandleFunc("/ppu/palette", d.ppuPalette)
	mux.HandleFunc("/ppu/mirrormode", d.ppuMirrorMode)
	mux.HandleFunc("/ppu/nametables", d.ppuNameTables)

	d.server = &http.Server{
		Addr:         listenAddress,
		Handler:      mux,
		ReadTimeout:  defaultWebserverTimeout,
		WriteTimeout: defaultWebserverTimeout,
	}

	return d
}

// Start the debugger webserver, this needs to be called in a goroutine.
func (d *Debugger) Start(ctx context.Context) {
	d.server.BaseContext = func(_ net.Listener) context.Context {
		return ctx
	}

	if err := d.server.ListenAndServe(); err != nil {
		panic(fmt.Errorf("listening on %s: %w", d.server.Addr, err))
	}
}
