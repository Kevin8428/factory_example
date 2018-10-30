package main

import (
	"net/http"
	"os"

	"github.com/kevin8428/factory_example/api"
)

func main() {
	config := api.Config{
		Env:  setConfigVar("ENV", "dev"),
		Port: setConfigVar("PORT", "8001"),
		Mode: setConfigVar("MODE", "factory"),
	}

	api := api.New()
	server := http.NewServeMux()
	switch config.Mode {
	case "factory":
		api.RegisterHandlers(server)
		api.InitializeHandlerFactory(server)
	default:
		api.InitializeHandler(server)
	}

	panic(http.ListenAndServe(":"+config.Port, server))
}

func setConfigVar(key, fallback string) string {
	val := os.Getenv(key)
	if len(val) == 0 {
		return fallback
	}
	return val
}

// 1. look at greg's PR for passing functions as struct values
// 2. look at thomas' search-service for passing functions around
