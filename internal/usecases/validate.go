package usecases

import (
	"context"
	"github.com/golang-jwt/jwt"
	"github.com/ssoql/auth-service/config"
	"github.com/ssoql/serviceutils/apierrors"
)

type ValidateUseCase interface {
	Handle(ctx context.Context, token string) (bool, apierrors.ApiError)
}

type validateUseCase struct{}

func NewValidateUseCase() *validateUseCase {
	return &validateUseCase{}
}

func (u validateUseCase) Handle(ctx context.Context, token string) (bool, apierrors.ApiError) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.SecurityKey), nil
	})

	if err != nil {
		return false, apierrors.NewUnauthorizedError("No token provided!")
	}

	if !parsedToken.Valid {
		return false, apierrors.NewUnauthorizedError("Unauthorized!")
	}

	if err != nil {
		return false, apierrors.NewInternalServerError("Authorization error!", err)
	}

	return true, nil
}
