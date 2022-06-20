# nesgodisasm - Disassembler for NES ROMs

nesgodisasm allows you to disassemble programs for the Nintendo Entertainment System (NES).


## Features

* Outputs ca65 compatible .asm files that can be used to reproduce the original NES ROM
* Translates known RAM addresses to aliases 
* Follows the program execution flow to differentiate between code and data
* Supports undocumented 6502 opcodes

## Installation

Your system needs to have a recent [Golang](https://go.dev/) version installed.

Install the latest stable version by running:

```
go install github.com/retroenv/nesgo/cmd/nesgo@latest
```

The latest development version can be installed using:

```
git clone https://github.com/retroenv/nesgo.git
cd nesgo
go build ./cmd/nesgodisasm
# use the dev version:
./nesgodisasm  
```

## Usage

Disassemble a ROM:

```
nesgodisasm -o example.asm example.nes
```

Assemble a changed .asm file back to a ROM:

```
ca65 example.asm -o example.o
ld65 example.o -t nes -o example.nes 
```

## Options

```
usage: nesgodisasm [options] <file to disassemble>

  -o string
    	name of the output .asm file, printed on console if no name given
  -q	perform operations quietly
  -v	verify using ca65 that the generated output matches the input
```
