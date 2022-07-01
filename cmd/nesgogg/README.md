# nesgogg - NES Game Genie decoder/encoder

## Installation

There are different options to install nesgogg, the binary releases do not have any dependencies, 
compiling the tool from source code needs to have a recent version of [Golang](https://go.dev/) installed.

1. Download and unpack a binary release from [Releases](https://github.com/retroenv/nesgo/releases)

2. Install the latest release from source: 

```
go install github.com/retroenv/nesgo/cmd/nesgogg@latest
```

3. Build the current development version:

```
git clone https://github.com/retroenv/nesgo.git
cd nesgo
go build ./cmd/nesgogg
# use the dev version:
./nesgogg  
```

## Usage

Decode a Game Genie code:

```
nesgogg <CODE>

# For example:
nesgogg PIGOAP
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
