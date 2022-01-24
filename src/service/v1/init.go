// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package v1

import (
	"ftx-omx/src/utils/dbg"
	"time"
)

func (s *Service) init() {
	s.state.startTime = time.Now()
	dbg.DbgCheck(dbgMain)
	s.runInit("Init Level 1: Configure Database!", s.initDB)
	s.runInit("Init Level 2: Configure Components!", s.initComponents)
	s.runInit("Init Level 3: Configure Http network!", s.initNetwork)
	s.runInit("Init Level 4: Configure Service monitoring!", s.initServiceMonitoring)
	s.runInit("Init Level 5: Configure API monitoring!", s.initApiMonitoring)
	s.verifyInit()
	DbgPrint("Init: Completed!")
}
