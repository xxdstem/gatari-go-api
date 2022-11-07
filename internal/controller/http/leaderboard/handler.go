package leaderboard

import (
	"api/internal/usecase"
	"api/pkg/logging"

	"github.com/julienschmidt/httprouter"
)

const (
	GlobalURL = "/leaderboard/:country"
)

type handler struct {
	logger *logging.Logger
	t      usecase.LeaderboardUseCase
}

func New(t usecase.LeaderboardUseCase, l *logging.Logger) Handler {
	return &handler{
		logger: l,
		t:      t,
	}
}

func (h *handler) Register(router *httprouter.Router) {

}
