package controller

import (
	"CliPorto/internal/service"
	"CliPorto/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ContactController struct {
	service *service.ContactService
}

func NewContactController(s *service.ContactService) *ContactController {
	return &ContactController{service: s}
}

func (c *ContactController) GetContactInfo(ctx *gin.Context) {
	data, err := c.service.GetContactInfo()
	if err != nil {
		utils.JSONError(ctx, err.(*utils.APIError))
		return
	}

	utils.JSONSuccess(ctx, http.StatusOK, data)
}
