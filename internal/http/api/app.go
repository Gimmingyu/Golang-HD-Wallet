package api

import (
	"github.com/gin-gonic/gin"
)

type (
	GinApp struct {
		router  *gin.Engine
		handler IHandler
	}

	IGinApp interface {
		SetRouter()
		Run() error
	}
)

func NewGinApp() *GinApp {
	r := gin.Default()
	h := NewHandler()

	return &GinApp{router: r, handler: h}
}

func (g *GinApp) SetRouter() {
	g.router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"result": "Success"})
	})
	g.router.GET("/public/:UserId", g.handler.GetChildPublicKey)
	g.router.GET("/private/:UserId", g.handler.GetChildPrivateKey)
}

func (g *GinApp) Run() error {
	///* connection Defer로 끊어주기 */
	//defer func(conn *grpc.ClientConn) {
	//	err := conn.Close()
	//	if err != nil {
	//		log.Panicf("Failed to close rpc connection: %v", err)
	//	}
	//}(g.handler.GetConnection())
	return g.router.Run()
}
