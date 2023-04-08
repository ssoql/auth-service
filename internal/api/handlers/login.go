package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/ssoql/auth-service/internal/usecases"
	"github.com/ssoql/serviceutils/apierrors"
	"github.com/ssoql/serviceutils/crypto"
)

type loginHandler struct {
	useCase usecases.LoginUseCase
}

func NewLoginHandler(loginUseCase usecases.LoginUseCase) *loginHandler {
	return &loginHandler{useCase: loginUseCase}
}

func (h *loginHandler) Handle(c *gin.Context) {
	username, password, err := retrieveValidAuthData(c)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	jwt, err := h.useCase.Handle(c, *username, *password)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(200, gin.H{"token": jwt})
}

func retrieveValidAuthData(c *gin.Context) (*string, *string, apierrors.ApiError) {
	username, password, ok := c.Request.BasicAuth()
	if !ok {
		return nil, nil, apierrors.NewBadRequestError("wrong auth data")
	}
	password = crypto.GetMd5(password)

	return &username, &password, nil
}
