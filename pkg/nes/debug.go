//go:build !nesgo

package nes

import (
	"context"
	"encoding/hex"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
)

const defaultWebserverTimeout = 5 * time.Second

type debugServer struct {
	server *http.Server
	sys    *System
}

func newDebugServer(listenAddress string, sys *System) *debugServer {
	d := &debugServer{
		sys: sys,
	}

	r := chi.NewRouter()

	r.Route("/ppu", func(r chi.Router) {
		r.Get("/palette", d.ppuPalette)
	})

	d.server = &http.Server{
		Addr:         listenAddress,
		Handler:      r,
		ReadTimeout:  defaultWebserverTimeout,
		WriteTimeout: defaultWebserverTimeout,
	}

	return d
}

func (d *debugServer) start(ctx context.Context) {
	d.server.BaseContext = func(_ net.Listener) context.Context {
		return ctx
	}

	if err := d.server.ListenAndServe(); err != nil {
		panic(fmt.Errorf("listening on %s: %w", d.server.Addr, err))
	}
}

func (d *debugServer) ppuPalette(w http.ResponseWriter, r *http.Request) {
	palette := d.sys.Bus.PPU.Palette()
	data := palette.Data()

	buf := strings.Builder{}
	fmt.Fprintf(&buf, "background color: %s\n", hex.EncodeToString(data[0:1]))
	fmt.Fprintf(&buf, "background palette 0: %s\n", hex.EncodeToString(data[1:4]))
	fmt.Fprintf(&buf, "background palette 1: %s\n", hex.EncodeToString(data[4:7]))
	fmt.Fprintf(&buf, "background palette 2: %s\n", hex.EncodeToString(data[7:10]))
	fmt.Fprintf(&buf, "sprite palette 0: %s\n", hex.EncodeToString(data[10:13]))
	fmt.Fprintf(&buf, "sprite palette 1: %s\n", hex.EncodeToString(data[13:16]))
	fmt.Fprintf(&buf, "sprite palette 2: %s\n", hex.EncodeToString(data[16:19]))
	fmt.Fprintf(&buf, "sprite palette 3: %s\n", hex.EncodeToString(data[19:22]))
	_, _ = w.Write([]byte(buf.String()))
}
