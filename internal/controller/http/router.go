package http

import (
	"api/internal/controller/http/users"
	"api/internal/usecase"
	"api/pkg/logging"

	"github.com/julienschmidt/httprouter"
)

func NewUsersRoute(r *httprouter.Router, l *logging.Logger, t usecase.UserUseCase, b usecase.BeatmapsUseCase) {
	userHandler := users.New(t, l)
	userHandler.Register(r)
}
