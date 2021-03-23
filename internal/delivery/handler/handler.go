package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Valeriy-Totubalin/test_project/internal/app/interfaces/factories_interfaces"
	"github.com/Valeriy-Totubalin/test_project/internal/app/interfaces/pkg_interfaces"
	"github.com/gin-gonic/gin"
)

const UnknowError = "unknown error"
const RegistrationSucces = "registration completed successfully"
const UserAlreadyExists = "user already exists"
const ItemCreatedSuccess = "item created successfully"
const ItemDeletedSuccess = "item deleted successfully"
const ItemNoCurrentUser = "item is not owned by the current user"
const NoGetItem = "you cannot get this item"
const ObjectReceived = "object received"
const YouOwner = "you are the owner of this item"
const ItemTransferedOrDeleted = "the item was transferred or deleted"

type Handler struct {
	TokenManager   pkg_interfaces.TokenManager
	ServiceFactory factories_interfaces.ServicesFactory
	LinkManager    pkg_interfaces.LinkManager
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	auth := router.Group("/auth")
	{
		v1 := auth.Group("/v1")
		{
			v1.POST("/registration", h.signUp)
			v1.POST("/login", h.signIn)
		}
	}

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		v1.Use(h.checkToken) // middleware
		{
			items := v1.Group("/items")
			{
				items.POST("/new", h.createItem)
				items.DELETE("/:id", h.deleteItem)
				items.GET("", h.getItems)
			}
			v1.POST("/send", h.sendItem)
			v1.GET("/get/:link", h.confirm)
		}
	}

	return router
}

func (h *Handler) GetCurrentUserId(c *gin.Context) (int, error) {
	userId := c.MustGet("user_id")
	if nil == userId {
		return 0, errors.New("no current user")
	}
	id, _ := strconv.Atoi(userId.(string))

	return id, nil
}
