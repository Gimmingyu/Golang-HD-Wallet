package api

import (
	"github.com/ethereum/go-ethereum/accounts"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

type (
	IChildWallet interface {
		GetPublicKey() string
		GetPrivateKey() (string, error)
	}

	ChildWallet struct {
		wallet  *hdwallet.Wallet
		account *accounts.Account
	}
)

func NewChildWallet(w *hdwallet.Wallet, a *accounts.Account) IChildWallet {
	return &ChildWallet{wallet: w, account: a}
}

func (w *ChildWallet) GetPublicKey() string {
	return w.account.Address.Hex()
}

func (w *ChildWallet) GetPrivateKey() (string, error) {
	return w.wallet.PrivateKeyHex(*w.account)
}
