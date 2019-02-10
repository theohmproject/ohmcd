ohmcd
====

[![Build Status](https://travis-ci.org/ohmcsuite/ohmcd.png?branch=master)](https://travis-ci.org/ohmcsuite/ohmcd)
[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](http://copyfree.org)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/ohmcsuite/ohmcd)

ohmcd is an alternative full node bitcoin implementation written in Go (golang).

This project is currently under active development and is in a Beta state.  It
is extremely stable and has been in production use since October 2013.

It properly downloads, validates, and serves the block chain using the exact
rules (including consensus bugs) for block acceptance as Bitcoin Core.  We have
taken great care to avoid ohmcd causing a fork to the block chain.  It includes a
full block validation testing framework which contains all of the 'official'
block acceptance tests (and some additional ones) that is run on every pull
request to help ensure it properly follows consensus.  Also, it passes all of
the JSON test data in the Bitcoin Core code.

It also properly relays newly mined blocks, maintains a transaction pool, and
relays individual transactions that have not yet made it into a block.  It
ensures all individual transactions admitted to the pool follow the rules
required by the block chain and also includes more strict checks which filter
transactions based on miner requirements ("standard" transactions).

One key difference between ohmcd and Bitcoin Core is that ohmcd does *NOT* include
wallet functionality and this was a very intentional design decision.  See the
blog entry [here](https://blog.conformal.com/ohmcd-not-your-moms-bitcoin-daemon)
for more details.  This means you can't actually make or receive payments
directly with ohmcd.  That functionality is provided by the
[ohmcwallet](https://github.com/ohmcsuite/ohmcwallet) and
[Paymetheus](https://github.com/ohmcsuite/Paymetheus) (Windows-only) projects
which are both under active development.

## Requirements

[Go](http://golang.org) 1.11 or newer.

## Installation

#### Windows - MSI Available

https://github.com/ohmcsuite/ohmcd/releases

#### Linux/BSD/MacOSX/POSIX - Build from Source

- Install Go according to the installation instructions here:
  http://golang.org/doc/install

- Ensure Go was installed properly and is a supported version:

```bash
$ go version
$ go env GOROOT GOPATH
```

NOTE: The `GOROOT` and `GOPATH` above must not be the same path.  It is
recommended that `GOPATH` is set to a directory in your home directory such as
`~/goprojects` to avoid write permission issues.  It is also recommended to add
`$GOPATH/bin` to your `PATH` at this point.

- Run the following commands to obtain ohmcd, all dependencies, and install it:

```bash
$ cd $GOPATH/src/github.com/ohmcsuite/ohmcd
$ GO111MODULE=on go install -v . ./cmd/...
```

- ohmcd (and utilities) will now be installed in ```$GOPATH/bin```.  If you did
  not already add the bin directory to your system path during Go installation,
  we recommend you do so now.

## Updating

#### Windows

Install a newer MSI

#### Linux/BSD/MacOSX/POSIX - Build from Source

- Run the following commands to update ohmcd, all dependencies, and install it:

```bash
$ cd $GOPATH/src/github.com/ohmcsuite/ohmcd
$ git pull
$ GO111MODULE=on go install -v . ./cmd/...
```

## Getting Started

ohmcd has several configuration options available to tweak how it runs, but all
of the basic operations described in the intro section work with zero
configuration.

#### Windows (Installed from MSI)

Launch ohmcd from your Start menu.

#### Linux/BSD/POSIX/Source

```bash
$ ./ohmcd
```

## IRC

- irc.freenode.net
- channel #ohmcd
- [webchat](https://webchat.freenode.net/?channels=ohmcd)

## Issue Tracker

The [integrated github issue tracker](https://github.com/ohmcsuite/ohmcd/issues)
is used for this project.

## Documentation

The documentation is a work-in-progress.  It is located in the [docs](https://github.com/ohmcsuite/ohmcd/tree/master/docs) folder.

## GPG Verification Key

All official release tags are signed by Conformal so users can ensure the code
has not been tampered with and is coming from the ohmcsuite developers.  To
verify the signature perform the following:

- Download the Conformal public key:
  https://raw.githubusercontent.com/ohmcsuite/ohmcd/master/release/GIT-GPG-KEY-conformal.txt

- Import the public key into your GPG keyring:
  ```bash
  gpg --import GIT-GPG-KEY-conformal.txt
  ```

- Verify the release tag with the following command where `TAG_NAME` is a
  placeholder for the specific tag:
  ```bash
  git tag -v TAG_NAME
  ```

## License

ohmcd is licensed under the [copyfree](http://copyfree.org) ISC License.
