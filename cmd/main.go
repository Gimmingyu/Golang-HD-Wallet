package main

import (
	"Golang-HD-Wallet/gen/hdwallet"
	"Golang-HD-Wallet/server"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Panic(err)
	}

	log.Println("HD-Wallet Server using Geth is not running!!")

	grpcServer := server.NewGRPCServer()
	hdwallet.RegisterHDWalletServer(grpcServer.Connection, grpcServer.HDWalletServer)

	if err := grpcServer.Connection.Serve(listener); err != nil {
		log.Panic(err)
	}
}
