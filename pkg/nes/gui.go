//go:build !nesgo
// +build !nesgo

package nes

import "github.com/retroenv/nesgo/pkg/bus"

type guiInitializer func(bus *bus.Bus) (guiRender func() (bool, error), guiCleanup func(), err error)

// GuiStarter will be set by the chosen and imported GUI renderer.
var GuiStarter guiInitializer
