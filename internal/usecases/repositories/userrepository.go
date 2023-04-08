package repositories

import (
	"context"

	"github.com/ssoql/auth-service/internal/app/entities"
	"github.com/ssoql/serviceutils/apierrors"
)

type UserReadRepository interface {
	GetByID(ctx context.Context, id int64) (*entities.User, apierrors.ApiError)
	GetByEmail(ctx context.Context, email string) (*entities.User, apierrors.ApiError)
	Exists(ctx context.Context, email string) (bool, apierrors.ApiError)
}

type UserWriteRepository interface {
	Insert(ctx context.Context, user *entities.User) apierrors.ApiError
	Update(ctx context.Context, user *entities.User) apierrors.ApiError
	Delete(ctx context.Context, user *entities.User) apierrors.ApiError
}
