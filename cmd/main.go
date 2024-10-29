package main

import (
	"os"

	"github.com/gofiber/fiber/v2/log"
	clients "github.com/n0o01lh/ml-url-shortener/internals/clients"
	"github.com/n0o01lh/ml-url-shortener/internals/core/services"
	"github.com/n0o01lh/ml-url-shortener/internals/data"
	"github.com/n0o01lh/ml-url-shortener/internals/handlers"
	"github.com/n0o01lh/ml-url-shortener/internals/repositories"
	"github.com/n0o01lh/ml-url-shortener/internals/server"
)

func main() {

	awsAccessKey := os.Getenv("AWS_ACCESS_KEY")
	awsSecretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")

	log.Info("access key " + awsAccessKey)
	log.Info("secret key " + awsSecretAccessKey)

	dynamoDbClient, error := clients.NewDynamoDbClient(awsAccessKey, awsSecretAccessKey)

	if error != nil {
		panic(error)
	}

	dynamodb := data.NewDynamoDb(dynamoDbClient)

	shortRepository := repositories.NewShortRepository(dynamodb)
	resolveRepository := repositories.NewResolverRepository(dynamodb)
	statsRepository := repositories.NewStatsRepository(dynamodb)

	shortService := services.NewShortService(shortRepository)
	statsService := services.NewStatsService(statsRepository)
	resolveService := services.NewResolverService(resolveRepository)
	serviceOrchestrator := services.NewServiceOrchestrator(shortService, resolveService, statsService)

	shortHandlers := handlers.NewShortHandlers(serviceOrchestrator)
	resolveHandlers := handlers.NewResolverHandlers(serviceOrchestrator)
	statsHandlers := handlers.NewStatsHandlers(serviceOrchestrator)

	server := server.NewServer(shortHandlers, resolveHandlers, statsHandlers)

	server.Initialize()
}
