package usecases

import (
	"context"
	"github.com/golang-jwt/jwt"
	"github.com/ssoql/auth-service/config"
	"github.com/ssoql/auth-service/internal/usecases/repositories"
	"github.com/ssoql/serviceutils/apierrors"
	"time"
)

type LoginUseCase interface {
	Handle(ctx context.Context, username string, password string) (string, apierrors.ApiError)
}

type loginUseCase struct {
	dbRead repositories.UserReadRepository
}

func NewLoginUseCase(readRepository repositories.UserReadRepository) *loginUseCase {
	return &loginUseCase{dbRead: readRepository}
}

func (u *loginUseCase) Handle(ctx context.Context, username string, password string) (string, apierrors.ApiError) {
	user, err := u.dbRead.GetByEmail(ctx, username)
	if err != nil {
		return "", err
	}

	if user.Password != password {
		return "", apierrors.NewBadRequestError("wrong password")
	}

	jwt, err := createJWT(ctx, username)
	if err != nil {
		return "", err
	}

	return jwt, nil
}

func createJWT(ctx context.Context, username string) (string, apierrors.ApiError) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(5 * time.Minute).Unix()
	claims["created"] = time.Now().Unix()
	claims["isAdmin"] = true

	tokenString, err := token.SignedString([]byte(config.SecurityKey))
	if err != nil {
		return "", apierrors.NewInternalServerError("token generation failed", err)
	}

	return tokenString, nil
}
