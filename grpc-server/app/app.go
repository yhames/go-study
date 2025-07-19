package app

import (
	"github.com/gin-gonic/gin"
	"grpc-server/app/config"
	network "grpc-server/app/handler"
	"grpc-server/app/repository"
	"grpc-server/app/service"
)

type App struct {
	config *config.Config
	engine *gin.Engine

	network *network.Router
}

func NewApp(config *config.Config) *App {
	r, err := repository.NewRepository(config)
	if err != nil {
		panic(err)
	}

	s, err := service.NewService(config, r)
	if err != nil {
		panic(err)
	}

	n, err := network.NewRouter(config, s)
	if err != nil {
		panic(err)
	}

	app := &App{
		config:  config,
		network: n,
	}
	return app
}

func (a *App) Run() error {
	a.engine = gin.New()
	a.network.Setup(a.engine)
	return a.engine.Run(*a.config.Server.Port)
}
