package controller

import (
	"CliPorto/internal/service"
	"CliPorto/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProfileController struct {
	service *service.ProfileService
}

func NewProfileController(s *service.ProfileService) *ProfileController {
	return &ProfileController{service: s}
}

func (c *ProfileController) GetAll(ctx *gin.Context) {
	data, err := c.service.GetAll()
	if err != nil {
		utils.JSONError(ctx, err.(*utils.APIError))
		return
	}

	utils.JSONSuccess(ctx, http.StatusOK, data)
}
