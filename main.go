package main

import (
	"log"
	"net/http"

	"github.com/aeramu/nim-finder-server/implementation/handler/mux"
	"github.com/aeramu/nim-finder-server/implementation/repository/mongodb"
	"github.com/aeramu/nim-finder-server/usecase/service"
)

func main() {
	repo := mongodb.New()
	interactor := service.New(repo)
	handler := mux.New(interactor)

	log.Println("server ready")
	log.Fatal(http.ListenAndServe(":8000", handler))
}
