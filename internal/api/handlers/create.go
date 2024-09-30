package handlers

import (
	"net/http"

	"github.com/ssoql/serviceutils/apierrors"

	"github.com/gin-gonic/gin"

	"github.com/ssoql/auth-service/internal/use_cases"
)

type userCreateHandler struct {
	useCase use_cases.SaveUserUseCase
}

func NewUserCreateHandler(saveUserUseCase use_cases.SaveUserUseCase) *userCreateHandler {
	return &userCreateHandler{useCase: saveUserUseCase}
}

func (h *userCreateHandler) Handle(c *gin.Context) {
	userInput, err := retrieveValidInput(c)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	result, err := h.useCase.Handle(c, userInput.Email, userInput.Password)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func retrieveValidInput(c *gin.Context) (*UserInput, apierrors.ApiError) {
	var input UserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		return nil, apierrors.NewBadRequestError("invalid json data")
	}
	// todo add validation of parameters

	return &input, nil
}
