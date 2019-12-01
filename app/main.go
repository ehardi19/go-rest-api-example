package main

import (
	"github.com/ehardi19/go-rest-api-example/handlers"
	"github.com/ehardi19/go-rest-api-example/services"
)

func main() {
	services := services.NewServices()
	router := handlers.NewRouter(services)

	router.Run(":8080")
}
