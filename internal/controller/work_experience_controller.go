package controller

import (
	"CliPorto/internal/service"
	"CliPorto/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type WorkExperienceController struct {
	service *service.WorkExperienceService
}

func NewWorkExperienceController(s *service.WorkExperienceService) *WorkExperienceController {
	return &WorkExperienceController{service: s}
}

func (c *WorkExperienceController) GetAll(ctx *gin.Context) {
	data, err := c.service.GetAll()
	if err != nil {
		utils.JSONError(ctx, err.(*utils.APIError))
		return
	}

	utils.JSONSuccess(ctx, http.StatusOK, data)
}

func (c *WorkExperienceController) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		utils.JSONError(ctx, err.(*utils.APIError))
		return
	}

	data, apiErr := c.service.GetByID(id)
	if apiErr != nil {
		utils.JSONError(ctx, apiErr.(*utils.APIError))
		return
	}

	utils.JSONSuccess(ctx, http.StatusOK, data)
}
