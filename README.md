# goNES - Golang for the NES

goNES allows you to write programs for the Nintendo Entertainment System (NES).

Benefits of writing NES programs in Golang:

- Code autocompletion support in IDEs
- You can debug code directly from an IDE, it will be executed in the
  built-in Emulator (currently in a rather simple state) 
- Easy unit testing for your code
- Simple code documentation generation

goNES is in an early state of development and needs more work before it
can be used to build a larger project for NES. 

## Installation

Install goNES using `go install github.com/retroenv/nesgo/cmd/nesgo@latest`

Install cc65, it is used for generating the final .nes file from
assembly output generated by goNES. It is planned to remove this
dependency in future versions. 

The integrated OpenGL GUI support is enabled by default. Debugging
your code will execute using the built-in Emulator and GUI.
The GUI modes can be selected using the following build flags:

* `nogui`: disables all GUI modules
* `noopengl` disables the OpenGL GUI
* `sdl` enables the SDL GUI

The following libraries need to be installed, 
depending on the operating system:

**macOS**: `Xcode or Command Line Tools for Xcode (xcode-select --install)`

**Ubuntu/Debian-like**: `build-essential libgl1-mesa-dev xorg-dev`

**CentOS/Fedora-like**: `@development-tools libX11-devel libXcursor-devel
 libXrandr-devel libXinerama-devel mesa-libGL-devel libXi-devel
 libXxf86vm-devel`

**Windows**:
  * Install [msys2](http://www.msys2.org/) 
  * Start msys2 and execute `pacman -S --needed base-devel
    mingw-w64-i686-toolchain mingw-w64-x86_64-toolchain
    mingw64/mingw-w64-x86_64-SDL2`
  * Add `c:\tools\msys64\mingw64\bin\` to user path environment variable

## Usage

See one of the examples on how to write code for the NES!

## Differences / Limitations

* `return` has to be used instead of `rts` - it will get automatically
  added at the end of functions that are not inlined
* `goto` has to be used instead of `jump` - it is limited to the labels in the
  current function as jump destination
* instructions that accept immediate addressing and zero page,
  addressing mode have an alias function for the addressing mode, for example:
  `Lda(i uint8)` and `LdaAddr(address interface{}, reg ...interface{}`
