# How to build Backpack binaries from source code

This document will show you how to build binaries to use Backpack.

## Requirements

Before proceeding ensure to have the following software installed:

* [git](https://git-scm.com)
* [GNU Make](https://www.gnu.org/software/make/)
* [Go](https://golang.org) version > 1.15

On **macOS** you can use [homebrew](https://brew.sh) to install quickly the
requirements:

```shell
brew install go make git
```

On most of Debian and Ubuntu-based OOS you can run:

```shell
apt update
apt install git make
```

Please follow [this guide](https://github.com/golang/go/wiki/Ubuntu) (pkg) or
[this guide](https://golang.org/doc/install) (src) to install Go with the latest
version, as both Debian and Ubuntu official repositories might provide an old
version.

## Building

First, we need to clone the repository and obtain the [latest version](https://gitlab.com/Qm64/backpack/-/tags/latest)
of the source code.

```shell
git clone https://gitlab.com/Qm64/backpack.git
git checkout latest
```

Now we are ready to clean the builds, get the dependencies and compile the
binaries.

```shell
cd backpack
make
```

The binary will be available in `build/backpack`

## Installing

To install `buildpack` command, you can run:

```shell
make install
```

This will install it under your `$GOBIN`, which defaults to `$GOPATH/bin`
and should be available in the `$PATH`.