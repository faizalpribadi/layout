package usecase

import (
	"context"
	"time"

	"github.com/faizalpribadi/layout/internal/constant"
	"github.com/faizalpribadi/layout/internal/domain"
	"github.com/faizalpribadi/layout/internal/model"
)

type FeedUseCase struct {
	feedRepository domain.FeedRepository
	Timeout        time.Duration
}

func NewFeedUseCase(repository domain.FeedRepository) domain.FeedUseCase {
	return &FeedUseCase{
		feedRepository: repository,
		Timeout:        constant.Timeout * time.Second,
	}
}

func (f *FeedUseCase) GetFeeds(ctx context.Context) (feeds model.Feed, err error) {
	c, cancel := context.WithTimeout(ctx, f.Timeout)
	defer cancel()

	feeds, err = f.feedRepository.GetFeeds(c)
	return
}
