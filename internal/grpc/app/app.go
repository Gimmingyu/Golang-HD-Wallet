package app

import (
	. "Golang-HD-Wallet/gen/hdwallet"
	"Golang-HD-Wallet/internal/grpc/wallet"
	context "context"
)

type HealthCheckRequest struct{}

type HealthCheckResponse struct{}

type WalletApp struct {
	masterWallet wallet.IMasterWallet
	UnimplementedHDWalletServer
}

func (walletApp *WalletApp) Generate(ctx context.Context, request *GenerateRequest) (*GenerateResponse, error) {
	done := make(chan wallet.GenerateChannelContext)
	go func() {
		done <- walletApp.masterWallet.GenerateChildWallet(request.UserId)
	}()

	select {
	case result := <-done:
		return &GenerateResponse{
			PublicKey: result.ChildWallet.GetPublicKey(),
		}, result.Err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (walletApp *WalletApp) GetPrivateKey(ctx context.Context, request *GetPrivateKeyRequest) (*GetPrivateKeyResponse, error) {
	done := make(chan wallet.GetPrivateKeyContext)
	go func() {
		done <- walletApp.masterWallet.GetPrivateKey(request.UserId, request.PublicKey)
	}()

	select {
	case result := <-done:
		return &GetPrivateKeyResponse{
			PrivateKey: result.PrivateKey,
		}, result.Err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (walletApp *WalletApp) mustEmbedUnimplementedHDWalletServer() {
	//TODO implement me
	panic("implement me")
}

type IAppServer interface {
	HDWalletServer
}
