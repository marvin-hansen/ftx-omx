// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package v1

import (
	"ftx-omx/src/types"
	"log"
)

func (c *DataManager) StoreOrderStatus(orderStatus *types.OrderStatus) (ok bool, msg string) {
	c.state.Lock()
	defer c.state.Unlock()

	db := c.dependencies.dbComp.DB()
	_, insertErr := db.Model(orderStatus).Insert()
	if insertErr != nil {
		log.Println(insertErr)
		msg = insertErr.Error()
		return false, msg
	} else {
		msg = "ok"
		return true, msg
	}
}
