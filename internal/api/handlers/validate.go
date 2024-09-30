package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/ssoql/serviceutils/apierrors"

	"github.com/ssoql/auth-service/internal/use_cases"
)

type validateHandler struct {
	useCase use_cases.ValidateUseCase
}

func NewValidateHandler(validateSseCase use_cases.ValidateUseCase) *validateHandler {
	return &validateHandler{useCase: validateSseCase}
}

func (h *validateHandler) Handle(c *gin.Context) {
	input, err := retrieveTokenInput(c)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	valid, err := h.useCase.Handle(c, input.Token)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(200, gin.H{"authenticated": valid})
}

func retrieveTokenInput(c *gin.Context) (*TokenInput, apierrors.ApiError) {
	var input TokenInput

	if err := c.ShouldBindJSON(&input); err != nil {
		return nil, apierrors.NewBadRequestError("invalid json data")
	}

	return &input, nil
}
