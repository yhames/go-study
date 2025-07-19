package network

import (
	"github.com/gin-gonic/gin"
	"grpc-server/app/config"
	"grpc-server/app/service"
)

type Router struct {
	config  *config.Config
	service *service.Service
}

func NewRouter(config *config.Config, service *service.Service) (*Router, error) {
	return &Router{
		config:  config,
		service: service,
	}, nil
}

func (r *Router) Setup(engine *gin.Engine) {
	// TODO
}
