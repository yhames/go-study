package network

import (
	"chat-ws-control/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Network struct {
	engine *gin.Engine

	service *service.Service

	port string
}

func NewNetwork(service *service.Service, port string) *Network {
	n := &Network{
		engine:  gin.New(),
		service: service,
		port:    port,
	}
	n.engine.Use(gin.Logger())
	n.engine.Use(gin.Recovery())
	n.engine.Use(cors.New(cors.Config{
		AllowWebSockets:  true,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	registerTowerApi(n)

	return n
}

func (n *Network) Start() error {
	return n.engine.Run(n.port)
}
