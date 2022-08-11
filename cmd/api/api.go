package main

import (
	"ginApi/internal/handler"
	"ginApi/internal/service"
	"ginApi/pkg/middleware"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	ginDump "github.com/tpkeeper/gin-dump"
)

var (
	videoService service.VideoService = service.New()
	videoHandler handler.VideoHandler = handler.New(videoService)
)

func setupLogOutput() {
	f, _ := os.Create("api.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	setupLogOutput()

	server := gin.New()
	server.Use(
		gin.Recovery(),
		middleware.Logger(),
		middleware.BasicAuth(),
		ginDump.Dump(),
		)

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoHandler.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoHandler.Save(ctx))
	})


	server.Run(":8080")

}
