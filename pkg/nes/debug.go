//go:build !nesgo

package nes

import (
	"context"
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
	fmt.Fprintf(&buf, "background color: %s\n", bytesToHex(data[0:1]))
	fmt.Fprintf(&buf, "background palette 0: %s\n", bytesToHex(data[1:4]))
	fmt.Fprintf(&buf, "background palette 1: %s\n", bytesToHex(data[4:7]))
	fmt.Fprintf(&buf, "background palette 2: %s\n", bytesToHex(data[7:10]))
	fmt.Fprintf(&buf, "sprite palette 0: %s\n", bytesToHex(data[10:13]))
	fmt.Fprintf(&buf, "sprite palette 1: %s\n", bytesToHex(data[13:16]))
	fmt.Fprintf(&buf, "sprite palette 2: %s\n", bytesToHex(data[16:19]))
	fmt.Fprintf(&buf, "sprite palette 3: %s\n", bytesToHex(data[19:22]))
	_, _ = w.Write([]byte(buf.String()))
}

func (d *debugServer) ppuNameTables(w http.ResponseWriter, r *http.Request) {
	tables := d.sys.Bus.NameTable.Data()

	buf := strings.Builder{}

	for table := 0; table < 4; table++ {
		fmt.Fprintf(&buf, "nametable %d\n", table)

		data := tables[table]
		for row := 0; row < 30; row++ {
			address := row * 32
			fmt.Fprintf(&buf, "%s\n", bytesToHex(data[address:address+32]))
		}
	}
	_, _ = w.Write([]byte(buf.String()))
}

func bytesToHex(data []byte) string {
	parts := make([]string, 0, len(data))
	for _, b := range data {
		parts = append(parts, fmt.Sprintf("%02X", b))
	}
	s := strings.Join(parts, ",")
	return s
}
