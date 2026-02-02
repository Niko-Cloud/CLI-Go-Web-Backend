package utils

import "github.com/gin-gonic/gin"

func JSONError(ctx *gin.Context, err *APIError) {
	ctx.JSON(err.Code, err)
}

func JSONSuccess(ctx *gin.Context, status int, data any) {
	ctx.JSON(status, data)
}
