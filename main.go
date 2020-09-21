package main

import (
	"log"
	"net/http"
	"os"

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

	port := port()

	log.Println("Server ready at ", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func port() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	return ":" + port
}
