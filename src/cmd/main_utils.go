// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package main

import (
	"encoding/json"
	"fmt"
	"web_socket/src/components/crypto_manager"
	"web_socket/src/types"
)

func generateAuthKey(length uint) {
	cm := crypto_manager.NewNewCryptoManager()

	keyText, ok := cm.GenerateKey(length)
	b64 := cm.EncodeBase64(keyText)
	if ok {
		println("New Auth key generated")
		println()
		println("Plaintext: " + keyText)
		println()
		println("Base64: " + b64)
		println()
	} else {
		println("Error generating new key!")
	}
}

func printPlaceOrderJsonReq() {
	buyOrder := getPlaceOrder(true, true)
	b, _ := json.Marshal(buyOrder)
	printJson(b, "BUY order JSON request")

	sellOrder := getPlaceOrder(false, false)
	b, _ = json.Marshal(sellOrder)
	printJson(b, "SELL order JSON request")
}

func printPlaceBookOrderJsonReq() {
	buyOrder := getPlaceBookOrder(true, true, false)
	b, _ := json.Marshal(buyOrder)
	printJson(b, "BUY order JSON request")

	sellOrder := getPlaceBookOrder(false, false, false)
	b, _ = json.Marshal(sellOrder)
	printJson(b, "SELL order JSON request")

	stopLossOrder := getPlaceBookOrder(false, false, true)
	b, _ = json.Marshal(stopLossOrder)
	printJson(b, "STOP LOSS order JSON request")
}

func printPlaceTriggerJsonReq() {
	sellOrder := getPlaceTriggerOrder(false, false)
	b, _ := json.Marshal(sellOrder)
	printJson(b, "SELL order JSON request")
}

func printPlaceTrailingJsonReq() {
	sellOrder := getPlaceTrailingStopOrder(false, false)
	b, _ := json.Marshal(sellOrder)
	printJson(b, "Trailing Stop Loss order JSON request")
}

func printJson(b []byte, msg string) {
	println()
	fmt.Println(msg)
	println("=========================")
	println()
	fmt.Println(string(b))
	println()
	println("=========================")
	println()
}

func getPlaceOrder(open, buy bool) (placeBookOrder types.PlaceOrder) {
	orderAction, orderSide, _ := getActionSide(open, buy, false)
	placeBookOrder = types.PlaceOrder{
		Auth:              "626aaf908056cc86b1d10ac355811497",
		ApiID:             "6tgrfg8wc2v8",
		Action:            orderAction,
		Market:            "ETH/USD",
		Type:              "limit",
		Side:              orderSide,
		PercentSize:       0.1,
		UsePercentSize:    true,
		ClientID:          "",
		ReduceOnly:        false,
		Ioc:               false,
		PostOnly:          true,
		RejectOnPriceBand: false,
	}
	return placeBookOrder
}

func getPlaceBookOrder(open, buy, stoploss bool) (placeOrder *types.PlaceBookOrder) {
	orderAction, orderSide, orderBookSide := getActionSide(open, buy, stoploss)
	bookPriceType := getBookPriceType(buy, stoploss)
	reduceOnly := getReduceOnly(buy, stoploss)

	placeOrder = &types.PlaceBookOrder{
		Auth:              "626aaf908056cc86b1d10ac355811497",
		ApiID:             "6tgrfg8wc2v8",
		Action:            orderAction,
		Market:            "ETH-PERP",
		Type:              "limit",
		Side:              orderSide,
		BookSide:          orderBookSide,
		BookPriceType:     bookPriceType,
		PercentSize:       0.1,
		ClientID:          "",
		ReduceOnly:        reduceOnly,
		Ioc:               false,
		PostOnly:          true,
		RejectOnPriceBand: false,
	}
	return placeOrder
}

func getPlaceTriggerOrder(open, buy bool) (placeOrder *types.PlaceTriggerOrder) {
	orderAction, orderSide, _ := getActionSide(open, buy, false)
	placeOrder = &types.PlaceTriggerOrder{
		Auth:             "626aaf908056cc86b1d10ac355811497",
		ApiID:            "6tgrfg8wc2v8",
		Action:           orderAction,
		Market:           "ETH-PERP",
		Type:             "limit",
		Side:             orderSide,
		Size:             0,
		PercentSize:      0.1,
		UsePercentSize:   true,
		TriggerPrice:     3335,
		LimitPrice:       3337,
		ReduceOnly:       false,
		PostOnly:         true,
		RetryUntilFilled: false,
	}
	return placeOrder
}

func getPlaceTrailingStopOrder(open, buy bool) (placeOrder *types.PlaceTrailingStopOrder) {
	orderAction, orderSide, _ := getActionSide(open, buy, false)

	placeOrder = &types.PlaceTrailingStopOrder{
		Auth:             "626aaf908056cc86b1d10ac355811497",
		ApiID:            "6tgrfg8wc2v8",
		Action:           orderAction,
		Market:           "ETH-PERP",
		Type:             "limit",
		Side:             orderSide,
		Size:             0,
		PercentSize:      0.1,
		UsePercentSize:   true,
		TrailValue:       15,
		PostOnly:         true,
		ReduceOnly:       true,
		RetryUntilFilled: false,
	}
	return placeOrder
}

func getActionSide(open, buy, stoploss bool) (orderAction, orderSide, orderBookSide string) {

	if open {
		orderAction = "open"
	} else {
		orderAction = "close"
	}

	if stoploss {
		orderAction = "stoploss"
	}

	if buy {
		orderSide = "buy"
		orderBookSide = "bid"
	} else {
		orderSide = "sell"
		orderBookSide = "ask"
	}

	return orderAction, orderSide, orderBookSide
}

func getBookPriceType(buySell, stoploss bool) (bookPriceType string) {
	if buySell {
		bookPriceType = "LargestOrderSizePrice"
	}
	if stoploss {
		bookPriceType = "MidOrderBookPrice"
	} else {
		bookPriceType = "LargestOrderSizePrice"
	}
	return bookPriceType
}

func getReduceOnly(buy, stoploss bool) (roc bool) {

	if stoploss {
		roc = true
	}
	if buy {
		roc = false
	} else {
		roc = true
	}
	return roc
}
