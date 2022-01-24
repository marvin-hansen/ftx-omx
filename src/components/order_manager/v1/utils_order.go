// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package v1

import (
	t "ftx-omx/src/types"
	"ftx-omx/src/utils/dbg"
	"github.com/go-numb/go-ftx/rest"
	"github.com/go-numb/go-ftx/rest/private/orders"
	"github.com/go-numb/go-ftx/types"
)

// getDynamicLimitOrder determines order size and order price dynamically based on the placed order parameters
func (c *OrderManager) getDynamicLimitOrder(clientID string, placeOrder *t.PlaceOrder) (o *orders.RequestForPlaceOrder, msg string, err error) {

	account, clErr := c.dependencies.accountManager.GetAccount(clientID, placeOrder.Market)
	defer account.HTTPC.CloseIdleConnections()
	if clErr != nil {
		msg = "Error getting client for account " + clientID
		DbgPrint(msg + clErr.Error())
		dbg.LogError(clErr)
		return o, msg, clErr
	}

	orderSize, sizeMsg, sizeErr := c.getLimitOrderSize(account, placeOrder)
	if sizeErr != nil {
		return o, sizeMsg, sizeErr
	}

	orderPrice, priceErr := getPrice(account, placeOrder.Market)
	if priceErr != nil {
		priceMsg := "Error: Cannot get order price from exchange. Abort!"
		orderSize = 0
		return o, priceMsg, priceErr
	}

	o = &orders.RequestForPlaceOrder{
		Type:       types.LIMIT,
		Market:     placeOrder.Market,
		Side:       placeOrder.Side,
		Price:      orderPrice.Last,
		Size:       orderSize,
		Ioc:        placeOrder.Ioc,
		ReduceOnly: placeOrder.ReduceOnly,
		PostOnly:   placeOrder.PostOnly,
	}

	msg = "ok!"
	return o, msg, nil
}

func (c *OrderManager) getFixedCloseOrder(placeOrder *t.PlaceOrder) (o *orders.RequestForPlaceOrder) {

	o = &orders.RequestForPlaceOrder{
		Type:              types.LIMIT,
		Market:            placeOrder.Market,
		Side:              placeOrder.Side,
		Price:             placeOrder.Price,
		Size:              placeOrder.Size,
		Ioc:               placeOrder.Ioc,
		ReduceOnly:        placeOrder.ReduceOnly,
		PostOnly:          placeOrder.PostOnly,
		RejectOnPriceBand: placeOrder.RejectOnPriceBand,
	}

	return o
}

func (c *OrderManager) getCloseOrderFromOpenOrder(placeOrder *t.PlaceOrder, openOrder *orders.ResponseForPlaceOrder) (o *orders.RequestForPlaceOrder) {

	closeSide, _ := getActionSide(openOrder.Side)
	o = &orders.RequestForPlaceOrder{
		Type:       types.LIMIT,
		Market:     openOrder.Market,
		Side:       closeSide,
		Price:      placeOrder.Price,
		Size:       openOrder.Size,
		ReduceOnly: placeOrder.ReduceOnly,
		PostOnly:   placeOrder.PostOnly,
	}

	return o
}

func (c *OrderManager) getTriggerOrderFromOpenOrder(placeOrder *t.PlaceTriggerOrder, openOrder *t.PlaceTriggerOrder) (o *orders.RequestForPlaceTriggerOrder, msg string, err error) {

	closeSide, _ := getActionSide(openOrder.Side)
	// PT & SL required both, trigger and (limit) order price
	o = &orders.RequestForPlaceTriggerOrder{
		Market:           openOrder.Market,
		Type:             openOrder.Type,
		Side:             closeSide,
		Size:             openOrder.Size,
		TriggerPrice:     placeOrder.TriggerPrice,
		OrderPrice:       placeOrder.LimitPrice,
		ReduceOnly:       placeOrder.ReduceOnly,
		RetryUntilFilled: placeOrder.RetryUntilFilled,
	}

	msg = "ok"
	return o, msg, nil
}

func (c *OrderManager) getTriggerOrder(client *rest.Client, placeOrder *t.PlaceTriggerOrder) (o *orders.RequestForPlaceTriggerOrder, msg string, err error) {

	orderSize, sizeMsg, sizeErr := c.getTriggerOrderSize(client, placeOrder)
	if sizeErr != nil {
		return o, sizeMsg, sizeErr
	}

	o = &orders.RequestForPlaceTriggerOrder{
		Market:           placeOrder.Market,
		Type:             placeOrder.Type,
		Side:             placeOrder.Side,
		Size:             orderSize,
		OrderPrice:       placeOrder.LimitPrice,
		TriggerPrice:     placeOrder.TriggerPrice,
		ReduceOnly:       placeOrder.ReduceOnly,
		RetryUntilFilled: placeOrder.RetryUntilFilled,
	}

	msg = "ok!"
	return o, msg, nil
}

func (c *OrderManager) getCloseBookOrderFromOpenOrder(placeOrder *t.PlaceBookOrder, openOrder *orders.ResponseForPlaceOrder) (o *t.PlaceBookOrder) {

	orderSide, orderBookSide := getActionSide(openOrder.Side)
	o = &t.PlaceBookOrder{
		Auth:          placeOrder.Auth,
		ApiID:         placeOrder.ApiID,
		Action:        placeOrder.Action,
		Market:        openOrder.Market,
		Type:          openOrder.Type,
		Side:          orderSide,
		BookSide:      orderBookSide,
		BookPriceType: placeOrder.BookPriceType,
		OrderSize:     openOrder.Size,
		PercentSize:   placeOrder.PercentSize,
		ClientID:      placeOrder.ClientID,
		ReduceOnly:    placeOrder.ReduceOnly,
		PostOnly:      placeOrder.PostOnly,
	}

	return o
}

func getActionSide(side string) (orderSide, orderBookSide string) {
	if side == "buy" {
		orderSide = "sell"
		orderBookSide = "ask"
	} else {
		orderSide = "buy"
		orderBookSide = "bid"
	}
	return orderSide, orderBookSide
}
