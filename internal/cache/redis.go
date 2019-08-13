package cache

import (
	"github.com/faizalpribadi/layout/internal/config"
)

type (
	Redis interface{}
	redis struct {
		cfg config.Configuration
	}
)

func NewRedisCache(cfg config.Configuration) *redis {
	return &redis{cfg}
}

//Add ur method here
