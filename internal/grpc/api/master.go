package api

import (
	"github.com/ethereum/go-ethereum/accounts"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"log"
)

type (
	MasterWallet struct {
		wallet  *hdwallet.Wallet
		account *accounts.Account
	}
	IMasterWallet interface {
		GetPrivateKey(userId uint64, pubKey string) GetPrivateKeyContext
		GenerateChildWallet(userId uint64) GenerateChannelContext
	}

	GenerateChannelContext struct {
		ChildWallet IChildWallet
	}

	GetPrivateKeyContext struct {
		PrivateKey string
	}
)

func (m *MasterWallet) GenerateChildWallet(userId uint64) GenerateChannelContext {
	mnemonic := GetMnemonicFromEnv()
	wallet, _ := GetWalletFromMnemonic(mnemonic)
	path := GetDerivePathForUser(userId)
	account, err := wallet.Derive(path, true)
	if err != nil {
		log.Printf("Failed to derive : %v\n", err)
		return GenerateChannelContext{}
	}
	return GenerateChannelContext{
		ChildWallet: NewChildWallet(wallet, &account),
	}
}

func (m *MasterWallet) GetPrivateKey(userId uint64, pubKey string) GetPrivateKeyContext {
	/* UserID에 맞는 Child 지갑 가져옴 */
	generateCtx := m.GenerateChildWallet(userId)
	/* private key 꺼내기 */
	privateKey, err := generateCtx.ChildWallet.GetPrivateKey()
	if err != nil {
		log.Println("Failed to get private key from wallet : ", err)
		return GetPrivateKeyContext{}
	}
	return GetPrivateKeyContext{
		PrivateKey: privateKey,
	}
}

func CreateNewMasterWallet() IMasterWallet {
	mnemonic := GetMnemonicFromEnv()
	wallet, _ := GetWalletFromMnemonic(mnemonic)
	path := GetDerivePathForUser(0)
	account, err := wallet.Derive(path, true)
	if err != nil {
		log.Panicf("Failed to derive : %v\n", err)
	}

	return &MasterWallet{
		wallet:  wallet,
		account: &account,
	}
}
