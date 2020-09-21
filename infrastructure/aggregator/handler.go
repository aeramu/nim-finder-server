package aggregator

import (
	"net/http"

	"github.com/friendsofgo/graphiql"

	"github.com/aeramu/nim-finder-server/implementation/handler/graphql"
	"github.com/aeramu/nim-finder-server/implementation/repository/mongodb"
	"github.com/aeramu/nim-finder-server/usecase/service"
)

//NewHandler return functional handler for aggregation
func NewHandler() http.Handler {
	repo := mongodb.New()
	interactor := service.New(repo)
	handler := graphql.New(interactor)
	http.Handle("/", handler)

	graphiql, _ := graphiql.NewGraphiqlHandler("/")
	http.Handle("/graphiql", graphiql)

	return http.DefaultServeMux
}
