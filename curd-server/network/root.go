package network

import "github.com/gin-gonic/gin"

type Network struct {
	engine *gin.Engine
}

func NewNetwork() *Network {
	return &Network{
		engine: gin.New(),
	}
}

func (n *Network) StartServer(port string) error {
	if port == "" {
		port = ":8080" // Default port
	}
	return n.engine.Run(port)
}
