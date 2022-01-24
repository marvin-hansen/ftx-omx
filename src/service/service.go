// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package service

import (
	v1 "ftx-omx/src/service/v1"
	"ftx-omx/src/types"
)

type Service interface {
	Serve()
	Stop()
}

func NewService(config *types.MainConfig) (service Service) {
	service = v1.NewService(config)
	return service
}
