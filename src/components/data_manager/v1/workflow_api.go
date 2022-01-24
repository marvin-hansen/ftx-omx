// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package v1

import (
	"ftx-omx/src/types"
	"log"
)

func (c *DataManager) StoreApi(api types.Api) (ok bool, msg string) {
	c.state.Lock()
	defer c.state.Unlock()

	db := c.dependencies.dbComp.DB()
	encApi, err := c.encryptApi(api)
	if err != nil {
		return false, "error decrypting API"
	}

	_, insertErr := db.Model(&encApi).Insert()
	if insertErr != nil {
		log.Println(insertErr)
		msg = insertErr.Error()
		return false, msg
	}

	// add api to local cache
	c.state.apiCache.Set(encApi.Id, encApi)
	msg = "ok"
	return true, msg
}

func (c *DataManager) GetAllApis() (api []types.Api, ok bool, msg string) {
	mtd := "GetAllApi"
	c.state.RLock()
	defer c.state.RUnlock()

	// Select all api's.
	db := c.dependencies.dbComp.DB()
	var apis []types.Api
	queryErr := db.Model(&apis).Select()
	if queryErr != nil {
		msg = queryErr.Error()
		log.Println(queryErr)
		return []types.Api{}, false, msg
	}

	var decApis []types.Api
	// Add API to cache if it is not there yet.
	for _, a := range apis {
		_, exists := c.state.apiCache.Get(a.Id)
		if !exists {
			DbgPrint(mtd, "add encrypted api to cache: ")
			c.state.apiCache.Set(a.Id, a)
		}

		DbgPrint(mtd, "decrypt API: ")
		decApi, err := c.decryptApi(a)
		if err != nil {
			return []types.Api{}, false, "error decrypting API"
		}

		decApis = append(decApis, decApi)
	}

	msg = "ok"
	return decApis, true, msg
}

func (c *DataManager) GetApi(apiID string) (api types.Api, ok bool, msg string) {
	mtd := "GetApi"
	c.state.RLock()
	defer c.state.RUnlock()

	// try cache...
	api, ok = c.state.apiCache.Get(apiID)
	if ok {
		DbgPrint(mtd, "Load API from Cache")

		DbgPrint(mtd, "decrypt cached API")
		decApi, err := c.decryptApi(api)
		if err != nil {
			return types.Api{}, false, "error decrypting API"
		}

		msg = "ok"
		return decApi, ok, msg
	}

	// try DB next
	db := c.dependencies.dbComp.DB()
	api = types.Api{}

	queryErr := db.Model(&api).Where("id = ?", apiID).Select()
	if queryErr != nil {
		msg = queryErr.Error()
		log.Println(queryErr)
		return types.Api{}, false, msg
	}

	DbgPrint(mtd, "add encrypted api to cache: ")
	c.state.apiCache.Set(api.Id, api)

	DbgPrint(mtd, "decrypt API: ")
	decApi, err := c.decryptApi(api)
	if err != nil {
		return types.Api{}, false, "error decrypting API"
	}

	msg = "ok"
	return decApi, true, msg
}

func (c *DataManager) DeleteApi(apiID string) (ok bool, msg string) {
	c.state.Lock()
	defer c.state.Unlock()

	db := c.dependencies.dbComp.DB()
	api := &types.Api{}
	_, delErr := db.Model(api).Where("id = ?", apiID).Delete()
	if delErr != nil {
		msg = delErr.Error()
		log.Println(delErr)
		return false, msg
	}
	// delete api from cache if exits.
	_, ok = c.state.apiCache.Get(apiID)
	if ok {
		c.state.apiCache.Delete(apiID)
	}

	msg = "ok"
	return true, msg
}
