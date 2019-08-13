package usecase

import (
	"context"
	"time"

	"github.com/faizalpribadi/layout/internal/constant"
	"github.com/faizalpribadi/layout/internal/domain"
	"github.com/faizalpribadi/layout/internal/model"
)

type UserUseCase struct {
	userRepository domain.UserRepository
	Timeout        time.Duration
}

func NewUserUseCase(repository domain.UserRepository) domain.UserUseCase {
	return &UserUseCase{
		userRepository: repository,
		Timeout:        constant.Timeout * time.Second,
	}
}

func (u *UserUseCase) GetUsers(ctx context.Context) (users model.User, err error) {
	c, cancel := context.WithTimeout(ctx, u.Timeout)
	defer cancel()

	users, err = u.userRepository.GetUsers(c)
	return
}
