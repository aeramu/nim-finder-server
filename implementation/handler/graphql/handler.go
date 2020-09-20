package graphql

import (
	"net/http"

	"github.com/aeramu/nim-finder-server/usecase/adapter"
	"github.com/graph-gophers/graphql-go/relay"
)

type handler struct {
	r     adapter.Repository
	relay *relay.Handler
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.relay.ServeHTTP(w, r)
}
