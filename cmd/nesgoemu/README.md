# nesgoemu - Emulator for NES ROMs

nesgoemu allows you to emulate ROMs for the Nintendo Entertainment System (NES).

## Features

* Offers the GUI in SDL or OpenGL mode
* Can be used headless without a GUI
* Supports outputting of CPU traces
* Supports undocumented 6502 CPU opcodes

Check the [issue tracker](https://github.com/retroenv/nesgo/labels/emulator) for planned features or known bugs.

## Installation

Your system needs to have a recent [Golang](https://go.dev/) version installed.

Check [GUI installation](https://github.com/retroenv/nesgo/blob/main/docs/gui.md) to set up the GUI dependencies. 

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

  -a string
    	listening address for the debug server to use (default "127.0.0.1:8080")
  -c	console mode, disable GUI
  -d	start built-in webserver for debug mode
  -e int
    	entrypoint to start the CPU (default -1)
  -s int
    	stop execution at address (default -1)
  -t	print CPU tracing
```
