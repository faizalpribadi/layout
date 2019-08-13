package handler

import (
	"net/http"

	"github.com/faizalpribadi/layout/pkg/helper"

	"github.com/faizalpribadi/layout/internal/domain"
	log "github.com/sirupsen/logrus"
)

type FeedHandler struct {
	usecase domain.FeedUseCase
}

func NewFeedHandler(usecase domain.FeedUseCase) domain.FeedHandler {
	return &FeedHandler{usecase}
}

func (h *FeedHandler) GetFeeds(rw http.ResponseWriter, r *http.Request) {
	feeds, err := h.usecase.GetFeeds(r.Context())
	if err != nil {
		log.Error(err)
		return
	}

	helper.APIResponse(rw, http.StatusOK, feeds)
}
