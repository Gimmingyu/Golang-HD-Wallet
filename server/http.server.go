package server

import (
	"Golang-HD-Wallet/internal/http/api"
	"github.com/joho/godotenv"
	"log"
)

type (
	HttpServer struct {
		app api.IGinApp
		cfg *api.Config
	}

	IHttpServer interface {
		Initialize()
		Run() error
	}
)

func (h *HttpServer) Initialize() {
	if err := godotenv.Load(".env"); err != nil {
		log.Panicf("Failed to load environment: %v", err)
	}

	h.cfg = api.NewAppConfig()
	h.app = api.NewGinApp()
	h.app.SetRouter()

}

func (h *HttpServer) Run() error {
	return h.app.Run()
}

func NewHttpServer() IHttpServer {
	return &HttpServer{}
}
