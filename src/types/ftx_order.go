// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package types

import (
	"fmt"
	"github.com/go-numb/go-ftx/rest/private/orders"
	"time"
)

type PlaceOrder struct {
	Auth              string  `json:"auth,omitempty"` // for the tradingview webhook that can't send auth headers
	ApiID             string  `json:"api_id,omitempty"`
	Action            string  `json:"action,omitempty"` // for the tradingview webhook that requires only one endpoint.
	Market            string  `json:"market"`
	Type              string  `json:"type"`
	Side              string  `json:"side"`
	Price             float64 `json:"price"`
	Size              float64 `json:"size"`
	PercentSize       float64 `json:"percentSize,omitempty"`
	UsePercentSize    bool    `json:"usePercentSize,omitempty"`
	ClientID          string  `json:"clientId,omitempty"`
	ReduceOnly        bool    `json:"reduceOnly,omitempty"`
	Ioc               bool    `json:"ioc,omitempty"`
	PostOnly          bool    `json:"postOnly,omitempty"`
	RejectOnPriceBand bool    `json:"rejectOnPriceBand,omitempty"`
}

func (o PlaceOrder) String() string {
	return fmt.Sprintf("[PlaceOrder]: Api ID: %v, Action: %v, Market: %v, Type: %v, Side: %v, Price: %v, Size: %v, PercentSize: %v, UsePercentSize: %v, ClientID: %v, ReduceOnly: %v, Ioc: %v, PostOnly: %v, RejectOnPriceBand: %v,",
		o.ApiID,
		o.Action,
		o.Market,
		o.Type,
		o.Side,
		o.Price,
		o.Size,
		o.PercentSize,
		o.UsePercentSize,
		o.ClientID,
		o.ReduceOnly,
		o.Ioc,
		o.PostOnly,
		o.RejectOnPriceBand,
	)
}

type OrderStatus struct {
	PK            int64     `pg:",pk,unique" json:",omitempty"` // PK for internal DB use
	AccountId     string    `json:"accountId"`
	AccountName   string    `json:"accountName"`
	ID            int64     `json:"id"`
	ClientID      string    `json:"clientId"`
	Market        string    `json:"market"`
	Type          string    `json:"type"`
	Side          string    `json:"side"`
	Size          float64   `json:"size"`
	Price         float64   `json:"price"`
	ReduceOnly    bool      `json:"reduceOnly"`
	Ioc           bool      `json:"ioc"`
	PostOnly      bool      `json:"postOnly"`
	Status        string    `json:"status"`
	FilledSize    float64   `json:"filledSize"`
	RemainingSize float64   `json:"remainingSize"`
	AvgFillPrice  float64   `json:"avgFillPrice"`
	CreatedAt     time.Time `json:"createdAt"`
}

func NewOrderStatus(api Api, order orders.Order) *OrderStatus {
	return &OrderStatus{
		AccountId:     api.Id,
		AccountName:   api.AccountName,
		ID:            order.ID,
		ClientID:      order.ClientID,
		Market:        order.Market,
		Type:          order.Type,
		Side:          order.Side,
		Size:          order.Size,
		Price:         order.Price,
		ReduceOnly:    order.ReduceOnly,
		Ioc:           order.Ioc,
		PostOnly:      order.PostOnly,
		Status:        order.Status,
		FilledSize:    order.FilledSize,
		RemainingSize: order.RemainingSize,
		AvgFillPrice:  order.AvgFillPrice,
		CreatedAt:     order.CreatedAt,
	}
}
