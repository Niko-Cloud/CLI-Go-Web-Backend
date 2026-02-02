package controller

import (
	"CliPorto/internal/service"
	"CliPorto/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SkillController struct {
	service *service.SkillService
}

func NewSkillController(service *service.SkillService) *SkillController {
	return &SkillController{service: service}
}

func (c *SkillController) GetAllSkills(ctx *gin.Context) {
	data, err := c.service.GetAll()
	if err != nil {
		utils.JSONError(ctx, err.(*utils.APIError))
		return
	}

	utils.JSONSuccess(ctx, http.StatusOK, data)
}
