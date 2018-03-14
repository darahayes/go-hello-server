package web

import "github.com/gorilla/mux"

func NewRouter() *mux.Router {
	return mux.NewRouter().StrictSlash(true)
}

func SetupHelloRoute(r *mux.Router, handler *helloHandler) {
	r.HandleFunc("/hello", handler.HelloEndpoint).Methods("GET")
}
