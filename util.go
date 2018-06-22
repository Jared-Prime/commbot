package commbot

import (
	"context"
	"fmt"
	"log"
	"os"
)

// EnvarContextKey is used to reference an environmet variable via context.Value
type EnvarContextKey string

// Setup is a utility function for fetching environment variables from the OS
func Setup(ctx context.Context, enVars []string) (context.Context, error) {
	live, err := envar("LIVE_MODE")
	if err != nil {
		// set test node to true
		ctx = context.WithValue(ctx, EnvarContextKey("TEST_MODE"), true)
		ctx = context.WithValue(ctx, EnvarContextKey("LIVE_MODE"), false)
	}

	for _, name := range enVars {
		value, err := envar(name)

		if err != nil {
			return ctx, err
		}

		if live == "" {
			log.Println(value)
		}

		ctx = context.WithValue(ctx, EnvarContextKey(name), value)
	}

	return ctx, nil
}

func envar(name string) (string, error) {
	envar := os.Getenv(name)
	if envar == "" {
		return envar, fmt.Errorf("the environment variable %s is missing", name)
	}

	return envar, nil
}
