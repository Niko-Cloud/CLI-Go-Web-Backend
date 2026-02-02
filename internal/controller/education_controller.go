package controller

import (
	"CliPorto/internal/service"
	"CliPorto/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EducationController struct {
	service *service.EducationService
}

func NewEducationController(s *service.EducationService) *EducationController {
	return &EducationController{service: s}
}

func (c *EducationController) GetAll(ctx *gin.Context) {
	data, err := c.service.GetAll()
	if err != nil {
		utils.JSONError(ctx, err.(*utils.APIError))
		return
	}

	utils.JSONSuccess(ctx, http.StatusOK, data)
}
