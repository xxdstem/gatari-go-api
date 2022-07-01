package redis

import (
	"api/internal/usecase"
	"api/pkg/logging"
	"api/pkg/redispubhandler"
	"gopkg.in/redis.v5"
)

func NewRouter(r *redis.Client, l *logging.Logger, t usecase.UserUseCase, b usecase.BeatmapsUseCase) {
	err := redispubhandler.Handle(r, "keeper:user_update", nil)
	if err != nil {
		l.Fatalln(err)
	}
}
