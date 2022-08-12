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
	server.Static("/css", "../../templates/css")
	server.LoadHTMLFiles(
		"/home/fg/go/src/ginApi/internal/templates/header.html",
		"/home/fg/go/src/ginApi/internal/templates/index.html",
		"/home/fg/go/src/ginApi/internal/templates/footer.html",
	)

	server.Use(
		gin.Recovery(),
		middleware.Logger(),
		middleware.BasicAuth(),
		ginDump.Dump(),
	)

	apiRoutes := server.Group("/api")
	{
		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, videoHandler.FindAll())
		})

		apiRoutes.POST("/videos", func(ctx *gin.Context) {
			err := videoHandler.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Ok"})
			}

		})

		viewRoutes := server.Group("/view")
		{
			viewRoutes.GET("/videos", videoHandler.ShowAll)
		}

	}

	server.Run(":8080")

}
