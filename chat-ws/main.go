package main

import "chat-ws/network"

func main() {
	server := network.NewServer()
	err := server.Start()
	if err != nil {
		panic(err)
	}
}
