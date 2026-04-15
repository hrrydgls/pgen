# PGen - Your password generator in your CLI

## Installation

### With Go

To install PGen using Go, run the following command:

```bash
go install github.com/hrrydgls/pgen@latest
```

This will download, build, and install the tool to your `$GOPATH/bin` or `$GOBIN` directory.

### Without Go

Download the pre-built binary for your platform from the [releases page](https://github.com/hrrydgls/pgen/releases):

#### Linux
```bash
sudo curl -L https://github.com/hrrydgls/pgen/releases/download/v1.0.0/pgen -o /usr/local/bin/pgen && sudo chmod +x /usr/local/bin/pgen
```

#### Windows
```bash
curl -L https://github.com/hrrydgls/pgen/releases/download/v1.0.0/pgen-windows.exe -o pgen.exe
```
Note: You may need to add the current directory to your PATH or move `pgen.exe` to a directory in your PATH to run it without specifying the path.

#### macOS
```bash
sudo curl -L https://github.com/hrrydgls/pgen/releases/download/v1.0.0/pgen-mac -o /usr/local/bin/pgen && sudo chmod +x /usr/local/bin/pgen
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
