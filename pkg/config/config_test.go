package config

import (
	"os"
	"reflect"
	"testing"
)

func TestConfig(t *testing.T) {
	expected := config{ListenAddress: ":4000"}
	config := GetConfig()

	if !reflect.DeepEqual(expected, config) {
		t.Fatal("GetConfig() did not return expected result")
	}
}

func TestConfigEnvironmentVariables(t *testing.T) {
	expected := config{ListenAddress: ":5000"}

	os.Setenv("PORT", "5000")

	config := GetConfig()

	if !reflect.DeepEqual(expected, config) {
		t.Fatal("GetConfig() did not return expected result")
	}
}
