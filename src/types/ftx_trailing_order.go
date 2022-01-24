// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package types

import "fmt"

type PlaceTrailingStopOrder struct {
	Auth             string  `json:"auth,omitempty"` // for the tradingview webhook that can't send auth headers
	ApiID            string  `json:"api_id,omitempty"`
	Action           string  `json:"action,omitempty"` // for the tradingview webhook that requires only one endpoint.
	Market           string  `json:"market"`
	Type             string  `json:"type"` // Type stop, trailingStop, takeProfit. default is stop
	Side             string  `json:"side"` //"buy" or "sell"
	Size             float64 `json:"size"`
	PercentSize      float64 `json:"percentSize,omitempty"`      // order size as percentage of available capital
	UsePercentSize   bool    `json:"usePercentSize,omitempty"`   // whether to use percent size. default false.
	TrailValue       float64 `json:"trailValue,omitempty"`       // negative for "sell"; positive for "buy"
	PostOnly         bool    `json:"postOnly,omitempty"`         // optional; default is false
	ReduceOnly       bool    `json:"reduceOnly,omitempty"`       // optional; default is false
	RetryUntilFilled bool    `json:"retryUntilFilled,omitempty"` // Whether to keep re-triggering until filled. optional, default true for market orders
}

func (o PlaceTrailingStopOrder) String() string {
	return fmt.Sprintf("[PlaceTrailingStopOrder]: Api ID: %v, Action: %v, Market: %v, Type: %v, Side: %v, Size: %v, PercentSize: %v, UsePercentSize: %v, TrailValue: %v, PostOnly: %v, ReduceOnly: %v, RetryUntilFilled: %v",
		o.ApiID,
		o.Action,
		o.Market,
		o.Type,
		o.Side,
		o.Size,
		o.PercentSize,
		o.UsePercentSize,
		o.TrailValue,
		o.PostOnly,
		o.ReduceOnly,
		o.RetryUntilFilled,
	)
}
