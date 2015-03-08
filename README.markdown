# path_helper.go

This is a Golang implementation of the OS X [`path_helper`][path_helper]
utility.

[path_helper]: https://developer.apple.com/library/mac/documentation/Darwin/Reference/ManPages/man8/path_helper.8.html

## Usage

First setup one or more path files. A path file contains one directory per line
that should be added to `PATH`. Eg:

```
/usr/local/bin
/usr/bin
/bin
```

Add the following to your `.bashrc` or `.zshrc` file:

```
eval $(path_helper ~/.paths)
```

If you have multiple path files specify them as arguments:

```
eval $(path_helper /etc/paths /etc/paths.d/*)
```

## Installation

Clone this repo:

```
git clone https://github.com/itspriddle/path_helper.git
```

Build and install:

```
make
make install
```

If you want to install somewhere other than `/usr/local/bin`, supply `PREFIX`:

```
make
PREFIX=$HOME/local make install
```
