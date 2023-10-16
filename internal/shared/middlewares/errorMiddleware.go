package middlewares

import (
	"awesomeProject2/internal/api"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorMiddleware(ctx *gin.Context) {
	ctx.Next()
	if len(ctx.Errors) == 0 {
		return
	}

	code := GetHttpCodeByError(ctx.Errors[0])
	ctx.JSON(code, ctx.Errors[0].Error())
}

func GetHttpCodeByError(err error) int {
	var code int

	switch {
	case
		errors.Is(err, api.ValidationErr):
		code = http.StatusUnprocessableEntity
	default:
		code = http.StatusInternalServerError
	}

	return code
}
