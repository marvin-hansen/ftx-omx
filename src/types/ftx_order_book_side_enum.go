// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package types

type OrderBookSide string

const (
	Ask OrderBookSide = "ask"
	Bid               = "bid"
)

func (s OrderBookSide) String() string {
	sides := [...]string{"ask", "bid"}
	x := string(s)
	for _, v := range sides {
		if v == x {
			return x
		}
	}
	return "UnknownOrderBookSide"
}
