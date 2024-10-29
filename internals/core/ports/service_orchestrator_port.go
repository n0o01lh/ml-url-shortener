package ports

import "github.com/n0o01lh/ml-url-shortener/internals/core/domain"

type ServiceOrchestrator interface {
	CreateShortUrl(shortRequest *domain.ShortRequest) (*domain.ShortedUrl, error)
	UpdateShortUrl(id string, shortRequest *domain.ShortRequest) (*domain.ShortedUrl, error)
	GetShortUrl(id string) (*domain.ShortedUrl, error)
	GetStats(id string) (*domain.Stats, error)
}
