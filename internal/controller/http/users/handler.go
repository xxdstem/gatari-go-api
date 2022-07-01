package users

import (
	"api/internal/usecase"
	"api/pkg/logging"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

const (
	userURL     = "/users/:uid"
	userModeURL = "/users/:uid/:mode"
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
	router.GET(userModeURL, h.GetUser)

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
