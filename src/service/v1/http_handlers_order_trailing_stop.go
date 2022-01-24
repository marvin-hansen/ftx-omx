// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package v1

import (
	"ftx-omx/src/cfg"
	"ftx-omx/src/types"
	"ftx-omx/src/utils/crypto"
	"ftx-omx/src/utils/dbg"
	"log"
	"net/http"
)

func (s *Service) openTrailingStopOrderHandler(w http.ResponseWriter, r *http.Request) {
	msg := ""
	authKey := r.Header.Get("auth")
	if authKey != crypto.DecodeKey(cfg.GetOrderAuthKey()) {
		msg = "API Access denied!"
		DbgPrint(msg + "Incorrect API key!")
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return
	}

	DbgPrint("check if Order  service is switched on!")
	if s.state.orderOn == false {
		msg = "API Disabled"
		DbgPrint(msg)
		http.StatusText(http.StatusServiceUnavailable)
		_, _ = w.Write([]byte(orderServiceUnavailableError))
		return
	}

	DbgPrint("decode JSON")
	var placeOrder *types.PlaceTrailingStopOrder
	err := decodeJSONBody(w, r, &placeOrder)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dbg.DbgStringerObject(dbgMain, "placeBookOrder struct", placeOrder)
	DbgPrint("place stop limit order!")
	s.openTrailingStopOrder(w, placeOrder)
}
