package config

import (
	"fmt"
	"os"
)

type config struct {
	ListenAddress string //will look like ":4000"
}

func GetConfig() config {
	return config{
		ListenAddress: fmt.Sprintf(":%v", getEnv("PORT", "4000")),
	}
}

func getEnv(key string, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}
