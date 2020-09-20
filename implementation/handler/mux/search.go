package mux

import (
	"encoding/json"
	"net/http"
)

func search(w http.ResponseWriter, r *http.Request) {
	users := i.Search("yud")
	json.NewEncoder(w).Encode(users)
}
