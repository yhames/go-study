package main

import (
	"chat-ws/config"
	"chat-ws/network"
	"chat-ws/repository"
	"chat-ws/service"
	"flag"
)

var configPath = flag.String("config", "./config.toml", "config set")
var port = flag.String("port", ":8080", "port set")

func main() {
	flag.Parse()
	c := config.NewConfig(*configPath)
	r, err := repository.NewRepository(c)
	if err != nil {
		panic(err)
	}

	sv := service.NewService(r)

	s := network.NewServer(sv, r, *port)
	err = s.Start()
	if err != nil {
		panic(err)
	}
}
