# nesgodisasm - Disassembler for NES ROMs

nesgodisasm allows you to disassemble programs for the Nintendo Entertainment System (NES).

## Features

* Outputs ca65 compatible .asm files that can be used to reproduce the original NES ROM
* Translates known RAM addresses to aliases
* Traces the program execution flow to differentiate between code and data
* Supports undocumented 6502 CPU opcodes
* Supports branching into opcode parts of an instruction
* Does not output trailing zero bytes of banks by default
* Flexible architecture that allows it to create output modules for other assemblers 

Check the [issue tracker](https://github.com/retroenv/nesgo/issues?q=is%3Aissue+is%3Aopen+label%3Adisassembler) for planned features or known bugs.

Currently, only ROMs that use mapper 0 are supported.

## Installation

There are different options to install nesgodisasm, the binary releases do not have any dependencies, 
compiling the tool from source code needs to have a recent version of [Golang](https://go.dev/) installed.

1. Download and unpack a binary release from [Releases](https://github.com/retroenv/nesgo/releases)

2. Install the latest release from source: 

```
go install github.com/retroenv/nesgo/cmd/nesgodisasm@latest
```

3. Build the current development version:

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
  -nohexcomments
    	do not output opcode bytes as hex values in comments
  -nooffsets
    	do not output offsets in comments
  -o string
    	name of the output .asm file, printed on console if no name given
  -q	perform operations quietly
  -verify
    	verify the generated output by assembling with ca65 and check if it matches the input
  -z	output the trailing zero bytes of banks
```
