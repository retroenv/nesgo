# nesgo - Golang to NES ROM compiler

nesgo allows you to write programs for the Nintendo Entertainment System (NES) using Golang.

## Features

- Code autocompletion in any IDE that supports Golang
- The code can be debugged directly from an IDE at source level, it will be executed in the built-in Emulator
- Easy unit testing of code
- Simple code documentation generation
- Alias functions for 6502 CPU instructions to allow full control over output
* Outputs a ca65 compatible .asm file to allow easy inspection of generated code

Check the [issue tracker](https://github.com/retroenv/nesgo/issues?q=is%3Aissue+is%3Aopen+label%3Acompiler) for planned features or known bugs.

**nesgo is in an early stage of development and needs more work before it
can be used to build a larger project for NES!**

## Installation

Your system needs to have a recent [Golang](https://go.dev/) version installed.

[cc65](https://github.com/cc65/cc65) needs to be installed, it is used for generating
the final .nes file from assembly output generated by nesgo.
It is planned to remove this dependency in future versions.

To use the GUI mode check [GUI installation](https://github.com/retroenv/nesgo/blob/main/docs/gui.md) to set up the GUI dependencies.

Install the latest stable version by running:

```
go install github.com/retroenv/nesgo/cmd/nesgo@latest
```

The latest development version can be installed using:

```
git clone https://github.com/retroenv/nesgo.git
cd nesgo
go build ./cmd/nesgo
# use the dev version:
./nesgo  
```

## Usage

See one of the examples on how to write code for the NES using Golang.

The first compilation can take a few minutes depending on the system,
this is due to Golang compiling the CGO GUI dependencies.

nesgo can be used in different ways:

1. Compile a project to a .nes file:
   `nesgo -f ./examples/blue/main.go -o ./examples/blue/main.nes`

2. Use `go build ./examples/blue/main.go` to compile the program as a
   static binary including the Emulator

3. Run the code in the Emulator using `go run ./examples/blue/main.go`

4. Debug the program using your IDE of choice using the Delve debugger,
   set breakpoints, watch memory or CPU register, execute the code step by step etc

5. Run the generated .nes file using the `go run ./cmd/nesgoemu -f examples/blue/main.nes`

## Options

```
usage: nesgo [options] <file to compile>

  -o string
    	name of the output .nes file
  -q	perform operations quietly
```

## Differences / Limitations

* `return` has to be used instead of `rts` - it will get automatically
  added at the end of functions that are not inlined
* `goto` has to be used instead of `jump` - it is limited to the labels in the
  current function as jump destination
* for instructions that accept multiple addressing modes, the parameters can be
  cast into a helper type to set the mode, using the identifiers
  `ZeroPage`, `Absolute` or `Indirect`. If the instruction supports
  an immediate parameter, it will be set by default
