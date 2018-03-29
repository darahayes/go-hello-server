package web

import "github.com/gorilla/mux"

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.Use(loggerMiddleWare)
	return router
}

func SetupHelloRoute(r *mux.Router, handler *helloHandler) {
	r.HandleFunc("/", handler.HelloEndpoint).Methods("GET")
}
