// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package v1

import (
	"errors"
	t "ftx-omx/src/types"
	"github.com/go-numb/go-ftx/rest"
	"github.com/go-numb/go-ftx/rest/private/account"
	"log"
)

type SizeParam struct {
	Market         string
	UsePercentSize bool
	PercentSize    float64
	FixedSize      float64
	BookPrice      bool
	BookSide       string
	BookPriceType  string
}

func (c *OrderManager) getBookOrderSize(client *rest.Client, placeOrder *t.PlaceBookOrder) (orderSize float64, msg string, err error) {
	sizeParam := &SizeParam{
		Market:         placeOrder.Market,
		UsePercentSize: true,
		PercentSize:    placeOrder.PercentSize,
		FixedSize:      placeOrder.OrderPrice,
		BookPrice:      true,
		BookSide:       placeOrder.BookSide,
		BookPriceType:  placeOrder.BookPriceType,
	}
	return c.getOrderSize(client, sizeParam)
}

func (c *OrderManager) getLimitOrderSize(client *rest.Client, placeOrder *t.PlaceOrder) (orderSize float64, msg string, err error) {
	sizeParam := &SizeParam{
		Market:         placeOrder.Market,
		UsePercentSize: placeOrder.UsePercentSize,
		PercentSize:    placeOrder.PercentSize,
		FixedSize:      placeOrder.Size,
		BookPrice:      false,
	}
	return c.getOrderSize(client, sizeParam)
}

func (c *OrderManager) getTriggerOrderSize(client *rest.Client, placeOrder *t.PlaceTriggerOrder) (triggerOrderSize float64, msg string, err error) {
	sizeParam := &SizeParam{
		Market:         placeOrder.Market,
		UsePercentSize: placeOrder.UsePercentSize,
		PercentSize:    placeOrder.PercentSize,
		FixedSize:      placeOrder.Size,
		BookPrice:      false,
	}
	return c.getOrderSize(client, sizeParam)
}

func (c *OrderManager) getTrailingOrderSize(client *rest.Client, placeOrder *t.PlaceTrailingStopOrder) (triggerOrderSize float64, msg string, err error) {
	sizeParam := &SizeParam{
		Market:         placeOrder.Market,
		UsePercentSize: placeOrder.UsePercentSize,
		PercentSize:    placeOrder.PercentSize,
		FixedSize:      placeOrder.Size,
		BookPrice:      false,
	}
	return c.getOrderSize(client, sizeParam)
}

func (c *OrderManager) getOrderSize(client *rest.Client, sizeParam *SizeParam) (orderSize float64, msg string, err error) {

	if sizeParam.UsePercentSize {
		var price float64
		if sizeParam.BookPrice {
			bookPrice, errMsg, priceErr := getOrderBookPrice(client, sizeParam.Market, sizeParam.BookSide, sizeParam.BookPriceType)
			if priceErr != nil {
				log.Println(errMsg)
				return orderSize, errMsg, priceErr
			}

			price = bookPrice

		} else {
			orderPrice, priceErr := getPrice(client, sizeParam.Market)
			if priceErr != nil {
				msg = "Error: Cannot get order price from exchange. Abort!"
				orderSize = 0
				return orderSize, msg, priceErr
			}
			price = orderPrice.Last
		}

		info, accErr := client.Information(&account.RequestForInformation{})
		if accErr != nil {
			msg = "Error: Cannot get account information from exchange. Abort!"
			orderSize = 0
			return orderSize, msg, accErr
		}

		percent := sizeParam.PercentSize
		lastPrice := price
		freeCollateral := info.FreeCollateral
		percentSize := (freeCollateral / lastPrice) * percent

		percentDollarValue := percentSize * lastPrice
		if percentDollarValue < 2.5 {
			msg = "Error: Position size to small. Abort"
			orderSize = 0
			return orderSize, msg, errors.New(msg)
		}

		msg = "ok"
		orderSize = percentSize
		return orderSize, msg, nil

	} else { // fixed size order from request
		msg = "ok"
		orderSize = sizeParam.FixedSize
		return orderSize, msg, nil
	}
}
