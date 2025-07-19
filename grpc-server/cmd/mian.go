package main

import (
	"flag"
	"grpc-server/app"
	"grpc-server/app/config"
)

var configFlag = flag.String("config", "./config.toml", "config path")

func main() {
	flag.Parse()
	c := config.NewConfig(*configFlag)
	a := app.NewApp(c)
	err := a.Run()
	if err != nil {
		panic(err)
	}
}
