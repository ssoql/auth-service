package use_cases

import (
	"context"

	"github.com/ssoql/auth-service/internal/use_cases/repositories"

	"github.com/ssoql/serviceutils/apierrors"

	"github.com/ssoql/auth-service/internal/app/entities"
)

type SaveUserUseCase interface {
	Handle(ctx context.Context, question, answer string) (*entities.User, apierrors.ApiError)
}

type saveUserUseCase struct {
	dbRead  repositories.UserReader
	dbWrite repositories.UserWriter
}

func NewSaveUserUseCase(readRepository repositories.UserReader, writeRepository repositories.UserWriter) *saveUserUseCase {
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
