package controller

import (
	"CliPorto/internal/service"
	"CliPorto/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommandController struct {
	service *service.CommandService
}

func NewCommandController(s *service.CommandService) *CommandController {
	return &CommandController{service: s}
}

func (c *CommandController) GetAll(ctx *gin.Context) {
	data, err := c.service.GetAll()
	if err != nil {
		utils.JSONError(ctx, err.(*utils.APIError))
		return
	}

	utils.JSONSuccess(ctx, http.StatusOK, data)
}

func (c *CommandController) GetByName(ctx *gin.Context) {
	name := ctx.Param("name")

	data, err := c.service.GetByName(name)
	if err != nil {
		utils.JSONError(ctx, err.(*utils.APIError))
		return
	}

	utils.JSONSuccess(ctx, http.StatusOK, data)
}
