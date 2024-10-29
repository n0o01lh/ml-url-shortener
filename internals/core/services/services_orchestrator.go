package services

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/n0o01lh/ml-url-shortener/internals/core/domain"
	"github.com/n0o01lh/ml-url-shortener/internals/core/ports"
)

type ServiceOrchestrator struct {
	shortService    ports.ShortService
	resolverService ports.ResolverService
	statsService    ports.StatsService
}

func NewServiceOrchestrator(shortService ports.ShortService, resolveService ports.ResolverService, statsService ports.StatsService) *ServiceOrchestrator {
	return &ServiceOrchestrator{
		shortService:    shortService,
		resolverService: resolveService,
		statsService:    statsService,
	}
}

var _ ports.ServiceOrchestrator = (*ServiceOrchestrator)(nil)

func (c *ServiceOrchestrator) CreateShortUrl(request *domain.ShortRequest) (*domain.ShortedUrl, error) {

	shortedUrl, err := c.shortService.Create(request)

	if err != nil {
		log.Error(err)
		return nil, err
	}

	if err = c.statsService.Create(shortedUrl.Id); err != nil {
		log.Error(err)
	}

	return shortedUrl, nil
}

func (c *ServiceOrchestrator) UpdateShortUrl(id string, request *domain.ShortRequest) (*domain.ShortedUrl, error) {

	shortedUrl, err := c.shortService.Update(id, request)

	if err != nil {
		log.Error(err)
		return nil, err
	}

	shortedUrlUpdated, _ := c.resolverService.Get(shortedUrl.Id)

	return shortedUrlUpdated, nil
}

func (c *ServiceOrchestrator) GetShortUrl(id string) (*domain.ShortedUrl, error) {

	shortedUrlUpdated, err := c.resolverService.Get(id)

	if err != nil {
		log.Error(err)
		return nil, err
	}

	if err = c.statsService.Update(shortedUrlUpdated.Id); err != nil {
		log.Error(err)
	}

	return shortedUrlUpdated, nil
}

func (c *ServiceOrchestrator) GetStats(id string) (*domain.Stats, error) {

	stats, err := c.statsService.Get(id)

	if err != nil {
		log.Error(err)
		return nil, err
	}

	return stats, nil
}
