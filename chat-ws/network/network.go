package network

import (
	"chat-ws/repository"
	"chat-ws/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

type Server struct {
	engine *gin.Engine

	service    *service.Service
	repository *repository.Repository

	port string
	ip   string
}

func NewServer(service *service.Service, repository *repository.Repository, port string) *Server {
	s := &Server{
		engine:     gin.New(),
		service:    service,
		repository: repository,
		port:       port,
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

	r := NewRoom()
	go r.RunInit()
	registerServer(s)
	s.engine.GET("/room", r.SocketServe)

	return s
}

func (n *Server) Start() error {
	log.Println("Starting server...")
	return n.engine.Run(n.port)
}
