package main

import (
	"awesomeProject2/internal/shared/middlewares"
	"github.com/gin-gonic/gin"
)

var (
	engine *gin.Engine
	group  *gin.RouterGroup
)

func init() {
	engine = gin.Default()
	engine.Use(middlewares.ErrorMiddleware)
	group = engine.Group("api/v1")
}
