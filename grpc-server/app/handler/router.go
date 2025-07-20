package handler

import (
	"github.com/gin-gonic/gin"
	"grpc-server/app/config"
	"grpc-server/app/service"
	"grpc-server/grpc/client"
)

type Router struct {
	config     *config.Config
	grpcClient *client.GrpcClient

	service *service.Service
}

func NewRouter(config *config.Config, service *service.Service, grpcClient *client.GrpcClient) (*Router, error) {
	return &Router{
		config:     config,
		service:    service,
		grpcClient: grpcClient,
	}, nil
}

func (r *Router) Setup(engine *gin.Engine) {
	engine.Use(gin.Recovery())
	engine.Use(gin.Logger())

	engine.POST("/login", r.Login)
	engine.GET("/verify", r.verifyLogin(), r.Verify)
}
