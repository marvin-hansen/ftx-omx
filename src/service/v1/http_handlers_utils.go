// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package v1

import (
	"ftx-omx/src/types"
	"log"
	"net/http"
)

// ############################ Open & Close Limit order ############################

func (s *Service) openLimitOrder(w http.ResponseWriter, placeOrder *types.PlaceOrder, auto bool) {

	DbgPrint("Place open limit order")
	orderID, ok, msg := s.state.orderManager.PlaceOpenLimitOrder(placeOrder.ApiID, placeOrder, auto)
	if ok {
		DbgPrint(" Order placed. ID: " + orderID)
		_, _ = w.Write([]byte(msg))
	} else {
		log.Println(msg)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func (s *Service) closeLimitOrder(w http.ResponseWriter, placeOrder *types.PlaceOrder, autoMode bool) {

	DbgPrint("Place close limit order")
	orderID, ok, msg := s.state.orderManager.PlaceCloseLimitOrder(placeOrder.ApiID, placeOrder, autoMode)
	if ok {
		DbgPrint("Close Order placed ID: " + orderID)
		_, _ = w.Write([]byte(msg))
	} else {
		log.Println(msg)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

// ############################ Open & Close Book order ############################

func (s *Service) openBookOrder(w http.ResponseWriter, placeOrder *types.PlaceBookOrder, auto bool) {

	DbgPrint("Place open book order")
	orderID, ok, msg := s.state.orderManager.PlaceOpenBookOrder(placeOrder.ApiID, placeOrder, auto)
	if ok {
		DbgPrint(" Order placed. ID: " + orderID)
		_, _ = w.Write([]byte(msg))
	} else {
		log.Println(msg)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func (s *Service) closeBookOrder(w http.ResponseWriter, placeOrder *types.PlaceBookOrder, auto bool) {

	DbgPrint("Place close book order")
	orderID, ok, msg := s.state.orderManager.PlaceCloseBookOrder(placeOrder.ApiID, placeOrder, auto)
	if ok {
		DbgPrint("Close Order placed ID: " + orderID)
		_, _ = w.Write([]byte(msg))
	} else {
		log.Println(msg)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func (s *Service) stopLossBookOrder(w http.ResponseWriter, placeOrder *types.PlaceBookOrder, auto bool) {
	DbgPrint("Place stop Loss book order")
	orderID, ok, msg := s.state.orderManager.PlaceStopLossBookOrder(placeOrder.ApiID, placeOrder, auto)
	if ok {
		DbgPrint("Close Order placed ID: " + orderID)
		_, _ = w.Write([]byte(msg))
	} else {
		log.Println(msg)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

// ############################ Stop Limit & Take profit order ############################

func (s *Service) openStopLimitOrder(w http.ResponseWriter, placeOrder *types.PlaceTriggerOrder) {

	DbgPrint("Place stopLimit order")
	orderID, ok, msg := s.state.orderManager.PlaceStopLossOrder(placeOrder.ApiID, placeOrder)
	if ok {
		DbgPrint(" stopLimit order placed. ID: " + orderID)
		_, _ = w.Write([]byte(msg))
	} else {
		log.Println(msg)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func (s *Service) openTakeProfitOrder(w http.ResponseWriter, placeOrder *types.PlaceTriggerOrder) {

	DbgPrint("Place takeProfit order")
	orderID, ok, msg := s.state.orderManager.PlaceStopLossOrder(placeOrder.ApiID, placeOrder)
	if ok {
		DbgPrint(" takeProfit order placed. ID: " + orderID)
		_, _ = w.Write([]byte(msg))
	} else {
		log.Println(msg)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

// ############################ Trailing Stop order ############################

func (s *Service) openTrailingStopOrder(w http.ResponseWriter, placeOrder *types.PlaceTrailingStopOrder) {

	DbgPrint("Place trailing stop order")
	orderID, ok, msg := s.state.orderManager.PlaceTrailingStopOrder(placeOrder.ApiID, placeOrder)
	if ok {
		DbgPrint(" trailing stop order placed. ID: " + orderID)
		_, _ = w.Write([]byte(msg))
	} else {
		log.Println(msg)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
