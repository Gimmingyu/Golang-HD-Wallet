package api

import (
	. "Golang-HD-Wallet/gen/hdwallet"
	context "context"
)

type WalletApp struct {
	masterWallet IMasterWallet
	UnimplementedHDWalletServer
}

func NewWalletApp(masterWallet IMasterWallet, unimplementedHDWalletServer UnimplementedHDWalletServer) *WalletApp {
	return &WalletApp{masterWallet: masterWallet, UnimplementedHDWalletServer: unimplementedHDWalletServer}
}

func (walletApp *WalletApp) Generate(ctx context.Context, request *GenerateRequest) (*GenerateResponse, error) {
	done := make(chan GenerateChannelContext)
	go func() {
		done <- walletApp.masterWallet.GenerateChildWallet(request.UserId)
	}()

	select {
	case result := <-done:

		return &GenerateResponse{
			PublicKey: result.ChildWallet.GetPublicKey(),
		}, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (walletApp *WalletApp) GetPrivateKey(ctx context.Context, request *GetPrivateKeyRequest) (*GetPrivateKeyResponse, error) {
	done := make(chan GetPrivateKeyContext)
	go func() {
		done <- walletApp.masterWallet.GetPrivateKey(request.UserId, request.PublicKey)
	}()
	select {
	case result := <-done:
		return &GetPrivateKeyResponse{
			PrivateKey: result.PrivateKey,
		}, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (walletApp *WalletApp) mustEmbedUnimplementedHDWalletServer() {
	//TODO implement me
	panic("implement me")
}
