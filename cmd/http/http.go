package main

import (
	"Golang-HD-Wallet/server"
	"log"
)

func main() {

	s := server.NewHttpServer()
	s.Initialize()
	if err := s.Run(); err != nil {
		log.Panicf("Failed to start server: %v", err)
	}
}
