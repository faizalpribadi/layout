package domain

import (
	"context"
	"net/http"

	"github.com/faizalpribadi/layout/internal/model"
)

type (
	//Specific domain contract interface for Feeds
	FeedRepository interface {
		GetFeeds(ctx context.Context) (model.Feed, error)
	}
	UserRepository interface {
		GetUsers(ctx context.Context) (model.User, error)
	}
	FeedUseCase interface {
		GetFeeds(ctx context.Context) (feeds model.Feed, err error)
	}
	UserUseCase interface {
		GetUsers(ctx context.Context) (users model.User, err error)
	}

	UserHandler interface {
		GetUsers(rw http.ResponseWriter, r *http.Request)
	}
	FeedHandler interface {
		GetFeeds(rw http.ResponseWriter, r *http.Request)
	}
	Server interface {
		//Add method here
		Handler() http.Handler
		Run()
	}
)
