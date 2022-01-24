// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package v1

import "web_socket/src/components/api_manager"

type Dependencies struct {
	apiManager api_manager.ApiManager
}

func newDependencies(apiManager api_manager.ApiManager) (deps *Dependencies) {
	deps = &Dependencies{apiManager: apiManager}
	return deps
}
