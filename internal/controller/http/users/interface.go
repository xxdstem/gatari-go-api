package users

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Handler interface {
	GetUser(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Register(router *httprouter.Router)
}
