//go:build !nesgo

// Package debugger provides a Debugger webserver.
package debugger

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
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

	r := chi.NewRouter()

	r.Route("/cpu", func(r chi.Router) {
		r.Get("/", d.cpuState)
		r.Post("/pause", d.cpuPause)
	})

	r.Route("/mapper", func(r chi.Router) {
		r.Get("/", d.mapperState)
	})

	r.Route("/ppu", func(r chi.Router) {
		r.Get("/palette", d.ppuPalette)
		r.Get("/mirrormode", d.ppuMirrorMode)
		r.Get("/nametables", d.ppuNameTables)
	})

	d.server = &http.Server{
		Addr:         listenAddress,
		Handler:      r,
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
