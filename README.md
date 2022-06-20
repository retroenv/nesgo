# nesgo - Golang based tooling for the NES

[![Build status](https://github.com/retroenv/nesgo/actions/workflows/go.yaml/badge.svg?branch=main)](https://github.com/retroenv/nesgo/actions)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/retroenv/nesgo)
[![Go Report Card](https://goreportcard.com/badge/github.com/retroenv/nesgo)](https://goreportcard.com/report/github.com/retroenv/nesgo)
[![codecov](https://codecov.io/gh/retroenv/nesgo/branch/main/graph/badge.svg?token=NS5UY28V3A)](https://codecov.io/gh/retroenv/nesgo)

nesgo offers tooling for the Nintendo Entertainment System (NES), written in Golang.

## Included tooling

[nesgo](https://github.com/retroenv/nesgo/tree/main/cmd/nesgo) Golang to NES compiler
[nesgodisasm](https://github.com/retroenv/nesgo/tree/main/cmd/nesgodisasm) Disassembler for NES ROMs
[nesgoemu](https://github.com/retroenv/nesgo/tree/main/cmd/nesgoemu) Emulator for NES ROMs

## Project layout

    ├─ cmd              Tools main directories
    ├─ example/         Program examples in Golang
    ├─ internal/        internal compiler code
    ├─ pkg/             libraries used by different packages
    ├─ pkg/neslib       helper useful for writing NES programs in Golang
