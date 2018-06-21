package commbot

import (
	"fmt"
	"os"
)

// Envar is a utility function for fetching environment variable from the OS
func Envar(name string) (string, error) {
	envar := os.Getenv(name)
	if envar == "" {
		return envar, fmt.Errorf("the environment variable %s is missing", name)
	}

	return envar, nil
}
