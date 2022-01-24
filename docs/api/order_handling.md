# OMX Order handling & routing

## Order handling

1) OMX does *not* cancel any unfilled order. Tracking order state is the sole responsibility of the issuing system. (
   Thus no OCO)
2) OMX does stores the configuration of the opening order to infer the matching close order
3) OMX does not allow multiple (long / short) positions on the same instrument
4) OMX, for the time being, only allows single positions on any given account
5) OMX only supports limit orders for the time being, but market orders may be added in the future

## Order routing

### Conventional order routing

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

### Dynamic order book routing

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
