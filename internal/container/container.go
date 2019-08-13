//+build wireinject

package container

import (
	"context"

	"github.com/faizalpribadi/layout/internal/app/handler"
	"github.com/faizalpribadi/layout/internal/app/usecase"
	"github.com/faizalpribadi/layout/internal/config"
	"github.com/faizalpribadi/layout/internal/http"

	"github.com/faizalpribadi/layout/internal/app/repository"
	"github.com/google/wire"

	"github.com/faizalpribadi/layout/internal/domain"
)

func NewContainer(ctx context.Context) (domain.Server, error) {
	wire.Build(
		repository.NewUserRepository,
		usecase.NewUserUseCase,
		handler.NewUserHandler,
		repository.NewFeedRepository,
		usecase.NewFeedUseCase,
		handler.NewFeedHandler,
		config.NewConfiguration,
		http.NewServer,
	)
	return &http.Server{}, nil
}
