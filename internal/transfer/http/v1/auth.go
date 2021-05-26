package v1

import (
	"ToDoRestApi/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) singIn(ctx *gin.Context) {

}

func (h *Handler) singUp(ctx *gin.Context) {
	var input domain.User

	if err := ctx.BindJSON(&input); err != nil {
		newErrorMessage(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)

	if err != nil {
		newErrorMessage(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
