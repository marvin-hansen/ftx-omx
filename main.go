// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package main

import (
	"web_socket/src/cfg"
	"web_socket/src/service"
)

func main() {
	service.NewService(cfg.GetMainConfig()).Serve()
}
