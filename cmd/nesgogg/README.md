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

Encode a Game Genie code:

```
nesgogg -a 0x94A7 -v 0x02
```

## Options

```
usage: nesgogg [options] <code>

  -a string
    	address to patch in decimal or hex with 0x prefix
  -c string
    	compare value in decimal or hex with 0x prefix
  -v string
    	value to write in decimal or hex with 0x prefix
```
