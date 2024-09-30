package repositories

import (
	"context"

	"github.com/ssoql/serviceutils/apierrors"

	"github.com/ssoql/auth-service/internal/app/entities"
)

type UserReader interface {
	GetByID(ctx context.Context, id int64) (*entities.User, apierrors.ApiError)
	GetByEmail(ctx context.Context, email string) (*entities.User, apierrors.ApiError)
	Exists(ctx context.Context, email string) (bool, apierrors.ApiError)
}

type UserWriter interface {
	Insert(ctx context.Context, user *entities.User) apierrors.ApiError
	Update(ctx context.Context, user *entities.User) apierrors.ApiError
	Delete(ctx context.Context, user *entities.User) apierrors.ApiError
}
