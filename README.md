# OMX Order Management & Execution for FTX.com

OMX is an open source order management & execution system for FTX crypto exchange. It takes an order specification,
constructs the order, and sends it to the configured FTX account. OMX supports multi-tenant order routing by default
that means you can add multiple FTX accounts, and the OMX will route the order to the configured exchange account.
Notice, OMX only supports FTX exchange, but it does so with an expanded features otherwise unavailable.

# Table of Contents

1. [Why OMX?](#para1)
2. [Features](#para2)
3. [Order handling](#para3)
4. [Supported platforms](#para4)
5. [Order routing](#para5)
6. [Supported platforms](#para6)
7. [Build Requirements](#para7)

## <a name="para1"/> Why OMX?

After having searched and tried a lot, I could not find a proper order management & execution service for FTX exchange
that allows me to send limit orders straight into the order book and that is actually open source. Therefore, I wrote
one.

When you want your orders routed the way you want at the conditions you want and on the terms you want, you actually
need an OMX that manages your orders and executions.

## Features <a name="para2"></a>

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

* No OCO orders - One cancel the other order
* No bracket orders (because of lacking OCO)
* No cancel, re-price & resend unfilled orders
* No market orders
* No additional or other exchanges.

Adding these functions would be part of a more comprehensive Execution Management System (EMX), which was and still
remains outside the scope of this project.

## <a name="para3"/>  Order handling

1) OMX does *not* cancel any unfilled order. Tracking order state is the sole responsibility of the issuing system. (
   Thus no OCO)
2) OMX does stores the configuration of the opening order to infer the matching close order
3) OMX does not allow multiple (long / short) positions on the same instrument
4) OMX, for the time being, only allows single positions on any given account
5) OMX only supports limit orders for the time being, but market orders may be added in the future

## Order routing  <a name="para4"></a>

#### Conventional order routing

Conventionally, when a trading system emits a directional signal i.e. buy or sell, the system sends an order request
with that direction at the current price.

This allows 4 types of conventional orders:

* Buy to Open Long
* Sell to Close Long
* Sell to Open Short
* Buy (back) to Close Short

These 4 types can be send out either as limit orders at a pre-defined price or better, or as market order which will be
any price and usually worse than the current market price. Most crypto exchanges name Limit orders "Maker" because they
make liquidity by adding orders in the order book which may sit there for a while, and name Market orders "Taker"
because these take liquidity immediately out of the order book. In terms of fee structure, maker orders are universally
cheaper as an incentive from the exchange to build up strong market liquidity through a well filled order book.

OMX supports all four of the conventional order types as Limit orders only, but cannot guarantee an order fill or even a
placement in the order book. However, dynamic order book routing described below guarantees order book placement, but
may not guarantee (complete) fill as it depends on market conditions.

#### Dynamic order book routing

When sending out an order on a volatile trading day, it might be possible that between the moment the order was send and
the moment the order will be placed, the market moved contrary to the order direction and a limit order will be
converted into market order because the limit price falls outside the order book range. This can be named adverse order
execution because the actual execution of a proper limit order was adverse to its actual intend. There are few ways to
handle adverse execution:

1) Verify that the limit price falls within the order book price range before sending the order
   * Send for execution if the limit price matches the order book
   * Cancel order and return and error otherwise
   * Implies any order may not get send
2) Adjust the limit price so that it falls within the order book price before sending the order
   * Implies that the actual order price may differ from the initial limit price which requires internal adjustment
3) Send the order request without an actual limit price, but a configuration that determines the limit price dynamically
   from the order book
   * Implies that the actual order price remains unknown until the order has been sent and confirmed by the exchange.

OMX applies the third approach by purpose because for all major instruments in the crypto markets traded on FTC i.e.
BTC-PERP, the high liquidity ensures a relatively low spread between BID & ASK so even if the next best price of the
order book gets selected dynamically, the actual difference will be between 10 cents and maybe $1. For comparison,
getting a limit order converted into a market order can lead at worst a substantial adverse price drop away from the
market price and a 7% base taker fee in case no fee discount is applied. In practice, a 3 - 7% shift away from market
price has been observed on multiple occasions and that was on the motivations to write OMX.

OMX supports the following modes to dynamically determine the limit price from the order book:

* LargestOrderSizePrice
* SmallestOrderSizePrice
* FirstOrderBookPrice
* MidOrderBookPrice
* LastOrderBookPrice
* LowestPrice
* HighestPrice

The default book size in OMX has been set to the first 20 entries of the order book to ensure speedy processing.
Therefore, LastOrderBookPrice actually refers to index position 19 (as the index starts with 0).

The largest order size price refers to the price in the order book with the largest order volume. It is often the case
that the largest buy / bid order volume has been placed at either a pivot point or otherwise defined retracement level.
At that price level, order volume is often ten times or more above the normal order book volume per price.

Selecting the largest order size price implies:

1) The order will be placed in the order book as adverse price movements are relatively unlikely (but not impossible)
2) Because of the guaranteed order book placement, all orders will be handled and billed as standard limit orders.
3) Because of the extreme order size, by the time the market price hits that price level, a nano bull run follows.
4) Depending on the actual market activity, it may take some time (several minutes or longer) until the fill occurs.
5) Fill rate observed in practice is relatively high and leads to a wider spread between buy & sell and overall lower
   fees.
6) However, there is no fill guarantee as adverse market moves can and will happen
7) Partial fills happen and by default the FTX flag "retry until filled" will be send to ensure an order remains open
   until completed although this requires the market to touch the limit price multiple times until fill has been
   completed.

Smallest order size refers to the price at which the order book volume is the smallest. There are certain corner cases
were that may be desired.

First and highest order size price are identical as the first order book is always the lowest price. These two names
only exists to allow for both conventions although the actual target price will always be the highest price in the order
book. Be careful, as this one changes within split-seconds and requires near zero latency to hit.

Similarly, last and lowest / best price are identical too. For the buy / bid side, it is always the lowest price. For
the sell / ask side, it is always the highest price that is the best. Be careful, as these are very far away from the
actual market price and may not be reached at all, but OMX will not cancel the open order so it is up to the sending
system to track order state and may cancel unfilled orders.

MidOrderBookPrice refers to the price at index position 9, which is the tenth price of the order book. In situations of
unstable latency, this may help to get a price that is reasonable and sufficiently away from the market price to ensure
order book placement regardless of latency induced delays.

More dynamic order book pricing can be added by modifying the workflow in the orderManager component.

## Supported platforms  <a name="para5"></a>

Supported:

* MacOS Intel & M1 (Arm64)
* Linux Intel, AMD64 & Arm64
* Windows 10 requires Windows Subsystem for Linux (WSL)

Untested platforms:

* Raspberry Pi (Might work as it actually is Linux Arm64)
* Windows 11 (Should technically exactly as Win 10 when using WSL, but was never tested)
* Windows on Arm (Anyone out there with a Win Arm64 device who is willing to do a test build?)

## Build Requirements  <a name="para6"></a>

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

## Getting started  <a name="para7"></a>

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

## Make commands <a name="para8"></a>

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

## API Guide <a name="para8"></a>

* [Account API](docs/api/account_api_guide.md)
* [Order API](docs/api/order_api_guide.md)

## Known issues: <a name="para9"></a>

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

## Development <a name="para10"></a>

* See ngrok guide for local request inspection
* See standard component guide for overall development style & best practices
* See CIRA guide for even more details on the used component model

## Licence: <a name="para11"></a>
* All content under MIT Licence.

## Author(s): <a name="para12"></a>
* Marvin Hansen 
