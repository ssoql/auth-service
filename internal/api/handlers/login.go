package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/ssoql/serviceutils/crypto"

	"github.com/ssoql/auth-service/internal/use_cases"
)

type loginHandler struct {
	useCase use_cases.LoginUseCase
}

func NewLoginHandler(loginUseCase use_cases.LoginUseCase) *loginHandler {
	return &loginHandler{useCase: loginUseCase}
}

func (h *loginHandler) Handle(c *gin.Context) {
	userInput, err := retrieveValidInput(c)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	jwt, err := h.useCase.Handle(c, userInput.Email, crypto.GetMd5(userInput.Password))
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(200, gin.H{"token": jwt})
}
