// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package v1

import (
	"github.com/go-numb/go-ftx/rest/private/orders"
	"web_socket/src/types"
)

func (c *OrderManager) ResetAllOrderMap() {
	c.state.Lock()
	defer c.state.Unlock()
	newState()
}

func (c *OrderManager) ResetLimitOrderMap() {
	c.state.Lock()
	defer c.state.Unlock()
	c.state.orderMap = types.NewSyncedOrderedMap[string, *orders.ResponseForPlaceOrder]()
}

func (c *OrderManager) ResetBookOrderMap() {
	c.state.Lock()
	defer c.state.Unlock()
	c.state.bookOrderMap = types.NewSyncedOrderedMap[string, *orders.ResponseForPlaceOrder]()
}
