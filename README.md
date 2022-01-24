# OMX Order Management & Execution for FTX.com

## Why OMX?

After having searched and tried a lot, I could not find a proper order management & execution service for FTX exchange
that allows me to send limit orders straight into the order book and is actually open source. Therefore, I wrote one.

When you want your orders routed the way you want at the conditions you want and on the terms you want, you actually
need an OMX that manages your orders and executions.

## Features

OMX provides a REST API with handful of features beyond the standard FTX API.

1) Tradingview integration - Connect OMX to Tradingview and let it execute your orders on FTX
2) REST integration - Connect to any other algo trading and let it execute your orders on FTX
3) Multiaccount support - Add different (sub) accounts and route orders to each of them.
4) Set / unset leverage per (sub) account
5) Orderbook limit orders - Send orders straight into the order book using dynamic pricing
6) Automatic order size - Define a percentage of account value i.e. 3% and OMX calculates the correct order size
7) Automatic order close - Just send a close or stop loss and OMX infers the order size from the opening order
8) End to end encryption of all sensitive data
9) Required API authentication for all orders
10) DB logging of order fill & status

OMX does NOT provide:

* OCO - One cancel the other order
* Cancels or re-prices & resend unfilled orders
* Market orders

## Supported platforms

Supported:

* MacOS Intel & M1 (Arm64)
* Linux Intel, AMD64 & Arm64
* Windows 10 requires Windows Subsystem for Linux (WSL)

Untested platforms:

* Raspberry Pi (Might work as it actually is Linux Arm64)
* Windows 11 (Should technically exactly as Win 10 when using WSL, but was never tested)
* Windows on Arm (Anyone out there with a Win Arm64 device who is willing to do a test build?)

## Requirements

* Bash
* clang 7 or higher (required for Bazel & CGO)
* Bazelisk (will download the Bazel version defined for the project)
* Docker

Please run:

```bash
make check  
```

The script will test all requirements and provide download & install links. Please install the missing bits and re-run
the script until all requirements are available. Expected outcome of the script after all requirements have been
installed:

```bash
* Bash installed
* Make installed
* Clang installed
* Bash configured clang as CC required by Bazel
* Curl installed
* WGet installed
* Docker installed
* Bazel installed

===============================
All OMX dependencies installed!
===============================
```

#### Golang

* No Golang installation is required.
* Any existing Golang installation will not be touched.
* The Golang SDK required to build OMX will be downloaded by BAZEL during the initial build.

#### Bazelisk

I strongly advised to install Bazelisk and let it manage Bazel installations instead of fiddling with Bazel
installation. The underlying reasoning is that, while Bazel 4 maintains long term support, the reality is that all bash
scripts assume functionality of the Bazel version defined for the project. Currently, the project defines Bazel 4.2.2
because only version 4.2 and later builds correctly on Apply Silicon Mac systems (Arm64) so in order to support this new
platform, no Bazel version prior to 4.2 should be used.

Install Bazelisk:

* Mac/Homebrew: brew install bazelisk
* npm: npm install -g @bazel/bazelisk
* Linux: see step below
* Windows: Use WSL and follow the linux steps below

```bash
wget https://github.com/bazelbuild/bazelisk/releases/download/v1.11.0/bazelisk-linux-amd64
chmod +x bazelisk-linux-amd64
sudo mv bazelisk-linux-amd64 /usr/local/bin/bazel
     
# make sure you get the binary available in $PATH
which bazel
/usr/local/bin/bazel
```

## Getting started

Setup requires three steps.

1) Add new crypto keys (important & mandatory)
    - Generate keys: make gen_keys
    - Add keys to crypto config (don't run yet, need container rebuild)
    - open src/cfg/crypto_config.go
    - Insert each Base64 key as A) masterKey B) restApiAuthKey C) restOrderAuthKey
    - Keep the plaintext key for B) restApiAuthKey C) restOrderAuthKey
    - The two plaintext API auth keys need to be sent for either Account or order requests. See API Doc for details
    - Important: If the crypto keys are missing, no data can be encrypted / decrypted, and API authentication will fail
      one way or the other.

