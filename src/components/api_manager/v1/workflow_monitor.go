// Copyright (c) 2022-2022. Marvin Hansen | marvin.hansen@gmail.com

// Goroutine safe channel close
// https://stackoverflow.com/questions/68996755/goroutine-safe-channel-close-doesnt-actually-close-webscoket
// FTX Websocket
// https://github.com/go-numb/go-ftx/

package v1

import (
	"context"
	"fmt"
	"github.com/go-numb/go-ftx/realtime"
	"web_socket/src/types"
	"web_socket/src/utils/dbg"
)

func (c *ApiManager) StartMonitorApi(api types.Api) {
	mtd := "StartMonitorApi: "

	DbgPrint(mtd + "create context")
	ctx, _ := context.WithCancel(context.Background())

	DbgPrint(mtd + "start monitoring")
	cancel, err := c.monitorApi(api, ctx)
	if err != nil {
		DbgPrint(mtd + "Error starting API monitoring: " + err.Error())
		return
	}

	DbgPrint(mtd + "store context")
	c.state.cMap[api.Id] = cancel
}

func (c *ApiManager) StopMonitorApi(key string) {
	mtd := "StopMonitorApi: "

	DbgPrint(mtd + "load context")
	cancel, ok := c.state.cMap[key]
	if ok {
		DbgPrint(mtd + "call cancel to stop monitoring")
		cancel()

		DbgPrint(mtd + "delete context")
		delete(c.state.cMap, key)
	} else {
		DbgPrint(mtd + "no context found. return!")
		return
	}
}

func (c *ApiManager) StartAllMonitorApis() {
	mtd := "StopMonitorApi: "
	apis, ok, msg := c.deps.dataManager.GetAllApis()
	if !ok {
		DbgPrint(mtd + "Could not load API's from DB")
		DbgPrint(mtd + msg)
		return
	}

	if len(apis) == 0 {
		DbgPrint(mtd + " No API's received from DB")
		return
	}

	for _, a := range apis {
		DbgPrint(mtd + " Start API monitoring for: " + a.Id)
		c.StartMonitorApi(a)
	}
}

func (c *ApiManager) StopAllMonitorApis() {
	mtd := "StopMonitorApi: "
	apis, ok, msg := c.deps.dataManager.GetAllApis()
	if !ok {
		DbgPrint(mtd + "Could not load API's from DB")
		DbgPrint(mtd + msg)
		return
	}

	if len(apis) == 0 {
		DbgPrint(mtd + " No API's received from DB")
		return
	}

	for _, a := range apis {
		DbgPrint(mtd + " Stop API monitoring for: " + a.Id)
		c.StopMonitorApi(a.Id)
	}
}

func (c *ApiManager) monitorApi(api types.Api, ctx context.Context) (context.CancelFunc, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	ch := make(chan realtime.Response)
	go func() {
		err := realtime.ConnectForPrivate(ctx, ch, api.Key, api.Secret, []string{"orders", "fills"}, nil, api.AccountName)
		if err != nil {
			println(err.Error())
		}
	}()

	go func() {
		defer func() {
			// make sure context is always cancelled to avoid goroutine leak
			cancel()
		}()

		for {
			select {
			case v := <-ch:
				switch v.Type {

				case realtime.ORDERS:
					if debug {
						fmt.Printf("%d	%+v\n", v.Type, v.Orders)
					}

					DbgPrint("Store orderStatus")
					orderStatus := types.NewOrderStatus(api, v.Orders)
					c.deps.dataManager.StoreOrderStatus(orderStatus)

				case realtime.FILLS:
					if debug {
						fmt.Printf("%d	%+v\n", v.Type, v.Fills)
					}

					DbgPrint("Store orderFill")
					orderFill := types.NewOrderFill(api, v.Fills)
					c.deps.dataManager.StoreOrderFill(orderFill)

				case realtime.ERROR:
					dbg.LogError(v.Results)
					err := types.NewError(v.Symbol, v.Results.Error())
					DbgPrint("Store Error:")
					c.deps.dataManager.StoreError(err)

				case realtime.UNDEFINED:
					dbg.LogError(v.Results)
					fmt.Printf("UNDEFINED %s	%s\n", v.Symbol, v.Results.Error())
					DbgPrint("Store Undefined: ")
					err := types.NewError(v.Symbol, v.Results.Error())
					c.deps.dataManager.StoreError(err)
				}
			}
		}
	}()

	return cancel, nil
}
