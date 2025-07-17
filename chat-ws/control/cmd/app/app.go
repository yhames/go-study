package app

import (
	"chat-ws-control/config"
	"chat-ws-control/network"
	"chat-ws-control/repository"
	"chat-ws-control/service"
)

type App struct {
	config     *config.Config
	repository *repository.Repository
	service    *service.Service
	network    *network.Network
}

func NewApp(config *config.Config) *App {
	app := &App{
		config: config,
	}
	var err error
	if app.repository, err = repository.NewRepository(app.config); err != nil {
		panic(err)
	}
	app.service = service.NewService(app.repository)
	app.network = network.NewNetwork(app.service, app.config.Info.Port)

	return app
}

func (a *App) Run() error {
	return a.network.Start()
}
