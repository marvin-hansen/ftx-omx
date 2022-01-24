// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package v1

import (
	"os"
	"os/signal"
	"syscall"
)

func (s *Service) monitor() { // Graceful shutdown with Go
	go func() { // https://medium.com/over-engineering/graceful-shutdown-with-go-http-servers-and-kubernetes-rolling-updates-6697e7db17cf
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)    // interrupt signal from terminal
		signal.Notify(sigint, os.Kill)         // interrupt signal from terminal
		signal.Notify(sigint, syscall.SIGTERM) // interrupt signal from kubernetes
		<-sigint
		DbgPrint(" Received an interrupt signal, shut down.")
		s.Stop()
	}()
}
