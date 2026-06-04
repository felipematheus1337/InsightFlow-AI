package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetIdFromQuery(ctx *gin.Context) (string, bool) {
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, "ID IS required")
		return "", false
	}

	return id, true
}

func ErrParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}
