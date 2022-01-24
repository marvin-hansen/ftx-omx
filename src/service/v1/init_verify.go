// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package v1

import (
	"ftx-omx/src/utils/dbg"
)

func (s *Service) verifyInit() {
	dbg.NilCheck(s.state.dbConf, "NPE: DB Config is nil. Fix init.")
	dbg.NilCheck(s.state.dbModel, "NPE: DB Model is nil. Fix init.")
	dbg.NilCheck(s.state.compManager, "NPE: component manager is nil. Fix init.")
	dbg.NilCheck(s.state.apiManager, "NPE: api manager is nil. Fix init.")
	dbg.NilCheck(s.state.clientManager, "NPE: client manager is nil. Fix init.")
	dbg.NilCheck(s.state.dbComp, "NPE: dbComp is nil. Fix init.")
	dbg.NilCheck(s.state.orderManager, "NPE: order manager is nil. Fix init.")
	dbg.NilCheck(s.state.config, "NPE: main config is nil. Fix init.")
	dbg.NilCheck(s.state.httpServer, "NPE: httpServer is nil. Fix init.")
	dbg.NilCheck(s.state.lis, "NPE: lis is nil. Fix init.")
	dbg.NilCheck(s.state.httpListener, "NPE: httpListener is nil. Fix init.")
	dbg.NilCheck(s.state.mux, "NPE: mux is nil. Fix init.")
	DbgPrint("Init: verified!")
}
