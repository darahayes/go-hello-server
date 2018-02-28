package main

import (
	"fmt"
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
	listenAddr := ":8000"
	router := newRouter()
	fmt.Println(fmt.Sprintf("starting server... going to listen on %v", listenAddr))
	log.Fatal(http.ListenAndServe(":8000", router))
}
