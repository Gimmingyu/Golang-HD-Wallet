package server

import (
	"Golang-HD-Wallet/gen/hdwallet"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	Connection     *grpc.Server
	HDWalletServer hdwallet.UnimplementedHDWalletServer
}

func NewGRPCServer() *GRPCServer {
	connection := grpc.NewServer()
	hdWalletServer := hdwallet.UnimplementedHDWalletServer{}
	return &GRPCServer{
		Connection:     connection,
		HDWalletServer: hdWalletServer,
	}
}
