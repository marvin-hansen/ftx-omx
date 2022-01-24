## Basic Order Types

FTX order documentation: https://help.ftx.com/hc/en-us/sections/4414741392916-Types-of-Order

### Market Order

A market order is an order sent as far through the book as possible. Normally this means that market orders (and stop
market orders) will get fully filled but that it could be at any price. FTX has certain price limits in place for
markets order to prevent a market order from accidentally moving a market 60% during an illiquid period.

The market order price limits are:

* 25% through the orderbook for spot markets (including leveraged tokens),
* 2% through the book for futures markets, and
* 2% of the underlying asset for MOVE markets.

#### Fill:

* When the order book contains matching or more liquidity, market orders will always be filled although in various parts
  and at any price.
* However, if there isn't sufficient size bid/offered within the price limits of the best bid/offer, a market order
  might not be fully filled.

#### Implications:

* Even when aiming for spot trading, use the matching PERP Future with leverage set to 1 to ensure any accidental market
  order may only move 2% through the book to prevent substantial slippage.
* PERP Futures usually have higher liquidity and a deeper order book so that increases chances of an at target price or
  very close to target fill from the book.
* For markets with low or inconsistent liquidity, then, market orders may ensure a fill although at an adverse price and
  at substantially higher fee.
* Market order fees can’t easily be reduced through staking, thus incur substantial expenses for actual adverse order
  execution. Thus, market orders should only be considered as action of last resort for example when a market crashes.

#### Application:

* Implement stop-market as last resort option required to add a sell-off protection mechanism.

### Limit Order

A limit order is an order sent that will buy/sell up only to a certain price. if you send a limit buy order with a limit
price on BTC with a limit price of $10,100, that means you are willing to pay up to $10,100 for BTC. Your order will
trade against any resting offers below $10,100, and if it is not fully filled it will leave out a providing bid at
$10,100 for the remaining size.

#### Buy Limit:

* The limit price for buy orders means buy either at the limit price or lower.
* Any unfilled parts of the order return to the book and may (or may not) get filled at the limit price or lower.

#### Sell Limit

* The limit price for a sell orders means sell either at the limit price or higher.
* Any unfilled parts of the order return to the book and may (or may not) get filled at a later stage at the limit price
  or higher.

#### Fill:

* Limit orders are not guaranteed to get filled
* If no one is willing to sell below your limit price (or buy above, in the case of a sell limit order) then the order
  will not be fully filled, and the remainder will be sent as a providing order on the orderbook at the limit price.
* However, limit orders do guarantee that to the extent you do get filled on it, the fills will be at a price no worse
  than your limit price.

#### Implications:

* For markets with high liquidity, the lack of execution guarantee might be a lot less of an issue as it appears on
  paper because of a deep order book. However, this does not guarantee complete execution as large orders may not get
  filled entirely at or better than limit price when the market is moving.
* For favorable buy execution, the order limit must be placed below the current price to prevent conversion into a
  market order.
* For favorable sell execution, the order limit must be placed above the current price to ensure order book placement.

#### Application:

* Minimize total latency from order issuance until order book placement.
* Limit buy: Inspect the order book for the normal difference between current price and best Bid price and add the
  difference as offset to calculate the

## Advanced Order Types

FTX offers the following advanced order types:  Stop-Loss, Take Profit, and Trailing Stop. Note that all the stop order
types are triggered by the mark price of the relevant market.Mark price is the median of bid, ask, and last. By
implication, this may lead to a price outside the order book, which can lead to a conversion and execution as a market
order independent of the advanced order configuration.

**Retry until filled**

