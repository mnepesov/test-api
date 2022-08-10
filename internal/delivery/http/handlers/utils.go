package handlers

import (
	"back/internal/domains"
	"github.com/gin-gonic/gin"
)

func sendErrorResponse(ctx *gin.Context, msg string, status int) {
	resp := domains.ErrorResponse{
		Message: msg,
	}
	
	ctx.JSON(status, resp)
}

func sendSuccessResponse(ctx *gin.Context, data interface{}, status int) {
	ctx.JSON(status, data)
}
