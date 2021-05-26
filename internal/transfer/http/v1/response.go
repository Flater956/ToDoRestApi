package v1

import (
	"github.com/gin-gonic/gin"
	"log"
)

type respMessage struct {
	Message string `json:"message"`
}

func newErrorMessage(ctx *gin.Context, statusCode int, message string)  {
	log.Fatal(message)

	ctx.AbortWithStatusJSON(statusCode, respMessage{message})
}