// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package main

import (
	"ftx-omx/src/cfg"
	"ftx-omx/src/service"
)

func main() {
	service.NewService(cfg.GetMainConfig()).Serve()
}
