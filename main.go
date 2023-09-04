package main

import (
	"furina/config"
	"furina/server"
)

func main() {
	srv := server.NewServer()
	if err := srv.Run(config.GetConfig().Server.HttpPort); err != nil {
		panic(any(err))
	}
}
