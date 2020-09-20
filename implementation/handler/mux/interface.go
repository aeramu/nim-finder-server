package mux

import (
	"net/http"

	"github.com/aeramu/nim-finder-server/usecase/service"
	"github.com/gorilla/mux"
)

var i service.Interactor

//New handler
func New(interactor service.Interactor) http.Handler {
	i = interactor
	r := mux.NewRouter()
	r.HandleFunc("/search", search)

	return r
}
