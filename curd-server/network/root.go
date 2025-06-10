package network

import "github.com/gin-gonic/gin"

type Network struct {
	engine *gin.Engine
}

func NewNetwork() *Network {
	router := &Network{
		engine: gin.New(),
	}

	newUserRouter(router)

	return router
}

func (n *Network) StartServer(port string) error {
	if port == "" {
		port = ":8080" // Default port
	}
	return n.engine.Run(port)
}
