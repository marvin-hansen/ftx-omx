// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package types

import (
	"github.com/go-numb/go-ftx/rest/private/fills"
	"time"
)

type OrderFill struct {
	PK            int64     `pg:",pk,unique" json:",omitempty"` // PK for internal DB use
	AccountId     string    `json:"accountId"`
	AccountName   string    `json:"accountName"`
	Future        string    `json:"future"`
	Market        string    `json:"market"`
	Type          string    `json:"type"`
	Liquidity     string    `json:"liquidity"`
	BaseCurrency  string    `json:"baseCurrency"`
	QuoteCurrency string    `json:"quoteCurrency"`
	FeeCurrency   string    `json:"feeCurrency"`
	Side          string    `json:"side"`
	Price         float64   `json:"price"`
	Size          float64   `json:"size"`
	Fee           float64   `json:"fee"`
	FeeRate       float64   `json:"feeRate"`
	Time          time.Time `json:"time"`
	ID            int       `json:"id"`
	OrderID       int       `json:"orderId"`
	TradeID       int       `json:"tradeId"`
}

func NewOrderFill(api Api, fill fills.Fill) *OrderFill {

	return &OrderFill{
		AccountId:     api.Id,
		AccountName:   api.AccountName,
		Future:        fill.Future,
		Market:        fill.Market,
		Type:          fill.Type,
		Liquidity:     fill.Liquidity,
		BaseCurrency:  fill.BaseCurrency,
		QuoteCurrency: fill.QuoteCurrency,
		FeeCurrency:   fill.FeeCurrency,
		Side:          fill.Side,
		Price:         fill.Price,
		Size:          fill.Size,
		Fee:           fill.Fee,
		FeeRate:       fill.FeeRate,
		Time:          fill.Time,
		ID:            fill.ID,
		OrderID:       fill.OrderID,
		TradeID:       fill.TradeID,
	}
}
