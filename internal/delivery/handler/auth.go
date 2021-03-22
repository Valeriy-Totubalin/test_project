package handler

import (
	"net/http"

	"github.com/Valeriy-Totubalin/test_project/internal/delivery/request"
	"github.com/Valeriy-Totubalin/test_project/internal/delivery/response"
	"github.com/Valeriy-Totubalin/test_project/internal/domain"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	var data request.SignUp
	err := c.ShouldBindJSON(&data)
	if nil != err {
		c.JSON(http.StatusBadRequest, response.Error{Error: err.Error()})
		return
	}

	service, err := h.ServiceFactory.NewAuthService()
	if nil != err {
		c.JSON(http.StatusInternalServerError, response.Error{Error: UnknowError})
		return
	}

	user := &domain.User{
		Login:    data.Login,
		Password: data.Password,
	}

	isExists := service.IsExists(user)
	if isExists {
		c.JSON(http.StatusBadRequest, response.Error{Error: UserAlreadyExists})
		return
	}

	err = service.SignUp(user)
	if nil != err {
		c.JSON(http.StatusInternalServerError, response.Error{Error: UnknowError})
		return
	}

	c.JSON(http.StatusOK, response.Message{Message: RegistrationSucces})
}

func (h *Handler) signIn(c *gin.Context) {
	var data request.SignIn
	err := c.ShouldBindJSON(&data)
	if nil != err {
		c.JSON(http.StatusBadRequest, response.Error{Error: err.Error()})
		return
	}

	service, err := h.ServiceFactory.NewAuthService()
	if nil != err {
		c.JSON(http.StatusInternalServerError, response.Error{Error: UnknowError})
		return
	}

	user := &domain.User{
		Login:    data.Login,
		Password: data.Password,
	}

	token, err := service.SignIn(user)
	if nil != err {
		c.JSON(http.StatusInternalServerError, response.Error{Error: UnknowError})
		return
	}

	c.JSON(http.StatusOK, response.Token{AccessToken: token})
}
