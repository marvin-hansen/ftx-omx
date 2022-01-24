// Copyright (c) 2022-2022. Marvin Hansen | marvin.hansen@gmail.com

package v1

import (
	t "ftx-omx/src/types"
	dbg2 "ftx-omx/src/utils/dbg"
	"github.com/go-numb/go-ftx/rest"
	"github.com/go-numb/go-ftx/rest/private/orders"
	"strconv"
)

func (c *OrderManager) PlaceTrailingStopOrder(clientID string, placeOrder *t.PlaceTrailingStopOrder) (orderID string, ok bool, msg string) {
	c.state.Lock()
	defer c.state.Unlock()

	accountClient, clErr := c.dependencies.accountManager.GetAccount(clientID, placeOrder.Market)
	defer accountClient.HTTPC.CloseIdleConnections()
	if clErr != nil {
		msg = "Error: Failed to get client for account " + clientID
		DbgPrint(msg + clErr.Error())
		dbg2.LogError(clErr)
		orderID = "Nil. See Error msg."
		return orderID, false, msg
	}

	DbgPrint("Construct Stop Loss order")
	orderReq, orderReqMsg, orderReqErr := c.getTrailingOrder(accountClient, placeOrder)
	if orderReqErr != nil {
		dbg2.DbgRequestForPlaceTrailingStopOrder(dbgErr, orderReq)
		DbgPrint(orderReqMsg)
		dbg2.LogError(orderReqErr)
		msg = "Error: failed to construct trailing stop order request"
		orderID = "Nil. See Error msg."
		return orderID, false, msg
	}

	DbgPrint("Place open order in the market")
	orderResponse, orderErr := accountClient.PlaceTriggerOrder(orderReq)
	if orderErr != nil {
		dbg2.DbgTriggerOrderResponse(dbgErr, orderResponse)
		msg = "Error: Failed sending order to exchange!"
		DbgPrint(msg)
		dbg2.LogError(orderErr)
		orderID = "Nil. See Error msg."
		return orderID, false, msg
	}

	msg = "ok!"
	orderID = strconv.Itoa(orderResponse.ID)
	return orderID, true, msg
}

func (c *OrderManager) getTrailingOrder(client *rest.Client, placeOrder *t.PlaceTrailingStopOrder) (o *orders.RequestForPlaceTriggerOrder, msg string, err error) {

	orderSize, sizeMsg, sizeErr := c.getTrailingOrderSize(client, placeOrder)
	if sizeErr != nil {
		return o, sizeMsg, sizeErr
	}

	o = &orders.RequestForPlaceTriggerOrder{
		Market:           placeOrder.Market,
		Type:             placeOrder.Type,
		Side:             placeOrder.Side,
		Size:             orderSize,
		TrailValue:       placeOrder.TrailValue,
		ReduceOnly:       placeOrder.ReduceOnly,
		RetryUntilFilled: placeOrder.RetryUntilFilled,
	}

	msg = "ok!"
	return o, msg, nil
}
