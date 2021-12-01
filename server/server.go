package server

import (
	"github.com/gin-gonic/gin"
)

type APIServer struct {
	r      *gin.Engine
	routes *APIRoutes
}

func NewAPIServer() (*APIServer, error) {
	r := gin.Default()

	routes, err := NewAPIRoutes(r)
	if err != nil {
		return nil, err
	}

	return &APIServer{
		r:      r,
		routes: routes,
	}, nil
}

func (server *APIServer) Run() error {
	server.routes.SetRoutes()

	err := server.r.Run()

	if err != nil {
		return err
	}

	return nil
}
