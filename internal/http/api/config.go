package api

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	Port uint64
}

func NewAppConfig() *Config {
	sPort := os.Getenv("PORT")
	if sPort == "" {
		log.Panicf("Failed to set port: %v", sPort)
	}
	uPort, err := strconv.ParseUint(sPort, 10, 64)
	if err != nil {
		log.Panicf("Failed to convert port to uint: %v", sPort)
	}
	return &Config{Port: uPort}
}
