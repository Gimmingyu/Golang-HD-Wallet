package wallet

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
		GetPublicKey() string
		GenerateChildWallet(userId uint64) GenerateChannelContext
	}

	GenerateChannelContext struct {
		ChildWallet IChildWallet
		Err         error
	}

	GetPrivateKeyContext struct {
		PrivateKey string
		Err        error
	}
)

func (m *MasterWallet) GenerateChildWallet(userId uint64) GenerateChannelContext {
	//TODO implement me
	panic("implement me")
}

func (m *MasterWallet) GetPrivateKey(userId uint64, pubKey string) GetPrivateKeyContext {
	privateKey, err := m.wallet.PrivateKeyHex(*m.account)
	if err != nil {

	}
	return GetPrivateKeyContext{
		PrivateKey: privateKey, Err: err,
	}
}

func (m *MasterWallet) GetPublicKey() string {
	return m.account.Address.Hex()
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
