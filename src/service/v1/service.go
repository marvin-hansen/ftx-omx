// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package v1

import (
	"ftx-omx/src/types"
)

type Service struct {
	state *State
}

func NewService(mainConfig *types.MainConfig) (service *Service) {
	service = &Service{
		state: NewState(mainConfig),
	}
	service.init()
	return service
}
