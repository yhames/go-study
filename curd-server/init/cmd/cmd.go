package cmd

import (
	"crud-server/config"
	"crud-server/network"
	"crud-server/repository"
	"crud-server/service"
)

type Cmd struct {
	config     *config.Config
	network    *network.Network
	repository *repository.Repository
	service    *service.Service
}

func NewCmd(filePath string) *Cmd {
	cmd := &Cmd{
		config: config.NewConfig(filePath),
	}
	cmd.repository = repository.NewRepository()
	cmd.service = service.NewService(cmd.repository)
	cmd.network = network.NewNetwork(cmd.service)
	err := cmd.network.StartServer(cmd.config.Server.Port)
	if err != nil {
		panic(err)
	}
	return cmd
}
