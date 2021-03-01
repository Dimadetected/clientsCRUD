package handler

import (
	"github.com/gin-gonic/gin"
	"log"
)

type ErrorResponse struct {
	message string
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	log.Print(message)
	c.AbortWithStatusJSON(statusCode, ErrorResponse{message: message})
}
