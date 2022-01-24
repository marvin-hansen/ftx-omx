// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package v1

import "ftx-omx/src/components/component_manager"

func (s *Service) initComponents(msg string) {
	PrintInitHeader(s.state.serviceID, msg)

	s.state.compManager = component_manager.NewComponentManager(s.state.dbComp)
	DbgPrint(" * component manager instantiated")

	s.state.apiManager = s.state.compManager.GetApiManager()
	DbgPrint(" * api manager instantiated")

	s.state.clientManager = s.state.compManager.GetAccountManager()
	DbgPrint(" * client manager instantiated")

	s.state.orderManager = s.state.compManager.GetOrderManager()
	DbgPrint(" * order manager instantiated")

	DbgPrint(" * Components init complete.")
}
