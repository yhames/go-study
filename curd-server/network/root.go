package network

import (
	"crud-server/service"
	"github.com/gin-gonic/gin"
)

type Network struct {
	engine  *gin.Engine
	service *service.Service
}

func NewNetwork(service *service.Service) *Network {
	router := &Network{
		engine: gin.New(),
	}
	newUserRouter(router, service.UserService)
	return router
}

func (network *Network) StartServer(port string) error {
	if port == "" {
		port = ":8080" // Default port
	}
	return network.engine.Run(port)
}
