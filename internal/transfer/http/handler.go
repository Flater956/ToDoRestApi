package http

import (
	v1 "ToDoRestApi/internal/transfer/http/v1"
	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	h.initApi(router)

	return router
}

func (h *Handler) initApi(router *gin.Engine) {
	v1Handler := v1.NewHandler()

	v1 := router.Group("/api/v1")
	{
		v1Handler.Init(v1)
	}
}
