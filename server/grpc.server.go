package server

import (
	"Golang-HD-Wallet/gen/hdwallet"
	"Golang-HD-Wallet/internal/grpc/api"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	Connection     *grpc.Server
	HDWalletServer *api.WalletApp
}

func NewGRPCServer() *GRPCServer {
	connection := grpc.NewServer()
	hdWalletServer := api.NewWalletApp(api.CreateNewMasterWallet(), hdwallet.UnimplementedHDWalletServer{})
	return &GRPCServer{
		Connection:     connection,
		HDWalletServer: hdWalletServer,
	}
}
