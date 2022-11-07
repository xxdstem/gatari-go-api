package leaderboard

import (
	"api/internal/entity"
	"api/internal/usecase"
	"api/pkg/logging"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

const (
	userURL      = "/user/:uid/:mode"
	userStatsURL = "/user/:uid/:mode/stats"
)

type handler struct {
	logger *logging.Logger
	t      usecase.UserUseCase
}

func New(t usecase.UserUseCase, l *logging.Logger) Handler {
	return &handler{
		logger: l,
		t:      t,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(userURL, h.GetUser)
	router.GET(userStatsURL, h.GetUserStats)

}