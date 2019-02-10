### Table of Contents
1. [About](#About)
2. [Getting Started](#GettingStarted)
    1. [Installation](#Installation)
        1. [Windows](#WindowsInstallation)
        2. [Linux/BSD/MacOSX/POSIX](#PosixInstallation)
          1. [Gentoo Linux](#GentooInstallation)
    2. [Configuration](#Configuration)
    3. [Controlling and Querying ohmcd via ohmcctl](#ohmcctlConfig)
    4. [Mining](#Mining)
3. [Help](#Help)
    1. [Startup](#Startup)
        1. [Using bootstrap.dat](#BootstrapDat)
    2. [Network Configuration](#NetworkConfig)
    3. [Wallet](#Wallet)
4. [Contact](#Contact)
    1. [IRC](#ContactIRC)
    2. [Mailing Lists](#MailingLists)
5. [Developer Resources](#DeveloperResources)
    1. [Code Contribution Guidelines](#ContributionGuidelines)
    2. [JSON-RPC Reference](#JSONRPCReference)
    3. [The ohmcsuite Bitcoin-related Go Packages](#GoPackages)

<a name="About" />

### 1. About

ohmcd is a full node bitcoin implementation written in [Go](http://golang.org),
licensed under the [copyfree](http://www.copyfree.org) ISC License.

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

<a name="GettingStarted" />

### 2. Getting Started

<a name="Installation" />

**2.1 Installation**

The first step is to install ohmcd.  See one of the following sections for
details on how to install on the supported operating systems.

<a name="WindowsInstallation" />

**2.1.1 Windows Installation**<br />

* Install the MSI available at: https://github.com/ohmcsuite/ohmcd/releases
* Launch ohmcd from the Start Menu

<a name="PosixInstallation" />

**2.1.2 Linux/BSD/MacOSX/POSIX Installation**


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
$ go get -u github.com/Masterminds/glide
$ git clone https://github.com/ohmcsuite/ohmcd $GOPATH/src/github.com/ohmcsuite/ohmcd
$ cd $GOPATH/src/github.com/ohmcsuite/ohmcd
$ glide install
$ go install . ./cmd/...
```

- ohmcd (and utilities) will now be installed in ```$GOPATH/bin```.  If you did
  not already add the bin directory to your system path during Go installation,
  we recommend you do so now.

**Updating**

- Run the following commands to update ohmcd, all dependencies, and install it:

```bash
$ cd $GOPATH/src/github.com/ohmcsuite/ohmcd
$ git pull && glide install
$ go install . ./cmd/...
```

<a name="GentooInstallation" />

**2.1.2.1 Gentoo Linux Installation**

* Install Layman and enable the Bitcoin overlay.
  * https://gitlab.com/bitcoin/gentoo
* Copy or symlink `/var/lib/layman/bitcoin/Documentation/package.keywords/ohmcd-live` to `/etc/portage/package.keywords/`
* Install ohmcd: `$ emerge net-p2p/ohmcd`

<a name="Configuration" />

**2.2 Configuration**

ohmcd has a number of [configuration](http://godoc.org/github.com/ohmcsuite/ohmcd)
options, which can be viewed by running: `$ ohmcd --help`.

<a name="ohmcctlConfig" />

**2.3 Controlling and Querying ohmcd via ohmcctl**

ohmcctl is a command line utility that can be used to both control and query ohmcd
via [RPC](http://www.wikipedia.org/wiki/Remote_procedure_call).  ohmcd does
**not** enable its RPC server by default;  You must configure at minimum both an
RPC username and password or both an RPC limited username and password:

* ohmcd.conf configuration file
```
[Application Options]
rpcuser=myuser
rpcpass=SomeDecentp4ssw0rd
rpclimituser=mylimituser
rpclimitpass=Limitedp4ssw0rd
```
* ohmcctl.conf configuration file
```
[Application Options]
rpcuser=myuser
rpcpass=SomeDecentp4ssw0rd
```
OR
```
[Application Options]
rpclimituser=mylimituser
rpclimitpass=Limitedp4ssw0rd
```
For a list of available options, run: `$ ohmcctl --help`

<a name="Mining" />

**2.4 Mining**

ohmcd supports the `getblocktemplate` RPC.
The limited user cannot access this RPC.


**1. Add the payment addresses with the `miningaddr` option.**

```
[Application Options]
rpcuser=myuser
rpcpass=SomeDecentp4ssw0rd
miningaddr=12c6DSiU4Rq3P4ZxziKxzrL5LmMBrzjrJX
miningaddr=1M83ju3EChKYyysmM2FXtLNftbacagd8FR
```

**2. Add ohmcd's RPC TLS certificate to system Certificate Authority list.**

`cgminer` uses [curl](http://curl.haxx.se/) to fetch data from the RPC server.
Since curl validates the certificate by default, we must install the `ohmcd` RPC
certificate into the default system Certificate Authority list.

**Ubuntu**

1. Copy rpc.cert to /usr/share/ca-certificates: `# cp /home/user/.ohmcd/rpc.cert /usr/share/ca-certificates/ohmcd.crt`
2. Add ohmcd.crt to /etc/ca-certificates.conf: `# echo ohmcd.crt >> /etc/ca-certificates.conf`
3. Update the CA certificate list: `# update-ca-certificates`

**3. Set your mining software url to use https.**

`$ cgminer -o https://127.0.0.1:8334 -u rpcuser -p rpcpassword`

<a name="Help" />

### 3. Help

<a name="Startup" />

**3.1 Startup**

Typically ohmcd will run and start downloading the block chain with no extra
configuration necessary, however, there is an optional method to use a
`bootstrap.dat` file that may speed up the initial block chain download process.

<a name="BootstrapDat" />

**3.1.1 bootstrap.dat**

* [Using bootstrap.dat](https://github.com/ohmcsuite/ohmcd/tree/master/docs/using_bootstrap_dat.md)

<a name="NetworkConfig" />

**3.1.2 Network Configuration**

* [What Ports Are Used by Default?](https://github.com/ohmcsuite/ohmcd/tree/master/docs/default_ports.md)
* [How To Listen on Specific Interfaces](https://github.com/ohmcsuite/ohmcd/tree/master/docs/configure_peer_server_listen_interfaces.md)
* [How To Configure RPC Server to Listen on Specific Interfaces](https://github.com/ohmcsuite/ohmcd/tree/master/docs/configure_rpc_server_listen_interfaces.md)
* [Configuring ohmcd with Tor](https://github.com/ohmcsuite/ohmcd/tree/master/docs/configuring_tor.md)

<a name="Wallet" />

**3.1 Wallet**

ohmcd was intentionally developed without an integrated wallet for security
reasons.  Please see [ohmcwallet](https://github.com/ohmcsuite/ohmcwallet) for more
information.


<a name="Contact" />

### 4. Contact

<a name="ContactIRC" />

**4.1 IRC**

* [irc.freenode.net](irc://irc.freenode.net), channel `#ohmcd`

<a name="MailingLists" />

**4.2 Mailing Lists**

* <a href="mailto:ohmcd+subscribe@opensource.conformal.com">ohmcd</a>: discussion
  of ohmcd and its packages.
* <a href="mailto:ohmcd-commits+subscribe@opensource.conformal.com">ohmcd-commits</a>:
  readonly mail-out of source code changes.

<a name="DeveloperResources" />

### 5. Developer Resources

<a name="ContributionGuidelines" />

* [Code Contribution Guidelines](https://github.com/ohmcsuite/ohmcd/tree/master/docs/code_contribution_guidelines.md)

<a name="JSONRPCReference" />

* [JSON-RPC Reference](https://github.com/ohmcsuite/ohmcd/tree/master/docs/json_rpc_api.md)
    * [RPC Examples](https://github.com/ohmcsuite/ohmcd/tree/master/docs/json_rpc_api.md#ExampleCode)

<a name="GoPackages" />

* The ohmcsuite Bitcoin-related Go Packages:
    * [ohmcrpcclient](https://github.com/ohmcsuite/ohmcd/tree/master/rpcclient) - Implements a
      robust and easy to use Websocket-enabled Bitcoin JSON-RPC client
    * [ohmcjson](https://github.com/ohmcsuite/ohmcd/tree/master/ohmcjson) - Provides an extensive API
      for the underlying JSON-RPC command and return values
    * [wire](https://github.com/ohmcsuite/ohmcd/tree/master/wire) - Implements the
      Bitcoin wire protocol
    * [peer](https://github.com/ohmcsuite/ohmcd/tree/master/peer) -
      Provides a common base for creating and managing Bitcoin network peers.
    * [blockchain](https://github.com/ohmcsuite/ohmcd/tree/master/blockchain) -
      Implements Bitcoin block handling and chain selection rules
    * [blockchain/fullblocktests](https://github.com/ohmcsuite/ohmcd/tree/master/blockchain/fullblocktests) -
      Provides a set of block tests for testing the consensus validation rules
    * [txscript](https://github.com/ohmcsuite/ohmcd/tree/master/txscript) -
      Implements the Bitcoin transaction scripting language
    * [ohmcec](https://github.com/ohmcsuite/ohmcd/tree/master/ohmcec) - Implements
      support for the elliptic curve cryptographic functions needed for the
      Bitcoin scripts
    * [database](https://github.com/ohmcsuite/ohmcd/tree/master/database) -
      Provides a database interface for the Bitcoin block chain
    * [mempool](https://github.com/ohmcsuite/ohmcd/tree/master/mempool) -
      Package mempool provides a policy-enforced pool of unmined bitcoin
      transactions.
    * [ohmcutil](https://github.com/ohmcsuite/ohmcutil) - Provides Bitcoin-specific
      convenience functions and types
    * [chainhash](https://github.com/ohmcsuite/ohmcd/tree/master/chaincfg/chainhash) -
      Provides a generic hash type and associated functions that allows the
      specific hash algorithm to be abstracted.
    * [connmgr](https://github.com/ohmcsuite/ohmcd/tree/master/connmgr) -
      Package connmgr implements a generic Bitcoin network connection manager.
