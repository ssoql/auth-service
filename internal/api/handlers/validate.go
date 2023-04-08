package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/ssoql/auth-service/internal/usecases"
	"github.com/ssoql/serviceutils/apierrors"
)

type validateHandler struct {
	useCase usecases.ValidateUseCase
}

func NewValidateHandler(validateSseCase usecases.ValidateUseCase) *validateHandler {
	return &validateHandler{useCase: validateSseCase}
}

func (h *validateHandler) Handle(c *gin.Context) {
	token, err := retrieveAuthToken(c)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	valid, err := h.useCase.Handle(c, token)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(200, gin.H{"authorized": valid})
}

func retrieveAuthToken(c *gin.Context) (string, apierrors.ApiError) {
	token := c.GetHeader("Token")
	if token == "" {
		return "", apierrors.NewBadRequestError("missing auth token")
	}

	return token, nil
}
