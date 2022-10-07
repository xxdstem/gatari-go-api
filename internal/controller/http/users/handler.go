package users

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

	//userModeURL  = "/user/:uid/:mode"
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

func (h *handler) GetUserStats(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	h.logger.Info("hi hitler!")
	param := params.ByName("uid")
	uid, err := strconv.Atoi(param)
	if err != nil {
		h.logger.Fatal(err)
	}
	var user *entity.UserStats
	mode := params.ByName("mode")
	if modeInt, err := strconv.Atoi(mode); err == nil {
		user = h.t.GetUserStatsByID(uid, int8(modeInt))
	} else {
		user = h.t.GetUserStatsByID(uid, -1)
	}
	j, err := json.Marshal(user)
	if err != nil {
		h.logger.Fatal(err)
	}
	if _, err := w.Write(j); err != nil {
		h.logger.Error(err)
	}
}

func (h *handler) GetUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	param := params.ByName("uid")

	h.logger.Infoln(param)
	uid, err := strconv.Atoi(param)
	if err != nil {
		h.logger.Fatal(err)
	}

	user := h.t.GetUserById(uid)
	j, err := json.Marshal(user)
	if err != nil {
		h.logger.Fatal(err)
	}
	if _, err := w.Write(j); err != nil {
		h.logger.Error(err)
	}

}
