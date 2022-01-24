// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package types

import "fmt"

type PlaceTriggerOrder struct {
	Auth           string  `json:"auth,omitempty"` // for the tradingview webhook that can't send auth headers
	ApiID          string  `json:"api_id,omitempty"`
	Action         string  `json:"action,omitempty"` // for the tradingview webhook that requires only one endpoint.
	Market         string  `json:"market"`
	Type           string  `json:"type"` // Type stop, trailingStop, takeProfit. default is stop
	Side           string  `json:"side"` //"buy" or "sell"
	Size           float64 `json:"size"`
	PercentSize    float64 `json:"percentSize,omitempty"`    // order size as percentage of available capital
	UsePercentSize bool    `json:"usePercentSize,omitempty"` // whether to use percent size. default false.
	// Additional parameters for stop loss / take profit orders
	TriggerPrice     float64 `json:"triggerPrice,omitempty"`     // Special Orders
	LimitPrice       float64 `json:"orderPrice,omitempty"`       // optional; order type is limit if this is specified; otherwise market
	ReduceOnly       bool    `json:"reduceOnly,omitempty"`       // optional; default is false
	PostOnly         bool    `json:"postOnly,omitempty"`         // optional; default is false
	RetryUntilFilled bool    `json:"retryUntilFilled,omitempty"` // Whether to keep re-triggering until filled. optional, default true for market orders
}

func (o PlaceTriggerOrder) String() string {
	return fmt.Sprintf("[PlaceTriggerOrder]: Api ID: %v, Action: %v, Market: %v, Type: %v, Side: %v, Size: %v, PercentSize: %v, UsePercentSize: %v, TriggerPrice: %v, LimitPrice: %v, ReduceOnly: %v, PostOnly: %v, RetryUntilFilled: %v",
		o.ApiID,
		o.Action,
		o.Market,
		o.Type,
		o.Side,
		o.Size,
		o.PercentSize,
		o.UsePercentSize,
		o.TriggerPrice,
		o.LimitPrice,
		o.ReduceOnly,
		o.PostOnly,
		o.RetryUntilFilled,
	)
}
