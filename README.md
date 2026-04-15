# PGen - Your password generator in your CLI

## Installation

### With Go

To install PGen using Go, run the following command:

```bash
go install github.com/hrrydgls/pgen@latest
```

This will download, build, and install the tool to your `$GOPATH/bin` or `$GOBIN` directory.

### Without Go

Download the pre-built binary for your platform:

#### Linux
```bash
curl -L https://github.com/hrrydgls/pgen/raw/master/pgen -o pgen && chmod +x pgen
```

#### Windows
```bash
curl -L https://github.com/hrrydgls/pgen/raw/master/pgen-windows.exe -o pgen.exe
```

#### macOS
```bash
curl -L https://github.com/hrrydgls/pgen/raw/master/pgen-mac -o pgen && chmod +x pgen
```

After downloading, you can run the binary directly.

## Usage

```bash
pgen length [options]

# length is 8 by default and options are ulns

# Options:
# u -> upper case
# l -> lower case
# n -> numbers
# s -> symbols

```

## Examples

You can even run it like this:

```sh
➜ pgen 20           
j3T72Y869XDM}8U9?d)v

➜ pgen 20 u
YGRMKFFNSPQFJHPWVNGZ

➜ pgen u   
TTVEXFTK

➜ pgen   
3/B{2NZe

```
