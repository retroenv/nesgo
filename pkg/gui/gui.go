//go:build !nesgo

// Package gui implements multiple emulator GUIs.
package gui

const (
	scaleFactor = 2.0
	windowTitle = "nesgo"
)

// reference to avoid unused linter warning
var _ = scaleFactor
var _ = windowTitle
