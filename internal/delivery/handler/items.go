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
		c.JSON(http.StatusInternalServerError, response.Error{Error: UnknowError})
		return
	}

	userId, err := h.GetCurrentUserId(c)
	if nil != err {
		c.JSON(http.StatusInternalServerError, response.Error{Error: UnknowError})
		return
	}

	item := &domain.Item{
		Name:   data.Name,
		UserId: userId,
	}

	err = service.Create(item)
	if nil != err {
		c.JSON(http.StatusInternalServerError, response.Error{Error: UnknowError})
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
		c.JSON(http.StatusInternalServerError, response.Error{Error: UnknowError})
		return
	}

	userId, err := h.GetCurrentUserId(c)
	if nil != err {
		c.JSON(http.StatusInternalServerError, response.Error{Error: UnknowError})
		return
	}

	isOwner, err := service.IsOwner(data.Id, userId)
	if nil != err {
		c.JSON(http.StatusInternalServerError, response.Error{Error: UnknowError})
		return
	}

	if !isOwner {
		c.JSON(http.StatusBadRequest, response.Error{Error: ItemNoCurrentUser})
		return
	}

	err = service.Delete(&domain.Item{Id: data.Id})
	if nil != err {
		c.JSON(http.StatusInternalServerError, response.Error{Error: UnknowError})
		return
	}

	c.JSON(http.StatusOK, response.Message{Message: ItemDeletedSuccess})
}

func (h *Handler) getItems(c *gin.Context) {
	service, err := h.ServiceFactory.NewItemService()
	if nil != err {
		c.JSON(http.StatusInternalServerError, response.Error{Error: UnknowError})
		return
	}

	userId, err := h.GetCurrentUserId(c)
	if nil != err {
		c.JSON(http.StatusInternalServerError, response.Error{Error: UnknowError})
		return
	}

	items, err := service.GetAll(userId)
	if nil != err {
		c.JSON(http.StatusInternalServerError, response.Error{Error: UnknowError})
		return
	}

	var responseItems []*response.Item

	for _, item := range items {
		responseItems = append(responseItems, &response.Item{
			Id:   item.Id,
			Name: item.Name,
		})
	}

	c.JSON(http.StatusOK, responseItems)
}

func (h *Handler) sendItem(c *gin.Context) {
	var data request.SendItem
	err := c.ShouldBindJSON(&data)
	if nil != err {
		c.JSON(http.StatusBadRequest, response.Error{Error: err.Error()})
		return
	}

	service, err := h.ServiceFactory.NewItemService()
	if nil != err {
		c.JSON(http.StatusInternalServerError, response.Error{Error: UnknowError})
		return
	}

	userId, err := h.GetCurrentUserId(c)
	if nil != err {
		c.JSON(http.StatusInternalServerError, response.Error{Error: UnknowError})
		return
	}

	isOwner, err := service.IsOwner(data.ItemId, userId)
	if nil != err {
		c.JSON(http.StatusInternalServerError, response.Error{Error: UnknowError})
		return
	}

	if !isOwner {
		c.JSON(http.StatusForbidden, response.Error{Error: ItemNoCurrentUser})
		return
	}

	link := &domain.Link{
		ItemId:    data.ItemId,
		UserLogin: data.UserLogin,
	}

	tempLink, err := service.GetTempLink(link)
	if nil != err {
		c.JSON(http.StatusInternalServerError, response.Error{Error: UnknowError})
		return
	}

	c.JSON(http.StatusOK, response.TempLink{Link: tempLink})
}

func (h *Handler) confirm(c *gin.Context) {
	var data request.Confirm
	err := c.ShouldBindUri(&data)
	if nil != err {
		c.JSON(http.StatusBadRequest, response.Error{Error: err.Error()})
		return
	}

	service, err := h.ServiceFactory.NewItemService()
	if nil != err {
		c.JSON(http.StatusInternalServerError, response.Error{Error: UnknowError})
		return
	}

	userId, err := h.GetCurrentUserId(c)
	if nil != err {
		c.JSON(http.StatusInternalServerError, response.Error{Error: UnknowError})
		return
	}

	CanConfirm, err := service.CanConfirm(data.Link, userId)
	if nil != err {
		c.JSON(http.StatusInternalServerError, response.Error{Error: UnknowError})
		return
	}

	if !CanConfirm {
		c.JSON(http.StatusForbidden, response.Error{Error: NoGetItem})
		return
	}

	link, err := h.LinkManager.Parse(data.Link)
	if nil != err {
		c.JSON(http.StatusInternalServerError, response.Error{Error: UnknowError})
		return
	}

	isDeleted, err := service.IsDeleted(link.ItemId)
	if nil != err {
		c.JSON(http.StatusInternalServerError, response.Error{Error: UnknowError})
		return
	}

	if isDeleted {
		c.JSON(http.StatusBadRequest, response.Error{Error: ItemTransferedOrDeleted})
		return
	}

	err = service.Confirm(data.Link, userId)
	if nil != err {
		c.JSON(http.StatusInternalServerError, response.Error{Error: UnknowError})
		return
	}

	c.JSON(http.StatusOK, response.Message{Message: ObjectReceived})
}
