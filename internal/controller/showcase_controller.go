package controller

import (
	"CliPorto/internal/service"
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
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, data)
}

func (c *ShowcaseController) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	data, err := c.service.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if data == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	ctx.JSON(http.StatusOK, data)
}
