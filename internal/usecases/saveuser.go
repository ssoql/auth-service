package usecases

import (
	"context"
	"github.com/ssoql/auth-service/internal/usecases/repositories"

	"github.com/ssoql/auth-service/internal/app/entities"
	"github.com/ssoql/serviceutils/apierrors"
)

type SaveUserUseCase interface {
	Handle(ctx context.Context, question, answer string) (*entities.User, apierrors.ApiError)
}

type saveUserUseCase struct {
	dbRead  repositories.UserReadRepository
	dbWrite repositories.UserWriteRepository
}

func NewSaveUserUseCase(readRepository repositories.UserReadRepository, writeRepository repositories.UserWriteRepository) *saveUserUseCase {
	return &saveUserUseCase{dbRead: readRepository, dbWrite: writeRepository}
}

func (u *saveUserUseCase) Handle(ctx context.Context, email, password string) (*entities.User, apierrors.ApiError) {
	newUser := &entities.User{
		Email:    email,
		Password: password,
	}

	if err := u.dbWrite.Insert(ctx, newUser); err != nil {
		return nil, err
	}

	return newUser, nil
}
