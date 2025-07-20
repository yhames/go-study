package main

import (
	"flag"
	"grpc-server/app"
	"grpc-server/app/config"
	"grpc-server/grpc/server"
	"time"
)

var configFlag = flag.String("config", "./config.toml", "config path")

func main() {
	flag.Parse()
	c := config.NewConfig(*configFlag)

	// Start gRPC server
	server.NewGrpcServer(c)
	time.Sleep(1e9) // Wait for gRPC server to start

	// Start the application
	app.NewApp(c)
}
