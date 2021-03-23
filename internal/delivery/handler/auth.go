package handler

import (
	"net/http"

	"github.com/Valeriy-Totubalin/test_project/internal/delivery/request"
	"github.com/Valeriy-Totubalin/test_project/internal/delivery/response"
	"github.com/Valeriy-Totubalin/test_project/internal/domain"
	"github.com/gin-gonic/gin"
)

// @Summary sign-up
// @Tags auth
// @Description Create new account
// @ID sign-up
// @Accept  json
// @Produce  json
// @Param input body request.SignUp true "account info"
// @Success 201 {object} response.Message
// @Failure 400 {object} response.Error
// @Failure 500 {object} response.Error
// @Router /auth/v1/registration [post]
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

	c.JSON(http.StatusCreated, response.Message{Message: RegistrationSucces})
}

// @Summary sign-in
// @Tags auth
// @Description Log in with an existing account
// @ID sign-in
// @Accept  json
// @Produce  json
// @Param input body request.SignIn true "login and password from the account"
// @Success 200 {object} response.Token
// @Failure 400 {object} response.Error
// @Failure 500 {object} response.Error
// @Router /auth/v1/login [post]
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
