package wallet

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"log"
	"os"

	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

type GetDerivePathForUserReq struct {
	userId uint
	wallet *hdwallet.Wallet
}

func GetWalletFromMnemonic(mnemonic string) (*hdwallet.Wallet, error) {
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		log.Panicf("Failed to create wallet : %v", err)
	}
	return wallet, err
}

func GetDerivePathForUser(userId uint) accounts.DerivationPath {
	basePath := os.Getenv("BASE_PATH")
	if basePath == "" {
		log.Panicf("Failed to set BASE_PATH environment variable")
	}
	path := fmt.Sprintf("%v%v", basePath, userId)
	return hdwallet.MustParseDerivationPath(path)
}
