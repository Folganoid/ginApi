package handler

import (
	"ginApi/internal/model"
	"ginApi/internal/service"
	"github.com/gin-gonic/gin"
)

type VideoHandler interface {
	FindAll() []model.Video
	Save(ctx *gin.Context) model.Video
}

type handler struct {
	service service.VideoService
}

func New(service service.VideoService) VideoHandler {
	return &handler {
		service: service,
	}
}

func (h *handler) FindAll() []model.Video {
	return h.service.FindAll()
}

func (h *handler) Save(ctx *gin.Context) model.Video {
	var video model.Video
	ctx.BindJSON(&video)
	h.service.Save(video)
	return video
}
