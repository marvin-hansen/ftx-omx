// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package types

import "fmt"

type PlaceBookOrder struct {
	Auth              string  `json:"auth,"`                       // for the tradingview webhook that can't send auth headers
	ApiID             string  `json:"api_id"`                      // ID of the API to which the order will be send
	Action            string  `json:"action,"`                     // for the tradingview webhook that requires only one endpoint.
	Market            string  `json:"market"`                      // ethusd
	Type              string  `json:"type"`                        // limit, market
	Side              string  `json:"side"`                        // buy, sell
	BookSide          string  `json:"bookSide"`                    // ask, bid
	BookPriceType     string  `json:"bookPriceType"`               //
	PercentSize       float64 `json:"percentSize"`                 // 0.03 = 3%
	OrderSize         float64 `json:"orderSize,omitempty"`         // 0.03 = 3%
	OrderPrice        float64 `json:"orderPrice,omitempty"`        // 0.03 = 3%
	ClientID          string  `json:"clientId,omitempty"`          // client order id. optional
	ReduceOnly        bool    `json:"reduceOnly,omitempty"`        // only fills if it reduces position
	Ioc               bool    `json:"ioc,omitempty"`               // market order only. "immediate or cancel" false
	PostOnly          bool    `json:"postOnly,omitempty"`          // limit order only "Post to order book" true
	RejectOnPriceBand bool    `json:"rejectOnPriceBand,omitempty"` //
}

func (o PlaceBookOrder) String() string {
	return fmt.Sprintf("[PlaceBookOrder]: Api ID: %v, Action: %v, Market: %v, Type: %v, Side: %v,  BookSide: %v,  BookPriceType: %v, PercentSize: %v,  OrderSize: %v,  OrderPrice: %v, ClientID: %v, ReduceOnly: %v, Ioc: %v, PostOnly: %v, RejectOnPriceBand: %v",
		o.ApiID,
		o.Action,
		o.Market,
		o.Type,
		o.Side,
		o.BookSide,
		o.BookPriceType,
		o.PercentSize,
		o.OrderSize,
		o.OrderPrice,
		o.ClientID,
		o.ReduceOnly,
		o.Ioc,
		o.PostOnly,
		o.RejectOnPriceBand,
	)
}
