package main

import (
	"net/http"

	"github.com/darahayes/go-hello-server/pkg/config"
	"github.com/darahayes/go-hello-server/pkg/hello"
	"github.com/darahayes/go-hello-server/pkg/web"
	log "github.com/sirupsen/logrus"
)

func main() {
	config := config.GetConfig()

	initLogger(config.LogLevel, config.LogFormat)

	router := web.NewRouter()

	{
		helloService := hello.NewHelloWorldService()
		helloHandler := web.NewHelloHandler(helloService)
		web.SetupHelloRoute(router, helloHandler)
	}

	log.WithFields(log.Fields{"listenAddress": config.ListenAddress}).Info("Starting application")
	log.Fatal(http.ListenAndServe(config.ListenAddress, router))
}

func initLogger(level, format string) {
	logLevel, err := log.ParseLevel(level)

	if err != nil {
		log.Fatalf("log level %v is not allowed. Must be one of [debug, info, warning, error, fatal, panic]", level)
		logLevel = log.InfoLevel
	}

	log.SetLevel(logLevel)

	switch format {
	case "json":
		log.SetFormatter(&log.JSONFormatter{})
	case "text":
		log.SetFormatter(&log.TextFormatter{DisableColors: true})
	default:
		log.Fatalf("log format %v is not allowed. Must be one of [text, json]", format)
	}
}
