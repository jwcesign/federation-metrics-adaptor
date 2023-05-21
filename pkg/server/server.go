package server

import (
	genericapiserver "k8s.io/apiserver/pkg/server"
)

type server struct {
	*genericapiserver.GenericAPIServer
}

func (s *server) RunUntil(stopCh <-chan struct{}) error {
	return s.GenericAPIServer.PrepareRun().Run(stopCh)
}
