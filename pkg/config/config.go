package config

import (
	"fmt"
	"os"
	"strings"
)

type config struct {
	ListenAddress string //will look like ":4000"
	LogLevel      string
	LogFormat     string
}

func GetConfig() config {
	return config{
		ListenAddress: fmt.Sprintf(":%v", getEnv("PORT", "4000")),
		LogLevel:      strings.ToLower(getEnv("LOG_LEVEL", "info")),
		LogFormat:     strings.ToLower(getEnv("LOG_FORMAT", "text")), //cann be text or json
	}
}

func getEnv(key string, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}
