# nesgoemu - Emulator for NES ROMs

nesgoemu allows you to emulator ROMs for the Nintendo Entertainment System (NES).

## Features

* Offers the GUI in SDL or OpenGL output mode
* Can be used headless without a GUI
* Supports outputting of CPU traces
* Supports undocumented 6502 CPU opcodes

Check the [issue tracker](https://github.com/retroenv/nesgo/labels/emulator) for planned features or known bugs.

## Installation

Your system needs to have a recent [Golang](https://go.dev/) version installed.

Install the latest stable version by running:

```
go install github.com/retroenv/nesgo/cmd/nesgoemu@latest
```

The latest development version can be installed using:

```
git clone https://github.com/retroenv/nesgo.git
cd nesgo
go build ./cmd/nesgoemu
# use the dev version:
./nesgoemu  
```

## Usage

Emulate a ROM:

```
nesgoemu example.nes
```

## Options

```
usage: nesgoemu [options] <file to emulate>

  -e int
    	entrypoint to start the CPU (default -1)
  -t	print CPU tracing
```
