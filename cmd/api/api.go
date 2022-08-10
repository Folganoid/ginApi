package main

import (
	"ginApi/internal/handler"
	"ginApi/internal/service"
	"github.com/gin-gonic/gin"
)

var (
	videoService service.VideoService = service.New()
	videoHandler handler.VideoHandler = handler.New(videoService)
)

func main() {
	server := gin.Default()

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoHandler.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoHandler.Save(ctx))
	})


	server.Run(":8080")

}