FTX offers the ability to retry all trigger market orders (Stop-Loss, Take Profit, and Trailing Stop) until they get
filled. When a trigger order becomes triggered, it's possible for the order it sends to fail: the account may not have
enough margin, price bands during sharp market moves might prevent market orders from matching against other orders,
etc. In these cases, it may be preferable to retry sending the triggered order until their overall triggered order size
is filled. Retried triggers will be sent when the standard conditions around mark price and trigger price are met. Note
that this might mean that if markets move, some of the fills on your retried orders might be at prices far away from
your trigger price. This is, in many cases, a feature: you might want to close your position regardless of price given
that your trigger condition is met. If this is not something you want, then instead use triggered limit orders, or turn
off the 'Retry' option.

**Reduce-only**

Reduce-only orders will stop retrying when your position is fully closed. If you send a reduce only order, it will only
trade if it would decrease your position size

### Stop Loss

When creating a stop-loss order, you directly input the desired trigger price. If you are buying, the order will get
sent when the market price exceeds your trigger price. If you are selling, the order will get sent when the market price
drops below your trigger price. It will be sent as a market order if you selected Stop Market. Otherwise, it will be
sent as a limit order at the limit price.

#### Buy Stop Loss (To close a short position)

* Stop-loss buy orders are sent when the market price exceeds their trigger price.
* Like a normal limit buy order, the limit price means buy at the limit price or lower.

#### Sell Stop Loss (To close a long position)

* Stop-loss sell orders are sent when the market price drops below their trigger price.
* Like any stop limit order, the limit price means sell either at the limit price or higher.

#### Fill

* Like normal limit orders, a complete or partial fill is not guaranteed and depends on the depth of the order book.

#### Implications

* Stop loss orders do not appear in the order book until the trigger price has been reached.
* When the trigger price has been reached, the configured order will be placed in the order book with the defined limit
  price.
* For a sell stop-loss order, the limit price must be above the trigger price to preserve execution as a limit order.
* When the trigger and limit price are close, (partial) fill may occur relatively quickly in a highly liquid market.
* When the trigger and limit price differ by a wider margin, (partial) fill may not occur at all in case of a market
  reversal
* When the trigger price equals the limit price, there is a real chance that the limit order may get converted and
  filled as a market order because the limit price might be already outside the bid order book. Therefore, in both
  cases, the limit price must be unequal to the trigger price to prevent execution as a market order.
* For a buy stop-loss, the limit price must be below the trigger price to preserve execution as a limit order

#### Application

* For long positions, a sell stop-loss order limits loss at the limit price.
* For short positions, a buy stop-loss order limits loss at the limit price.

### Take Profit

A take profit order gets triggered in opposite to a stop-loss order. That means, if you aim to buy, the order will get
sent when the market price drops below your trigger price. If you aim to sell at a specific price, the order will get
sent when the market price exceeds above your trigger price. The order will be sent as a market order if you selected
Take profit. Otherwise, if you selected Take profit limit, it will be sent as a limit order at the limit price.

#### Buy Take Profit (To close a short position)

* For an open short position, the matching buy (back) take profit order will be placed when the market price drops below
  your trigger price.
* Like a normal buy limit order, the limit price means buy at the limit price or lower.

#### Sell Take Profit (To close a long position)

* For an open long position, the matching sell take profit order will be place with the limit price when the market
  price exceeds the trigger price.
* Like a normal sell limit order, the limit price means sell either at the limit price or higher.

#### Fill:

* Like normal limit orders, a complete or partial fill of any take profit limit order is not guaranteed and depends on
  the depth of the order book.

#### Implications:

* Take profit orders do not appear in the order book until the trigger price has been reached.
* When the trigger price has been reached, the configured order will be placed in the order book with the defined limit
  price.
* Similar to stop-loss orders, the buy-limit price must be below the trigger price and the sell-limit price must be
  above the trigger price to ensure order book placement and execution as a limit order.

#### Application

* For long positions, a sell take-profit closes the position.
* For short positions, a buy (back) take profit closes the position.

### Trailing Stop

