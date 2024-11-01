package main

import (
	"os"

	clients "github.com/n0o01lh/ml-url-shortener/internals/clients"
	"github.com/n0o01lh/ml-url-shortener/internals/core/services"
	"github.com/n0o01lh/ml-url-shortener/internals/data"
	"github.com/n0o01lh/ml-url-shortener/internals/handlers"
	"github.com/n0o01lh/ml-url-shortener/internals/repositories"
	"github.com/n0o01lh/ml-url-shortener/internals/server"
	"github.com/n0o01lh/ml-url-shortener/internals/workers"
	"golang.org/x/time/rate"
)

func main() {

	awsAccessKey := os.Getenv("AWS_ACCESS_KEY")
	awsSecretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")

	dynamoDbClient, error := clients.NewDynamoDbClient(awsAccessKey, awsSecretAccessKey)

	if error != nil {
		panic(error)
	}

	dynamodb := data.NewDynamoDb(dynamoDbClient)
	jobQueue := workers.NewJobQueue(5)
	limiter := rate.NewLimiter(rate.Limit(25), 2)

	shortRepository := repositories.NewShortRepository(dynamodb)
	resolveRepository := repositories.NewResolverRepository(dynamodb)
	statsRepository := repositories.NewStatsRepository(dynamodb)

	shortService := services.NewShortService(shortRepository)
	statsService := services.NewStatsService(statsRepository)
	resolveService := services.NewResolverService(resolveRepository)
	serviceOrchestrator := services.NewServiceOrchestrator(shortService, resolveService, statsService, jobQueue)

	shortHandlers := handlers.NewShortHandlers(serviceOrchestrator)
	resolveHandlers := handlers.NewResolverHandlers(serviceOrchestrator, limiter)
	statsHandlers := handlers.NewStatsHandlers(serviceOrchestrator)

	server := server.NewServer(shortHandlers, resolveHandlers, statsHandlers)
	server.Initialize()

	defer close(jobQueue.JobChannel)
	jobQueue.Wg.Wait()
	close(jobQueue.ResultChannel)
}
