package repository

import (
	"context"
	"time"

	"github.com/faizalpribadi/layout/internal/constant"
	"github.com/faizalpribadi/layout/internal/domain"

	"github.com/faizalpribadi/layout/pkg/helper"

	"github.com/faizalpribadi/layout/internal/model"
)

type (
	UserRepository struct {
		Timeout time.Duration
	}
)

func NewUserRepository() domain.UserRepository {
	return &UserRepository{
		Timeout: constant.Timeout * time.Second,
	}
}

func (r *UserRepository) GetUsers(ctx context.Context) (model.User, error) {
	user := model.User{}
	uid := helper.GenerateUID()
	user.ID = uid
	return user, nil
}