2) DB setup; run:
    - make db_deploy
    - make db_setup

3) Build container
    1) make build_docker
    2) make run_docker

Verify that both DB & OMX are up & running:

```bash
docker ps

CONTAINER ID   IMAGE                               COMMAND                  CREATED          STATUS         PORTS                          NAMES
dd348e68ab69   omx:latest                          "/service"               22 seconds ago   Up 7 seconds   0.0.0.0:80->80/tcp, 9090/tcp   omx
69d76ea97678   timescale/timescaledb:latest-pg14   "docker-entrypoint.s…"   8 hours ago      Up 8 hours     0.0.0.0:5432->5432/tcp         timescaledb
```

Start & Stop OMX container

```bash
docker container start omx
docker container stop omx
```

Important details:

* When adding an FTX account to OMX, order monitoring starts immediately
* When restarting the OMX container, all added accounts will be restored and order monitoring resumes. For many
  accounts, this may take a moment.
* When deleting the OMX container, all added accounts will remain fully encrypted in the database unless manually
  deleted.
* When starting a new OMX container, all accounts will be restored.
* However, when changing or losing the master key, all stored data are lost because no account can be decrypted anymore.
* When resetting the entire database, all data are lost, obviously.

## Make commands

```bash
Setup:
make check        	    Checks whether all project requirements are present.
make gen_keys           Generates new API access keys and a master key for end to end encryption.
make db_deploy        	Deploys & starts the DB container. Run just once to create. Then use docker container start/stop timescaledb.
make db_configure   	Configures the initial DB. Run once to first initialize DB or run again to run a hard DB reset which deletes all data!

Dev:
make build   		Builds the code base incrementally (fast). Use when coding.
make rebuild   		Rebuilds all dependencies & the code base (slow). Use after go mod changes.
make run   	        Runs the default target defined in dev/run script. Use to run default binary.

Docker:
make build_docker   	Builds a docker image locally.
make run_docker   	Run docker images locally.
make publish_docker  	Publishes docker images to registry.
make remove_docker	Removes OMX container & image. DB & data will remain intact and carry over.
make replace_docker  	Replaces running OMX image with latest published image. DB & data will carry over to new version.
make reset_docker    	ALL DATA WILL BE LOST: Removes running OMX container, replaces it with latest local build, AND destroys & rebuilds DB.
```

## API Guide

* [Account API](docs/api/account_api_guide.md)
* [Order API](docs/api/order_api_guide.md)

## Known issues:

### Order API needs documentation

True. It's on my todo list. In the meantime, look at the curl examples for each API function in
src/service/v1/http_handlers_order_*. Feel free to contribute documentation.

### Docker warning on Apple/M1 macs

Issue: When starting the OMX container on OSX on Apple/M1, docker issues the follwing warning.

Warning/Error:

WARNING: The requested image's platform (linux/amd64) does not match the detected host platform (linux/arm64/v8) and no
specific platform was requested

Workaround:

* Just Ignore. Image runs just fine.
* Rebuild image on your Mac, run it, and the warning will be gone

### Failed DB connection when starting OMX binary locally without docker

Issue: When starting OMX locally (without docker) start process aborts because of no DB connection despite an otherwise
working database.

Warning/Error:
dbClient/CreateDataBase: Can't create or update DB schema dial tcp: lookup timescaledb: no such host exit status 1

Cause:

* OMX selects the wrong DB host
* Currently, environment flags are set manually in config, and by default the docker db host is used.

Solution:

* Open src/cfg/main_config.go
* Set 'const env = t.Prod' to 'const env = t.Dev'
* Rebuild & re-run container

## Development

* See ngrok guide for local request inspection
* See standard component guide for overall development style & best practices
* See CIRA guide for even more details on the used component model

Licence:

* All content under MIT Licence.

Author:

* Marvin Hansen 
