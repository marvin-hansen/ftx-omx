// Copyright (c) 2022-2022. Marvin Hansen | marvin.hansen@gmail.com

package v1

import (
	"ftx-omx/src/cfg"
	"ftx-omx/src/types"
	"ftx-omx/src/utils/crypto"
	dbg2 "ftx-omx/src/utils/dbg"
	"log"
	"net/http"
)

// order is a single webhook for Tradingview integration
func (s *Service) limitOrderHandler(w http.ResponseWriter, r *http.Request) {
	msg := ""
	DbgPrint("check if Order  service is switched on!")
	if s.state.orderOn == false {
		msg = "API Disabled"
		DbgPrint(msg)
		http.StatusText(http.StatusServiceUnavailable)
		_, _ = w.Write([]byte(orderServiceUnavailableError))
		return
	}

	// DO NOT DBG or PRINT JSON BODY PRIOR TO DECODE OTHERWISE DECODE FAILS WITH AN EMPTY BODY ERROR
	DbgPrint("decode JSON")
	var placeOrder *types.PlaceOrder
	err := decodeJSONBody(w, r, &placeOrder)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Tradingview can't send an auth http header, thus the auth key must be extracted from the JSON message
	authKey := placeOrder.Auth
	if authKey == "" {
		msg = "API KEY MISSING!"
		DbgPrint(msg)
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return
	}

	if authKey != crypto.DecodeKey(cfg.GetOrderAuthKey()) {
		msg = "API Access denied!"
		DbgPrint(msg + "Incorrect API key!")
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return
	}

	dbg2.DbgStringerObject(dbgMain, "placeBookOrder struct", placeOrder)

	DbgPrint("Determine Action: " + placeOrder.Action)

	if placeOrder.Action == "open" {
		DbgPrint(" * Place open order:")
		s.openLimitOrder(w, placeOrder, true)
		return
	}

	if placeOrder.Action == "close" {
		DbgPrint(" * Place close order:")
		s.closeLimitOrder(w, placeOrder, true)
		return
	}

	// catch all in case the request contains anything else than open or close.
	msg = "No order action given! Only use open or close as action!"
	DbgPrint(msg)
	http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	_, _ = w.Write([]byte(msg))
	return
}

func (s *Service) bookOrderHandler(w http.ResponseWriter, r *http.Request) {
	mtd := "bookOrderHandler: "
	msg := ""
	DbgPrint(mtd + "check if Order  service is switched on!")
	if s.state.orderOn == false {
		msg = "API Disabled"
		DbgPrint(mtd + msg)
		http.StatusText(http.StatusServiceUnavailable)
		_, _ = w.Write([]byte(orderServiceUnavailableError))
		return
	}

	// DO NOT DBG or PRINT JSON BODY PRIOR OTHERWISE DECODE FAILS WITH AN EMPTY BODY ERROR
	DbgPrint(mtd + "decode JSON")
	var placeOrder *types.PlaceBookOrder
	err := decodeJSONBody(w, r, &placeOrder)
	if err != nil {
		dbg2.DbgRequest(dbgMain, *r)
		println("")
		log.Println(err.Error())
		msg = "JSON decoding failed. Abort!"
		DbgPrint(mtd + msg)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dbg2.DbgStringerObject(dbgMain, mtd+" PlaceOrder decoded from JSON", placeOrder)

	// Tradingview can't send an auth http header, thus the auth key must be extracted from the JSON message
	authKey := placeOrder.Auth
	if authKey == "" {
		msg = "API KEY MISSING!"
		DbgPrint(mtd + msg)
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return
	}

	if authKey != crypto.DecodeKey(cfg.GetOrderAuthKey()) {
		msg = "API Access denied!"
		DbgPrint(mtd + msg + "Incorrect API key!")
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return
	}

	DbgPrint(mtd + "Determine Action: " + placeOrder.Action)
	if placeOrder.Action == "open" {
		DbgPrint(mtd + " * Place open order:")
		s.openBookOrder(w, placeOrder, true)
		return
	}

	if placeOrder.Action == "close" {
		DbgPrint(mtd + " * Place close order:")
		s.closeBookOrder(w, placeOrder, true)
		return
	}

	if placeOrder.Action == "stoploss" {
		DbgPrint(mtd + " * Place stop loss order:")
		s.stopLossBookOrder(w, placeOrder, true)
		return
	}

	// catch all in case the request contains anything else than open or close.
	msg = "No order action given! Only use open or close as action!"
	DbgPrint(msg)
	http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	_, _ = w.Write([]byte(msg))
	return
}
