package handler

import (
	"net/http"

	"github.com/faizalpribadi/layout/pkg/helper"

	"github.com/faizalpribadi/layout/internal/domain"
	log "github.com/sirupsen/logrus"
)

type UserHandler struct {
	usecase domain.UserUseCase
}

func NewUserHandler(usecase domain.UserUseCase) domain.UserHandler {
	return &UserHandler{usecase}
}

func (u *UserHandler) GetUsers(rw http.ResponseWriter, r *http.Request) {
	feeds, err := u.usecase.GetUsers(r.Context())
	if err != nil {
		log.Error(err)
		return
	}

	helper.APIResponse(rw, http.StatusOK, feeds)
}
