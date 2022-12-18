package api

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
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

func NewGinApp() IGinApp {
	r := gin.Default()
	h := NewHandler()

	return &GinApp{router: r, handler: h}
}

func (g *GinApp) SetRouter() {

	//_ = g.router.SetTrustedProxies([]string{"http://localhost:8080"})
	g.router.Use(cors.New(
		cors.Config{
			//AllowOrigins:    []string{"http://127.0.0.1:3000"},
			AllowAllOrigins: true,
			AllowMethods:    []string{"GET", "POST"},
			MaxAge:          12 * time.Hour,
		}))
	g.router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	api := g.router.Group("api")
	api.Use(gin.Recovery())
	api.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"result": "Success"})
	})

	deposit := api.Group("deposit")
	deposit.Use(SetRequestUUID())
	{
		deposit.GET("/public/:UserId", g.handler.GetChildPublicKey)
		deposit.GET("/private/:UserId", g.handler.GetChildPrivateKey)
	}
}

func (g *GinApp) Run() error {
	return g.router.Run()
}
