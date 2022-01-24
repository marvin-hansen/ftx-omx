// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package order_manager

import (
	"ftx-omx/src/components/account_manager"
	"ftx-omx/src/components/data_manager"
	v1 "ftx-omx/src/components/order_manager/v1"
	t "ftx-omx/src/types"
)

type OrderManager interface {
	CancelOrder(clientID, ticker string, orderID int) (ok bool, msg string)
	// PlaceOpenBookOrder opens a position by determining price from the order book. Set auto to true to enable auto-close.
	PlaceOpenBookOrder(clientID string, placeOrder *t.PlaceBookOrder, auto bool) (orderID string, ok bool, msg string)
	// PlaceCloseBookOrder closes position from openBookOrder. Set auto to true to construct close from open order
	PlaceCloseBookOrder(clientID string, placeOrder *t.PlaceBookOrder, auto bool) (orderID string, ok bool, msg string)
	// PlaceStopLossBookOrder closes position from openBookOrder. NO OCO: Call either PlaceClose or PlaceStopLoss, but never both. Set auto to true to construct close from open order
	PlaceStopLossBookOrder(clientID string, placeOrder *t.PlaceBookOrder, auto bool) (orderID string, ok bool, msg string)

	PlaceOpenLimitOrder(clientID string, placeOrder *t.PlaceOrder, auto bool) (orderID string, ok bool, msg string)
	PlaceCloseLimitOrder(clientID string, placeOrder *t.PlaceOrder, auto bool) (orderID string, ok bool, msg string)
	PlaceTrailingStopOrder(clientID string, placeOrder *t.PlaceTrailingStopOrder) (orderID string, ok bool, msg string)
	PlaceProfitTakerOrder(clientID string, placeOrder *t.PlaceTriggerOrder) (orderID string, ok bool, msg string)
	PlaceStopLossOrder(clientID string, placeOrder *t.PlaceTriggerOrder) (orderID string, ok bool, msg string)
	ResetAllOrderMap()
	ResetLimitOrderMap()
	ResetBookOrderMap()
}

func NewOrderManager(clientManager account_manager.AccountManager, dataManager data_manager.DataManager) (orderManager OrderManager) {
	orderManager = v1.NewOrderManager(clientManager, dataManager)
	return orderManager
}
