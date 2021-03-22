package handler

import (
	"net/http"

	"github.com/Valeriy-Totubalin/test_project/internal/delivery/request"
	"github.com/Valeriy-Totubalin/test_project/internal/delivery/response"
	"github.com/Valeriy-Totubalin/test_project/internal/domain"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createItem(c *gin.Context) {
	var data request.NewItem
	err := c.ShouldBindJSON(&data)
	if nil != err {
		c.JSON(http.StatusBadRequest, response.Error{Error: err.Error()})
		return
	}

	service, err := h.ServiceFactory.NewItemService()
	if nil != err {
		c.JSON(http.StatusInternalServerError, UnknowError)
		return
	}

	userId, err := h.GetCurrentUserId(c)
	if nil != err {
		c.JSON(http.StatusInternalServerError, UnknowError)
		return
	}

	item := &domain.Item{
		Name:   data.Name,
		UserId: userId,
	}

	err = service.Create(item)
	if nil != err {
		c.JSON(http.StatusInternalServerError, UnknowError)
		return
	}

	c.JSON(http.StatusCreated, response.CreatedItem{
		Message: ItemCreatedSuccess,
		Item: &response.Item{
			Id:   item.Id,
			Name: item.Name,
		},
	})
}

func (h *Handler) deleteItem(c *gin.Context) {
	var data request.DeleteItem
	err := c.ShouldBindUri(&data)
	if nil != err {
		c.JSON(http.StatusBadRequest, response.Error{Error: err.Error()})
		return
	}

	service, err := h.ServiceFactory.NewItemService()
	if nil != err {
		c.JSON(http.StatusInternalServerError, UnknowError)
		return
	}

	userId, err := h.GetCurrentUserId(c)
	if nil != err {
		c.JSON(http.StatusInternalServerError, UnknowError)
		return
	}

	isOwner, err := service.IsOwner(data.Id, userId)
	if nil != err {
		c.JSON(http.StatusInternalServerError, UnknowError)
		return
	}

	if !isOwner {
		c.JSON(http.StatusBadRequest, response.Error{Error: ItemNoCurrentUser})
		return
	}

	err = service.Delete(&domain.Item{Id: data.Id})
	if nil != err {
		c.JSON(http.StatusInternalServerError, UnknowError)
		return
	}

	c.JSON(http.StatusOK, response.Message{Message: ItemDeletedSuccess})
}
