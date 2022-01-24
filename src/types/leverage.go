// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package types

import "fmt"

type LeverageRequest struct {
	ApiID    string `json:"api_id"`
	Market   string `json:"market"`
	Leverage int    `json:"leverage"`
}

func (o LeverageRequest) String() string {
	return fmt.Sprintf("[LeverageRequest]: Api ID: %v, Market: %v, Leverage: %v ",
		o.ApiID,
		o.Market,
		o.Leverage,
	)
}
