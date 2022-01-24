// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package v1

import (
	"github.com/go-numb/go-ftx/rest/private/orders"
	"strconv"
	t "web_socket/src/types"
	dbg2 "web_socket/src/utils/dbg"
)

func (c *OrderManager) CancelOrder(clientID, ticker string, orderID int) (ok bool, msg string) {
	mtd := "CancelOrder: "
	c.state.Lock()
	defer c.state.Unlock()

	account, clErr := c.dependencies.accountManager.GetAccount(clientID, ticker)
	defer account.HTTPC.CloseIdleConnections()
	if clErr != nil {
		msg = mtd + "Error getting client for account " + clientID
		DbgPrint(msg)
		dbg2.LogError(clErr)
		return false, msg
	}

	_, err := account.CancelByID(&orders.RequestForCancelByID{OrderID: orderID})
	if err != nil {
		msg = mtd + "Error canceling order "
		DbgPrint(msg)
		dbg2.LogError(err)
		return false, msg
	} else {
		msg = "OK!"
		DbgPrint(mtd + msg)
		return true, msg
	}
}

func (c *OrderManager) PlaceOpenLimitOrder(clientID string, placeOrder *t.PlaceOrder, auto bool) (orderID string, ok bool, msg string) {
	mtd := "PlaceOpenLimitOrder: "
	c.state.Lock()
	defer c.state.Unlock()

	DbgPrint(mtd + "Construct open order")
	orderReq, orderReqMsg, orderReqErr := c.getDynamicLimitOrder(clientID, placeOrder)
	if orderReqErr != nil {
		dbg2.DbgRequestForPlaceOrder(dbgErr, mtd+orderReqMsg, orderReq)
		DbgPrint(mtd + orderReqMsg)
		dbg2.LogError(orderReqErr)
		orderID = "Nil. See Error msg."
		ok = false
		msg = orderReqMsg
		DbgPrint(mtd + msg)
		return orderID, ok, msg
	}

	DbgPrint(mtd + "Place open order in the market")
	orderResponse, orderRespMsg, orderErr := c.placeLimitOrder(clientID, orderReq)
	if orderErr != nil {
		DbgPrint(orderRespMsg)
		dbg2.LogError(orderErr)
		dbg2.DbgOrderResponse(dbgErr, orderResponse)
		orderID = "Nil. See Error msg."
		ok = false
		msg = orderRespMsg
		return orderID, ok, msg

	} else {
		if auto {
			DbgPrint(mtd + "Auto mode.")
			DbgPrint(mtd + "Add the opening order to the map")
			key := clientID + placeOrder.Market
			c.state.orderMap.Set(key, orderResponse)
		} // if it's a manual order, we don't save the open order

		orderID = strconv.Itoa(orderResponse.ID)
		ok = true
		msg = "ok!"
		DbgPrint(mtd + msg)
		return orderID, ok, msg
	}
}

func (c *OrderManager) PlaceCloseLimitOrder(clientID string, placeOrder *t.PlaceOrder, auto bool) (orderID string, ok bool, msg string) {
	mtd := "PlaceCloseLimitOrder: "
	c.state.Lock()
	defer c.state.Unlock()

	var orderRequest *orders.RequestForPlaceOrder
	key := clientID + placeOrder.Market

	if auto {
		DbgPrint(mtd + "Auto mode.")

		DbgPrint(mtd + "Load the opening order.")
		openOrder, exits := c.state.orderMap.Get(key)
		if !exits {
			orderID = "Nil. See Error msg."
			ok = false
			msg = mtd + "Error: No matching opening order found; can't construct auto-close. Error! "
			DbgPrint(msg)
			return orderID, ok, msg
		} else {
			// open order exits
			DbgPrint(mtd + "Construct close order from matching open order")
			orderRequest = c.getCloseOrderFromOpenOrder(placeOrder, openOrder)
		}

	} else { // manual mode
		DbgPrint(mtd + "Manual mode.")
		DbgPrint(mtd + "Construct fixed close order")
		orderRequest = c.getFixedCloseOrder(placeOrder)
	}

	DbgPrint(mtd + "Place close order")
	orderResponse, orderMsg, orderErr := c.placeLimitOrder(clientID, orderRequest)
	if orderErr != nil { // error
		dbg2.DbgRequestForPlaceOrder(dbgErr, mtd+orderMsg, orderRequest)
		dbg2.DbgOrderResponse(dbgErr, orderResponse)
		DbgPrint(mtd + orderMsg)
		dbg2.LogError(orderErr)

		orderID = "Nil. See Error msg."
		ok = false
		msg = orderMsg
		return orderID, ok, msg

	} else {
		if auto {
			DbgPrint(mtd + " Delete the opening order from the map")
			c.state.orderMap.Delete(key)
		}

		orderID = strconv.Itoa(orderResponse.ID)
		ok = true
		msg = "ok!"
		DbgPrint(mtd + msg)
		return orderID, ok, msg
	}
}

func (c *OrderManager) placeLimitOrder(clientID string, order *orders.RequestForPlaceOrder) (resp *orders.ResponseForPlaceOrder, msg string, err error) {
	mtd := "placeLimitOrder: "
	accountClient, clErr := c.dependencies.accountManager.GetAccount(clientID, order.Market)
	defer accountClient.HTTPC.CloseIdleConnections()
	if clErr != nil {
		msg = mtd + "Error getting client for account " + clientID
		DbgPrint(msg)
		dbg2.LogError(clErr)
		return nil, msg, clErr
	}

	DbgPrint(mtd + "Place book order in the market")
	orderResponse, OrderErr := accountClient.PlaceOrder(order)
	if OrderErr != nil {
		msg = mtd + "Error sending order "
		DbgPrint(msg)
		dbg2.LogError(OrderErr)
		dbg2.DbgOrderResponse(dbgErr, orderResponse)
		return nil, msg, OrderErr

	} else {
		msg = "ok!"
		DbgPrint(mtd + msg)
		return orderResponse, msg, nil
	}
}
