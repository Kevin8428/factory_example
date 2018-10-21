package main

import (
	"fmt"
	"net/http"

	"github.com/kevin8428/factory_example/api"
)

func main() {
	config := api.Config{
		Env:  "dev",
		Port: "8001",
	}
	// fmt.Println("listening on server ", config.Port)

	api := api.New()
	server := http.NewServeMux()
	api.InitializeHandler(server)
	err := http.ListenAndServe(":"+config.Port, server)
	fmt.Println("listening on server ", config.Port)
	if err != nil {
		fmt.Println("error starting server: ", err)
	} else {
		fmt.Println("listening on server ")
	}
}
