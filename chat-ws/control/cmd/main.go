package main

import (
	"chat-ws-control/cmd/app"
	"chat-ws-control/config"
	"flag"
)

var configPath = flag.String("config", "./config.toml", "config set")

func main() {
	flag.Parse()
	c := config.NewConfig(*configPath)
	newApp := app.NewApp(c)
	if err := newApp.Run(); err != nil {
		panic(err)
	}
}
