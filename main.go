package main

import (
	"log"
	"net/http"

	"github.com/aeramu/nim-finder-server/implementation/handler/graphql"
	"github.com/aeramu/nim-finder-server/implementation/repository/mongodb"
	"github.com/aeramu/nim-finder-server/usecase/service"
	"github.com/friendsofgo/graphiql"
)

//TODOL implement pagination
func main() {
	repo := mongodb.New()
	interactor := service.New(repo)
	handler := graphql.New(interactor)
	http.Handle("/", handler)

	graphiql, _ := graphiql.NewGraphiqlHandler("/")
	http.Handle("/graphiql", graphiql)

	log.Println("server ready")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
