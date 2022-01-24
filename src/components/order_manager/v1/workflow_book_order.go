// Copyright (c) 2022-2022. Marvin Hansen | marvin.hansen@gmail.com

package v1

import (
	"fmt"
	"github.com/go-numb/go-ftx/rest"
	"github.com/go-numb/go-ftx/rest/private/orders"
	"github.com/go-numb/go-ftx/types"
	"log"
	"strconv"
	t "web_socket/src/types"
	"web_socket/src/utils/dbg"
)

func (c *OrderManager) PlaceOpenBookOrder(clientID string, placeOrder *t.PlaceBookOrder, auto bool) (orderID string, ok bool, msg string) {
	c.state.Lock()
	defer c.state.Unlock()

	orderResponse, ordRespMsg, ordRspErr := c.placeBookOrder(clientID, placeOrder, auto)
	if ordRspErr != nil {
		dbg.DbgOrderResponse(dbgErr, orderResponse)
		DbgPrint(ordRespMsg)
		dbg.LogError(ordRspErr)
		//
		orderID = "nil. See error!"
		ok = false
		msg = "Failed to place order in market. Abort!"
		return orderID, ok, msg
	}

	if auto {
		DbgPrint("Auto mode.")
		DbgPrint("Add the opening order to the map")
		//
		key := clientID + placeOrder.Market
		c.state.bookOrderMap.Set(key, orderResponse)
	}

	orderID = strconv.Itoa(orderResponse.ID)
	ok = true
	msg = "ok!"
	return orderID, ok, msg
}

func (c *OrderManager) PlaceStopLossBookOrder(clientID string, placeOrder *t.PlaceBookOrder, auto bool) (orderID string, ok bool, msg string) {
	// the only real difference here is the order price, which needs to be determined in the placeOrder anyway.
	// Set auto always to true to delete open order
	// Either PlaceClose or PlaceStopLoss gets called, but never both.
	return c.PlaceCloseBookOrder(clientID, placeOrder, auto)
}

func (c *OrderManager) PlaceCloseBookOrder(clientID string, placeOrder *t.PlaceBookOrder, auto bool) (orderID string, ok bool, msg string) {
	c.state.Lock()
	defer c.state.Unlock()

	if auto { // auto mode constructs close order from opening order
		DbgPrint("Auto mode.")
		DbgPrint("Load the opening order.")
		key := clientID + placeOrder.Market
		openOrder, exits := c.state.bookOrderMap.Get(key)
		if !exits {
			orderID = "Nil. See Error msg."
			ok = false
			msg = "Error: No matching opening order found; can't construct auto-close. Abort! "
			DbgPrint(msg)
			return orderID, ok, msg

		} else {
			closeOrder := c.getCloseBookOrderFromOpenOrder(placeOrder, openOrder)
			orderResponse, ordMsg, ordErr := c.placeBookOrder(clientID, closeOrder, true)
			if ordErr != nil {
				dbg.DbgOrderResponse(dbgErr, orderResponse)
				dbg.LogError(ordErr)
				//
				orderID = "nil. See error!"
				ok = false
				msg = ordMsg
				return orderID, ok, msg
			}

			DbgPrint(" Delete the opening order from the map")
			c.state.bookOrderMap.Delete(key)

			orderID = strconv.Itoa(orderResponse.ID)
			msg = "ok!"
			ok = true
			return orderID, ok, msg
		}

	} else { // manual mode constructed from place order only. Don't need to load & delete open order
		orderResponse, ordMsg, ordErr := c.placeBookOrder(clientID, placeOrder, false)
		if ordErr != nil {
			dbg.DbgOrderResponse(dbgErr, orderResponse)
			DbgPrint(ordMsg)
			dbg.LogError(ordErr)
			//
			orderID = "nil. See error!"
			msg = ordMsg
			ok = false
			return orderID, ok, msg
		}
		orderID = strconv.Itoa(orderResponse.ID)
		msg = "ok!"
		ok = true
		return orderID, ok, msg
	}
}

func (c *OrderManager) placeBookOrder(clientID string, placeOrder *t.PlaceBookOrder, autoSize bool) (resp *orders.ResponseForPlaceOrder, msg string, err error) {
	mtd := "[placeBookOrder]: "

	accountClient, clErr := c.dependencies.accountManager.GetAccount(clientID, placeOrder.Market)
	defer accountClient.HTTPC.CloseIdleConnections()

	if clErr != nil {
		msg = mtd + "Error getting client for account " + clientID
		DbgPrint(msg + clErr.Error())
		dbg.LogError(clErr)
		return nil, msg, clErr
	}
	order, reqMsg, reqErr := c.getBookOrder(accountClient, placeOrder, autoSize)
	if reqErr != nil {
		dbg.DbgRequestForPlaceOrder(dbgErr, mtd+reqMsg, order)
		msg = reqMsg
		return nil, msg, reqErr
	}

	// The price needs an update before order placement because the market may have moved since the arrival of the initial order request
	DbgPrint(mtd + " Pull latest order book price")
	limitPrice, priceErrMsg, priceErr := getOrderBookPrice(accountClient, placeOrder.Market, placeOrder.BookSide, placeOrder.BookPriceType)
	if priceErr != nil {
		log.Println(priceErrMsg)
		return nil, priceErrMsg, priceErr
	}

	DbgPrint(mtd + " update order to the latest order book price: " + strconv.Itoa(int(limitPrice)))
	order.Price = limitPrice

	DbgPrint(mtd + "Place book order in the market")
	dbg.DbgRequestForPlaceOrder(debug, mtd+priceErrMsg, order)
	orderResponse, orderErr := accountClient.PlaceOrder(order)
	if orderErr != nil { // error
		dbg.DbgOrderResponse(dbgErr, orderResponse)
		msg = mtd + "Error sending order. Abort! "
		DbgPrint(msg)
		dbg.LogError(orderErr)
		return nil, msg, orderErr

	} else { // no error
		dbg.DbgOrderResponse(debug, orderResponse)
		msg = "ok!"
		return orderResponse, msg, nil
	}
}

func (c *OrderManager) getBookOrder(client *rest.Client, placeOrder *t.PlaceBookOrder, autoSize bool) (o *orders.RequestForPlaceOrder, msg string, err error) {
	mtd := "getBookOrder: "

	DbgPrint(mtd + " Determines the exact order size dynamically! ")
	// Determines the exact order size dynamically i.e. when opening percentage size
	orderSize, sizeMsg, sizeErr := c.getBookOrderSize(client, placeOrder)
	if sizeErr != nil {
		return o, sizeMsg, sizeErr
	}

	DbgPrint(mtd + "orderSize: " + fmt.Sprintf("%.2f", orderSize))

	DbgPrint(mtd + " Construct request for Place order! ")
	o = &orders.RequestForPlaceOrder{
		Type:              types.LIMIT,
		Market:            placeOrder.Market,
		Side:              placeOrder.Side,
		Size:              orderSize,
		Ioc:               placeOrder.Ioc,
		ReduceOnly:        placeOrder.ReduceOnly,
		PostOnly:          placeOrder.PostOnly,
		RejectOnPriceBand: placeOrder.RejectOnPriceBand,
	}

	dbg.DbgRequestForPlaceOrder(debug, mtd+"dynamically sized order", o)

	msg = "ok!"
	return o, msg, nil
}
