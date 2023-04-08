package infrastructure

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/ssoql/auth-service/internal/app/entities"
	"github.com/ssoql/auth-service/internal/infrastructure/db"
	"github.com/ssoql/serviceutils/apierrors"
)

type userReadRepository struct {
	db *db.ClientDB
}

func NewUserReadRepository(db *db.ClientDB) *userReadRepository {
	return &userReadRepository{db: db}
}

func (r *userReadRepository) GetByID(ctx context.Context, id int64) (*entities.User, apierrors.ApiError) {
	faq := &entities.User{}

	if err := r.db.WithContext(ctx).Where("id = ?", id).First(faq).Error; err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "record not found") {
			return nil, apierrors.NewNotFoundError("faq with given id does not exists")
		}
		return nil, apierrors.NewInternalServerError(
			"error when tying to fetch faq",
			fmt.Errorf("database error: %s", err.Error()),
		)
	}

	return faq, nil
}

func (r *userReadRepository) GetByEmail(ctx context.Context, email string) (*entities.User, apierrors.ApiError) {
	user := &entities.User{}

	if err := r.db.WithContext(ctx).Where("email = ?", email).First(user).Error; err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "record not found") {
			return nil, apierrors.NewNotFoundError("user does not exists")
		}
		return nil, apierrors.NewInternalServerError(
			"error when tying to fetch user",
			fmt.Errorf("database error: %s", err.Error()),
		)
	}

	return user, nil
}

func (r *userReadRepository) Exists(ctx context.Context, email string) (bool, apierrors.ApiError) {
	var exists bool
	err := r.db.WithContext(ctx).Model(&entities.User{}).
		Select("count(*) > 0").
		Where("email = ?", email).
		Find(&exists).
		Error

	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "record not found") {
			return false, apierrors.NewNotFoundError("user with given email does not exists")
		}
		return false, apierrors.NewInternalServerError(
			"error when tying to fetch user",
			errors.New("database error"),
		)
	}

	return exists, nil
}
