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
	FeedRepository struct {
		Timeout time.Duration
	}
)

func NewFeedRepository() domain.FeedRepository {
	return &FeedRepository{
		Timeout: constant.Timeout * time.Second,
	}
}

func (f *FeedRepository) GetFeeds(ctx context.Context) (model.Feed, error) {
	feed := model.Feed{}
	uid := helper.GenerateUID()
	feed.ID = uid
	return feed, nil
}
