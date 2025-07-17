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
	ip   string
}

func NewNetwork(service *service.Service, port string) *Network {
	s := &Network{
		engine:  gin.New(),
		service: service,
		port:    port,
	}
	s.engine.Use(gin.Logger())
	s.engine.Use(gin.Recovery())
	s.engine.Use(cors.New(cors.Config{
		AllowWebSockets:  true,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	return s
}

func (n *Network) Start() error {
	return n.engine.Run(n.port)
}
