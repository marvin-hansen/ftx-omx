// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package v1

func (s *Service) initServiceMonitoring(msg string) {
	PrintInitHeader(s.state.serviceID, msg)
	s.monitor()
	DbgPrint(" * Service Monitoring configured.")
}

func (s *Service) initApiMonitoring(msg string) {
	PrintInitHeader(s.state.serviceID, msg)
	s.state.apiManager.StartAllMonitorApis()
	DbgPrint(" * API Monitoring configured.")

}
