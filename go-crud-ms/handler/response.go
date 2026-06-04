package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, code int, message string) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(code, gin.H{
		"message":   message,
		"errorCode": code,
	})
}

func sendSuccess(ctx *gin.Context, op string, data interface{}, code int) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(code, gin.H{
		"message": fmt.Sprintf("operation from handler : %s successfull", op),
		"data":    data,
	})
}
