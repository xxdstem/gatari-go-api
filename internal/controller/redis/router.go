package redis

import (
	"api/internal/usecase"
	"api/pkg/redispubhandler"
	"log"

	"gopkg.in/redis.v5"
)

func NewRouter(r *redis.Client, t usecase.UserRepository) {
	err := redispubhandler.Handle(r, "api:user_update", NewUserHandler(t))
	if err != nil {
		log.Fatal(err)
	}
}
