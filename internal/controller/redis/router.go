package redis

import (
	"api/internal/usecase"
	"api/pkg/logging"
	"api/pkg/redispubhandler"

	"gopkg.in/redis.v5"
)

func NewRouter(r *redis.Client, l *logging.Logger, t usecase.UserUseCase, b usecase.BeatmapsUseCase) {
	// Not used for now. Just being here to not forget how to do this
	h := NewUserHandler(t, l)
	err := redispubhandler.Handle(r, "keeper:user_update", h)
	if err != nil {
		l.Fatalln(err)
	}
}
