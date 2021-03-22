package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/Valeriy-Totubalin/test_project/internal/delivery/response"
	"github.com/gin-gonic/gin"
)

const authorizationHeader = "Authorization"

func (h *Handler) checkToken(c *gin.Context) {
	id, err := h.parseAuthHeader(c)
	if nil != err {
		c.AbortWithStatusJSON(http.StatusUnauthorized, response.Error{Error: err.Error()})
	}
	c.Set("user_id", id)
}

func (h *Handler) parseAuthHeader(c *gin.Context) (int, error) {
	header := c.GetHeader(authorizationHeader)
	if "" == header {
		return 0, errors.New("empty auth header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return 0, errors.New("invalid auth header")
	}

	if len(headerParts[1]) == 0 {
		return 0, errors.New("token is empty")
	}

	id, err := h.TokenManager.Parse(headerParts[1])
	if nil != err {
		return 0, err
	}

	return id, nil
}
