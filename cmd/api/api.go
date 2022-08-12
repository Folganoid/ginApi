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
	"path/filepath"
	"runtime"
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

	var (
		_, b, _, _ = runtime.Caller(0)
		basepath   = filepath.Dir(b)
	)

	setupLogOutput()

	server := gin.New()
	server.Static("/css", "../../templates/css")
	server.LoadHTMLFiles(
		basepath + "/../../internal/templates/header.html",
		basepath + "/../../internal/templates/index.html",
		basepath + "/../../internal/templates/footer.html",
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

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	server.Run(":" + port)

}
