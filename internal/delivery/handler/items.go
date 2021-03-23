package handler

import (
	"net/http"

	"github.com/Valeriy-Totubalin/test_project/internal/delivery/request"
	"github.com/Valeriy-Totubalin/test_project/internal/delivery/response"
	"github.com/Valeriy-Totubalin/test_project/internal/domain"
	"github.com/gin-gonic/gin"
)

// @Summary create
// @Security ApiKeyAuth
// @Tags items
// @Description Create new item for current user
// @ID createItem
// @Accept  json
// @Produce  json
// @Param input body request.NewItem true "New item data"
// @Success 201 {object} response.CreatedItem
// @Failure 400 {object} response.Error
// @Failure 401 {object} response.Error
// @Failure 500 {object} response.Error
// @Router /api/v1/items/new [post]
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

// @Summary delete
// @Security ApiKeyAuth
// @Tags items
// @Description Delete item by Id
// @ID deleteItem
// @Accept  json
// @Produce  json
// @Param input body request.DeleteItem true "Item Id"
// @Success 200 {object} response.Message
// @Failure 400 {object} response.Error
// @Failure 401 {object} response.Error
// @Failure 500 {object} response.Error
// @Router /api/v1/items/{id} [delete]
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

// @Summary get
// @Security ApiKeyAuth
// @Tags items
// @Description Get items fo current user
// @ID getItems
// @Accept  json
// @Produce  json
// @Success 200 {object} []response.Item
// @Failure 400 {object} response.Error
// @Failure 401 {object} response.Error
// @Failure 500 {object} response.Error
// @Router /api/v1/items [get]
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

// @Summary send
// @Security ApiKeyAuth
// @Tags items
// @Description Send item
// @ID sendItem
// @Accept  json
// @Produce  json
// @Param input body request.SendItem true "Item data for send"
// @Success 200 {object} response.TempLink
// @Failure 400 {object} response.Error
// @Failure 401 {object} response.Error
// @Failure 500 {object} response.Error
// @Router /api/v1/send [post]
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

// @Summary get
// @Security ApiKeyAuth
// @Tags items
// @Description Confirm send item
// @ID confirm
// @Accept  json
// @Produce  json
// @Param input body request.SendItem true "link"
// @Success 200 {object} response.Message
// @Failure 400 {object} response.Error
// @Failure 401 {object} response.Error
// @Failure 500 {object} response.Error
// @Router /api/v1/get [get]
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
