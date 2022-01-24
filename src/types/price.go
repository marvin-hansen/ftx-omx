// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package types

type Price struct {
	Ask  float64
	Bid  float64
	Last float64
}

func NewPrice(last, bid, ask float64) *Price {
	return &Price{
		Ask:  ask,
		Bid:  bid,
		Last: last,
	}
}
