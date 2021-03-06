package main

import (
	"log"
	"net/http"
	"os"

	"github.com/aeramu/nim-finder-server/internal/handler/graphql"
	"github.com/aeramu/nim-finder-server/internal/repository/elasticsearch"
	"github.com/aeramu/nim-finder-server/internal/search"
	"github.com/friendsofgo/graphiql"
)

//TODOL implement pagination
func main() {
	repo := elasticsearch.NewStudentRepo()
	service := search.NewService(repo)
	handler := graphql.NewHandler(service)
	http.Handle("/", enableCORS(handler))

	graphiql, _ := graphiql.NewGraphiqlHandler("/")
	http.Handle("/graphiql", graphiql)

	port := port()

	log.Println("Server ready at", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func port() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	return ":" + port
}

func enableCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		h.ServeHTTP(w, r)
	})
}
