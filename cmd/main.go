package main

import (
	"log"
	"net/http"

	"github.com/git-Suyash/go-BookStore/pkg/routes"
	"github.com/gorilla/mux"
)

const PORT = ":3000"

func main() {

	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(PORT, r))

}
