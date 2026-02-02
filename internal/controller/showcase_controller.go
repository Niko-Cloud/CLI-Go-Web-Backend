package controller

import (
	"CliPorto/internal/service"
	"CliPorto/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ShowcaseController struct {
	service *service.ShowcaseService
}

func NewShowcaseController(s *service.ShowcaseService) *ShowcaseController {
	return &ShowcaseController{service: s}
}

func (c *ShowcaseController) GetAll(ctx *gin.Context) {
	data, err := c.service.GetAll()
	if err != nil {
		utils.JSONError(ctx, utils.ErrInternal)
		return
	}

	utils.JSONSuccess(ctx, http.StatusOK, data)
}

func (c *ShowcaseController) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		utils.JSONError(ctx, utils.ErrBadRequest)
		return
	}

	data, apiErr := c.service.GetByID(id)
	if apiErr != nil {
		utils.JSONError(ctx, apiErr.(*utils.APIError))
		return
	}

	utils.JSONSuccess(ctx, http.StatusOK, data)
}
