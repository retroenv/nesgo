//go:build !nesgo

package nes

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

// Context returns a context that is cancelled automatically when a SIGINT,
// SIGQUIT or SIGTERM signal is received.
func Context() context.Context {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-sig
		cancel()
	}()

	return ctx
}
