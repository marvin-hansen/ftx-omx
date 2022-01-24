// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package v1

import (
	"errors"
	"web_socket/src/types"
)

func (c *ApiManager) AddApi(value types.Api) {
	c.state.Lock()
	defer c.state.Unlock()

	c.deps.dataManager.StoreApi(value)
}
func (c *ApiManager) GetApi(key string) (types.Api, error) {
	c.state.RLock()
	defer c.state.RUnlock()

	api, ok, _ := c.deps.dataManager.GetApi(key)
	if ok {
		return api, nil
	} else {
		return types.Api{}, errors.New("No Api Found for key: " + key)
	}
}

func (c *ApiManager) UpdateApi(key string, value types.Api) {
	c.state.Lock()
	defer c.state.Unlock()

	_, ok, _ := c.deps.dataManager.GetApi(key)
	if ok {
		c.deps.dataManager.DeleteApi(key)

		c.deps.dataManager.StoreApi(value)
	} else {
		c.deps.dataManager.StoreApi(value)
	}
}

func (c *ApiManager) CheckApiExists(key string) (ok bool) {
	c.state.RLock()
	defer c.state.RUnlock()

	_, ok, _ = c.deps.dataManager.GetApi(key)
	return ok
}

func (c *ApiManager) DeleteApi(key string) {
	c.state.Lock()
	defer c.state.Unlock()

	c.deps.dataManager.DeleteApi(key)
}
