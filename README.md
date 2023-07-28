# Vortexus - A tool to create working projects in seconds

## Overview

Vortexus is a cross-compatible CLI tool that can be used to create working projects in minutes. Its built in `.go` and can be run on any platform as a binary.

## Demo

https://github.com/FluffySnowman/vortexus/assets/51316255/db2ca024-8c93-446f-9bd2-589803539937

## Usage

The binaries can be downloaded from the [releases page](https://github.com/fluffysnowman/vortexus/releases) or can be built from source.

#### Easiest method-

1. Download the binaries from the [releases page](https://github.com/fluffysnowman/vortexus/releases) 
2. Double click the file to run (windows exe)

### Using the binaries (recommended)

#### Linux (curl)

```bash
# download the binary to a file
curl https://raw.githubusercontent.com/FluffySnowman/vortexus/binaries/vortexus_linux_amd64 > vortexus

# change made to executable
chmod +x vortexus

# run the binary
./vortexus
```

#### Windows (curl)

```bash
# download the binary to a file
curl https://raw.githubusercontent.com/FluffySnowman/vortexus/binaries/vortexus_windows_x86.exe > vortexus.exe

# run the binary
.\vortexus.exe
```

### Building from source

Make sue you have `go` version `1.20+` installed on your system. If not, you can download it from [here](https://golang.org/dl/).

```bash
# clone the repo
git clone https://github.com/fluffysnowman/vortexus.git --depth 1 --branch mojo
cd vorutexus

# build the binary
go build main.go

# run the binary
./main

# or on windows -
.\main.exe
```

## Current list of features

- `Express.js` API (customizable with `MongoDB` connection along with some other options)
- Static site server (made in `Express.js`)
- Basic `Go` API
- Simple `Python Flask` API

## Upcoming features

- Express API with more options
- React app
- React Native app
- Discord bot
- Vue.js Site
- Go API with more options
- multiple versions of each with db's connocted along with many other things along the way

## Contributing

I don't think anyone will but if anyone does then I'm sorry for you having to go through this messy codebase
