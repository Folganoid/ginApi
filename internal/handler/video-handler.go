package handler

import (
	"ginApi/internal/model"
	"ginApi/internal/service"
	"ginApi/pkg/validators"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

type VideoHandler interface {
	FindAll() []model.Video
	Save(ctx *gin.Context) error
}

type handler struct {
	service service.VideoService
}

func New(service service.VideoService) VideoHandler {
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	return &handler {
		service: service,
	}
}

func (h *handler) FindAll() []model.Video {
	return h.service.FindAll()
}

func (h *handler) Save(ctx *gin.Context) error {
	var video model.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}
	err = validate.Struct(video)
	if err != nil {
		return err
	}

	h.service.Save(video)
	return nil
}
