package handler

import "github.com/gin-gonic/gin"

type ErrorResponse struct {
	message string
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, ErrorResponse{message: message})
}
