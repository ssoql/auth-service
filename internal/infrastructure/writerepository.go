package infrastructure

import (
	"context"
	"errors"
	"log"
	"strings"

	"github.com/ssoql/auth-service/internal/app/entities"
	"github.com/ssoql/auth-service/internal/infrastructure/db"
	"github.com/ssoql/serviceutils/apierrors"
	"github.com/ssoql/serviceutils/crypto"
)

type userWriteRepository struct {
	db *db.ClientDB
}

func NewUserWriteRepository(db *db.ClientDB) *userWriteRepository {
	return &userWriteRepository{db: db}
}

func (r *userWriteRepository) Insert(ctx context.Context, user *entities.User) apierrors.ApiError {
	user.Password = crypto.GetMd5(user.Password)

	if err := r.db.WithContext(ctx).Create(user).Error; err != nil {
		log.Println("error when trying to prepare save user statement", err.Error())

		if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
			return apierrors.NewBadRequestError("this user already exists")
		}
		return apierrors.NewInternalServerError("error when tying to save user", errors.New("database error"))
	}

	return nil
}

func (r *userWriteRepository) Update(ctx context.Context, user *entities.User) apierrors.ApiError {
	user.Password = crypto.GetMd5(user.Password)

	if err := r.db.WithContext(ctx).Updates(&user).Error; err != nil {
		log.Println("error when trying to prepare save faq statement", err.Error())

		if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
			return apierrors.NewBadRequestError("this question already exists")
		}
		return apierrors.NewInternalServerError("error when tying to update faq", errors.New("database error"))
	}

	return nil
}

func (r *userWriteRepository) Delete(ctx context.Context, user *entities.User) apierrors.ApiError {
	if err := r.db.WithContext(ctx).Delete(user).Error; err != nil {
		return apierrors.NewInternalServerError("error when tying to delete faq", errors.New("database error"))
	}

	return nil
}
