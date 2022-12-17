package wallet

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GetMnemonicFromEnv() string {
	if err := godotenv.Load(".env"); err != nil {
		log.Panicf("Failed to load env file : %v\n", err)
	}

	mnemonic := os.Getenv("MNEMONIC")
	if mnemonic == "" {
		panic("FAILED TO GET MNEMONIC")
	}

	return mnemonic
}
