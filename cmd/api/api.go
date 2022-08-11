package main

import (
	"ginApi/internal/handler"
	"ginApi/internal/service"
	"ginApi/pkg/middleware"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
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
		ctx.JSON(http.StatusOK, videoHandler.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		err := videoHandler.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "Ok"})
		}

	})


	server.Run(":8080")

}
