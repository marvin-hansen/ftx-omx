// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package v1

import (
	"ftx-omx/src/cfg"
	"ftx-omx/src/clients/pgdb"
)

func (s *Service) initDB(msg string) {
	PrintInitHeader(s.state.serviceID, msg)
	s.state.dbComp = pgdb.NewDBComponent(s.state.dbConf, cfg.Prod)
	DbgPrint(" * DB Init Complete!")
}
