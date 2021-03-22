package handler

import (
	"net/http"

	"github.com/Valeriy-Totubalin/test_project/internal/app/interfaces/pkg_interfaces"
	"github.com/gin-gonic/gin"
)

const UnknowError = "Unknown error"

type Handler struct {
	TokenManager pkg_interfaces.TokenManager
	// ServiceFactory interfaces.ServiceFactory
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	// auth := router.Group("/auth")
	// {
	// 	v1 := auth.Group("/v1")
	// 	{
	// 		v1.POST("/registration", h.signUp)
	// 	}
	// }

	return router
}
