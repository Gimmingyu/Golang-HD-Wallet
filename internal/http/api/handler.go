package api

import (
	"Golang-HD-Wallet/gen/hdwallet"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type (
	Handler struct {
		connection *grpc.ClientConn
		client     hdwallet.HDWalletClient
	}

	IHandler interface {
		GetChildPublicKey(ctx *gin.Context)
		GetChildPrivateKey(ctx *gin.Context)
		GetConnection() *grpc.ClientConn
		GetClient() hdwallet.HDWalletClient
	}
)

func (h *Handler) GetChildPublicKey(ctx *gin.Context) {
	var req *hdwallet.GenerateRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		log.Println("HERE!!!")
		ctx.JSON(400, gin.H{"message": err.Error()})
	}
	log.Println(req)
	r, err := h.client.Generate(ctx, req)
	if err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
	}
	ctx.JSON(200, gin.H{"result": r})
}

func (h *Handler) GetChildPrivateKey(ctx *gin.Context) {
	var req *hdwallet.GetPrivateKeyRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		log.Println("HERE!!!")
		ctx.JSON(400, gin.H{"message": err.Error()})
	}

	log.Println(req)
	r, err := h.client.GetPrivateKey(ctx, req)
	if err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
	}
	ctx.JSON(200, gin.H{"result": r})
}

func (h *Handler) GetConnection() *grpc.ClientConn {
	return h.connection
}

func (h *Handler) GetClient() hdwallet.HDWalletClient {
	return h.client
}

func NewHandler() *Handler {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Panicf("Failed to dial: %v", err)
	}

	rpcClient := hdwallet.NewHDWalletClient(conn)
	return &Handler{connection: conn, client: rpcClient}
}
