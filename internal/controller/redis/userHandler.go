package redis

import (
	"api/internal/usecase"
	"api/pkg/logging"
	"api/pkg/redispubhandler"
)

type handler struct {
	t usecase.UserUseCase
	l *logging.Logger
}

// type result struct {
// 	UserID int `json:"user_id"`
// }

func NewUserHandler(t usecase.UserUseCase, l *logging.Logger) *handler {
	return &handler{
		l: l,
		t: t,
	}
}

func (b *handler) Response(r *redispubhandler.Context) {
	if r.Error != nil {
		b.l.Errorln(r.Error)
	}

}
