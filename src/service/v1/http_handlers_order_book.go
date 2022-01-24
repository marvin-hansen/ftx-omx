// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package v1

import (
	"ftx-omx/src/cfg"
	"ftx-omx/src/types"
	"ftx-omx/src/utils/crypto"
	"ftx-omx/src/utils/dbg"
	"log"
	"net/http"
)

func (s *Service) openBookOrderHandler(w http.ResponseWriter, r *http.Request) {
	mtd := "openBookOrderHandler: "
	msg := ""
	authKey := r.Header.Get("auth")
	if authKey != crypto.DecodeKey(cfg.GetOrderAuthKey()) {
		msg = "API Access denied!"
		DbgPrint(msg + "Incorrect API key!")
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return
	}

	DbgPrint(mtd + "check if Order  service is switched on!")
	if s.state.orderOn == false {
		msg = "API Disabled"
		DbgPrint(msg)
		http.StatusText(http.StatusServiceUnavailable)
		_, _ = w.Write([]byte(orderServiceUnavailableError))
		return
	}

	DbgPrint(mtd + "decode JSON")
	var placeOrder *types.PlaceBookOrder
	err := decodeJSONBody(w, r, &placeOrder)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	dbg.DbgStringerObject(dbgMain, mtd+" PlaceOrder decoded from JSON", placeOrder)

	DbgPrint(mtd + "place open book order!")
	s.openBookOrder(w, placeOrder, false)
}

func (s *Service) closeBookOrderHandler(w http.ResponseWriter, r *http.Request) {
	mtd := "closeBookOrderHandler: "
	msg := ""
	authKey := r.Header.Get("auth")
	if authKey != crypto.DecodeKey(cfg.GetOrderAuthKey()) {
		msg = "API Access denied!"
		DbgPrint(mtd + msg + "Incorrect API key!")
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return
	}

	DbgPrint(mtd + "check if Order  service is switched on!")
	if s.state.orderOn == false {
		msg = "API Disabled"
		DbgPrint(msg)
		http.StatusText(http.StatusServiceUnavailable)
		_, _ = w.Write([]byte(orderServiceUnavailableError))
		return
	}

	DbgPrint(mtd + "decode JSON")
	var placeOrder *types.PlaceBookOrder
	err := decodeJSONBody(w, r, &placeOrder)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	DbgPrint(mtd + "place close book order!")
	s.closeBookOrder(w, placeOrder, false)
}

func (s *Service) stopLossBookOrderHandler(w http.ResponseWriter, r *http.Request) {
	mtd := "stopLossBookOrderHandler: "
	msg := ""
	authKey := r.Header.Get("auth")
	if authKey != crypto.DecodeKey(cfg.GetOrderAuthKey()) {
		msg = "API Access denied!"
		DbgPrint(mtd + msg + "Incorrect API key!")
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return
	}

	DbgPrint(mtd + "check if Order  service is switched on!")
	if s.state.orderOn == false {
		msg = "API Disabled"
		DbgPrint(msg)
		http.StatusText(http.StatusServiceUnavailable)
		_, _ = w.Write([]byte(orderServiceUnavailableError))
		return
	}

	DbgPrint(mtd + "decode JSON")
	var placeOrder *types.PlaceBookOrder
	err := decodeJSONBody(w, r, &placeOrder)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	DbgPrint(mtd + "place stop loss book order!")
	s.stopLossBookOrder(w, placeOrder, false)
}
