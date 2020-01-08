package utils

import (
	"balena-factoryreset-example/constants"
	"errors"
	"os"
)

// GetEnv Get value of environmemnt variable. When empty fallback to defaultValue
func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

// GetSupervisorURI Get env value to the supervisor uri
func GetSupervisorURI() (string, error) {
	value := os.Getenv(constants.BalenaSupervisorURIEnv)
	if len(value) > 0 {
		return value, nil
	}
	return "", errors.New("Could not get supervisor address")
}

// GetSupervisorAPIKey Get env value Api Key to the supoervisor
func GetSupervisorAPIKey() (string, error) {
	value := os.Getenv(constants.BalenaSupervisorAPIKeyEnv)
	if len(value) > 0 {
		return value, nil
	}
	return "", errors.New("Could not get supervisor api key")
}
