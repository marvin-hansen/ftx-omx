// Copyright (c) 2022-2022. Marvin Hansen | marvin.hansen@gmail.com
package v1

// https://docs.ftx.com/#place-trigger-order

import (
	t "ftx-omx/src/types"
	dbg2 "ftx-omx/src/utils/dbg"
	"strconv"
)

func (c *OrderManager) PlaceProfitTakerOrder(clientID string, placeOrder *t.PlaceTriggerOrder) (orderID string, ok bool, msg string) {
	c.state.Lock()
	defer c.state.Unlock()
	return c.placeTriggerOrder(clientID, placeOrder)
}

func (c *OrderManager) PlaceStopLossOrder(clientID string, placeOrder *t.PlaceTriggerOrder) (orderID string, ok bool, msg string) {
	c.state.Lock()
	defer c.state.Unlock()
	return c.placeTriggerOrder(clientID, placeOrder)
}

func (c *OrderManager) placeTriggerOrder(clientID string, placeOrder *t.PlaceTriggerOrder) (orderID string, ok bool, msg string) {
	accountClient, clErr := c.dependencies.accountManager.GetAccount(clientID, placeOrder.Market)
	defer accountClient.HTTPC.CloseIdleConnections()
	if clErr != nil {
		msg = "Error: Failed to get client for account " + clientID
		DbgPrint(msg)
		dbg2.LogError(clErr)
		orderID = "Nil. See Error msg."
		ok = false
		return orderID, ok, msg
	}

	DbgPrint("Construct Stop Loss order request")
	orderReq, orderReqMsg, orderReqErr := c.getTriggerOrder(accountClient, placeOrder)
	if orderReqErr != nil {
		dbg2.DbgRequestForPlaceTriggerOrder(dbgErr, orderReq)
		DbgPrint(orderReqMsg)
		dbg2.LogError(orderReqErr)
		orderID = "Nil. See Error msg."
		ok = false
		msg = "Error: failed to construct trigger order request."
		return orderID, ok, msg
	}

	DbgPrint("Place open order in the market")
	orderResponse, err := accountClient.PlaceTriggerOrder(orderReq)
	if err != nil {
		msg = "Error: Failed sending order to exchange!"
		DbgPrint(msg + err.Error())
		dbg2.LogError(err)
		dbg2.DbgTriggerOrderResponse(dbgErr, orderResponse)
		orderID = "Nil. See Error msg."
		ok = false
		return orderID, ok, msg
	}

	orderID = strconv.Itoa(orderResponse.ID)
	ok = true
	msg = "ok!"
	return orderID, ok, msg
}
