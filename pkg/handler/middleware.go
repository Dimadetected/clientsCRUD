package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) UserIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		NewErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		NewErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userId)
}

func (h *Handler) getUserId(c *gin.Context) (int, error) {
	id, err := c.Get(userCtx)
	if !err {
		NewErrorResponse(c, http.StatusUnauthorized, "user not found")
		return 0, errors.New("user not found")
	}

	userId, err := id.(int)
	if !err {
		NewErrorResponse(c, http.StatusUnauthorized, "user not found")
		return 0, errors.New("user not found")
	}

	return userId, nil
}
