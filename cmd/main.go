package main

import (
	"github.com/n0o01lh/ml-url-shortener/internals/core/services"
	"github.com/n0o01lh/ml-url-shortener/internals/handlers"
	"github.com/n0o01lh/ml-url-shortener/internals/repositories"
	"github.com/n0o01lh/ml-url-shortener/internals/server"
)

func main() {

	shortRepository := repositories.NewShortRepository()
	shortService := services.NewShortService(shortRepository)
	shortHandlers := handlers.NewShortHandlers(shortService)

	server := server.NewServer(shortHandlers)

	server.Initialize()
}
