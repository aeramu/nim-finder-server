package mux

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func search(w http.ResponseWriter, r *http.Request) {
	keyword := mux.Vars(r)["keyword"]
	users := i.Search(keyword, 20)
	json.NewEncoder(w).Encode(users)
}
