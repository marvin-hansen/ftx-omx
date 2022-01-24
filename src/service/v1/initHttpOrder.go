// Copyright (c) 2022-2022. Marvin Hansen | marvin.hansen@gmail.com

package v1

import "net/http"

func (s *Service) initOrderApi() {

	// ********************
	// Tradingview Endpoint
	// ********************

	// Combined limit order endpoint to open / close positions using only best effort limit orders
	// limit order handler. Note, execution as a limit order is not guaranteed in case of an adverse market move!
	//
	// BUY TO OPEN
	// curl -X http://localhost/order/limit POST -H 'Content-Type: application/json; charset=utf-8' -d '{"auth":"626aaf908056cc86b1d10ac355811497","api_id":"6tgrfg8wc2v8","action":"open","market":"ETH/USD","type":"limit","side":"buy","price":3132,"size":0,"percentSize":0.1,"usePercentSize":true}'
	//
	// SELL TO  CLOSE
	// curl -X POST http://localhost/order/limit -H 'Content-Type: application/json; charset=utf-8' -d {"auth":"626aaf908056cc86b1d10ac355811497","api_id":"6tgrfg8wc2v8","action":"close","market":"ETH/USD","type":"limit","side":"sell","price":3132,"size":0,"percentSize":0.1,"usePercentSize":true}'
	http.HandleFunc("/order/limit", s.limitOrderHandler)

	// Combined  book order endpoint to open / close positions using only guaranteed limit orders placed in the order book
	// Note, while the order is guaranteed to be placed as a limit order,
	// its execution is not guaranteed in case the market does not touch the limit price
	//
	// BUY TO OPEN
	// curl -X POST http://localhost/order/book -H 'Content-Type: application/json; charset=utf-8' -d  '{"auth":"626aaf908056cc86b1d10ac355811497","api_id":"6tgrfg8wc2v8","action":"open","market":"ETH-PERP","type":"limit","side":"buy","bookSide":"bid","bookPriceType":"LargestOrderSizePrice","percentSize":0.1,"reduceOnly":false,"ioc":false,"postOnly":true}'
	//
	// SELL TO  CLOSE
	// curl -X POST http://localhost/order/book -H 'Content-Type: application/json; charset=utf-8' -d  '{"auth":"626aaf908056cc86b1d10ac355811497","api_id":"6tgrfg8wc2v8","action":"close","market":"ETH-PERP","type":"limit","side":"sell","bookSide":"ask","bookPriceType":"LargestOrderSizePrice","percentSize":0.1,"reduceOnly":true,"ioc":false,"postOnly":true}'
	//
	// STOP LOSS TO CLOSE
	// curl -X POST http://localhost/order/book -H 'Content-Type: application/json; charset=utf-8' -d '{"auth":"626aaf908056cc86b1d10ac355811497","api_id":"6tgrfg8wc2v8","action":"stoploss","market":"ETH-PERP","type":"limit","side":"sell","bookSide":"ask","bookPriceType":"MidOrderBookPrice","percentSize":0.1,"reduceOnly":true,"ioc":false,"postOnly":true}'
	http.HandleFunc("/order/book", s.bookOrderHandler)

	// ********************
	// Limit Order Endpoint
	// ********************

	// Place open limit order
	// curl -X POST http://localhost/order/limit/open  -H 'auth: KEY' -H 'Content-Type: application/json' -d '{ \"api_id\": \"API_ID\", \"market\": \"ETH/USD\", \"type\": \"limit\", \"side\": \"buy\", \"price\": 3982,\"size\": 0.004}'
	http.HandleFunc("/order/limit/open", s.openLimitOrderHandler)

	// Place close limit order
	// curl -X POST  http://localhost/order/limit/close -H 'auth: KEY' -H 'Content-Type: application/json' -d '{ \"api_id\": \"API_ID\", \"market\": \"ETH/USD\", \"type\": \"limit\", \"side\": \"sell\", \"price\": 3982,\"size\": 0.004}'
	http.HandleFunc("/order/limit/close", s.closeLimitOrderHandler)

	// ********************
	// Stop Order Endpoint
	// ********************

	// Place take profit order
	// curl -X POST http://localhost/order/stop/takeprofit -H 'auth: KEY' -H 'Content-Type: application/json; charset=utf-8' -d '{"api_id":"6tgrfg8wc2v8","action":"close","market":"ETH-PERP","type":"limit","side":"sell","size":0,"percentSize":0.1,"usePercentSize":true,"triggerPrice":3335,"orderPrice":3337,"postOnly":true}'
	http.HandleFunc("/order/stop/takeprofit", s.openTakeProfitOrderHandler)

	// Place stop limit order
	// curl -X POST http://localhost/order/stop/stoplimit -H 'auth: KEY' -H 'Content-Type: application/json; charset=utf-8' -d '{"api_id":"6tgrfg8wc2v8","action":"close","market":"ETH-PERP","type":"limit","side":"sell","size":0,"percentSize":0.1,"usePercentSize":true,"triggerPrice":3335,"orderPrice":3337,"postOnly":true}'
	http.HandleFunc("/order/stop/stoplimit", s.openStopLimitOrderHandler)

	// Place trailing stop loss order
	// curl -X POST http://localhost/order/stop/trailingstop -H 'auth: KEY'-H 'Content-Type: application/json; charset=utf-8' --d '{"api_id":"6tgrfg8wc2v8","action":"close","market":"ETH-PERP","type":"limit","side":"sell","size":0,"percentSize":0.1,"usePercentSize":true,"trailValue":15,"postOnly":true,"reduceOnly":true}'
	http.HandleFunc("/order/stop/trailingstop", s.openTrailingStopOrderHandler)

	// ********************
	// Book Order Endpoint
	// ********************

	// Place open limit order in the order book
	// curl -X POST http://localhost/order/book/open -H 'auth: KEY' -H 'Content-Type: application/json; charset=utf-8' -d  '{"api_id":"6tgrfg8wc2v8","action":"open","market":"ETH-PERP","type":"limit","side":"buy","bookSide":"bid","bookPriceType":"LargestOrderSizePrice","percentSize":0.1,"reduceOnly":false,"ioc":false,"postOnly":true}'
	http.HandleFunc("/order/book/open", s.openBookOrderHandler)

	// Place close limit order in the order book
	// curl -X POST http://localhost/order/book/close -H 'auth: KEY' -H 'Content-Type: application/json; charset=utf-8' -d  '{"api_id":"6tgrfg8wc2v8","action":"close","market":"ETH-PERP","type":"limit","side":"sell","bookSide":"ask","bookPriceType":"LargestOrderSizePrice","percentSize":0.1,"postOnly":true}' http://localhost/order/book
	http.HandleFunc("/order/book/close", s.closeBookOrderHandler)

	// Place stop loss limit order in the order book
	// curl -X POST http://localhost/order/book/stoploss -H 'auth: KEY' -H 'Content-Type: application/json; charset=utf-8' -d '{"api_id":"6tgrfg8wc2v8","action":"stoploss","market":"ETH-PERP","type":"limit","side":"sell","bookSide":"ask","bookPriceType":"MidOrderBookPrice","percentSize":0.1,"reduceOnly":true,"ioc":false,"postOnly":true}'
	http.HandleFunc("/order/book/stoploss", s.stopLossBookOrderHandler)

	// ********************
	// Reset Endpoint
	// ********************

	// Resets all order cache to resolve problem of missing open order
	//curl -X POST http://localhost/order/reset  -H 'auth: KEY'
	http.HandleFunc("/order/reset", s.resetAllOrderMapHandler)

	// Resets limit order cache to resolve problem of missing open order
	// curl -X POST  http://localhost/order/limit/reset -H 'auth: KEY'
	http.HandleFunc("/order/limit/reset", s.resetLimitOrderMapHandler)

	// Resets book order cache to resolve problem of missing open order
	// curl -X POST http://localhost/order/book/reset  -H 'auth: KEY'
	http.HandleFunc("/order/book/reset", s.resetBookOrderMapHandler)
}
