package main

import (
	"furina/config"
	"furina/server"
	"log"
)

func main() {
	srv := server.NewServer()
	if err := srv.Run(":" + config.GetConfig().Server.HttpPort); err != nil {
		log.Fatalln(err)
	}
}
