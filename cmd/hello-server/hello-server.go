package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/darahayes/go-hello-server/pkg/config"
	"github.com/darahayes/go-hello-server/pkg/hello"
	"github.com/darahayes/go-hello-server/pkg/web"
)

func main() {
	config := config.GetConfig()

	router := web.NewRouter()

	{
		helloService := hello.NewHelloWorldService()
		helloHandler := web.NewHelloHandler(helloService)
		web.SetupHelloRoute(router, helloHandler)
	}

	fmt.Println(fmt.Sprintf("starting server... going to listen on %v", config.ListenAddress))
	log.Fatal(http.ListenAndServe(config.ListenAddress, router))
}
