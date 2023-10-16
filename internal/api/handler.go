package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler interface {
	GetByDate(ctx *gin.Context)
	GetBetweenDateRange(ctx *gin.Context)
	GetFileFromStorage(ctx *gin.Context)
}

type handler struct {
	controller Controller
}

func NewHandler(group *gin.RouterGroup, controller Controller) Handler {
	handlerImpl := &handler{
		controller: controller,
	}

	handlerImpl.registerEndpoints(group)
	return handlerImpl
}

func (h *handler) registerEndpoints(group *gin.RouterGroup) {
	group.GET(":date", h.GetByDate)
	group.GET("", h.GetBetweenDateRange)
	group.GET("storage/:path", h.GetFileFromStorage)
}

func (h *handler) GetByDate(ctx *gin.Context) {
	date := ctx.Param("date")
	event, err := h.controller.GetByDate(ctx, date)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, event)
}

func (h *handler) GetBetweenDateRange(ctx *gin.Context) {
	fromDate := ctx.Query("fromDate")
	toDate := ctx.Query("toDate")
	events, err := h.controller.GetBetweenDateRange(ctx, fromDate, toDate)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, events)
}

func (h *handler) GetFileFromStorage(ctx *gin.Context) {
	path := ctx.Param("path")
	file, err := h.controller.GetFileFromStorage(ctx, path)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	if _, err = ctx.Writer.Write(file); err != nil {
		_ = ctx.Error(err)
		return
	}
}
