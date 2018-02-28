package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func helloEndpoint(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Hello World!"))
}

func newRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/hello", helloEndpoint).Methods("GET")
	return router
}

func main() {
	router := newRouter()
	log.Fatal(http.ListenAndServe(":8000", router))
}
