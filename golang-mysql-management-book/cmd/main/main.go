package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/relucat1/go-bookstore/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookstores(r)
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)

}
