package app

import (
	"github.com/gin-gonic/gin"
	"grpc-server/app/config"
	"grpc-server/app/handler"
	"grpc-server/app/repository"
	"grpc-server/app/service"
	"grpc-server/grpc/client"
)

type App struct {
	config  *config.Config
	engine  *gin.Engine
	handler *handler.Router
}

func NewApp(config *config.Config) {
	c, err := client.NewGrpcClient(config)
	if err != nil {
		panic(err)
	}

	r, err := repository.NewRepository(config, c)
	if err != nil {
		panic(err)
	}

	s, err := service.NewService(config, r)
	if err != nil {
		panic(err)
	}

	n, err := handler.NewRouter(config, s, c)
	if err != nil {
		panic(err)
	}

	app := &App{
		config:  config,
		handler: n,
	}

	app.engine = gin.New()
	app.handler.Setup(app.engine)
	err = app.engine.Run(*app.config.Server.Port)
	if err != nil {
		panic(err)
	}
}
