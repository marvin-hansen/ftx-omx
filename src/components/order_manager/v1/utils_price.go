// Copyright (c) 2022-2022. Marvin Hansen | marvin.hansen@gmail.com

package v1

import (
	"github.com/go-numb/go-ftx/rest"
	"github.com/go-numb/go-ftx/rest/public/markets"
	"log"
	t "web_socket/src/types"
	"web_socket/src/utils/dbg"
)

func getPrice(client *rest.Client, productCode string) (price *t.Price, err error) {
	market, err := client.Markets(&markets.RequestForMarkets{ProductCode: productCode})
	if err != nil {
		dbg.LogError(err)
		return nil, err
	}
	m := (*market)[0]
	return t.NewPrice(m.Ask, m.Bid, m.Last), err
}

func getOrderBookPrice(client *rest.Client, productCode, bookSide, priceType string) (orderBookPrice float64, msg string, err error) {

	orderBook, err := client.Orderbook(&markets.RequestForOrderbook{ProductCode: productCode})
	if err != nil {
		dbg.LogError(err)
		msg = "error, cannot get order book from exchange. Abort!"
		log.Println(msg)
		return 0, msg, err
	}

	// Determining order book side remains invariant of determining the actual price
	var book [][]float64
	if bookSide == t.Bid {
		book = orderBook.Bids
	} else {
		book = orderBook.Asks
	}
	//dbg.DbgOrderBook(debug, book)

	// max idx of the order book
	bookLength := len(book)

	switch priceType {

	case t.LargestOrderSizePrice:
		var sizes []float64
		for _, v := range book {
			sizes = append(sizes, v[1])
		}

		_, maxIdx, _ := getMaxValue(sizes) // find the entry with the largest bid / ask size
		orderBookPrice = book[maxIdx][0]   // determine the price corresponding to the largest size
		msg = "ok"
		return orderBookPrice, msg, nil

	case t.SmallestOrderSizePrice:
		var sizes []float64
		for _, v := range book {
			sizes = append(sizes, v[1])
		}

		_, minIdx, _ := getMinValue(sizes) // find the entry with the smallest bid / ask size
		orderBookPrice = book[minIdx][0]   // determine the price corresponding to the smallest size

		msg = "ok"
		return orderBookPrice, msg, nil

	case t.LowestOrderBookPrice:
		var prices []float64
		for _, v := range book {
			prices = append(prices, v[0])
		}
		orderBookPrice, _, _ = getMinValue(prices)
		msg = "ok"
		return orderBookPrice, msg, nil

	case t.MidOrderBookPrice:
		midIdx := bookLength / 2
		orderBookPrice = book[midIdx][0]
		msg = "ok"
		return orderBookPrice, msg, nil

	case t.HighestOrderBookPrice:
		var prices []float64
		for _, v := range book {
			prices = append(prices, v[0])
		}
		orderBookPrice, _, _ = getMaxValue(prices)
		msg = "ok"
		return orderBookPrice, msg, nil

	case t.FirstOrderBookPrice:
		orderBookPrice = book[0][0]
		msg = "ok"
		return orderBookPrice, msg, nil

	case t.LastOrderBookPrice:
		orderBookPrice = book[bookLength][0]
		msg = "ok"
		return orderBookPrice, msg, nil

	default:
		// by default or if unknown, we return the price at the middle of the order book to ensure favorable execution
		midIdx := bookLength / 2
		orderBookPrice = book[midIdx][0]
		msg = "no or unknown order book price. Returning fifth entry by default"
		return orderBookPrice, msg, nil //errors.New(msg)
	}
}