Trailing stop orders are like stop-losses, but their trigger prices change as the market moves. Instead of directly
supplying the trigger price, you give a trail value. The trail value must be in fixed dollar and appropriate to the
market. Suppose you are buying. If the market price moves up by the trail value, your order will trigger. If the market
price moves down past the lowest point seen since you entered your order, then it'll only trigger if the price moves up
by trail value from that new lowest point.

If you are selling, the trail value must be negative.

#### Buy Trailing Stop

* If the market price moves up by the trail value, your order will trigger.
* If the market price moves down past the lowest point seen since you entered your order, then it'll only trigger if the
  price moves up by trail value from that new lowest point.

#### Sell Trailing Stop

* If the market price moves down by the trail value, your order will trigger.
* If the market price moves up past the highest point seen since you entered your order, then it'll only trigger if the
  price moves down again by trail value from that new highest point.

#### Fill

* It seems trailing stop gets executed as a market order because the trigger price gets determined automatically and as
  such a potential limit price would be equal the trigger price which usually converts into a market order.

#### Implications

* Trailing stop orders usually lead to adverse execution and excessive taker fees.

#### Application

* None. Market orders should not be used.

## Special Order Parameters

### Post Only (To Orderbook)

If you send a post only order, your order will not be allowed to become a market order.  
If the order would provide liquidity to the order book, it will be sent as a normal limit order, but if it would land
outside the order book it will be canceled instead.

This means that executed post-only orders only pay (lower) maker fees; they can never be charged (higher) taker fees.

### Immediate or cancel (IOC)

Immediate or cancel (IOC) orders are the opposite of post only: they can only take. If you send an IOC that would not
immediately trade, it will be canceled.

This means that IOC orders only pay (higher) taker fees; they can never be charged (lower) maker fees.

The FTX fee Table: https://help.ftx.com/hc/en-us/articles/360024479432-Fees

## Order Execution

An order that reached the exchange either executes at a favorable price and order type, or at an adverse price and order
type. Adverse execution may happen for any of the following reasons:

* Latency caused a delay so that the actual order price is already unfavorable and will be converted into a market order
* The order price was determined too close to the current price and the market has already moved into the opposite
  direction so an otherwise valid limit order will be converted into a market order.
* The order price was determined incorrectly and lead to a correct order type but with unfavorable execution

More specifically, the exact circumstances in which a buy or sell order executes adversely are outlined next.

### Adverse Order Execution:

* Buy: Above current price regardless of limit.
    - Cause: Market pulled back after the order was send but before execution thus lead to an unfavorable fill i.e.
      conversion into market order.
* Sell: Below current price regardless of limit
    - Cause: Market rose after the order was send but before the order was filled thus lead to adverse execution i.e.
      conversion into market order
* Latency between order issuance and execution may cause the discrepancy between order and execution price

In all three cases, the actual cause of adverse execution really is the discrepancy between current price and the actual
order price.

### Favorable Order Execution:

**Favorable Buy Order:**

* Below current price AND at or below the limit price
* That is any price in the buy order book below the zero index (first entry & closest to current price)
* Order must be Post Only

**Favorable Sell Order:**

* Price must be above current AND at or above the limit price
* That is any price below the zero index of the sell order book.
* Order must be Post Only

**Favorable Price:**

* Because not all prices from the order book can get filled, the best price that can be filled needs to be determined
* The best fillable price is usually the order book price with the largest order size b/c these get cleared out soon.

**Favorable Fill:**

* Order placement immediately after order price determination
* Limit price still a bit away from current price when it reaches the exchange
* Near zero latency required to prevent unfavorable market moves
* Because order is post only, it will be issued to the order book
* If the order gets canceled due to unfavorable market move, it needs to be retried at the next best order book price
  until it gets filled.
* Do not use stop-loss or take profit orders as these do not guarantee proper post-only limit-only execution.

**Favorable Long Position Orders**

* Buy favorable to open long
* Sell favorable to close long

**Favorable Short Position Orders**

* Sell favorable to open short
* Buy (back) favorable to close short