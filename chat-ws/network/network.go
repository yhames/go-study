package network

import (
	"chat-ws/repository"
	"chat-ws/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

type Network struct {
	engin *gin.Engine

	service    *service.Service
	repository *repository.Repository

	port string
	ip   string
}

func NewServer(service *service.Service, repository *repository.Repository, port string) *Network {
	n := &Network{
		engin:      gin.New(),
		service:    service,
		repository: repository,
		port:       port,
	}
	n.engin.Use(gin.Logger())
	n.engin.Use(gin.Recovery())
	n.engin.Use(cors.New(cors.Config{
		AllowWebSockets:  true,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	r := NewRoom()
	go r.RunInit()
	n.engin.GET("/room", r.SocketServe)

	return n
}

func (n *Network) Start() error {
	log.Println("Starting server...")
	return n.engin.Run(n.port)
}
