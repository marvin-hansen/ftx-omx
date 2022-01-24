// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package v1

import (
	"os"
	"time"
)

func (s *Service) Stop() {
	s.state.startTime = time.Now()
	DbgPrint("")
	DbgPrint("Shutdown Service:")

	s.state.apiManager.StopAllMonitorApis()
	DbgPrint("* Stop API monitoring!")

	_ = s.state.dbComp.Shutdown()
	DbgPrint("* DB connection closed!")

	_ = s.state.httpListener.Close()
	DbgPrint("* Http connection closed!")

	PrintStopHeader(time.Since(s.state.startTime))
	os.Exit(0)
}
