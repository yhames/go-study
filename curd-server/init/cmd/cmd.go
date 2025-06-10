package cmd

import (
	"crud-server/config"
	"crud-server/network"
)

type Cmd struct {
	config  *config.Config
	network *network.Network
}

func NewCmd(filePath string) *Cmd {
	cmd := &Cmd{
		config:  config.NewConfig(filePath),
		network: network.NewNetwork(),
	}
	err := cmd.network.StartServer(cmd.config.Server.Port)
	if err != nil {
		panic(err)
	}
	return cmd
}
