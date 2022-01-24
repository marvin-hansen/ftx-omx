// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package types

type OrderBookPriceType string

const (
	LargestOrderSizePrice  = "LargestOrderSizePrice"
	SmallestOrderSizePrice = "SmallestOrderSizePrice"
	FirstOrderBookPrice    = "FirstOrderBookPrice"
	MidOrderBookPrice      = "MidOrderBookPrice"
	LastOrderBookPrice     = "LastOrderBookPrice"
	LowestOrderBookPrice   = "LowestOrderBookPrice"
	HighestOrderBookPrice  = "HighestOrderBookPrice"
)

func (s OrderBookPriceType) String() string {
	types := [...]string{"LargestOrderSizePrice", "SmallestOrderSizePrice", "FirstOrderBookPrice", "LastOrderBookPrice", "LowestPrice", "HighestPrice"}
	x := string(s)
	for _, v := range types {
		if v == x {
			return x
		}
	}
	return "UnknownBookPriceType"
}
